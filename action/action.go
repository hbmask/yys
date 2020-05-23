package action

import (
	"fmt"
	"github.com/lxn/win"
	"math/rand"
	"time"
	"yys/getyyshwnd"
	"yys/me_win32"
	"yys/yys_find_img"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

//常用图像匹配点击
func Dj_click(r *yys_find_img.Result) {
	hwnd := getyyshwnd.Get_expvar_hwnd()
	//r := yys_find_img.Result{}
	xy :=r.Result_img_topleft //目标坐标
	cxy :=r.Clickrangevalue  //随机点击值
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	x :=xy[0]+rand.Intn(cxy[0])//屏幕坐标+游戏窗口内容坐标
	y :=xy[1]+rand.Intn(cxy[1])
	fmt.Printf("正确位置:%d 随机偏移范围:%d 偏移后点击位置:%d,%d 相似度:%.2f \n",xy,cxy,x,y,r.Confidence)
	tmp :=me_win32.MAKELPARAM(uint16(x),uint16(y))//将两个16位的数联合成一个无符号的32位数
	//win.SetCursorPos(int32(x+pt.X),int32(y))
	win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	//win.SetCursorPos(int32(x),int32(y))
	win.SendMessage(hwnd,win.WM_MOUSEMOVE,win.MK_LBUTTON,tmp)//移动位置
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Duration(rand.Intn(100)+30))
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开

}

//图像点击便宜
func Dj_click_imgpy(r *yys_find_img.Result,xw ,yh int) {
	hwnd := getyyshwnd.Get_expvar_hwnd()
	//xw yh指定点击随机值
	xy :=r.Result_img_topleft //目标坐标
	cxy :=r.Clickrangevalue  //随机点击值

	//xy 图像点击目标
	//cxy 图像点击时候添加随机值

	x :=xy[0]+xw+rand.Intn(cxy[0])//屏幕坐标+游戏窗口内容坐标
	y :=xy[1]+yh+rand.Intn(cxy[1])
	fmt.Printf("正确位置:%d 随机偏移范围:%d 偏移后点击位置:%d,%d 相似度:%.2f \n",xy,cxy,x,y,r.Confidence)
	tmp :=me_win32.MAKELPARAM(uint16(x),uint16(y))//将两个16位的数联合成一个无符号的32位数
	//win.SetCursorPos(int32(x+pt.X),int32(y))
	win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	//win.SetCursorPos(int32(x),int32(y))
	win.SendMessage(hwnd,win.WM_MOUSEMOVE,win.MK_LBUTTON,tmp)//移动位置
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Duration(rand.Intn(100)+30))
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开

}

//点击厕纸
func Dj_clicks(r []*yys_find_img.Result) {
	hwnd := getyyshwnd.Get_expvar_hwnd()
	for i,_ :=range r{
		xy :=r[i].Result_img_topleft //目标坐标
		cxy :=r[i].Clickrangevalue   //随机点击值
		fmt.Println("目标位置:",xy,cxy)
		//xy 图像点击目标
		//cxy 图像点击时候添加随机值

		//pt :=win.POINT{}
		//win.ClientToScreen(hwnd,&pt)//获取一个正确的屏幕坐标
		//win.ScreenToClient(hwnd,&pt)
		x :=xy[0]+rand.Intn(cxy[0])//屏幕坐标+游戏窗口内容坐标
		//x :=xy[0]+int(pt.X)//屏幕坐标+游戏窗口内容坐标
		//y :=xy[1]+int(pt.Y)//	x :=xy[0]+int(pt.X)//屏幕坐标+游戏窗口内容坐标
		y :=xy[1]+rand.Intn(cxy[1])
		//fmt.Println("Dj_click PT:",pt,x,y,r.Confidence)
		fmt.Println("Dj_click PT:",x,y,r[i].Confidence)
		tmp :=me_win32.MAKELPARAM(uint16(x),uint16(y))//将两个16位的数联合成一个无符号的32位数
		win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
		//win.SetCursorPos(int32(x),int32(y))
		win.SendMessage(hwnd,win.WM_MOUSEMOVE,win.MK_LBUTTON,tmp)//移动位置
		win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
		time.Sleep(time.Duration(rand.Intn(100)+30))
		win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	}
}

//退出专用
func DJ_Click_TuiChu() {
	hwnd := getyyshwnd.Get_expvar_hwnd()
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	x :=10+rand.Intn(1100)//屏幕坐标+游戏窗口内容坐标
	y :=610+rand.Intn(20)
	fmt.Printf("偏移后点击位置:%d,%d  \n",x,y)
	tmp :=me_win32.MAKELPARAM(uint16(x),uint16(y))//将两个16位的数联合成一个无符号的32位数
	win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	win.SendMessage(hwnd,win.WM_MOUSEMOVE,win.MK_LBUTTON,tmp)//移动位置
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Duration(rand.Intn(100)+30))
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	time.Sleep(time.Millisecond*500)

}

//指定点击范围
func DJ_Click_Range(x,y,xr,yr int) {
	hwnd := getyyshwnd.Get_expvar_hwnd()
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	cx :=x+rand.Intn(xr)//屏幕坐标+游戏窗口内容坐标
	cy :=y+rand.Intn(yr)
	fmt.Printf("偏移后点击位置:%d,%d  \n",cx,cy)
	tmp :=me_win32.MAKELPARAM(uint16(cx),uint16(cy))//将两个16位的数联合成一个无符号的32位数
	win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	win.SendMessage(hwnd,win.WM_MOUSEMOVE,win.MK_LBUTTON,tmp)//移动位置
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Duration(rand.Intn(100)+30))
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	time.Sleep(time.Millisecond*500)
}