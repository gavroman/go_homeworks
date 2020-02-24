package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func equal(l, r *[]string) bool {
	if len(*l) != len(*r) {
		return false
	}
	for i := range *l {
		if (*l)[i] != (*r)[i] {
			return false
		}
	}
	return true
}

func TestOK(t *testing.T) {
	stringsToSort1 := []string{
		"ddddddddddddddd",
		"aaaaaaaaaaaaa",
		"ccccccccccccc",
		"aaaaaaaaaaaaa",
		"aaaaaaaaaaaaa",
		"bbbbbbbbbbbbbb",
		"aaaaaaaaaaaaa",
	}
	stringsToSort2 := []string{
		"123213",
		"123213",
		"123213",
		"34324",
		"123213",
		"5453",
		"65432",
		"355",
	}
	stringsToSort3 := []string{
		"Copyr ighta 2009 TheG Auth",
		"useoa lrigh tsrt aaaa bbbb",
		"uSeOa thisa sour aAaa aaaa",
		"govrn edbya 323  shid cccc",
		"licen setha atca eeee ound",
		"inteL ICENS Efit sfdg frgh",
	}
	var testCase1 = []struct {
		in    []string
		flags Flags
		out   []string
	}{
		{stringsToSort1, Flags{false, true, false, true, -1, ""},
			[]string{
				"ddddddddddddddd",
			},
		},
		{stringsToSort1, Flags{false, false, false, false, -1, ""},
			[]string{
				"aaaaaaaaaaaaa",
				"aaaaaaaaaaaaa",
				"aaaaaaaaaaaaa",
				"aaaaaaaaaaaaa",
				"bbbbbbbbbbbbbb",
				"ccccccccccccc",
				"ddddddddddddddd",
			},
		}, {stringsToSort1, Flags{false, true, false, false, -1, ""},
			[]string{
				"aaaaaaaaaaaaa",
				"bbbbbbbbbbbbbb",
				"ccccccccccccc",
				"ddddddddddddddd",
			},
		}, {stringsToSort1, Flags{false, true, true, false, -1, ""},
			[]string{
				"ddddddddddddddd",
				"ccccccccccccc",
				"bbbbbbbbbbbbbb",
				"aaaaaaaaaaaaa",
			},
		}, {stringsToSort2, Flags{false, false, false, false, -1, ""},
			[]string{
				"123213",
				"123213",
				"123213",
				"123213",
				"34324",
				"355",
				"5453",
				"65432",
			},
		}, {stringsToSort2, Flags{false, false, true, true, -1, ""},
			[]string{
				"123213",
				"123213",
				"123213",
				"123213",
				"65432",
				"34324",
				"5453",
				"355",
			},
		}, {stringsToSort2, Flags{false, true, false, false, -1, ""},
			[]string{
				"123213",
				"34324",
				"355",
				"5453",
				"65432",
			},
		}, {stringsToSort2, Flags{false, true, false, true, -1, ""},
			[]string{
				"355",
				"5453",
				"34324",
				"65432",
				"123213",
			},
		}, {stringsToSort2, Flags{false, true, true, true, -1, ""},
			[]string{
				"123213",
				"65432",
				"34324",
				"5453",
				"355",
			},
		}, {stringsToSort3, Flags{false, false, false, false, -1, ""},
			[]string{
				"Copyr ighta 2009 TheG Auth",
				"govrn edbya 323  shid cccc",
				"inteL ICENS Efit sfdg frgh",
				"licen setha atca eeee ound",
				"uSeOa thisa sour aAaa aaaa",
				"useoa lrigh tsrt aaaa bbbb",
			},
		}, {[]string{
			"Copyr ighta 2009 TheG Auth",
			"useoa lrigh tsrt aaaa bbbb",
			"uSeOa thisa sour aAaa aaaa",
			"govrn edbya 323  shid cccc",
			"licen setha atca eeee ound",
			"inteL ICENS Efit sfdg frgh",},
			Flags{true, true, false, false, 3, ""},
			[]string{
				"useoa lrigh tsrt aaaa bbbb",
				"licen setha atca eeee ound",
				"inteL ICENS Efit sfdg frgh",
				"govrn edbya 323  shid cccc",
				"Copyr ighta 2009 TheG Auth",
			},
		}, {stringsToSort3, Flags{false, false, false, false, 2, ""},
			[]string{
				"Copyr ighta 2009 TheG Auth",
				"govrn edbya 323  shid cccc",
				"inteL ICENS Efit sfdg frgh",
				"licen setha atca eeee ound",
				"uSeOa thisa sour aAaa aaaa",
				"useoa lrigh tsrt aaaa bbbb",
			},
		}, {stringsToSort3, Flags{true, false, false, false, 2, ""},
			[]string{
				"Copyr ighta 2009 TheG Auth",
				"govrn edbya 323  shid cccc",
				"licen setha atca eeee ound",
				"inteL ICENS Efit sfdg frgh",
				"uSeOa thisa sour aAaa aaaa",
				"useoa lrigh tsrt aaaa bbbb",
			},
		},
	}

	for i, testData := range testCase1 {
		t.Run(fmt.Sprint("Test case ", i), func(t *testing.T) {
			err := sortStrings(&testData.in, testData.flags)
			if err != nil {
				t.Errorf("Test %d failed: %s", i, err)
			}
			assert.Equal(t, testData.in, testData.out, fmt.Sprintf("Test case %d failed : wrong sort", i))
		})
	}
}
