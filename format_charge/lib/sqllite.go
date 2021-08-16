package lib

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type SqlLite struct {
	Db *sql.DB
}

func NewSqlLiteClient() *SqlLite {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.31.229:3306)/count")
	if err != nil {
		log.Fatalln(err)
	}
	return &SqlLite{
		Db: db,
	}
}

func (s *SqlLite) ImportAllRows(data []Row) {
	importNo := time.Now().Unix()
	prepare, err := s.Db.Prepare(`INSERT INTO count 
			(ChargeTime,ChargeType,ChargeTarget,Product,AddOrRemove,Count,PayType,ChargeStatus,Channel,ImportNo)
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatalln(err)
	}
	for _, row := range data {
		_, err := prepare.Exec(
			row.ChargeTime,
			row.ChargeType,
			row.ChargeTarget,
			row.Product,
			row.AddOrRemove,
			row.Count,
			row.PayType,
			row.ChargeStatus,
			row.Channel,
			importNo,
		)
		if err != nil {
			d, _ := s.Db.Prepare("delete from count where ImportNo = ?")
			d.Exec(importNo)
			log.Fatalln(err)
		}
	}
}

func (s *SqlLite) Close(){
	s.Db.Close()
}
