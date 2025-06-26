package main

import "main/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
