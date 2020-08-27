package yys_find_img

import (
	"fmt"
	"github.com/lxn/win"
	"math"
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

//处理GetPixel 误差范围值
func (r *Result)Find_Pixels_Tolerance_Scope(x,y int)int{
	d := x-y //范围误差值 -1 0 1
	d =int(math.Abs(float64(d)))//取绝对值
	return d
}

//减少系统开销,特殊特定场景下 使用 像素取色代替
func (r *Result)Find_Pixels(colorxy [][]int)bool{
	hwnd := getyyshwnd.Get_expvar_hwnd()
	hdc :=win.GetDC(hwnd)
	defer win.DeleteDC(hdc)
	for i,_ :=range colorxy{
		colorPixel :=win.GetPixel(hdc,int32(colorxy[i][0]),int32(colorxy[i][1]))
		c :=int(colorPixel)//取窗口中的某个像素色值
		b :=int(colorxy[i][2])//传参列表第二个色值
		d :=r.Find_Pixels_Tolerance_Scope(c,b)
		if d >1{
			//if d <0{fmt.Println("Find_Pixels",d)}
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
	c :=int(colorPixel)
	d :=r.Find_Pixels_Tolerance_Scope(c,colorrfe)
	if d>1 {
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