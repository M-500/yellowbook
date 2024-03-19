package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-19 15:58

type MiddlewareBuilder struct {
	NameSpace  string
	SubSystem  string
	Name       string
	Help       string
	InstanceId string
}

func (m *MiddlewareBuilder) Build() gin.HandlerFunc {
	labels := []string{"method", "pattern", "status"}
	summary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: m.NameSpace,
			Subsystem: m.SubSystem,
			Name:      m.Name + "_resp_time",
			Help:      m.Help,
			ConstLabels: map[string]string{
				"instance_id": m.InstanceId,
			},
			Objectives: map[float64]float64{
				0.5:   0.01,
				0.9:   0.01,
				0.99:  0.001,
				0.999: 0.0001,
			},
		},
		labels)
	prometheus.MustRegister(summary)

	gauge := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: m.NameSpace,
			Subsystem: m.SubSystem,
			Name:      m.Name + "_active_req",
			Help:      m.Help,
			ConstLabels: map[string]string{
				"instance_id": m.InstanceId,
			},
		})
	prometheus.MustRegister(gauge)
	return func(ctx *gin.Context) {
		start := time.Now()
		gauge.Inc()
		defer func() {
			dur := time.Since(start)
			gauge.Dec()
			// 考虑到404的问题
			pattern := ctx.FullPath()
			if pattern == "" {
				pattern = "unknow"
			}
			summary.WithLabelValues(
				ctx.Request.Method,
				pattern, strconv.Itoa(ctx.Writer.Status()),
			).Observe(float64(dur.Milliseconds()))
		}()
		ctx.Next() // 这里执行到业务 考虑到可能panic 所以放到defer中

	}
}
