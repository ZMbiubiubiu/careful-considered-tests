package session

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"geeorm/dialect"
)

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	db, err := sql.Open("sqlite3", "gee.db")
	assert.Nil(t, err)
	dial, ok := dialect.GetDialect("sqlite3")
	assert.Equal(t, true, ok)
	s := New(db, dial).Model(&User{})
	err = s.DropTable()
	assert.Nil(t, err)
	err = s.CreateTable()
	assert.Nil(t, err)
	_, err = s.Insert(user1, user2)
	assert.Nil(t, err)
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
}
