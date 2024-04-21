package session

import (
	"database/sql"
	"strings"

	"geeorm/clause"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/schema"
)

type Session struct {
	db       *sql.DB // 数据库连接
	dialect  dialect.Dialect
	refTable *schema.Schema // struct 与 db的映射
	clause   clause.Clause  // sql 子句
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

// Clear 如此可重复利用session
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ") // why
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 下面是对Exec、Query、QueryRow的封装
// 统一打日志
// 执行完成后，调用Clear方法，清空Session.sql和Session.sqlVars两个变量

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err.Error())
	}
	return
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err.Error())
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}
