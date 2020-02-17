package main

import (
    "math"
    "testing"
)

func TestOK(t *testing.T) {
    EPS := 0.01
    inputData := []string{
        "1+2-34",
        "4+100/4-2*9",
        "11-123*(10-2+1)",
        "14-34+34*(543-32)/1928*333-(13-3)",
        "(1-2-3-4-16/2+10+4-48)*24/2*(15-2*(17+4)-28+5)",
    }
    expects := []float64{
        -31,
        11,
        -1096,
        2970.7997,
        30000,
    }

    for i := range inputData {
        out, err := reversePolishNotation(inputData[i])
        if err != nil {
            t.Errorf("Test case %d failed: %s", i, err)
        }
        result, err := calcFromPolishNotation(out)
        if err != nil {
            t.Errorf("Test case %d failed: %s", i, err)
        }
        if math.Abs(result-expects[i]) > EPS {
            t.Errorf("Test case %d failed: result mistmach\n Got: %f\n Expected: %f,", i, result, expects[i])
        }
    }
}

func TestFailCase1(t *testing.T) {
    _, err := reversePolishNotation("1+2-34*&")
    if err == nil {
        t.Errorf("Test failed, expected error")
    }
}
