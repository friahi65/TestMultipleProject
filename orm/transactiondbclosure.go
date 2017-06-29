package orm

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var skip bool

func GetPostgressConnectionString() string {
	e := ""
	var pgHost []byte
	if e = os.Getenv("POSTGRES_HOST"); e == "" {
		cmdStr := "docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' pg "
		out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
		fmt.Println("here", out)
		if err != nil {
			errors.New("ekh")
		}
		pgHost = out
	} else {
		pgHost = []byte(e)
	}

	pgPort := "5432"
	psUser := "postgres"
	psPassword := "postgres"
	sslMode := "disable"
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s", pgHost, pgPort, psUser, psPassword, sslMode)
	fmt.Println(connStr)
	return connStr
}

type Gormd struct {
	ID   string `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
type FatemehTest struct {
	ID   string `gorm:"primary_key"`
	Name string
}

var DB *gorm.DB

func OpenTestConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", GetPostgressConnectionString())
	db.LogMode(true)
	return
}

func OpenSQLConnection() (db *gorm.DB, err error) {
	dbhost := "172.17.0.3"
	db, err = gorm.Open("mysql", fmt.Sprintf("host=%s ", dbhost))
	db.LogMode(true)
	return
}
func main() {
	var err error
	if DB, err = OpenTestConnection(); err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to test database, but got err=%+v", err))
	}

	DB.DropTable(&Gormd{})
	DB.CreateTable(&Gormd{})
	dat1 := Gormd{ID: "1", Name: "one"}
	DB.Create(&dat1)
	dat2 := Gormd{ID: "2", Name: "two"}
	DB.Create(&dat2)

}
func DeleteTable(DB *gorm.DB, table Gormd) error {
	if err := DB.DropTableIfExists(table).Error; err != nil {
		return err
	}
	return nil
}
