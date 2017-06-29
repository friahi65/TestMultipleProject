package orm

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightConn(t *testing.T) {

	assert := assert.New(t)
	constr := GetRightPostgressConnectionString()
	if constr == "" {
		t.Skip("skipping")
	}
	DB, err := OpenTestConnection(constr)
	assert.Nil(err, "oh no")
	err = DeleteTable(DB, Gormd{})
	assert.Nil(err, "o  no")
	err = DB.CreateTable(&Gormd{}).Error
	assert.Nil(err, "o  no")
}
func TestWrongConn(t *testing.T) {

	assert := assert.New(t)
	constr := GetWrongPostgressConnectionString()
	if constr == "" {
		t.Skip("skipping")
	}
	DB, err := OpenTestConnection(constr)
	assert.Nil(err, "oh no")
	err = DeleteTable(DB, Gormd{})
	assert.Nil(err, "o  no")
	err = DB.CreateTable(&Gormd{}).Error
	assert.Nil(err, "o  no")
}

func GetRightPostgressConnectionString() string {
	e := ""
	var pgHost []byte
	if e = os.Getenv("POSTGRES_HOST"); e == "" {
		cmdStr := "docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' pg "
		out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
		fmt.Println("here", out, len(out))
		if err != nil || len(out) == 0 {
			connStr := ""
			errors.New("Errors")
			return connStr
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

func GetWrongPostgressConnectionString() string {
	e := ""
	var pgHost []byte
	if e = os.Getenv("POSTGRES_HOST"); e == "" {
		cmdStr := "dockerr inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' pg "
		out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
		fmt.Println("here", out, len(out))
		if err != nil || len(out) == 0 {
			connStr := ""
			errors.New("Errors")
			return connStr
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
