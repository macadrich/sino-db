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

func NewDatabase(user, pass, host, dbname, dir string) Database {
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	log.Debugln(URL)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		log.Fatalln("Failed to connect to mariadb database!")
	}

	log.Debugln("MariaDB Database connection established")

	// load sql script file for migrations
	migrations := &migrate.FileMigrationSource{
		Dir: dir + "/",
	}

	// validate sqlDB
	mdb, err := db.DB()
	if err != nil {
		log.Fatalln("Error on sqlDB:", err)
	}

	// apply sql migration
	n, err := migrate.Exec(mdb, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatalln("Error occcured:", err, migrations)
	}

	// show migration results
	log.Debugf("Applied %d migrations!\n", n)

	return Database{
		DB: db,
	}
}
