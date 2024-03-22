package ioc

import (
	"gin-web/internal/repository/dao"
	prometheus2 "github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 13:46

func InitDB() *gorm.DB {
	config := &gorm.Config{}
	db, err := gorm.Open(mysql.Open("admin:123456@tcp(192.168.1.52:3306)/xhs?charset=utf8&parseTime=true"), config)
	if err != nil {
		panic(err)
	}
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(20)
	}
	// 初始化数据库
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}

	err = db.Use(prometheus.New(prometheus.Config{
		DBName:          "",
		RefreshInterval: 15,    // 插件采集数据的评率
		StartServer:     false, // 因为已经启动过了。所以要设置为false
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"thread_runing"},
			},
		},
	}))

	if err != nil {
		panic(err)
	}

	vector := prometheus2.NewSummaryVec(prometheus2.SummaryOpts{
		Namespace: "test_gin_web",
		Name:      "gorm_query_time",
		Subsystem: "damn1",
		Help:      "统计Gorm的查询时间",
		//Objectives: map[float64]float64{
		//	0.5:   0.01,
		//	0.9:   0.01,
		//	0.99:  0.001,
		//	0.999: 0.0001,
		//},
	}, []string{"type", "table"})

	// 监控查询的执行时间
	err = db.Callback().Create().Before("*").Register("prometheus_create_before", func(db *gorm.DB) {
		startTime := time.Now()
		db.Set("start_time", startTime)
	}) // 作用于Insert语句
	if err != nil {
		panic(err)
	}
	// 监控查询的执行时间
	err = db.Callback().Create().After("*").Register("prometheus_create_before", func(db *gorm.DB) {
		val, _ := db.Get("start_time")
		startTime, ok := val.(time.Time)
		if !ok {
			return
		}
		// 上报普罗米修斯
		table := db.Statement.Table
		if len(table) == 0 {
			table = "unknown" // 调用 原生SQL的时候 就没有table ROW查询页没有
		}
		vector.WithLabelValues("create", table).Observe(time.Since(startTime).Seconds())

	}) // 作用于Insert语句
	return db
}

type Callbacks struct {
	vector prometheus2.SummaryVec
}

func (c *Callbacks) before() func(db *gorm.DB) {
	return func(db *gorm.DB) {
		startTime := time.Now()
		db.Set("start_time", startTime)
	}
}
func (c *Callbacks) after(typ string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		val, _ := db.Get("start_time")
		startTime, ok := val.(time.Time)
		if !ok {
			return
		}
		// 上报普罗米修斯
		table := db.Statement.Table
		if len(table) == 0 {
			table = "unknown" // 调用 原生SQL的时候 就没有table ROW查询页没有
		}
		c.vector.WithLabelValues(typ, table).Observe(time.Since(startTime).Seconds())
	}
}
