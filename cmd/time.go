package cmd

import (
	"gcli/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var calcuteTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var NowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowtime := timer.GetNowTime()
		log.Printf("输出结果: %s,%d", nowtime.Format("2006-01-02 15:15:14"), nowtime.Unix())
	},
}

var CalculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var ctime time.Time
		var layout = "2006-05-04 15:14:13"
		if calcuteTime == "" {
			ctime = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calcuteTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			ctime, err = time.Parse(layout, calcuteTime)
			if err != nil {
				t, _ := strconv.Atoi(calcuteTime)
				ctime = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(ctime, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果: %s,%d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(NowTimeCmd)
	timeCmd.AddCommand(CalculateTimeCmd)
	CalculateTimeCmd.Flags().StringVarP(&calcuteTime, "calcute", "c", "", `需要计算的时间，有效单位为时间戳或一个是化后的时间`)
	CalculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间,有效时间单位"ns","us","ms","s","m","h"`)

}
