package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getProductByIdHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s),
		getProductByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	getProductHandler := kithttp.NewServer(makeGetProductEndPoint(s), getProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductHandler)
	return r

}

func getProductByIdRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	var ID int
	ID, _ = strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDrequest{
		ProductID: ID}, nil

}

func getProductRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}
