package gWinControl

import (
	"testing"
	"fmt"
)


func TestMouseMove(t *testing.T) {
	err := MouseMove("left", 1, 1, 1, -300, 0)
	if err != nil {
		t.Errorf("报错信息(%v)", err)
	}
}

func TestGetMousePos(t *testing.T) {
	x, y := GetMousePos()
	fmt.Printf("当前鼠标位置: X轴(%v), Y轴(%v) \n", x, y)
}

func TestGetScreenInfo(t *testing.T) {
	s := GetScreenInfo()
	fmt.Printf("当前显示器信息: \n")
	fmt.Printf("\tX轴长度: %v, Y轴长度: %v \n", s.Xsize, s.Ysize)
}
