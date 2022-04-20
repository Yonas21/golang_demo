package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (conn *gorm.DB, error error) {

	config, err := getDbConfig(".env")

	if err != nil {
		return nil, err
	}
	dbServer := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	conn, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbServer, // data source name
		DefaultStringSize:         256,      // default size for string fields
		DisableDatetimePrecision:  true,     // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,     // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,     // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return conn, nil
}
