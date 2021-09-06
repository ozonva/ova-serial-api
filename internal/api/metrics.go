package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	incSuccessCreateSerialCounter()
	incFailCreateSerialCounter()
	incSuccessRemoveSerialCounter()
	incFailRemoveSerialCounter()
	incSuccessUpdateSerialCounter()
	incFailUpdateSerialCounter()
}

const (
	successResultLabel = "success"
	failResultLabel    = "fail"
)

var labels = []string{"result"}

type metrics struct {
	createSerialCounter *prometheus.CounterVec
	removeSerialCounter *prometheus.CounterVec
	updateSerialCounter *prometheus.CounterVec
}

func newApiMetrics() Metrics {
	return &metrics{
		createSerialCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "create_serial_request_count",
			Help: "number of created serials",
		},
			labels),

		removeSerialCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "remove_serial_request_count",
			Help: "number of removed serials",
		},
			labels),
		updateSerialCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "update_serial_request_count",
			Help: "number of updated serials",
		},
			labels),
	}
}

func (m *metrics) incSuccessCreateSerialCounter() {
	m.createSerialCounter.WithLabelValues(successResultLabel).Inc()
}

func (m *metrics) incFailCreateSerialCounter() {
	m.createSerialCounter.WithLabelValues(failResultLabel).Inc()
}

func (m *metrics) incSuccessRemoveSerialCounter() {
	m.removeSerialCounter.WithLabelValues(successResultLabel).Inc()
}

func (m *metrics) incFailRemoveSerialCounter() {
	m.removeSerialCounter.WithLabelValues(failResultLabel).Inc()
}

func (m *metrics) incSuccessUpdateSerialCounter() {
	m.updateSerialCounter.WithLabelValues(successResultLabel).Inc()
}

func (m *metrics) incFailUpdateSerialCounter() {
	m.updateSerialCounter.WithLabelValues(failResultLabel).Inc()
}
