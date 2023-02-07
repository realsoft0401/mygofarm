package Rom

import (
	"fmt"
	"mygofarm/Settings"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init(cfg *Settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		zap.L().Error("connect ORM failed", zap.Error(err))
		return
	}
	return
}