package gWinControl

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	//"github.com/vcaesar/keycode"
	hook "github.com/robotn/gohook"
)

type ScreenInfo struct {
	Xsize int
	Ysize int
}

// 等待按键的触发事件 【超时有bug不能用】
//  @param timeout 	时间限制(单位:秒): 0-无限制
func WaitButtonEvent(timeout int) {
	evChan := hook.Start()
	defer hook.End()

timeoutBreak :
	for {
		select {
		case ev := <-evChan:
			fmt.Println("hook: ", ev)
		case <-time.After(time.Duration(timeout) * time.Second):
			fmt.Println("超时,准备退出...")
			break timeoutBreak
		}
	}
}

// 注册指定按键
func RegisterButton(key string) {
	res := robotgo.AddEvent(key)
	fmt.Printf("成功注册键盘事件: %v \n", key)
	if res {
		fmt.Printf("RegisterButton: %v \n", key)
	}
}

// 捕获所有的按键
//  @param timeout 	时间限制(单位:秒): 0-无限制
//  @param wait  	是否阻塞执行, true-阻塞
//  @param callback func(hook.Event): 回调函数
func CatchAllButton(timeout int, wait bool, callback func(hook.Event)) {
	robotgo.EventHook(hook.KeyDown, []string{}, callback)
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)

	if wait {
		WaitButtonEvent(timeout)
	} else {
		go WaitButtonEvent(timeout)
	}
}


// 捕获指定的按键


// 移动鼠标
//  @param button 	按键类型:left(左键),right(右键)
//  @param free 	移动时按住按键 0:松开,1:按住
//  @param smooth 	是否平滑移动 0:瞬移,1:平滑
//  @param relative 鼠标相对/绝对移动 0:绝对,1:相对
//  @param xpos 	x轴坐标
//  @param ypos 	y轴坐标
//  @return err
func MouseMove(button string, free,smooth,relative,xpos,ypos int) error {
	if button != "left" && button != "right" {
		return fmt.Errorf("按键类型只支持: left, right. 不支持(%s)", button)
	}
	if free != 0 {
		robotgo.Toggle(button, "down")
		defer robotgo.Toggle(button, "up")
	}
	if smooth != 0 {	// 平滑移动
		if relative != 0 {	// 相对移动
			robotgo.MoveSmoothRelative(xpos, ypos, 1.0, 2.0)
		} else {	// 绝对移动
			robotgo.MoveSmooth(xpos, ypos)
		}
	} else {	// 瞬移
		if relative != 0 {	// 相对移动
			robotgo.MoveRelative(xpos, ypos)
		} else {	// 绝对移动
			robotgo.Move(xpos, ypos)
		}
	}
	return nil
}

// 获取当前鼠标坐标位置
//  @return xpos 	x轴坐标
//  @return ypos 	y轴坐标
func GetMousePos() (x, y int) {
	return robotgo.GetMousePos()
}

// 获取当前显示器信息
//  @return Xsize	X轴长度
//  @return Ysize	Y轴长度
func GetScreenInfo() ScreenInfo {
	x, y := robotgo.GetScreenSize()
	return ScreenInfo{x, y}
}

// 点击鼠标
//  @param button 	按键类型:left(左键),right(右键)
//  @param double 	0:单击,1:双击
//  @return err
func Mouse(button string, double int) error {
	if button != "left" && button != "right" {
		return fmt.Errorf("按键类型只支持: left, right. 不支持(%s)", button)
	}
	if double != 0 {	// 双击
		robotgo.Click(button, true)
	} else {
		robotgo.Click(button, false)
	}
	return nil
}

// 休眠一段时间
//  @param stype 	休眠类型:s(秒),ms(毫秒),us(微秒)
//  @param tm 		休眠时间
func Ctlsleep(stype string, tm int) {
	if stype == "s" {
		time.Sleep(time.Duration(tm) * time.Second)
	} else if stype == "ms" {
		time.Sleep(time.Duration(tm) * time.Millisecond)
	} else {
		time.Sleep(time.Duration(tm) * time.Microsecond)
	}
}

