package models

import (
	"time"
	"fmt"
)

type Model interface {
	member_fields(withID bool) string
	member_values(withID bool) string
}

func insert_query(model Model, table string, withID bool) string {
	fields := model.member_fields(withID)
	values := model.member_values(withID)
	return fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, fields, values)
}

func add_ID(str string, id int) string {
	return fmt.Sprintf("(%d, %s", id, str)
}

func choose_string(att1, att2 string) string {
	if att2 == "" || att1 == att2 {
		return att1
	}
	return att2
}

func choose_int(i1, i2 int) int {
	if i2 == 0 || i1 == i2 {
		return i1
	}
	return i2
}

func choose_time(t1, t2 time.Time) time.Time {
	if t2.IsZero() || t1.Equal(t2) {
		return t1
	}
	return t2
}
