package yys_screenshot

import (
	"encoding/base64"
	"fmt"
	"github.com/lxn/win"
	"gocv.io/x/gocv"
	"image/png"
	"os"
	"testing"
	"yys/data"
	"yys/yys_find_img"
)



func Test_JieTuCaptureRect(T *testing.T){
	hd :=new(Yys_windows_screenshot)
	img, _ := hd.YYS_Capture()//得到一个go类型的窗口句柄图像
	//ioutil.WriteFile("./output.png", img, 0666) //还原图像
	file, _ := os.Create("test.png")
	png.Encode(file,img)
}


func Test_CaptureRect(T *testing.T){
	hd :=new(Yys_windows_screenshot)
	r :=yys_find_img.Result{}
	rt :=win.RECT{}
	rect :=win.GetWindowRect(YYS_HWND,&rt) //获取阴阳师窗口在桌面的位置
	//fmt.Println(rt.Left, rt.Bottom, rt.Right, rt.Top)
	w := rt.Right- rt.Left  //得到阴阳师窗口的宽带
	h := rt.Bottom - rt.Top //得到阴阳师窗口的高度
	hd.Get_yys_hwnd()
	fmt.Println("rect:",rect,w,h)

	//var StdEncoding = base64.NewEncoding(data.Jiaceng)
	imgdata,_ :=base64.StdEncoding.DecodeString(data.Jiaceng)//解码 得到一个image byte切片

	//ioutil.WriteFile("./output.png", imgdata, 0666) //还原图像
	img, _ := hd.YYS_Capture()//得到一个go类型的窗口句柄图像
	im,_ :=gocv.ImageToMatRGBA(img)//得到一个Mat类型的图像
	im_search,_ := gocv.IMDecode(imgdata,gocv.IMReadAnyColor)//从内存中读取 读取 image byte切片
	//im_search :=gocv.IMRead(`D:\go_work\yys\img\3.png`, gocv.IMReadAnyColor)
	defer im_search.Close()
	rs :=r.Find_all_template(im,im_search,0.9)
	fmt.Println(rs[0])
	for i,_ :=range rs  {
		fmt.Println(rs[i])
	}
	//fileName := fmt.Sprintf("%d_%dx%d.png", 1, 1136, 642)
	//file, _ := os.Create(fileName)
	//defer file.Close()
	//png.Encode(file, img)
	//fmt.Printf("%T:",img)
	//fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	//im :=gocv.NewMat()
	//im.ToBytes()
	//show(img)
}

func Benchmark_CaptureRect(T *testing.B){
	hd :=new(Yys_windows_screenshot)
	//n := screenshot.NumActiveDisplays()
	//var none *uint16
	//var yname *uint16
	//yname= wd.StringToUTF16Ptr("阴阳师-网易游戏")
	////fmt.Println(*yname)
	//YYS_HWND :=win.FindWindow(none,yname)//获得指定窗口句柄
	//fmt.Println("窗口句柄:",YYS_HWND)
	r:=yys_find_img.Result{}
	rt :=win.RECT{}
	rect :=win.GetWindowRect(YYS_HWND,&rt) //获取阴阳师窗口在桌面的位置
	fmt.Println(rt.Left, rt.Bottom, rt.Right, rt.Top)
	w := rt.Right- rt.Left  //得到阴阳师窗口的宽带
	h := rt.Bottom - rt.Top //得到阴阳师窗口的高度

	fmt.Println("rect:",rect,w,h)

	//bounds := screenshot.GetDisplayBounds(i)
	//fmt.Println(bounds)
	//img, err := screenshot.CaptureRect(bounds)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%T ",img)
	img, _ := hd.YYS_Capture()
	im,_ :=gocv.ImageToMatRGBA(img)
	im_search :=gocv.IMRead(`D:\go_work\yys\img\2.png`, gocv.IMReadAnyColor)
	rs :=r.Find_template(im,im_search,0.85)
	fmt.Println(rs)

	//fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	//im :=gocv.NewMat()
	//im.ToBytes()
	//show(img)

}