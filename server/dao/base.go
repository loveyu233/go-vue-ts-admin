package dao

import "database/sql"

var mysqlClient *sql.DB

func InitMysqlClient(db *sql.DB) {
	mysqlClient = db
}
