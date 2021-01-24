package endpoints

import (
	"context"

	"bonnystudio.com/taskmanager/pkg/users"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetEndpoint    endpoint.Endpoint
	GetAllEndpoint endpoint.Endpoint
	CreateEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc users.Service) Set {
	return Set{
		GetEndpoint:    MakeGetEndpoint(svc),
		GetAllEndpoint: MakeGetAllEndpoint(svc),
		CreateEndpoint: MakeCreateEndpoint(svc),
		UpdateEndpoint: MakeUpdateEndpoint(svc),
		DeleteEndpoint: MakeDeleteEndpoint(svc),
	}
}

func MakeGetEndpoint(svc users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		user, err := svc.Get(ctx, req.UserID)
		if err != nil {
			return GetResponse{User: nil, Err: err.Error()}, nil
		}
		return GetResponse{User: user, Err: ""}, nil
	}
}

func MakeGetAllEndpoint(svc users.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		users, err := svc.GetAll(ctx)
		if err != nil {
			return GetAllResponse{Users: nil, Err: err.Error()}, nil
		}
		return GetAllResponse{Users: users, Err: ""}, nil
	}
}

func MakeCreateEndpoint(svc users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		user, err := svc.Create(ctx, req.Name)
		if err != nil {
			return CreateResponse{User: user, Err: err.Error()}, nil
		}
		return CreateResponse{User: user, Err: ""}, nil
	}
}

func MakeUpdateEndpoint(svc users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		err := svc.Update(ctx, req.User)
		if err != nil {
			return UpdateResponse{Err: err.Error()}, nil
		}
		return UpdateResponse{Err: ""}, nil
	}
}

func MakeDeleteEndpoint(svc users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := svc.Delete(ctx, req.UserID)
		if err != nil {
			return DeleteResponse{Err: err.Error()}, nil
		}
		return DeleteResponse{Err: ""}, nil
	}
}
