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

type getBestEmployeeRequest struct{}
type addEmployeeRequest struct {
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	LastName      string
	MobilePhone   string
	JobTitle      string
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

func makeGetBestEmployeeEndPoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndPoint := func(_ context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestEmployee()
		helper.Catch(err)
		return result, nil
	}
	return getBestEmployeeEndPoint
}

func MakeInsertEmployeeEndpoint(s Service) endpoint.Endpoint {
	getInsertEmployeeEndPoint := func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeeRequest)
		result, err := s.InsertEmployee(&req)
		helper.Catch(err)
		return result, nil
	}
	return getInsertEmployeeEndPoint

}
