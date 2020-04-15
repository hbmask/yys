package main

import (
	"expvar"
	"testing"
	"yys/getyyshwnd"
)

func Test_XuanShang(t *testing.T){
	YYSHWND := getyyshwnd.YYSHWND{}
	hwnd:=YYSHWND.Get_yys_hwnd()
	e:=expvar.NewInt("erhwnd")
	e.Set(int64(hwnd))
	tf :=new(TFMain)
	tf.XuanShang()
}