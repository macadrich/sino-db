package mariadb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)

type MDatabase struct {
	DB *sql.DB
}

func CreateNewDatabase(user, pass, host, dbname string) {
	URL := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8&parseTime=True&loc=Local", user, pass, host)
	db, err := sql.Open("mysql", URL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + ";")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Create database:", dbname, "success!")
}

func ConnectMYSQL(user, pass, host, dbname, dir string) (*MDatabase, error) {
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	db, err := sql.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
	}

	r, err := db.Exec(`CREATE TABLE IF NOT EXISTS sino (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,email VARCHAR(100),name VARCHAR(20),age INT)`)
	if err != nil {
		log.Fatalln("failed create table!", err)
	}
	ra, _ := r.RowsAffected()
	log.Infoln("RowsAffected:", ra)
	return &MDatabase{db}, nil
}

func (mdb *MDatabase) GetUsers() {
	res, err := mdb.DB.Query("SELECT * FROM sino")
	if err != nil {
		log.Fatalln("query failed: ", err)
	}
	for res.Next() {
		user := struct {
			ID    string `db:"id"`
			Email string `db:"email"`
			Name  string `db:"name"`
			Age   int    `db:"age"`
		}{}
		err := res.Scan(&user.ID, &user.Email, &user.Name, &user.Age)
		if err != nil {
			log.Fatalln("scan failed: ", err)
		}
		log.Infof("%+v\n", user)
	}
}

func (mdb *MDatabase) InsertUser(item interface{}) {
	res, err := mdb.DB.Exec("INSERT INTO sino(email,name,age) VALUES(?,?,?)", "ccc@gmail.com", "ccc", 33)
	if err != nil {
		log.Fatalln("exec insert failed: ", err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("get lastId failed: ", err)
	}
	log.Infoln("Last inserted row id: %d\n", lastId)
}
