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
    var flags Flags
    stringsToSort := []string{
        "ddddddddddddddd",
        "aaaaaaaaaaaaa",
        "ccccccccccccc",
        "aaaaaaaaaaaaa",
        "aaaaaaaaaaaaa",
        "bbbbbbbbbbbbbb",
        "aaaaaaaaaaaaa",
    }
    flags.F = false
    flags.U = false
    flags.R = false
    flags.N = false
    flags.K = -1

    sortedStrings := []string{
        "aaaaaaaaaaaaa",
        "aaaaaaaaaaaaa",
        "aaaaaaaaaaaaa",
        "aaaaaaaaaaaaa",
        "bbbbbbbbbbbbbb",
        "ccccccccccccc",
        "ddddddddddddddd",
    }

    err := sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 0 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 0 failed: wrong sort")
    }

    flags.F = false
    flags.U = true
    flags.R = false
    flags.N = false
    sortedStrings = []string{
        "aaaaaaaaaaaaa",
        "bbbbbbbbbbbbbb",
        "ccccccccccccc",
        "ddddddddddddddd",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 1 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 1 failed: wrong sort")
    }

    flags.F = false
    flags.U = true
    flags.R = true
    flags.N = false
    sortedStrings = []string{
        "ddddddddddddddd",
        "ccccccccccccc",
        "bbbbbbbbbbbbbb",
        "aaaaaaaaaaaaa",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 2 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 2 failed: wrong sort")
    }

    flags.F = false
    flags.U = true
    flags.R = true
    flags.N = false
    sortedStrings = []string{
        "ddddddddddddddd",
        "ccccccccccccc",
        "bbbbbbbbbbbbbb",
        "aaaaaaaaaaaaa",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 3 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 3 failed: wrong sort")
    }

    flags.F = false
    flags.U = true
    flags.R = false
    flags.N = true
    sortedStrings = []string{
        "ddddddddddddddd",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 4 failed: %s", err)
    }
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
    var flags Flags
    flags.F = false
    flags.U = false
    flags.R = false
    flags.N = false
    flags.K = -1

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
    err := sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 0 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 0 failed: wrong sort")
    }

    flags.F = false
    flags.U = false
    flags.R = true
    flags.N = true
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

    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 1 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 1 failed: wrong sort")
    }

    flags.F = false
    flags.U = true
    flags.R = false
    flags.N = false
    sortedStrings = []string{
        "123213",
        "34324",
        "355",
        "5453",
        "65432",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 2 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 2 failed: wrong sort")
    }

    flags.F = false
    flags.U = true
    flags.R = false
    flags.N = true
    sortedStrings = []string{
        "355",
        "5453",
        "34324",
        "65432",
        "123213",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 3 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 3 failed: wrong sort")
    }

    flags.F = false
    flags.U = true
    flags.R = true
    flags.N = true
    sortedStrings = []string{
        "123213",
        "65432",
        "34324",
        "5453",
        "355",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 4 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 4 failed: wrong sort")
    }
}

func TestOK3(t *testing.T) {
    var flags Flags
    stringsToSort := []string{
        "Copyr ighta 2009 TheG Auth",
        "useoa lrigh tsrt aaaa bbbb",
        "uSeOa thisa sour aAaa aaaa",
        "govrn edbya 323  shid cccc",
        "licen setha atca eeee ound",
        "inteL ICENS Efit sfdg frgh",
    }
    flags.F = false
    flags.U = false
    flags.R = false
    flags.N = false
    flags.K = -1
    sortedStrings := []string{
        "Copyr ighta 2009 TheG Auth",
        "govrn edbya 323  shid cccc",
        "inteL ICENS Efit sfdg frgh",
        "licen setha atca eeee ound",
        "uSeOa thisa sour aAaa aaaa",
        "useoa lrigh tsrt aaaa bbbb",
    }

    err := sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 0 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 0 failed: wrong sort")
    }

    flags.F = true
    flags.U = false
    flags.R = false
    flags.N = false
    sortedStrings = []string{
        "Copyr ighta 2009 TheG Auth",
        "govrn edbya 323  shid cccc",
        "inteL ICENS Efit sfdg frgh",
        "licen setha atca eeee ound",
        "useoa lrigh tsrt aaaa bbbb",
        "uSeOa thisa sour aAaa aaaa",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 1 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 1 failed: wrong sort")
    }

    flags.F = false
    flags.U = false
    flags.R = false
    flags.N = false
    flags.K = 2
    sortedStrings = []string{
        "Copyr ighta 2009 TheG Auth",
        "govrn edbya 323  shid cccc",
        "inteL ICENS Efit sfdg frgh",
        "licen setha atca eeee ound",
        "uSeOa thisa sour aAaa aaaa",
        "useoa lrigh tsrt aaaa bbbb",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 2 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 2 failed: wrong sort")
    }

    flags.F = true
    flags.U = false
    flags.R = false
    flags.N = false
    flags.K = 2
    sortedStrings = []string{
        "Copyr ighta 2009 TheG Auth",
        "govrn edbya 323  shid cccc",
        "licen setha atca eeee ound",
        "inteL ICENS Efit sfdg frgh",
        "uSeOa thisa sour aAaa aaaa",
        "useoa lrigh tsrt aaaa bbbb",
    }
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 3 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 3 failed: wrong sort")
    }

    flags.F = true
    flags.U = true
    flags.R = false
    flags.N = false
    flags.K = 3
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
    err = sortStrings(&stringsToSort, flags)
    if err != nil {
        t.Errorf("Test 4 failed: %s", err)
    }
    if !equal(&sortedStrings, &stringsToSort) {
        t.Errorf("Test 4 failed: wrong sort")
    }
}
