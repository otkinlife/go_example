package lib

import "strings"

func LoadAli(fileName string) []Row {
	var res []Row
	csvData := ReadCsv(fileName)
	start := false
	for _, line := range csvData {
		if find := strings.Contains(ConvertToUtf8(line[0]), "-----------------"); find {
			start = true
			continue
		}
		if find := strings.Contains(ConvertToUtf8(line[0]), "-----------------"); find && start {
			start = false
			break
		}
		if !start {
			continue
		}
		if len(line) != 12 {
			continue
		}
		res = append(res, Row{
			AddOrRemove:  ConvertToUtf8(line[0]),
			ChargeTarget: ConvertToUtf8(line[1]),
			Product:      ConvertToUtf8(line[3]),
			PayType:      ConvertToUtf8(line[4]),
			Count:        ConvertToUtf8(line[5]),
			ChargeStatus: ConvertToUtf8(line[6]),
			ChargeType:   ConvertToUtf8(line[7]),
			ChargeTime:   ConvertToUtf8(line[10]),
			Channel:      "alipay",
		})
	}
	if res == nil || len(res) <= 1 {
		return nil
	}
	return res
}

