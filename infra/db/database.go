package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/kchxng/go-oracle-api/infra"
	"github.com/kchxng/go-oracle-api/utils/logs"

	go_ora "github.com/sijms/go-ora/v2"
)

func InitDatabase() *sqlx.DB {

	host, port_int, service_name, username, password := infra.DatabaseParams()

	connStr := go_ora.BuildUrl(host, port_int, service_name, username, password, nil)

	db, err := sqlx.Open("oracle", connStr)
	if err != nil {
		logs.Error(err)
		logs.Error("Failed to connect oracle database")
	}
	// go_ora.SetStringConverter(db, charset, nCharset)

	logs.Info("Connected to database")
	return db

}
