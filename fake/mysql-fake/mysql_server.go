package main

import (
	"context"
	"fmt"

	"github.com/dolthub/vitess/go/vt/proto/query"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
)

// This is an example of how to implement a MySQL server.
// After running the example, you may connect to it using the following:
//
// > mysql --host=localhost --port=3306 --user=root mydb --execute="SELECT * FROM mytable;"
// The included MySQL client is used in this example, however any MySQL-compatible client will work.

var (
	dbName    = "chaos"
	tableName = "grayscale_set"
	address   = "localhost"
	port      = 3307
)

func ServerRun() {
	pro := createTestDatabase()
	engine := sqle.NewDefault(pro)

	session := memory.NewSession(sql.NewBaseSession(), pro)
	ctx := sql.NewContext(context.Background(), sql.WithSession(session))
	ctx.SetCurrentDatabase(dbName)

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", address, port),
	}
	s, err := server.NewServer(config, engine, memory.NewSessionBuilder(pro), nil)
	if err != nil {
		panic(err)
	}
	if err = s.Start(); err != nil {
		panic(err)
	}
}

func createTestDatabase() *memory.DbProvider {
	db := memory.NewDatabase(dbName)
	db.BaseDatabase.EnablePrimaryKeyIndexes()

	pro := memory.NewDBProvider(db)

	table := memory.NewTable(db, tableName, sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "id", Type: types.Int64, Nullable: false, Source: tableName, PrimaryKey: true, AutoIncrement: true},
		{Name: "scene", Type: types.Int8, Nullable: false, Source: tableName, PrimaryKey: false},
		{Name: "collection_id", Type: types.Int64, Nullable: false, Source: tableName},
		{Name: "status", Type: types.Int64, Nullable: false, Source: tableName},
		{Name: "percent", Type: types.Text, Nullable: false, Source: tableName},
		{Name: "wait_time", Type: types.Int64, Nullable: false, Source: tableName},
		{Name: "fail_continue_fault", Type: types.Int8, Nullable: false, Source: tableName},
		{Name: "modifier", Type: types.Text, Nullable: false, Source: tableName},
		{Name: "modify_time", Type: types.MustCreateDatetimeType(query.Type_DATETIME, 6), Nullable: false, Source: tableName},
	}), db.GetForeignKeyCollection())
	db.AddTable(tableName, table)

	//session := memory.NewSession(sql.NewBaseSession(), pro)
	//ctx := sql.NewContext(context.Background(), sql.WithSession(session))

	//creationTime := time.Unix(0, 1667304000000001000).UTC()
	//_ = table.Insert(ctx, sql.NewRow("Jane Deo", "janedeo@gmail.com", types.MustJSON(`["556-565-566", "777-777-777"]`), creationTime))
	//_ = table.Insert(ctx, sql.NewRow("Jane Doe", "jane@doe.com", types.MustJSON(`[]`), creationTime))
	//_ = table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", types.MustJSON(`["555-555-555"]`), creationTime))
	//_ = table.Insert(ctx, sql.NewRow("John Doe", "johnalt@doe.com", types.MustJSON(`[]`), creationTime))

	return pro
}
