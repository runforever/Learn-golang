package main

import (
        "testing"
)

func TestAdd(t *testing.T) {
        if x := Add(1, 3); x != 4 {
                t.Error("error in test Add")
        }
        if x := Add(1, 3); x != 5 {
                t.Error("error in test Add 5")
        }
}
