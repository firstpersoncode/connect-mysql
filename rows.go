package connmysql

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func (c *creds) get_rows(query string) {
  db := c.connect()

  rows, err := db.Query(query)

  if err != nil {
    panic(err.Error())
  }

  // Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

  // Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			fmt.Println(columns[i] + ":", value)
		}
		fmt.Println("-----------------------------------")
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

  defer rows.Close()
}

func (c *creds) add_rows(table string, fields string, values string) {
  db := c.connect()

  query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", table, fields, values)

  fmt.Println(query)

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  defer rows.Close()
}
