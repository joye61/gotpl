package core

type MySQLDB struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	Charset  string
}

type MySQLTable struct {
	Prepares map[string]string
}

func NewMySQLDB(config *MySQLDB) *MySQLDB {
	return &MySQLDB{}
}

func (mysql *MySQLDB) Close() {

}

func (db *MySQLDB) Table(name string) *MySQLTable {
	return &MySQLTable{}
}

func (t *MySQLTable) Exec() {}

func (t *MySQLTable) Get() {}
func (t *MySQLTable) Del() {}
func (t *MySQLTable) Add() {}
func (t *MySQLTable) Set() {}

func (t *MySQLTable) Gets() {}
func (t *MySQLTable) Dels() {}
func (t *MySQLTable) Adds() {}
func (t *MySQLTable) Sets() {}
