package main

import (
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

func TestOK1(t *testing.T) {
    stringsToSort := []string{
        "ddddddddddddddd",
        "aaaaaaaaaaaaa",
        "ccccccccccccc",
        "aaaaaaaaaaaaa",
        "aaaaaaaaaaaaa",
        "bbbbbbbbbbbbbb",
        "aaaaaaaaaaaaa",
    }
    fFlag := false
    uFlag := false
    rFlag := false
    nFlag := false
    sortFlags := map[byte]*bool{
        'f': &fFlag,
        'u': &uFlag,
        'r': &rFlag,
        'n': &nFlag,
    }
    sortedStrings := []string{
        "aaaaaaaaaaaaa",
        "aaaaaaaaaaaaa",
        "aaaaaaaaaaaaa",
        "aaaaaaaaaaaaa",
        "bbbbbbbbbbbbbb",
        "ccccccccccccc",
        "ddddddddddddddd",
    }
    column := -1

    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 0 failed: wrong sort")
    }

    fFlag = false
    uFlag = true
    rFlag = false
    nFlag = false
    sortedStrings = []string{
        "aaaaaaaaaaaaa",
        "bbbbbbbbbbbbbb",
        "ccccccccccccc",
        "ddddddddddddddd",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 1 failed: wrong sort")
    }

    fFlag = false
    uFlag = true
    rFlag = true
    nFlag = false
    sortedStrings = []string{
        "ddddddddddddddd",
        "ccccccccccccc",
        "bbbbbbbbbbbbbb",
        "aaaaaaaaaaaaa",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 2 failed: wrong sort")
    }

    fFlag = false
    uFlag = true
    rFlag = true
    nFlag = false
    sortedStrings = []string{
        "ddddddddddddddd",
        "ccccccccccccc",
        "bbbbbbbbbbbbbb",
        "aaaaaaaaaaaaa",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 3 failed: wrong sort")
    }

    fFlag = false
    uFlag = true
    rFlag = false
    nFlag = true
    sortedStrings = []string{
        "ddddddddddddddd",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 4 failed: wrong sort")
    }
}

func TestOK2(t *testing.T) {
    stringsToSort := []string{
        "123213",
        "123213",
        "123213",
        "34324",
        "123213",
        "5453",
        "65432",
        "355",
    }
    fFlag := false
    uFlag := false
    rFlag := false
    nFlag := false
    sortFlags := map[byte]*bool{
        'f': &fFlag,
        'u': &uFlag,
        'r': &rFlag,
        'n': &nFlag,
    }
    column := -1

    sortedStrings := []string{
        "123213",
        "123213",
        "123213",
        "123213",
        "34324",
        "355",
        "5453",
        "65432",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 0 failed: wrong sort")
    }

    fFlag = false
    uFlag = false
    rFlag = true
    nFlag = true
    sortedStrings = []string{
        "123213",
        "123213",
        "123213",
        "123213",
        "65432",
        "34324",
        "5453",
        "355",
    }

    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 1 failed: wrong sort")
    }

    fFlag = false
    uFlag = true
    rFlag = false
    nFlag = false
    sortedStrings = []string{
        "123213",
        "34324",
        "355",
        "5453",
        "65432",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 2 failed: wrong sort")
    }

    fFlag = false
    uFlag = true
    rFlag = false
    nFlag = true
    sortedStrings = []string{
        "355",
        "5453",
        "34324",
        "65432",
        "123213",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 3 failed: wrong sort")
    }

    fFlag = false
    uFlag = true
    rFlag = true
    nFlag = true
    sortedStrings = []string{
        "123213",
        "65432",
        "34324",
        "5453",
        "355",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 4 failed: wrong sort")
    }
}

func TestOK3(t *testing.T) {
    stringsToSort := []string{
        "Copyr ighta 2009 TheG Auth",
        "useoa lrigh tsrt aaaa bbbb",
        "uSeOa thisa sour aAaa aaaa",
        "govrn edbya 323  shid cccc",
        "licen setha atca eeee ound",
        "inteL ICENS Efit sfdg frgh",
    }
    fFlag := false
    uFlag := false
    rFlag := false
    nFlag := false
    sortFlags := map[byte]*bool{
        'f': &fFlag,
        'u': &uFlag,
        'r': &rFlag,
        'n': &nFlag,
    }
    sortedStrings := []string{
        "Copyr ighta 2009 TheG Auth",
        "govrn edbya 323  shid cccc",
        "inteL ICENS Efit sfdg frgh",
        "licen setha atca eeee ound",
        "uSeOa thisa sour aAaa aaaa",
        "useoa lrigh tsrt aaaa bbbb",
    }
    column := -1

    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 0 failed: wrong sort")
    }

    fFlag = true
    uFlag = false
    rFlag = false
    nFlag = false
    sortedStrings = []string{
        "Copyr ighta 2009 TheG Auth",
        "govrn edbya 323  shid cccc",
        "inteL ICENS Efit sfdg frgh",
        "licen setha atca eeee ound",
        "useoa lrigh tsrt aaaa bbbb",
        "uSeOa thisa sour aAaa aaaa",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 1 failed: wrong sort")
    }

    fFlag = false
    uFlag = false
    rFlag = false
    nFlag = false
    column = 2
    sortedStrings = []string{
        "Copyr ighta 2009 TheG Auth",
        "govrn edbya 323  shid cccc",
        "inteL ICENS Efit sfdg frgh",
        "licen setha atca eeee ound",
        "uSeOa thisa sour aAaa aaaa",
        "useoa lrigh tsrt aaaa bbbb",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 2 failed: wrong sort")
    }

    fFlag = true
    uFlag = false
    rFlag = false
    nFlag = false
    column = 2
    sortedStrings = []string{
        "Copyr ighta 2009 TheG Auth",
        "govrn edbya 323  shid cccc",
        "licen setha atca eeee ound",
        "inteL ICENS Efit sfdg frgh",
        "uSeOa thisa sour aAaa aaaa",
        "useoa lrigh tsrt aaaa bbbb",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 3 failed: wrong sort")
    }

    fFlag = true
    uFlag = true
    rFlag = false
    nFlag = false
    column = 3
    stringsToSort = []string{
        "Copyr ighta 2009 TheG Auth",
        "useoa lrigh tsrt aaaa bbbb",
        "uSeOa thisa sour aAaa aaaa",
        "govrn edbya 323  shid cccc",
        "licen setha atca eeee ound",
        "inteL ICENS Efit sfdg frgh",
    }
    sortedStrings = []string{
        "useoa lrigh tsrt aaaa bbbb",
        "licen setha atca eeee ound",
        "inteL ICENS Efit sfdg frgh",
        "govrn edbya 323  shid cccc",
        "Copyr ighta 2009 TheG Auth",
    }
    sortStrings(&stringsToSort, sortFlags, column)
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 4 failed: wrong sort")
    }
}
