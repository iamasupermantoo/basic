package crontab

import (
	"fmt"
	"time"
)

// InitCrontab 初始化定时任务
func InitCrontab() {
	// go Template(60 * time.Second)
}

// Template 定时任务模版
func Template(second time.Duration) {
	ch := time.NewTicker(second)

	for {
		fmt.Println("定时任务模版运行....")
		<-ch.C
	}
}
