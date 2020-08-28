package employee

import (
	"context"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}
type getEmployeeByIdRequest struct {
	EmployeeID string
}

func makeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployeesEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		result, err := s.GetEmployees(&req)
		helper.Catch(err)
		return result, nil
	}

	return getEmployeesEndPoint
}

func makeGetEmployeeByIdEndPoint(s Service) endpoint.Endpoint {
	getEmployeeByIdRequest := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIdRequest)
		result, err := s.GetEmployeeById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeeByIdRequest
}
