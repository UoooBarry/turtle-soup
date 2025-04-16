package main

import (
	"uooobarry/soup/cmd"
	"uooobarry/soup/config"
)

func main() {
	config.InitDB()
	cmd.Execute()
}
