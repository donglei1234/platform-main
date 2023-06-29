package fxsvcapp

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gstones/platform/services/common/access"
	"github.com/gstones/platform/services/common/service"
)

type MetricsHttpService struct {
	service.TcpTransport
}

func (s *MetricsHttpService) RegisterWithHttpServer(server service.HasHttpServeMux) {
	h := server.HttpServeMux()
	h.Handle("/metrics", promhttp.Handler())
}

func NewMetricsHttpService() (out service.HttpServiceFactory) {
	out.HttpService = &MetricsHttpService{}
	return
}

func (s *MetricsHttpService) AccessLevel() access.AccessLevel {
	return access.AccessUndefined
}
