package lib

import "strings"

func LoadWechat(fileName string) []Row {
	var res []Row
	csvData := ReadCsv(fileName)
	start := false
	for _, line := range csvData {
		if find := strings.Contains(line[0], "-----------------"); find {
			start = true
			continue
		}
		if find := strings.Contains(line[0], "-----------------"); find && start {
			start = false
			break
		}
		if !start {
			continue
		}
		if len(line) != 11 {
			continue
		}
		res = append(res, Row{
			ChargeTime:   line[0],
			ChargeType:   line[1],
			ChargeTarget: line[2],
			Product:      line[3],
			AddOrRemove:  line[4],
			Count:        strings.Trim(line[5], "￥"),
			PayType:      line[6],
			ChargeStatus: line[7],
			Channel:      "wechat",
		})

	}
	if res == nil || len(res) <= 1 {
		return nil
	}
	return res
}

