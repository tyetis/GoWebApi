package data

import (
	"fmt"

	"github.com/spf13/viper"
	sqlserver "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewConnectionDb(config *viper.Viper) (*gorm.DB, error) {
	var dialect gorm.Dialector
	gormConfig := &gorm.Config{}
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", config.GetString("DB_USER"), config.GetString("DB_PASS"), config.GetString("DB_URL"), config.GetString("DB_NAME"))
	fmt.Println(dsn)
	dialect = sqlserver.Open(dsn)
	return gorm.Open(dialect, gormConfig)
}
