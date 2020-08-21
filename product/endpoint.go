package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDrequest struct {
	ProductID int
}
type getProductsRequest struct {
	Limit  int
	Offset int
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDrequest)
		product, err := s.GetProductById(&req)

		if err != nil {
			panic(err)
		}	
		return product, nil
	}
	return getProductByIdEndPoint
}

func makeGetProductEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}

	return getProductsEndPoint
}
