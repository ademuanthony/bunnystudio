package endpoints

import (
	"context"

	"bonnystudio.com/taskmanager/internal"
	"bonnystudio.com/taskmanager/pkg/tasks"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetByUserIDEndpoint endpoint.Endpoint
	CreateEndpoint      endpoint.Endpoint
	UpdateEndpoint      endpoint.Endpoint
}

func NewEndpointSet(svc tasks.Service) Set {
	return Set{
		GetByUserIDEndpoint: MakeGetAllEndpoint(svc),
		CreateEndpoint:      MakeCreateEndpoint(svc),
		UpdateEndpoint:      MakeUpdateEndpoint(svc),
	}
}

func MakeGetAllEndpoint(svc tasks.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByUserIDRequest)
		tasks, err := svc.GetByUserID(ctx, req.UserID)
		if err != nil {
			return GetByUserIDResponse{Tasks: nil, Err: err.Error()}, nil
		}
		return GetByUserIDResponse{Tasks: tasks, Err: ""}, nil
	}
}

func MakeCreateEndpoint(svc tasks.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		task, err := svc.Create(ctx, internal.Task{Description: req.Description, State: req.State, UserID: req.UserID})
		if err != nil {
			return CreateResponse{Task: task, Err: err.Error()}, nil
		}
		return CreateResponse{Task: task, Err: ""}, nil
	}
}

func MakeUpdateEndpoint(svc tasks.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		err := svc.Update(ctx, req.Task)
		if err != nil {
			return UpdateResponse{Err: err.Error()}, nil
		}
		return UpdateResponse{Err: ""}, nil
	}
}
