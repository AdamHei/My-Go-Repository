package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	arr := []int{0}

	for ; len(arr) < 1000; {
		appended := make([]int, len(arr))
		copy(appended, arr)

		for i, v := range appended {
			appended[i] = v ^ 1
		}

		arr = append(arr, appended...)

		fmt.Println(arr)
	}

	q := nextInt(reader)

	for i := 0; i < q; i++ {
		fmt.Println(arr[nextInt(reader)])
	}
}

func nextInt(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	i, _ := strconv.ParseInt(strings.TrimSpace(str), 10, 64)
	return int(i)
}

func nextIntArray(reader *bufio.Reader) []int {
	str, _ := reader.ReadString('\n')
	return toIntArray(strings.Fields(str))
}

func toIntArray(slice []string) []int {
	arr := make([]int, len(slice))
	for i, numStr := range slice {
		res, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
		if err != nil {
			panic(err)
		}
		arr[i] = int(res)
	}
	return arr
}
