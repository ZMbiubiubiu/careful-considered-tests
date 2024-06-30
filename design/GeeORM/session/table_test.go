package session

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"geeorm/dialect"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestTable(t *testing.T) {
	db, err := sql.Open("sqlite3", "gee.db")
	assert.Nil(t, err)
	dail, ok := dialect.GetDialect("sqlite3")
	assert.Equal(t, true, ok)

	session := New(db, dail)

	session.Model(&User{})

	err = session.CreateTable()
	assert.Nil(t, err)
	hasTable := session.HasTable()
	assert.Equal(t, true, hasTable)

	err = session.DropTable()
	assert.Nil(t, err)
	hasTable = session.HasTable()
	assert.Equal(t, false, hasTable)
}
