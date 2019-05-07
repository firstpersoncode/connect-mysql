package connmysql

import (
  "fmt"
  // "encoding/json"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// type resp []map[string]interface{}

func (c *creds) get_rows(table string, fields string, options string) []map[string]interface{} {
  db := c.connect()

  query := fmt.Sprintf("SELECT %s FROM `%s` %s", fields, table, options)

  rows, err := db.Query(query)

  if err != nil {
    panic(err.Error())
  }

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

  var arr []map[string]interface{}

	for rows.Next() {
    dataStr := make(map[string]interface{})
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

      dataStr[columns[i]] = value

			// fmt.Println(columns[i] + ":", value)
		}
		// fmt.Println("-----------------------------------")
    arr = append(arr, dataStr)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

  defer rows.Close()

  // jbyte, err := json.Marshal(resp)
  //
  // if err != nil {
  //   fmt.Println(err.Error())
  // }
  //
  // jsonStr := string(jbyte)

  return arr
}

func (c *creds) add_rows(table string, fields string, values string) {
  db := c.connect()

  query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", table, fields, values)

  // fmt.Println(query)

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  defer rows.Close()
}
