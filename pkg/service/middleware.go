package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(TaxicalculationserviceService) TaxicalculationserviceService

type loggingMiddleware struct {
	logger log.Logger
	next   TaxicalculationserviceService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TaxicalculationserviceService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TaxicalculationserviceService) TaxicalculationserviceService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetPrice(ctx context.Context, from, to, rateName, rateType string) (price int32, err error) {
	defer func() {
		l.logger.Log("method", "GetPrice", "from", from, "to", to, "rateName", rateName, "rateType", rateType, "err", err)
	}()
	return l.next.GetPrice(ctx, from, to, rateName, rateType)
}
