package core

import (
	"database/sql"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

type MySQLConnInfo struct {
	User     string
	Password string
	Host     string
	Port     uint16
	Name     string
}

type MySQLDB struct {
	DB     *sql.DB
	Tables map[string]*MySQLTable
}

type MySQLTable struct {
	Name     string
	DB       *sql.DB
	Prepares map[string]string
}

func NewMySQLConn(config *MySQLConnInfo) *MySQLDB {
	myConfig := mysql.NewConfig()
	myConfig.User = config.User
	myConfig.Passwd = config.Password
	myConfig.Addr = config.Host + ":" + strconv.FormatUint(uint64(config.Port), 10)
	myConfig.DBName = config.Name
	myConfig.Params["charset"] = "utf8mb4,utf8"
	myConfig.Params["collation"] = "utf8mb4_general_ci"

	// 生成mysql连接字符串
	db, err := sql.Open("mysql", myConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	return &MySQLDB{
		DB:     db,
		Tables: map[string]*MySQLTable{},
	}
}

func (mysql *MySQLDB) Close() {
	mysql.DB.Close()
}

func (mysql *MySQLDB) Table(name string) *MySQLTable {
	if table, ok := mysql.Tables[name]; ok {
		return table
	}
	return &MySQLTable{
		Name:     name,
		DB:       mysql.DB,
		Prepares: map[string]string{},
	}
}

func (t *MySQLTable) Exec() {}

func (t *MySQLTable) Get() {}
func (t *MySQLTable) Del() {}
func (t *MySQLTable) Add() {}
func (t *MySQLTable) Set() {}

type GetsInfo struct {
	Field   string
	OrderBy map[string]string
	Where   map[string]string
	Page    uint
	Size    uint
}

func (t *MySQLTable) Gets(info *GetsInfo) {

}
func (t *MySQLTable) Dels() {}
func (t *MySQLTable) Adds() {}
func (t *MySQLTable) Sets() {}
