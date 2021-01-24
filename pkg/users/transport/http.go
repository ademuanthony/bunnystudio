package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"bonnystudio.com/taskmanager/internal/util"
	"bonnystudio.com/taskmanager/pkg/users/endpoints"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()
	m.Handle("/create", httptransport.NewServer(
		ep.CreateEndpoint,
		decodeHTTPCreateRequest,
		encodeResponse,
	))
	m.Handle("/update", httptransport.NewServer(
		ep.UpdateEndpoint,
		decodeHTTPUpdateRequest,
		encodeResponse,
	))
	m.Handle("/get", httptransport.NewServer(
		ep.GetEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))
	m.Handle("/getall", httptransport.NewServer(
		ep.GetAllEndpoint,
		decodeHTTPGetAllRequest,
		encodeResponse,
	))
	m.Handle("/remove", httptransport.NewServer(
		ep.DeleteEndpoint,
		decodeHTTPRemoveRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetRequest
	if r.ContentLength == 0 {
		logger.Log("Get request with no body")
		return req, nil
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoints.GetAllRequest{}, nil
}

func decodeHTTPUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPRemoveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
