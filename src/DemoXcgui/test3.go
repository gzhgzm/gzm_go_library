/*
 * 测试数据
 *
 *
 */
package main

import (
	"common"
	_ "embed"
	"fmt"
	"os"

	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
)

var (// dll动态库资源
	DLLFileName = "xcgui.dll"

	//go:embed lib/xcgui.dll
	xcguiDLL []byte
)

var (// gui界面
	//go:embed gui/test3_GUI.zip
	guiZip []byte
)

var ipAddr string

func main() {

	// ---- 将打包进二进制程序的 xcgui.dll 动态库还原为文件 ----
	if ok, _ := common.PathExists(DLLFileName); !ok {
		// 将二进制的dll内容写入指定文件
		os.WriteFile(DLLFileName, xcguiDLL, 0666)
		fmt.Printf("生成dll文件: %s \n", DLLFileName)
	}

	// xcgui库 初始化, 参数填true是启用D2D硬件加速, 填false关闭D2D
	a := app.New(true)

	// 创建窗口从内存压缩包中的布局文件
	w := window.NewByLayoutZipMem(guiZip, "main.xml", "", 0, 0)


	// ---- 注册编辑框 ----
	edit1 := widget.NewEditByName("edit1")


	// ---- 注册按钮 ----
	// 获取窗口布局文件中的按钮
	btn1 := widget.NewButtonByName("btn1")

	btn1.Event_BnClick(func(pbHandled *bool) int {
		ipAddr = edit1.GetTextEx()

		a.Alert("提示", btn1.GetText() + "：" + ipAddr)
		return 0
	})



	// 调整布局, 必须
	w.AdjustLayout()

	// 显示窗口
	w.Show(true)

	// 运行消息循环, 程序会被阻塞不退出, 当炫彩窗口数量为0时退出
	a.Run()

	// 退出界面库释放资源
	a.Exit()
}

