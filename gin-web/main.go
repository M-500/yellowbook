package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 10:12

func main() {
	server := InitWebServer()
	initPrometheus() // 启动Prometheus
	server.Run(":8372")
}

func initPrometheus() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		// 监听端口，可以做成可配置的
		http.ListenAndServe(":8899", nil)
	}()
}
