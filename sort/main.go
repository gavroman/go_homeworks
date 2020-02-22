package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	conv "strconv"
	strs "strings"
)

type Flags struct {
	F bool
	U bool
	R bool
	O string
	N bool
	K int
}

var flags Flags

func readStringsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var str string
	var result []string
	for scanner.Scan() {
		str = scanner.Text()
		result = append(result, str)
	}

	return result
}

func writeStrings(stringsToWrite *[]string, file *os.File) error {
	if stringsToWrite == nil {
		return errors.New("no strings to write")
	}

	writer := io.Writer(file)
	for i := range *stringsToWrite {
		_, err := writer.Write([]byte((*stringsToWrite)[i] + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func removeDuplicates(slice []string, equal func(l, r int) bool) []string {
	n := len(slice)
	if n <= 1 {
		return slice
	}
	j := 1
	for i := 1; i != n; i++ {
		if !equal(i, i-1) {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:j]
}

func getWordFromString(s *string, position int) string {
	result := ""
	words := strs.Fields(*s)
	if position < len(words) {
		result = words[position]
	}
	return result
}

func sortStrings(stringSlice *[]string, flags Flags) error {
	if stringSlice == nil {
		return errors.New("no data to sort")
	}

	processFlags := func(sign byte, i, j int) bool {
		lString, rString := (*stringSlice)[i], (*stringSlice)[j]
		if flags.K != -1 {
			lString = getWordFromString(&lString, flags.K)
			rString = getWordFromString(&rString, flags.K)
		}
		if flags.N {
			lNumber, _ := conv.Atoi(lString)
			rNumber, _ := conv.Atoi(rString)
			if sign == '<' {
				return lNumber < rNumber
			} else {
				return lNumber == rNumber
			}
		}
		if flags.F {
			lString, rString = strs.ToLower(lString), strs.ToLower(rString)
		}
		if sign == '<' {
			return lString < rString
		} else {
			return lString == rString
		}
	}

	sort.Slice(*stringSlice, func(i, j int) bool {
		result := processFlags('<', i, j)
		if flags.R {
			return !result
		}
		return result
	})

	if flags.U {
		equal := func(i, j int) bool {
			return processFlags('=', i, j)
		}
		*stringSlice = removeDuplicates(*stringSlice, equal)
	}

	return nil
}

func main() {
	flag.BoolVar(&flags.F, "f", false, "Case insensitive")
	flag.BoolVar(&flags.U, "u", false, "Remove duplicates")
	flag.BoolVar(&flags.R, "r", false, "Reverse result")
	flag.BoolVar(&flags.N, "n", false, "Sort as numbers")
	flag.StringVar(&flags.O, "o", "", "Output filename")
	flag.IntVar(&flags.K, "k", -1, "Sort by column")
	flag.Parse()
	inputFilename := flag.Arg(0)

	stringsFromFile := readStringsFromFile(inputFilename)
	if len(stringsFromFile) == 0 {
		log.Fatal("no strings read")
	}

	err := sortStrings(&stringsFromFile, flags)
	if err != nil {
		log.Fatal(err)
	}

	if flags.O != "" {
		outputFile, err := os.Create(flags.O)
		if err != nil {
			log.Fatal(err)
		}
		defer outputFile.Close()
		err = writeStrings(&stringsFromFile, outputFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = writeStrings(&stringsFromFile, os.Stdin)
	}
}
