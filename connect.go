package connmysql

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type creds struct {
  db_driver string
  user string
  pass string
  db_name string
}

// type resp []map[string]interface{}

func (c *creds) connect() (db *sql.DB) {
  db, err := sql.Open(c.db_driver, c.user+":"+c.pass+"@/"+c.db_name)

  if err != nil {
      panic(err.Error())
  }

  return db
}

func (c *creds) SetTable(table string, configs string) {
  c.drop_table(table)
  c.create_table(table, configs)
}

func (c *creds) DropTable(table string) {
  c.drop_table(table)
}

func (c *creds) CreateTable(table string, configs string) {
  c.create_table(table, configs)
}

func (c *creds) AddRows(table string, fields string, values string) {
  c.add_rows(table, fields, values)
}

func (c *creds) GetRows(table string, fields string, options string) []map[string]interface{} {
  rows := c.get_rows(table, fields, options)
  return rows
}

func Connect(db_driver string, user string, pass string, db_name string) creds {
  connect := creds{db_driver, user, pass, db_name}
  return connect
}
