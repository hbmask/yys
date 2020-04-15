package yys_find_img

import (
	"fmt"
	"github.com/lxn/win"
	"yys/getyyshwnd"
)

func (r *Result)Find_Pixel(colorxy [][]int){
	//hd :=yys_screenshot.Yys_windows_screenshot{}
	hwnd := getyyshwnd.Get_expvar_hwnd()
	hdc :=win.GetDC(hwnd)
	defer win.DeleteDC(hdc)
	for i,_ :=range colorxy{
		colorPixel :=win.GetPixel(hdc,int32(colorxy[i][0]),int32(colorxy[i][1]))
		fmt.Printf("{%d,%d,%d},\n",colorxy[i][0],colorxy[i][1],colorPixel)
		c :=uint32(colorPixel)
		r.ColorrfeToRGB(c)
	}
}

type ColorPiexxy struct {
}

//减少系统开销,特殊特定场景下 使用 像素取色代替
func (r *Result)Find_Pixels(colorxy [][]int)bool{

	//hd :=yys_screenshot.Yys_windows_screenshot{}
	//hwnd :=hd.Get_yys_hwnd()
	//hd :=Hwnd.YYSHWND{}
	hwnd := getyyshwnd.Get_expvar_hwnd()
	hdc :=win.GetDC(hwnd)
	defer win.DeleteDC(hdc)
	for i,_ :=range colorxy{
		colorPixel :=win.GetPixel(hdc,int32(colorxy[i][0]),int32(colorxy[i][1]))
		c :=uint32(colorPixel)
		b :=uint32(colorxy[i][2])
		if c != b {
			return false
		}
		//r.ColorrfeToRGB(c)
	}
	return true
}


//减少系统开销,特殊特定场景下 使用 像素取色代替
func (r *Result)Find_Pixels_jjtp9num(x,y,colorrfe int)bool{
	//hd :=Hwnd.YYSHWND{}
	hwnd := getyyshwnd.Get_expvar_hwnd()
	hdc :=win.GetDC(hwnd)
	defer win.DeleteDC(hdc)
	colorPixel :=win.GetPixel(hdc,int32(x),int32(y))
	c :=uint32(colorPixel)
	if c != uint32(colorrfe) {
		return false
		//r.ColorrfeToRGB(c)
	}
	return true
}



//Colorrfe 转RGB 方便查看是否取色正确
func (r *Result)ColorrfeToRGB(c uint32){
	rs := byte(c)
	b := byte(c>>8)
	g := byte(c>>16)
	a := byte(c>>24)
	fmt.Println(rs,b,g,a)
}