package main

import (
	"flag"
	"go_example/format_charge/lib"
)

var (
	a     string
	w     string
	title *lib.Row
)

func init() {
	flag.StringVar(&a, "a", "", "ali file path")
	flag.StringVar(&w, "w", "", "wechat file path")
	flag.Parse()
}

func main() {
	var data []lib.Row
	// load ali
	if a != "" {
		aliData := lib.LoadAli(a)
		title = &aliData[0]
		data = append(data, aliData[1:]...)
	}
	if w != "" {
		weChatData := lib.LoadWechat(w)
		if title == nil {
			title = &weChatData[0]
		}
		data = append(data, weChatData[1:]...)
	}
	sqlClient := lib.NewSqlLiteClient()
	defer sqlClient.Close()
	sqlClient.ImportAllRows(data)
	lib.WriteCsv(data, title)

}


