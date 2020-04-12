package main

import (
	"fmt"
	"github.com/lxn/win"
	"math/rand"
	"time"
	"yys/GetYYShwnd"
	"yys/me_win32"
	"yys/yys_find_img"
)


//鼠标拖动到指定地点
func (f *TFMain)move_click(flagman []int,r []*yys_find_img.Result,xr int,yr int,s string) {
	hwnd :=GetYYShwnd.Get_expvar_hwnd()
	rand.Seed(time.Now().UnixNano())
	xy :=[]int{}
	cxy :=[]int{}
	ssd :=[]float32{}
	for i,_:=range r  {
		if r[i].Result_img_centen[0]<475&&r[i].Result_img_centen[1]<584{
			xy =r[i].Result_img_centen //目标坐标
			cxy =r[i].Clickrangevalue  //随机点击值
			ssd =r[i].Confidence
			break
		}
	}
	//r := yys_find_img.Result{}
	mvx :=uint16(flagman[0])//鼠标移动到满字范围
	mvy :=uint16(flagman[1]+yr)
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	x :=uint16(xy[0]+10)//+rand.Intn(cxy[0])//游戏窗口内容坐标
	y :=uint16(xy[1]+10)//+rand.Intn(cxy[1])
	fmt.Printf("正确位置:%d 随机偏移范围:%d 偏移后点击位置:%d,%d 相似度:%.2f %s \n",xy,cxy,x,y,ssd,s)
	tmp :=me_win32.MAKELPARAM(x,y)//将两个16位的数联合成一个无符号的32位数 按下位置
	tmp_mv :=me_win32.MAKELPARAM(mvx,mvy)//将两个16位的数联合成一个无符号的32位数 按下位置
	//win.SetCursorPos(int32(x+pt.X),int32(y))
	//win.GetCaretPos(&ptmv)
	//pt :=win.POINT{}
	//win.ClientToScreen(hwnd,&pt)//左面到游戏窗口左上角坐标
	win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	win.SendMessage(hwnd,win.WM_MOUSEMOVE,win.MK_LBUTTON,tmp_mv)//移动位置
	time.Sleep(time.Millisecond*100)
	win.SendMessage(hwnd,win.WM_MOUSEMOVE,win.MK_LBUTTON,tmp_mv)//移动位置
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	//fmt.Println("鼠标最后移动松开的位置",int32(mvx)+pt.X,int32(mvy)+pt.Y)

	//time.Sleep(time.Millisecond*500)
	f.YYSLos(s)
	time.Sleep(time.Millisecond*500)

}
//常用图像匹配点击
func (f *TFMain)Dj_click(r *yys_find_img.Result,s string) {
	hwnd :=GetYYShwnd.Get_expvar_hwnd()
	rand.Seed(time.Now().UnixNano())
	//r := yys_find_img.Result{}
	xy :=r.Result_img_topleft //目标坐标
	cxy :=r.Clickrangevalue  //随机点击值
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	x :=xy[0]+rand.Intn(cxy[0])//屏幕坐标+游戏窗口内容坐标
	y :=xy[1]+rand.Intn(cxy[1])
	fmt.Printf("正确位置:%d 随机偏移范围:%d 偏移后点击位置:%d,%d 相似度:%.2f %s \n",xy,cxy,x,y,r.Confidence,s)
	tmp :=me_win32.MAKELPARAM(uint16(x),uint16(y))//将两个16位的数联合成一个无符号的32位数
	//win.SetCursorPos(int32(x+pt.X),int32(y))
	//win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	//win.SetCursorPos(int32(x),int32(y))
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Millisecond*200)
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	f.YYSLos(s)

}

//图像点击便宜
func (f *TFMain)Dj_click_imgpy(r *yys_find_img.Result,xw ,yh int,s string) {
	hwnd :=GetYYShwnd.Get_expvar_hwnd()
	rand.Seed(time.Now().UnixNano())
	//xw yh指定点击随机值
	xy :=r.Result_img_topleft //目标坐标
	cxy :=r.Clickrangevalue  //随机点击值

	//xy 图像点击目标
	//cxy 图像点击时候添加随机值

	x :=xy[0]+xw+rand.Intn(cxy[0])//屏幕坐标+游戏窗口内容坐标
	y :=xy[1]+yh+rand.Intn(cxy[1])
	fmt.Printf("正确位置:%d 随机偏移范围:%d 偏移后点击位置:%d,%d 相似度:%.2f %s \n",xy,cxy,x,y,r.Confidence,s)
	tmp :=me_win32.MAKELPARAM(uint16(x),uint16(y))//将两个16位的数联合成一个无符号的32位数
	//win.SetCursorPos(int32(x+pt.X),int32(y))
	//win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	//win.SetCursorPos(int32(x),int32(y))
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Millisecond*200)
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	f.YYSLos(s)
}

