package main

import (
	"bufio"
	"strconv"
	"strings"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_ = nextInt(reader)
	arr := nextIntArray(reader)

	total := 1
	for _, val := range arr {
		total *= val
	}

	total *= 4
	total /= 1024
	fmt.Println(total)
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
