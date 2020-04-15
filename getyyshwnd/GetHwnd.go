package getyyshwnd

import (
	"expvar"
	"github.com/lxn/win"
	wd "golang.org/x/sys/windows"
	"strconv"
	"yys/me_win32"
)
//func (yh YYSHWND)init(){
//	yh.YYSHWND =yh.Get_yys_hwnd()
//}

//得到一个句柄切片
var list_hwnd YYSHWND //接受 遍历的hwnd

type YYSHWND struct {
	YYSHWND win.HWND
	ALL_YYSHWND []win.HWND
}
//获取单个窗口句柄
func (yh YYSHWND)Get_yys_hwnd()win.HWND{
	//c :=expvar.NewInt("ab")
	//c.Set(9999)
	var none [20]uint16
	var yys *uint16
	retname :=&none[0]//给空指针一个地址接受值
	yys= wd.StringToUTF16Ptr("阴阳师-网易游戏")
	yh.YYSHWND =win.FindWindow(nil,yys) //获得指定窗口句柄
	me_win32.GetWindowText(yh.YYSHWND, retname,100)
	return yh.YYSHWND
}

//获取所有同名字窗口的HWND句柄
func (yh *YYSHWND)Get_list_hwnd(){
	//遍历所有顶级窗口,返回需求的句柄
	me_win32.EnumWindows(wd.NewCallback(yh.Get_yys_all_hwnd),0)
	//fmt.Println(len(list_hwnd.all_hwnd))
}

//活得句柄列表的一个回调函数
func (yh *YYSHWND)Get_yys_all_hwnd(yys_HWND win.HWND)int{
	var none [200]uint16
	retname :=&none[0]
	yys_window_name :="阴阳师-网易游戏"
	me_win32.GetWindowText(yys_HWND, retname,100)
	ret_yys_window_name:=win.UTF16PtrToString(retname)
	if me_win32.IsWindowVisible(yys_HWND) && me_win32.IsWindow(yys_HWND)&&me_win32.IsWindow(yys_HWND)&&
		ret_yys_window_name ==yys_window_name{
		list_hwnd.ALL_YYSHWND = append(list_hwnd.ALL_YYSHWND,yys_HWND)
		return 2
	}
	return 1
}

func Get_expvar_hwnd()win.HWND{
	hd :=expvar.Get("erhwnd")
	i,_ := strconv.Atoi(hd.String())
	d :=win.HWND(i)
	return d
}

func Getdc()win.HDC{
	hwnd :=Get_expvar_hwnd()
	hdc :=win.GetDC(hwnd)
	defer win.DeleteDC(hdc)
	return hdc
}