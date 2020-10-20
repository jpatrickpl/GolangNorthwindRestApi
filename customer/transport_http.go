package customer

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

	getCustomerHandler := kithttp.NewServer(makeGetCustomersEndPoint(s),
		getCustomerRequestDecoder, kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginated", getCustomerHandler)
	return r
}
func getCustomerRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getCustomerRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil

}
