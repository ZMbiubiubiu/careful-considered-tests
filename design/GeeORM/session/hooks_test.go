package session

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"geeorm/dialect"
	"geeorm/log"
)

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before inert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	db, err := sql.Open("sqlite3", "gee.db")
	assert.Nil(t, err)
	dial, ok := dialect.GetDialect("sqlite3")
	assert.Equal(t, true, ok)
	s := New(db, dial).Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})

	u := &Account{}

	err = s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("Failed to call hooks after query, got", u)
	}
}
