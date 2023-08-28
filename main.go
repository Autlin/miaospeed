package main

import (
	"github.com/Autlin/miaospeed/utils"

	"net/http"
	_ "net/http/pprof"
)

var COMPILATIONTIME string
var BUILDCOUNT string
var COMMIT string
var BRAND string

func main() {
	utils.COMPILATIONTIME = COMPILATIONTIME
	utils.BUILDCOUNT = BUILDCOUNT
	utils.COMMIT = COMMIT
	utils.BRAND = BRAND

	go func() {
		http.ListenAndServe("localhost:8081", nil)
	}()

	RunCli()
}
