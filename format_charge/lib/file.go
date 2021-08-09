package lib

import (
	"encoding/csv"
	"log"
	"os"
)

type Row struct {
	ChargeTime   string
	ChargeType   string
	ChargeTarget string
	Product      string
	AddOrRemove  string
	Count        string
	PayType      string
	ChargeStatus string
	Channel      string
}

func ReadCsv(fileName string) [][]string {
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	r.FieldsPerRecord = -1
	res, err := r.ReadAll()
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	return res
}

func WriteCsv(data []Row, title *Row) {
	fs, err := os.Create("./charge_res.csv")
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	w := csv.NewWriter(fs)
	data = append([]Row{*title}, data...)
	for _, row := range data {
		line := []string {
			row.ChargeTime,
			row.AddOrRemove,
			row.PayType,
			row.ChargeType,
			row.Count,
			row.ChargeTarget,
			row.ChargeStatus,
			row.Product,
			row.Channel,
		}
		err = w.Write(line)
		if err != nil {
			log.Fatalf("can not open the file, err is %+v", err)
		}
		w.Flush()
	}

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
