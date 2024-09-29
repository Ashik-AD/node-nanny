package service

import (
	"fmt"
	"testing"
)

func TestDbConnection(t *testing.T) {
	got := InitConnection()
	if got == nil {
        t.Error("Failed to connect database")
    } else {
        fmt.Println(got)
    }
}
