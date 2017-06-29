package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheCreateAndSet(t *testing.T) {
	assert := assert.New(t)
	DB, err := OpenTestConnection()
	assert.Nil(err, "oh no")

	err = DB.DropTable(&Gormd{}).Error
	assert.Nil(err, "o  no")
	err = DB.CreateTable(&Gormd{}).Error
	assert.Nil(err, "o  no")
	//dat1 := Gormd{ID: "1", Name: "one"}
	//DB.Create(&dat1)
	//dat2 := Gormd{ID: "2", Name: "two"}
	//DB.Create(&dat2)
}
