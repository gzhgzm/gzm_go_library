package main

import (
	"common"
	_ "embed"
	"fmt"
	"os"

	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (// dll动态库资源
	DLLFileName = "xcgui.dll"

	//go:embed lib/xcgui.dll
	xcguiDLL []byte
)

var (// ioc图片资源
	
	//go:embed ioc/BiaoTiMain.ico
	BiaoTiMainIcon []byte
	//go:embed ioc/BiaoTiSub.ico
	BiaoTiSubIcon1 []byte
	
	// 窗口图标句柄
	hMainIcon = 0
	hSubIcon1 = 0
)


func main() {

	// ---- 将打包进二进制程序的 xcgui.dll 动态库还原为文件 ----
	if ok, _ := common.PathExists(DLLFileName); !ok {
		// 将二进制的dll内容写入指定文件
		os.WriteFile(DLLFileName, xcguiDLL, 0666)
		fmt.Printf("生成dll文件: %s \n", DLLFileName)
	}


	// ------------------- 创建窗口 -------------------
	// xcgui库 初始化, 参数填true是启用D2D硬件加速, 填false关闭D2D
	a := app.New(true)

	// 创建普通窗口, 宽500, 高300, 窗口标题是"xxx"
	// xcc.Window_Style_Default 使用默认的窗口风格
	// xcc.Window_Style_Drag_Window 允许拖动窗口
	w := window.New(0, 0, 500, 300, "demo xcgui", 0,
			xcc.Window_Style_Default | xcc.Window_Style_Drag_Window)


	// ------------------- 设置窗口属性 -------------------
	// 设置窗口边框大小：窗口标题栏高度35
	w.SetBorderSize(0, 35, 0, 0)

	// 设置窗口透明类型：阴影窗口, 带透明通道, 边框阴影, 窗口透明或半透明
	w.SetTransparentType(xcc.Window_Transparent_Shadow)

	// 设置窗口透明度：255就是不透明
	w.SetTransparentAlpha(255)

	// 设置窗口阴影：阴影大小8, 深度255, 圆角内收大小10, 是否强制直角, 阴影颜色0是黑色
	w.SetShadowInfo(8, 255, 10, false, 0)


	// ------------------- 设置窗口标题栏ioc图标 -------------------
	// 从内存加载图片自适应大小
	hMainIcon = xc.XImage_LoadMemoryAdaptive(BiaoTiMainIcon, 0, 0, 0, 0)

	// 禁止图片句柄自动销毁, 这样就可以复用了, 否则用过之后它会自动释放掉
	xc.XImage_EnableAutoDestroy(hMainIcon, false)

	// 设置窗口图标
	w.SetIcon(hMainIcon)


	// ------------------- 创建新的按钮 -------------------
	// 创建按钮
	btn := widget.NewButton(20, 50, 100, 30, "IP按钮", w.Handle)

	hSubIcon1 = xc.XImage_LoadMemoryAdaptive(BiaoTiSubIcon1, 0, 0, 0, 0)
	xc.XImage_EnableAutoDestroy(hSubIcon1, false)

	// 注册按钮事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		// 创建信息框, 本质是一个模态窗口
		hWindow := a.Msg_Create("IP", "请输入新的IP",
				xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info,
				w.GetHWND(), xcc.Window_Style_Modal)
		
		// 设置窗口边框大小
		xc.XWnd_SetBorderSize(hWindow, 1, 35, 1, 1)
		// 设置窗口图标
		xc.XWnd_SetIcon(hWindow, hSubIcon1)
		// 显示模态窗口
		xc.XModalWnd_DoModal(hWindow)
		return 0
	})


	// 显示窗口
	w.Show(true)

	// 运行消息循环, 程序会被阻塞不退出, 当炫彩窗口数量为0时退出
	a.Run()

	// 退出界面库释放资源
	a.Exit()
}

