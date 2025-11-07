package config

const (
	Sqlite = "sqlite"
	Mysql  = "mysql"
)

type DB struct {
	DbType string `json:"dbType"`
	Dsn    string `json:"dsn"`
}

var defaultDB = &DB{
	DbType: "sqlite",
	Dsn:    "./db/lolo.db",
}

func GetDB() *DB {
	return GetConfig().DB
}

func (x *DB) GetDbType() string {
	return x.DbType
}

func (x *DB) GetDsn() string {
	return x.Dsn
}
