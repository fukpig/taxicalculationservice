package service

import (
	"context"
	"encoding/json"

	"github.com/fukpig/geoservice/proto/tripInfo"
	pb "github.com/fukpig/geoservice/proto/tripInfo"
	"github.com/fukpig/taxicalculationservice/pkg/rate"
	"github.com/go-pg/pg"
	"go.opencensus.io/trace"
)

const carPrice = 50
const minTripPrice = 150

type TaxicalculationserviceServiceRepository interface {
	FindRateByName(name string) (rate *rate.Rate, err error)
}

// TaxicalculationserviceService describes the service.
type TaxicalculationserviceService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetPrice(ctx context.Context, from, to, rateName, rateType string) (price int32, err error)
}

type basicTaxicalculationserviceService struct {
	geoServiceClient tripInfo.GeoServiceClient
	repository       TaxicalculationserviceServiceRepository
}

func (b *basicTaxicalculationserviceService) GetPrice(ctx context.Context, from, to, rateName, rateType string) (price int32, err error) {
	ctx, span := trace.StartSpan(context.Background(), "get-price")
	defer span.End()
	span.AddAttributes(
		trace.StringAttribute("method", "GET"),
	)

	spanContextJson, err := json.Marshal(span.SpanContext())
	rate, err := b.repository.FindRateByName(rateName)
	if err != nil {
		return 0, err
	}

	route := &pb.Route{From: from, To: to, SpanContext: string(spanContextJson)}

	tripInfo, err := b.geoServiceClient.GetTripInfo(context.Background(), route)
	if err != nil {
		return 0, err
	}

	price, err = b.calculatePrice(rate, rateType, tripInfo)
	return price, err
}

//With rate and distance and time data from geoservice calculate price of trip
func (b *basicTaxicalculationserviceService) calculatePrice(rate *rate.Rate, rateType string, data *pb.Response) (price int32, err error) {

	var ratePrice int32
	if rateType == "per-minute" {
		ratePrice = rate.PerMinute * data.Duration
	} else {
		ratePrice = rate.PerKm * data.Distance
	}

	total := carPrice + ratePrice
	if total < minTripPrice {
		total = 150
	}
	return total, nil
}

// NewBasicTaxicalculationserviceService returns a naive, stateless implementation of TaxicalculationserviceService.
func NewBasicTaxicalculationserviceService(db *pg.DB, geoServiceClient tripInfo.GeoServiceClient) TaxicalculationserviceService {
	return &basicTaxicalculationserviceService{repository: &rate.PgRepository{Db: db}, geoServiceClient: geoServiceClient}
}

// New returns a TaxicalculationserviceService with all of the expected middleware wired in.
func New(middleware []Middleware, db *pg.DB, geoServiceClient tripInfo.GeoServiceClient) TaxicalculationserviceService {
	var svc TaxicalculationserviceService = NewBasicTaxicalculationserviceService(db, geoServiceClient)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
