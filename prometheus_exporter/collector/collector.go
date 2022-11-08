package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"strings"
	"sync"
)

// Metrics 指标结构体
type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

/**
 * 函数：newGlobalMetric
 * 功能：创建指标描述符
 */
func newGlobalMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	if len(namespace) != 0 {
		return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
	} else {
		return prometheus.NewDesc(metricName, docString, labels, nil)
	}
}

// NewMetrics
// 工厂方法：NewMetrics
// 功能：初始化指标信息，即Metrics结构体
func NewMetrics(namespace string) *Metrics {
	labels := []string{"domain"}
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"https.certificate.check": newGlobalMetric(namespace, replace("https.certificate.check"), "https证书检查", labels),
		},
	}
}

// Describe
// 接口：Describe
// 功能：传递结构体中的指标描述符到channel
func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

// Collect
// 接口：Collect
// 功能：抓取最新的数据，传递给channel
func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock() // 加锁
	defer c.mutex.Unlock()
	labels := []string{"www.baidu.com"}
	ch <- prometheus.MustNewConstMetric(c.metrics["https.certificate.check"], prometheus.GaugeValue, float64(1), append(labels)...)
}

func replace(s string) string {
	return strings.Replace(s, ".", "_", -1)
}
