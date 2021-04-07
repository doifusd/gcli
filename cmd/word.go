package cmd

import (
	"gcli/pkg/word"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

//var wordCmd = &cobra.Command{
//	Use:   "word",
//	Short: "单词转换",
//	Long:  "支持多种格式单词",
//	Run:   func(cmd *cobra.Command, args []string) {},
//}

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderToUpperCamelCase
	ModeUnderToLowerCamelCase
	ModeCamelCaseToUnderUpper
)

var desc = strings.Join([]string{
	"该命令支持各种格式转转",
	"1:全部为大写",
	"2:全部为小写",
	"3: 下划线为大写驼峰",
	"4:　下划线为小写驼峰",
	"5:　驼峰转为下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转化",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderToUpperCamelCase:
			content = word.UnderToUpperCamelCase(str)
		case ModeUnderToLowerCamelCase:
			content = word.UnderToLowerCamelCase(str)
		case ModeCamelCaseToUnderUpper:
			content = word.CameCaseToUnder(str)
		default:
			log.Fatalf("暂不支持")
		}
		log.Printf("输出结果:%s", content)
	},
}

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
