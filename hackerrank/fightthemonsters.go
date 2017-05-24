package main

import (
	"bufio"
	"strconv"
	"strings"
	"os"
	"sort"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n_hit_t := nextIntArray(reader)
	hit := n_hit_t[1]
	t := n_hit_t[2]

	monsters := nextIntArray(reader)

	sort.Ints(monsters)

	killed := 0
	index := 0
	for t > 0 {
		if index >= len(monsters) {
			break
		}
		if monsters[index] <= 0 {
			index++
		} else if monsters[index] <= hit {
			killed++
			index++
		} else {
			monsters[index] -= hit
		}
		t--
	}

	fmt.Println(killed)
}

//func nextInt(reader *bufio.Reader) int {
//	str, _ := reader.ReadString('\n')
//	i, _ := strconv.ParseInt(strings.TrimSpace(str), 10, 64)
//	return int(i)
//}
//
//func nextIntArray(reader *bufio.Reader) []int {
//	str, _ := reader.ReadString('\n')
//	return toIntArray(strings.Fields(str))
//}
//
//func toIntArray(slice []string) []int {
//	arr := make([]int, len(slice))
//	for i, numStr := range slice {
//		res, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
//		if err != nil {
//			panic(err)
//		}
//		arr[i] = int(res)
//	}
//	return arr
//}