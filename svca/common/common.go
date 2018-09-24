package common

import (
	"log"

	zipkin "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

var Tracer *zipkin.Tracer

const endpointURL = "http://zipkin:9411/api/v2/spans"

func GetTracer() *zipkin.Tracer {
	if Tracer == nil {
		t, err := newTracer()
		if err != nil {
			log.Fatal(err)
		}
		Tracer = t
	}
	return Tracer

}

func newTracer() (*zipkin.Tracer, error) {
	// The reporter sends traces to zipkin server
	reporter := reporterhttp.NewReporter(endpointURL)

	// Local endpoint represent the local service information
	localEndpoint := &model.Endpoint{ServiceName: "svca", Port: 8080}

	// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, err
	}

	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return nil, err
	}

	return t, err
}
