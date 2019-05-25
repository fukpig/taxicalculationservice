package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	endpoint "github.com/fukpig/taxicalculationservice/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeGetPriceHandler creates the handler logic
func makeGetPriceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-price", http1.NewServer(endpoints.GetPriceEndpoint, decodeGetPriceRequest, encodeGetPriceResponse, options...))
}

// decodeGetPriceRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetPriceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetPriceRequest{}
	//err := json.NewDecoder(r.Body).Decode(&req)
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing request: %s", err)
		return req, err
	}
	req.From = r.Form.Get("from")
	req.To = r.Form.Get("to")
	req.RateName = r.Form.Get("rate-name")
	req.RateType = r.Form.Get("rate-type")

	if req.From == "" {
		return req, errors.New("invalid from address")
	}

	if req.To == "" {
		return req, errors.New("invalid to address")
	}

	if req.RateName == "" {
		req.RateName = "basic"
	}

	if req.RateType == "" {
		req.RateType = "per-km"
	}

	return req, nil
}

// encodeGetPriceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetPriceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
