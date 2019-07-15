package main

import (
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {

	app.UseRouter()
	run := m.Run()
	os.Exit(run)

}
