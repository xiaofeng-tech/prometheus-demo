package monitoring

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{

		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
