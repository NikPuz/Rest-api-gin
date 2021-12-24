package infrastructures

import (
	"RESTful_API_Gin/interfaces"
	"database/sql"
	"fmt"
)

type MySQLHandler struct {
	Conn *sql.DB
}

func (h *MySQLHandler) Execute(statement string) {
	h.Conn.Exec(statement)
}

func (h *MySQLHandler) Query(statement string) (interfaces.IRow, error) {
	fmt.Println(statement)

	rows, err := h.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(MySQLRow), err
	}

	row := new(MySQLRow)
	row.Rows = rows

	return row, nil
}

type MySQLRow struct {
	Rows *sql.Rows
}

func (r MySQLRow) Scan(dest ...interface{}) error {

	err := r.Rows.Scan(dest...)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r MySQLRow) Next() bool {
	return r.Rows.Next()
}
