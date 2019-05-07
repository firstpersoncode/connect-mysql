package connmysql

import (
  "fmt"
)

func (c *creds) drop_table(table string) {
  db := c.connect()
  query := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table)

  fmt.Println(query)

  drop, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  defer drop.Close()
}

func (c *creds) create_table(table string, configs string) {
  db := c.connect()
  query := fmt.Sprintf(
    "CREATE TABLE IF NOT EXISTS `%s` (%s)", table, configs)

  fmt.Println(query)

  insert, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  defer insert.Close()
}
