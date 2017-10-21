package main

import (
	"fmt"
	"github.com/adamhei/hackisuhonors/exchangeApis"
)

func main() {
	ch := make(chan []string)
	exchangeApis.FetchAllExchanges(ch)
	for i := 0; i < exchangeApis.NUMEXCHANGES; i++ {
		fmt.Println(<-ch)
	}
}
