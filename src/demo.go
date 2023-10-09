package main

import (
	"common"
	"fmt"
)

func main() {
	var cts common.TimeStat
	fmt.Printf("----------- begin ------------\n")
	cts.TimeStatInit()

	test5()

	cts.TimeStatShow()
	fmt.Printf("------------ end -------------\n")
	common.PressKeyExit()
}

func test5() {
	print("----------------------\n")
}
