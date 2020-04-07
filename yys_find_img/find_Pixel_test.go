package yys_find_img

import (
	"testing"
	"yys/GetYYShwnd"
)


//手动取色
func Test_find_Pixel(t *testing.T){
	h := GetYYShwnd.YYSHWND{}
	r :=new(Result)
	xyp:=[][]int{{960,37},{1025,37},{1094,37}}
	r.Find_Pixel(h.Get_yys_hwnd(),xyp)
}

func Test_find_Pixels(t *testing.T){
	h := GetYYShwnd.YYSHWND{}
	r :=new(Result)
	xyp:=[][]int{{408,136,1054842},{408,184,1121186},{450,184,1121187}}
	r.Find_Pixels(h.Get_yys_hwnd(),xyp)
}
