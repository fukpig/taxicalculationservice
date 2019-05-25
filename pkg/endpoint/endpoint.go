package endpoint

import (
	"context"

	service "github.com/fukpig/taxicalculationservice/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetPriceRequest collects the request parameters for the GetPrice method.
type GetPriceRequest struct {
	From     string `schema:"from"`
	To       string `schema:"to"`
	RateName string `schema:"rate-name"`
	RateType string `schema:"rate-type"`
}

// GetPriceResponse collects the response parameters for the GetPrice method.
type GetPriceResponse struct {
	Price int32 `json:"price"`
	Error error `json:"error"`
}

// MakeGetPriceEndpoint returns an endpoint that invokes GetPrice on the service.
func MakeGetPriceEndpoint(s service.TaxicalculationserviceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPriceRequest)
		price, err := s.GetPrice(ctx, req.From, req.To, req.RateName, req.RateType)
		return GetPriceResponse{Price: price, Error: err}, nil
	}
}

// Failed implements Failer.
func (r GetPriceResponse) Failed() error {
	return r.Error
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetPrice implements Service. Primarily useful in a client.
func (e Endpoints) GetPrice(ctx context.Context, from, to, rateName, rateType string) (e0 error) {
	request := GetPriceRequest{
		From:     from,
		To:       to,
		RateName: rateName,
		RateType: rateType,
	}
	response, err := e.GetPriceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetPriceResponse).Error
}
