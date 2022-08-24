package mariadb

import (
	"log"
	"os"
	"testing"

	"github.com/macadrich/sino-db/mariadb"
)

func TestCreateDatabase(t *testing.T) {
	log.Println("Create database in mariadb.")
	mariadb.CreateNewDatabase(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
}
