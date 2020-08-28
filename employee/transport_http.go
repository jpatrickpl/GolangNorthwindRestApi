package employee

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getEmployeesHandler := kithttp.NewServer(makeGetEmployeesEndPoint(s), getEmployeesRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	getEmployeeByIdHandler := kithttp.NewServer(makeGetEmployeeByIdEndPoint(s), getEmployeeByIDRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getEmployeeByIdHandler)

	return r
}
func getEmployeesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getEmployeeByIDRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getEmployeeByIdRequest{EmployeeID: chi.URLParam(r, "id")}, nil
}
