package main

import "testing"

func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Error("Add(2,3) should be 5")
    }
}

func TestSub(t *testing.T) {
    if Sub(5, 2) != 3 {
        t.Error("Sub(5,2) should be 3")
    }
}

func TestMul(t *testing.T) {
    if Mul(3, 4) != 12 {
        t.Error("Mul(3,4) should be 12")
    }
}
