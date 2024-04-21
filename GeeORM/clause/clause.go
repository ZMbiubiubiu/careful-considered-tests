package clause

import "strings"

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

// Set 生成一个子句的sql
func (c *Clause) Set(name Type, values ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}

	sql, vars := generators[name](values...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build 方法根据传入的 Type 的顺序，构造出最终的 SQL 语句。
func (c *Clause) Build(orders ...Type) (entireSql string, vars []interface{}) {
	var sqls []string
	for _, orderType := range orders {
		if sql, ok := c.sql[orderType]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[orderType]...)
		}
	}

	return strings.Join(sqls, " "), vars
}
