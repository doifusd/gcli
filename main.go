package main

import (
	"gcli/cmd"
	"log"
	"time"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.execute err:%v", err)
	}
	// loacltime()
}

func loacltime() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	input := "2021-04-08 12:02:03"
	layout := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(layout, input, location)
	dateTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
	log.Printf("输入时间: %s,输出时间　:%s", input, dateTime)
}
