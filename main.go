package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readStringsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("%s", err)
		return []string{}
	}

	reader := bufio.NewReader(file)

	var str string
	result := make([]string, 0, 8)
	for err == nil {
		str, err = reader.ReadString('\n')
		str = strings.TrimRight(str, "\n")
		result = append(result, str)
	}

	return result
}

func removeDuplicates(slice []string) []string {
	n := len(slice)
	if n <= 1 {
		return slice
	}
	j := 1
	for i := 1; i != n; i++ {
		if slice[i] != slice[i-1] {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:j]
}

func reverseSlice(slice *[]string) {
	for i := len(*slice)/2 - 1; i >= 0; i-- {
		j := len(*slice) - 1 - i
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

func sortStrings(stringSlice *[]string, flags map[byte]bool) {
	// no flags
	if len(flags) == 0 {
		sort.Strings(*stringSlice)
	}
	// numbers
	if flags[byte('n')] {
		fmt.Println("n")
		numberSlice := make([]int, len(*stringSlice))
		for idx := range *stringSlice {
			numberSlice[idx], _ = strconv.Atoi((*stringSlice)[idx])
		}
		sort.Ints(numberSlice)
		for idx := range *stringSlice {
			(*stringSlice)[idx] = strconv.Itoa(numberSlice[idx])
		}
	}
	// unique
	if flags[byte('u')] {
		fmt.Println("u")
		*stringSlice = removeDuplicates(*stringSlice)
	}
	if flags[byte('f')] {
		fmt.Println("f")
	}
	if flags[byte('o')] {
		fmt.Println("o")
	}
	if flags[byte('k')] {
		fmt.Println("k")
	}
	// reverse
	if flags[byte('r')] {
		fmt.Println("r")
		reverseSlice(stringSlice)
	}
}

func main() {
	//filename := "test_data/in1.txt"
	filename := "test_data/in2.txt"
	stringsFromFile := readStringsFromFile(filename)
	if len(stringsFromFile) == 0 {
		return
	}

	flags := make(map[byte]bool, 6)
	//flags[byte('r')] = false
	//flags[byte('n')] = false
	//flags[byte('u')] = false
	sortStrings(&stringsFromFile, flags)

	for idx := range stringsFromFile {
		fmt.Println(idx, ")", stringsFromFile[idx])
	}
}
