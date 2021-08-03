package lib

import (
	"github.com/axgle/mahonia"
	"strings"
)

func ConvertToUtf8(param string) string {
	param = strings.Trim(param, " ")
	return mahonia.NewDecoder("GBK").ConvertString(param)
}
