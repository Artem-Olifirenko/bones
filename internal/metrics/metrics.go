package metrics

import "github.com/prometheus/client_golang/prometheus"

// Metrics метрики микросервиса
type Metrics struct {
	// TODO: Удалите эту метрику из реального проекта
	// Пример регистрации метрики для prometheus
	// Создайте свои собственные метрики
	PrometheusTestCounter prometheus.Counter
}

// New возвращает новый экземпляр метрик prometheus
// TODO prefix у каждого микросервиса должен быть универсальный(нужно вписать название своего микросервиса)
func New() *Metrics {
	prefix := "test"
	metrics := &Metrics{
		PrometheusTestCounter: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name:      "test_counter",
				Namespace: prefix,
				Help:      "Some test counter for prometheus",
			}),
	}

	prometheus.MustRegister(metrics.PrometheusTestCounter)

	return metrics
}

// IncPrometheusTestCounter увеличивает счетчик тестовой метрики prometheus
func (m *Metrics) IncPrometheusTestCounter() {
	m.PrometheusTestCounter.Inc()
}
