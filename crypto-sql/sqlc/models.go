// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"database/sql"
)

type Tblbusd10minpercent struct {
	ID        int32
	Time      sql.NullTime
	Prevtime  sql.NullTime
	Symbol    sql.NullString
	Price     sql.NullFloat64
	Prevprice sql.NullFloat64
	Percent   sql.NullFloat64
}

type Tblbusd15minpercent struct {
	ID        int32
	Time      sql.NullTime
	Prevtime  sql.NullTime
	Symbol    sql.NullString
	Price     sql.NullFloat64
	Prevprice sql.NullFloat64
	Percent   sql.NullFloat64
}

type Tblbusd1minpercent struct {
	ID        int32
	Time      sql.NullTime
	Prevtime  sql.NullTime
	Symbol    sql.NullString
	Price     sql.NullFloat64
	Prevprice sql.NullFloat64
	Percent   sql.NullFloat64
}

type Tblbusd30minpercent struct {
	ID        int32
	Time      sql.NullTime
	Prevtime  sql.NullTime
	Symbol    sql.NullString
	Price     sql.NullFloat64
	Prevprice sql.NullFloat64
	Percent   sql.NullFloat64
}

type Tblbusd5minpercent struct {
	ID        int32
	Time      sql.NullTime
	Prevtime  sql.NullTime
	Symbol    sql.NullString
	Price     sql.NullFloat64
	Prevprice sql.NullFloat64
	Percent   sql.NullFloat64
}

type Tblbusd60minpercent struct {
	ID        int32
	Time      sql.NullTime
	Prevtime  sql.NullTime
	Symbol    sql.NullString
	Price     sql.NullFloat64
	Prevprice sql.NullFloat64
	Percent   sql.NullFloat64
}

type Tblbusdprice struct {
	ID    int32
	Time  sql.NullTime
	Price sql.NullString
}

type Tblbusdtopcoinhistory struct {
	ID      int32
	Time    sql.NullTime
	Topcoin sql.NullString
}
