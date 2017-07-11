// package models provides methods for interfacing with the database and all CRUD related actions
package models

import (
	"time"
	"fmt"
	"database/sql"
)

// type My_error is a simple implementation of the error interface
type My_error struct {
	Err string
	Error_code int
}

func (e *My_error) Error() string {
	return e.Err
}

func (e *My_error) ErrorCode() int {
	return e.Error_code
}

// type Model defines an interface for models that correspond to a table schema
// These methods are used for all database related queries
type Model interface {
	// model_values returns all member values of the Model schema, with or without the ID primary key
	model_values(withID bool) []interface{}
}

func insert_model(model Model, insert_statement string, db *sql.DB) (sql.Result, error) {
	return db.Exec(insert_statement, model.model_values(false)...)
}

func update_model(model Model, update_statement string, db *sql.DB) (sql.Result, error) {
	return db.Exec(update_statement, model.model_values(true)...)
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
