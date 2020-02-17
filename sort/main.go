package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	conv "strconv"
	strs "strings"
)

func readStringsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	reader := bufio.NewReader(file)

	var str string
	result := make([]string, 0, 8)
	for err == nil {
		str, err = reader.ReadString('\n')
		str = strs.TrimRight(str, "\n")
		result = append(result, str)
	}

	if result[len(result)-1] == "" {
		result = result[:len(result)-1]
	}

	return result
}

func writeStringsToFile(stringsToWrite *[]string, filename string) error {
	if stringsToWrite == nil {
		return nil
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)
	for i := range *stringsToWrite {
		(*stringsToWrite)[i] += "\n"
		_, err := writer.WriteString((*stringsToWrite)[i])
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func removeDuplicates(slice []string, equal func(l, r string) bool) []string {
	n := len(slice)
	if n <= 1 {
		return slice
	}
	j := 1
	for i := 1; i != n; i++ {
		if !equal(slice[i], slice[i-1]) {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:j]
}

func reverseSlice(slice *[]string) {
	for i := len(*slice)/2 - 1; i != -1; i-- {
		j := len(*slice) - 1 - i
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

func getWordFromString(s *string, position int) string {
	result := ""
	words := strs.Fields(*s)
	if position < len(words) {
		result = words[position]
	}
	return result
}

func sortStrings(stringSlice *[]string, flags map[byte]*bool, kFlagVal int) {
	less := func(i, j int) bool {
		lString, rString := (*stringSlice)[i], (*stringSlice)[j]
		if kFlagVal != -1 {
			lString = getWordFromString(&lString, kFlagVal)
			rString = getWordFromString(&rString, kFlagVal)
		}
		if *flags['n'] {
			lNumber, _ := conv.Atoi(lString)
			rNumber, _ := conv.Atoi(rString)
			return lNumber < rNumber
		}
		if *flags['f'] {
			lString, rString = strs.ToLower(lString), strs.ToLower(rString)
		}
		return lString < rString
	}
	sort.Slice(*stringSlice, less)

	if *flags['u'] {
		equal := func(l, r string) bool {
			if kFlagVal != -1 {
				l = getWordFromString(&l, kFlagVal)
				r = getWordFromString(&r, kFlagVal)
			}
			if *flags['n'] {
				lNumber, _ := conv.Atoi(l)
				rNumber, _ := conv.Atoi(r)
				return lNumber == rNumber
			}
			if *flags['f'] {
				l, r = strs.ToLower(l), strs.ToLower(r)
			}
			return l == r
		}
		*stringSlice = removeDuplicates(*stringSlice, equal)
	}
	if *flags['r'] {
		reverseSlice(stringSlice)
	}
}

func main() {
	sortFlags := map[byte]*bool{
		'f': flag.Bool("f", false, "Case insensitive"),
		'u': flag.Bool("u", false, "Remove duplicates"),
		'r': flag.Bool("r", false, "Reverse result"),
		'n': flag.Bool("n", false, "Sort as numbers"),
	}
	outputFilename := flag.String("o", "", "Output filename")
	kFlagValue := flag.Int("k", -1, "Sort by column")
	flag.Parse()
	inputFilename := flag.Arg(0)

	stringsFromFile := readStringsFromFile(inputFilename)
	if len(stringsFromFile) == 0 {
		return
	}

	sortStrings(&stringsFromFile, sortFlags, *kFlagValue)

	if outputFilename != nil && *outputFilename != "" {
		err := writeStringsToFile(&stringsFromFile, *outputFilename)
		if err != nil {
			return
		}
	} else {
		for idx := range stringsFromFile {
			fmt.Println(stringsFromFile[idx])
		}
	}
}