//点击厕纸
func (f *TFMain)Dj_clicks(r []*yys_find_img.Result,s string) {
	hwnd :=GetYYShwnd.Get_expvar_hwnd()
	rand.Seed(time.Now().UnixNano())
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
		fmt.Println("Dj_click PT:",x,y,r[i].Confidence,s)
		tmp :=me_win32.MAKELPARAM(uint16(x),uint16(y))//将两个16位的数联合成一个无符号的32位数
		//win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
		//win.SetCursorPos(int32(x),int32(y))
		win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
		time.Sleep(time.Millisecond*200)
		win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	}
	f.YYSLos(s)
}

//退出专用
func (f *TFMain)DJ_Click_TuiChu() {
	hwnd :=GetYYShwnd.Get_expvar_hwnd()
	rand.Seed(time.Now().UnixNano())
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	x :=10+rand.Intn(1100)//屏幕坐标+游戏窗口内容坐标
	y :=610+rand.Intn(20)
	fmt.Printf("偏移后点击位置:%d,%d  \n",x,y)
	tmp :=me_win32.MAKELPARAM(uint16(x),uint16(y))//将两个16位的数联合成一个无符号的32位数
	//win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Millisecond*200)
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	//f.YYSLos("<退出战斗>")
}

//指定点击范围
func (f *TFMain)DJ_Click_Range(x,y,xr,yr int,s string) {
	hwnd :=GetYYShwnd.Get_expvar_hwnd()
	rand.Seed(time.Now().UnixNano())
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	cx :=x+rand.Intn(xr)//屏幕坐标+游戏窗口内容坐标
	cy :=y+rand.Intn(yr)
	fmt.Printf("偏移后点击位置:%d,%d %s \n",cx,cy,s)
	tmp :=me_win32.MAKELPARAM(uint16(cx),uint16(cy))//将两个16位的数联合成一个无符号的32位数
	//win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Millisecond*100)
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	f.YYSLos(s)
}
//双击指定点击范围
func (f *TFMain)SJ_Click_Range(x,y,xr,yr int,s string) {
	hwnd :=GetYYShwnd.Get_expvar_hwnd()
	rand.Seed(time.Now().UnixNano())
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	cx :=uint16(x+rand.Intn(xr))//屏幕坐标+游戏窗口内容坐标
	cy :=uint16(y+rand.Intn(yr))
	fmt.Printf("偏移后点击位置:%d,%d %s \n",cx,cy,s)
	//CS_DBLCLKS
	tmp :=me_win32.MAKELPARAM(cx,cy)
	win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,tmp)//激活窗口
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Millisecond*100)
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	time.Sleep(time.Millisecond*100)
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON,tmp)//松开
	f.YYSLos(s)
	time.Sleep(time.Millisecond*500)
}


//鼠标向下拉动
func (f *TFMain)mv_mouse_Range(x,y,xr,yr int,s string) {
	hwnd :=GetYYShwnd.Get_expvar_hwnd()
	//rand.Seed(time.Now().UnixNano())
	//xy 图像点击目标
	//cxy 图像点击时候添加随机值
	cx :=x//+rand.Intn(xr)//屏幕坐标+游戏窗口内容坐标
	cy :=y//+rand.Intn(yr)
	fmt.Printf("偏移后点击位置:%d,%d %s \n",cx,cy,s)
	pt :=win.POINT{}
	win.ClientToScreen(hwnd,&pt)//左面到游戏窗口左上角坐标
	//win.SetCursorPos(int32(x)+pt.X,int32(y)+pt.Y)
	tmp :=me_win32.MAKELPARAM(uint16(cx),uint16(cy))//将两个16位的数联合成一个无符号的32位数
	tmp_mv :=me_win32.MAKELPARAM(uint16(0),uint16(yr))//将两个16位的数联合成一个无符号的32位数
	win.SendMessage(hwnd,win.WM_ACTIVATE,win.WA_ACTIVE,0)//激活窗口
	win.SendMessage(hwnd,win.WM_LBUTTONDOWN,win.MK_LBUTTON,tmp)//按下
	win.SendMessage(hwnd,win.WM_MOUSEMOVE,win.MK_LBUTTON,tmp_mv)//移动位置
	win.SendMessage(hwnd,win.WM_LBUTTONUP,win.MK_LBUTTON ,tmp)//按下
	f.YYSLos(s)
	time.Sleep(time.Millisecond*500)
}