package mariadb

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

type MDatabase struct {
	DB *sql.DB
}

func ConnectMYSQL(user, pass, host, dbname, dir string) (*sql.DB, error) {
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	db, err := sql.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func (mdb *MDatabase) Query() {
	res, err := mdb.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		user := struct {
			ID    string `db:"id"`
			Name  string `db:"name"`
			Email string `db:"email"`
		}{}
		err := res.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("%+v\n", user)
	}
}

// GORM
func NewDatabase(user, pass, host, dbname, dir string) Database {
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	log.Infoln(URL)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		log.Fatalln("Failed to connect to mariadb database!")
	}
	log.Infoln("MariaDB Database connection established")

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
	log.Infof("Applied %d migrations!\n", n)

	return Database{
		DB: db,
	}
}

// GORM
func (mdb *Database) Create(item interface{}) {
	r := mdb.DB.Create(item)
	if r.Error != nil && r.RowsAffected != 1 {
		log.Error("create failed", r.Error)
		return
	}

	log.Infoln("create user successfully!")
}
