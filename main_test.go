package main

import (
	"expvar"
	"testing"
	"yys/GetYYShwnd"
)

func Test_XuanShang(t *testing.T){
	YYSHWND := GetYYShwnd.YYSHWND{}
	hwnd:=YYSHWND.Get_yys_hwnd()
	e:=expvar.NewInt("erhwnd")
	e.Set(int64(hwnd))
	tf :=new(TFMain)
	tf.XuanShang()
}