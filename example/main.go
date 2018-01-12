package main

import (
	curr_dir "github.com/go-steven/curr-dir"
)

func main() {
	currDir := curr_dir.GetCurrDir()
	println(currDir)
}