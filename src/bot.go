package main

import (
	"common"
	"fmt"
	"gWinControl"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/keycode"
	hook "github.com/robotn/gohook"
)

func main() {
	var cts common.TimeStat
	fmt.Printf("----------- begin ------------\n")
	cts.TimeStatInit()

	test4()

	cts.TimeStatShow()
	fmt.Printf("------------ end -------------\n")
	common.PressKeyExit()
}

func test5() {
	gWinControl.CatchAllButton(0, false, func(e hook.Event) {
		fmt.Printf("hook4: %v \n", e.Keychar)
		//robotgo.EventEnd()
	})

	gWinControl.Ctlsleep("s", 1)

	for {
		gWinControl.Ctlsleep("s", 3)
		gWinControl.MouseMove("left", 1, 1, 1, -500, 0)
	}
}

func test4() {
	ch := make(chan int, 10)
	
	go func(ch chan int) {
		enable := false

		ErrOut:
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					break ErrOut
				}
				enable = !enable
				fmt.Printf("开关状体: %v \n", enable)
			default:
				if enable {
					gWinControl.Ctlsleep("ms", 50)
					gWinControl.MouseMove("left", 1, 1, 1, -50, 0)
					gWinControl.Ctlsleep("ms", 50)
					gWinControl.MouseMove("left", 1, 1, 1, 50, 0)
				}
			}
		}
	}(ch)

	gWinControl.CatchAllButton(10, false, func(e hook.Event) {
		//fmt.Printf("hook4: %v \n", e.Keychar)
		if e.Keychar == 92 {
			ch <- 1
		}
		//robotgo.EventEnd()
	})


	for {
		gWinControl.Ctlsleep("s", 10)
		//fmt.Printf("--- sleep 10s --- \n")
	}
}

func test3() {
	gWinControl.Ctlsleep("us", 100)
	robotgo.MoveSmoothRelative(100, 0, 1.0, 2.0)
}

func test2() {
	s := gWinControl.GetScreenInfo()
	fmt.Printf("当前显示器信息: \n")
	fmt.Printf("\tX轴长度: %v, Y轴长度: %v \n", s.Xsize, s.Ysize)
}

func test1() {
	fmt.Printf("robotgo Version 【%s】 \n", robotgo.Version)

	keyCode := keycode.Keycode
	mouseMap := keycode.MouseMap
	special := keycode.Special

	fmt.Printf("\n --- keyCode --- \n")
	for k, v := range keyCode {
		fmt.Printf(" (%v)-(%v) ", k, v)
	}
	fmt.Printf("\n --- keyCode --- \n")

	fmt.Printf("\n --- mouseMap --- \n")
	for k, v := range mouseMap {
		fmt.Printf(" (%v)-(%v) ", k, v)
	}
	fmt.Printf("\n --- mouseMap --- \n")

	fmt.Printf("\n --- special --- \n")
	for k, v := range special {
		fmt.Printf(" (%v)-(%v) ", k, v)
	}
	fmt.Printf("\n --- special --- \n")

	/* 	length, hight := robotgo.GetScreenSize()
	   	fmt.Printf("length(%v) hight(%v) \n", length, hight)

	   	robotgo.MoveMouse(1300, 400)
	   	robotgo.MouseClick(`left`, true)

	   	i, j := robotgo.GetMousePos()
	   	fmt.Printf("i(%v) j(%v) \n", i, j) */

	//robotgo.MoveMouseSmooth(800, 400, 1000.0)
	//robotgo.MoveMouse(800, 400)
}
