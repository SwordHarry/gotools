package cmd

import (
	"github.com/spf13/cobra"
	"gotools/internal/word"
	"log"
	"strings"
)

// 用于单词格式转换命令的子命令 word

const (
	ModeUpper                      = iota + 1 // 大写
	ModeLower                                 // 小写
	ModeUnderscoreToUpperCamelcase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelcase            // 下划线转小写驼峰
	ModeCamelcaseToUnderscore                 // 驼峰转下划线
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部单词转为小写",
	"2：全部单词转为小写",
	"3：下划线单词转为大写驼峰单词",
	"4：下划线单词转为小写驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelcase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelcaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果: %s\n", content)
	},
}

var str string
var mode int8

// 命令参数设置
func init() {
	fs := wordCmd.Flags()
	fs.StringVarP(&str, "str", "s", "", "请输入单词内容")
	fs.Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的格式")
}
