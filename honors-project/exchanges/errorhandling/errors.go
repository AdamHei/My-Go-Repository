package errorhandling

import (
	"github.com/adamhei/honors-project/exchanges/models"
	"log"
)

type MyError struct {
	Err  string
	Code int
}

func (e *MyError) Error() string {
	return e.Err
}

func (e *MyError) ErrorCode() int {
	return e.Code
}

// Print the error message and send an empty response through the channel
func ErrorHandler(errorMsg string, ch chan<- models.LimitedJson) {
	log.Print(errorMsg)
	ch <- make(models.LimitedJson)
}
