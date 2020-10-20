package customer

import (
	"context"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getCustomerRequest struct {
	Limit  int
	Offset int
}

func makeGetCustomersEndPoint(s Service) endpoint.Endpoint {
	getCustomerEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomerRequest)
		result, err := s.GetCustomers(&req)
		helper.Catch(err)
		return result, nil
	}
	return getCustomerEndPoint
}
