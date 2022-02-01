package mariadb

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(user, pass, host, dbname string) Database {
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	log.Infoln(URL)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to mariadb database!")
	}

	log.Infoln("MariaDB Database connection established")

	// load sql script file for migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/",
	}

	// validate sqlDB
	mdb, err := db.DB()
	if err != nil {
		log.Infof("Error on sqlDB:", err)
	}

	// apply sql migration
	n, err := migrate.Exec(mdb, "sinodb", migrations, migrate.Up)
	if err != nil {
		log.Infof("Error occcured:", err)
	}

	// show migration results
	log.Infof("Applied %d migrations!\n", n)

	return Database{
		DB: db,
	}
}
