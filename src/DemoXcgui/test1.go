package main

import (
	"common"
	_ "embed"
	"fmt"
	"os"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed lib/xcgui.dll
var xcguiDLL []byte

var DLLFileName = "xcgui.dll"

func main() {
	if ok, _ := common.PathExists(DLLFileName); !ok {
		// 将二进制的dll内容写入指定文件
		os.WriteFile(DLLFileName, xcguiDLL, 0666)
		fmt.Printf("生成dll文件: %s \n", DLLFileName)
	}

	// xcgui库 初始化, 参数填true是启用D2D硬件加速, 填false关闭D2D
	a := app.New(true)

	// 创建普通窗口, 宽500, 高300, 窗口标题是"xxx"
	// xcc.Window_Style_Default 使用默认的窗口风格
	w := window.New(0, 0, 500, 300, "demo xcgui", 0, xcc.Window_Style_Default)

	// 显示窗口
	w.Show(true)

	// 运行消息循环, 程序会被阻塞不退出, 当炫彩窗口数量为0时退出
	a.Run()

	// 退出界面库释放资源
	a.Exit()
}


