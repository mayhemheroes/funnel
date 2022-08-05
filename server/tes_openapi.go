package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/ohsu-comp-bio/funnel/tes"
	"github.com/ohsu-comp-bio/funnel/tes/openapi"
)

// TaskServiceApiService is a service that implents the logic for the TaskServiceApiServicer
// This service should implement the business logic for every endpoint for the TaskServiceApi API.
// Include any external packages or services that will be required by this service.
type TaskServiceApiService struct {
	server *TaskService
}

func NewOpenApiServer(srv *TaskService) *TaskServiceApiService {
	return &TaskServiceApiService{srv}
}

// NewTaskServiceApiService creates a default api service
//func NewTaskServiceApiService() openapi.TaskServiceApiServicer {
//	return &TaskServiceApiService{}
//}

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
	task := &tes.Task{}
	tes.OpenApi2Proto(body, task)

	res, err := s.server.CreateTask(ctx, task)

	if err != nil {
		return openapi.Response(500, err), nil
	}

	return openapi.Response(200, openapi.TesCreateTaskResponse{Id: res.Id}), nil
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

	req := &tes.GetTaskRequest{Id: id, View: view}

	task, err := s.server.GetTask(ctx, req)

	if err != nil {
		return openapi.Response(http.StatusNotFound, nil), nil
	}

	o := openapi.TesTask{}
	tes.Proto2OpenApi(task, &o)
	return openapi.Response(200, o), err
}

// ListTasks - ListTasks
func (s *TaskServiceApiService) ListTasks(ctx context.Context, namePrefix string,
	state openapi.TesState, tagKey []string, tagValue []string,
	pageSize int64, pageToken string, view string) (openapi.ImplResponse, error) {

	req := &tes.ListTasksRequest{
		NamePrefix: namePrefix,
		//State:      state,
	}

	out, err := s.server.ListTasks(ctx, req)
	if err != nil {
		return openapi.Response(http.StatusNotFound, nil), nil
	}

	resp := openapi.TesListTasksResponse{}
	tes.Proto2OpenApi(out, resp)

	return openapi.Response(200, resp), err
}
