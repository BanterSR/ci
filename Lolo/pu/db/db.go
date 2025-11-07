package db

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/glebarez/sqlite"

	"gucooing/lolo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gromlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db       *gorm.DB
	noDbType = errors.New("不正确的DbType")
)

func NewDB() error {
	cfg := config.GetDB()
	var err error
	switch cfg.GetDbType() {
	case config.Mysql:
		db, err = newMysql(cfg.GetDsn())
	case config.Sqlite:
		db, err = newSqlite(cfg.GetDsn())
	default:
		return noDbType
	}
	if err != nil {
		return err
	}
	err = db.AutoMigrate(
		&OFGame{},
		&BlackDevice{},
		&OFUser{},
	)

	db.Create(&OFUser{
		UserId: 999999,
		SdkUid: 0,
	})

	return err
}

func newMysql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), getGormConfig())
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(1000)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(100 * time.Millisecond) // 0.1 秒

	return db, nil
}

func newSqlite(dsn string) (*gorm.DB, error) {
	if _, err := os.Stat(filepath.Dir(dsn)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(dsn), 0777)
	}
	db, err := gorm.Open(sqlite.Open(dsn), getGormConfig())
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(1000)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(100 * time.Millisecond) // 0.1 秒

	return db, nil
}

func getGormConfig() *gorm.Config {
	info := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	info.Logger = gromlogger.Default.LogMode(gromlogger.Silent)
	return info
}
