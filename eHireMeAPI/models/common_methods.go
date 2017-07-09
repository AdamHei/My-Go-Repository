// package models provides methods for interfacing with the database and all CRUD related actions
package models

import (
	"time"
	"fmt"
)

// type my_error is a simple implementation of the error interface
type my_error struct {
	err string
}

func (e *my_error) Error() string {
	return e.err
}

// type Model defines an interface for models that correspond to a table schema
// These methods are used for all database related queries
type Model interface {
	// member_fields returns all columns related to the Model schema, with or without the ID primary key
	member_fields(withID bool) string

	// member_values returns all member values of the Model schema, with or without the ID primary key
	member_values(withID bool) string
}

// insert_query will format the INSERT query for a given Model in a given table, by name
func insert_query(model Model, table string, withID bool) string {
	fields := model.member_fields(withID)
	values := model.member_values(withID)
	return fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, fields, values)
}

// add_ID will prepend a primary key ID to a given string
// Usually used in member_fields and member_values implementations
func add_ID(str string, id int) string {
	return fmt.Sprintf("(%d, %s", id, str)
}

// choose_string is a helper function used when merging structs that will return one of two strings given
// Priority is given to the first string, as it comes from the database while the second, potentially nil,
// comes from the client when updating a model
func choose_string(att1, att2 string) string {
	if att2 == "" || att1 == att2 {
		return att1
	}
	return att2
}

// choose_int (see #choose_string) will return one of two integers with priority given to the first
func choose_int(i1, i2 int) int {
	if i2 == 0 || i1 == i2 {
		return i1
	}
	return i2
}
// choose_time (see #choose_string) will return one of two given Times with priority given to the first
func choose_time(t1, t2 time.Time) time.Time {
	if t2.IsZero() || t1.Equal(t2) {
		return t1
	}
	return t2
}
