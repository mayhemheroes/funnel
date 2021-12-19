package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ohsu-comp-bio/funnel/events"
	"github.com/ohsu-comp-bio/funnel/logger"
	"github.com/ohsu-comp-bio/funnel/tes"
	"github.com/ohsu-comp-bio/funnel/tes/openapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// TaskServiceApiService is a service that implents the logic for the TaskServiceApiServicer
// This service should implement the business logic for every endpoint for the TaskServiceApi API.
// Include any external packages or services that will be required by this service.
type TaskServiceApiService struct {
	Name    string
	Event   events.Writer
	Compute events.Writer
	Read    tes.ReadOnlyServer
	Log     *logger.Logger
}

// NewTaskServiceApiService creates a default api service
func NewTaskServiceApiService() openapi.TaskServiceApiServicer {
	return &TaskServiceApiService{}
}

// CancelTask - CancelTask
func (s *TaskServiceApiService) CancelTask(ctx context.Context, id string) (openapi.ImplResponse, error) {
	// TODO - update CancelTask with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("CancelTask method not implemented")
}

// CreateTask - CreateTask
func (s *TaskServiceApiService) CreateTask(ctx context.Context, body openapi.TesTask) (openapi.ImplResponse, error) {
	// TODO - update CreateTask with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TesCreateTaskResponse{}) or use other options such as http.Ok ...
	//return Response(200, TesCreateTaskResponse{}), nil

	task := convert.O2PTask(body)

	if err := tes.InitTask(task, true); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	if err := ts.Event.WriteEvent(ctx, events.NewTaskCreated(task)); err != nil {
		return nil, fmt.Errorf("error creating task: %s", err)
	}

	// dispatch to compute backend
	go func() {
		err := ts.Compute.WriteEvent(ctx, events.NewTaskCreated(task))
		if err != nil {
			ts.Log.Error("error submitting task to compute backend", "taskID", task.Id, "error", err)
		}
	}()

	return &tes.CreateTaskResponse{Id: task.Id}, nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("CreateTask method not implemented")
}

// GetServiceInfo - GetServiceInfo
func (s *TaskServiceApiService) GetServiceInfo(ctx context.Context) (openapi.ImplResponse, error) {
	// TODO - update GetServiceInfo with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TesServiceInfo{}) or use other options such as http.Ok ...
	//return Response(200, TesServiceInfo{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetServiceInfo method not implemented")
}

// GetTask - GetTask
func (s *TaskServiceApiService) GetTask(ctx context.Context, id string, view string) (openapi.ImplResponse, error) {
	// TODO - update GetTask with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TesTask{}) or use other options such as http.Ok ...
	//return Response(200, TesTask{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetTask method not implemented")
}

// ListTasks - ListTasks
func (s *TaskServiceApiService) ListTasks(ctx context.Context, namePrefix string, state openapi.TesState, tags map[string]string, pageSize int64, pageToken string, view string) (openapi.ImplResponse, error) {
	// TODO - update ListTasks with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TesListTasksResponse{}) or use other options such as http.Ok ...
	//return Response(200, TesListTasksResponse{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ListTasks method not implemented")
}
