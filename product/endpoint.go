package product

import (
	"context"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getProductByIDrequest struct {
	ProductID int
}
type getProductsRequest struct {
	Limit  int
	Offset int
}
type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type updateProductRequest struct {
	Id           int64
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductId string
}
type getBestSellerRequest struct {
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDrequest)
		product, err := s.GetProductById(&req)

		helper.Catch(err)
		return product, nil
	}
	return getProductByIdEndPoint
}

func makeGetProductEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		helper.Catch(err)
		return result, nil
	}

	return getProductsEndPoint
}

func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	AddProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)
		helper.Catch(err)
		return productId, nil
	}

	return AddProductsEndPoint
}

func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	UpdateProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		productId, err := s.UpdateProduct(&req)
		helper.Catch(err)
		return productId, nil
	}
	return UpdateProductEndPoint
}

func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	DeleteProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		productId, err := s.DeleteProduct(&req)
		helper.Catch(err)
		return productId, nil
	}

	return DeleteProductEndPoint
}

func makeBestSellersEndpoint(s Service) endpoint.Endpoint {
	getbestSellersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestSellers()
		helper.Catch(err)
		return result, nil
	}
	return getbestSellersEndpoint
}
