package main

import (
	"fmt"
	"testing"
)

func TestLatestLoads(t *testing.T) {
	getconfig()
	dbconnect()

	latestloads := getlatestloads()
	fmt.Println(string(latestloads))
}
