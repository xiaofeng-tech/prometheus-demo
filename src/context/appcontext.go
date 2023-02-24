package context

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xiaofeng-tech/prometheus-demo/src/monitoring"
)

type appContext struct {

}

func NewAppContexxt() *appContext {
	return &appContext{}
}


func (appContext appContext) Server() {
	monitoring.RecordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}