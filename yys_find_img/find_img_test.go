package yys_find_img

import (
	"encoding/base64"
	"fmt"
	"github.com/lxn/win"
	"gocv.io/x/gocv"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
	"yys/data"
	"yys/yys_screenshot"
)



func Test_find_template(t *testing.T){
	r:=new(Result)

	go func() {
		ticker := time.NewTicker(time.Second)

		for {
			<-ticker.C

			time.Sleep(time.Millisecond*200)

			//御魂->业原火->上锁->挑战
			//Yuhun_3_suo_flag:=r.Recognition(data.Yuhun_3_suo_flag,0.9)
			//if Yuhun_3_suo_flag!=nil {
			Yuhun_4_suo_tiaozhan_click:=r.Recognition(data.Yuhun_4_suo_tiaozhan_click,0.9)
			//if Yuhun_4_suo_tiaozhan_click!=nil {
			//	action.Dj_click(Yuhun_4_suo_tiaozhan_click)
			//	time.Sleep(time.Second*1)
			//}
			//}
			fmt.Println(Yuhun_4_suo_tiaozhan_click)
			//gocv.MatProfile.Count()
		}

	}()

	http.ListenAndServe("localhost:6060", nil)




	//fmt.Println(find_template(im_source,im_search,0.85))
	//got := r.Find_template(im_source,im_search,0.88)

	//find_all_template(im_source,im_search,0.88)

	//want :=got
	//if !reflect.DeepEqual(want,got){
	//	t.Errorf("excepted:%v,got:%v",want,got)
	//}
}

func (r *Result)Test_run_find_screenshot(t *testing.T){
	hd :=yys_screenshot.Yys_windows_screenshot{}
	hwnd :=hd.Get_yys_hwnd()
	rt :=win.RECT{}
	rect :=win.GetWindowRect(hwnd,&rt) //获取阴阳师窗口在桌面的位置
	fmt.Println(rt.Left, rt.Bottom, rt.Right, rt.Top)
	w := rt.Right- rt.Left  //得到阴阳师窗口的宽带
	h := rt.Bottom - rt.Top //得到阴阳师窗口的高度

	fmt.Println("rect:",rect,w,h)

	//var StdEncoding = base64.NewEncoding(data.Jiaceng)
	imgdata,_ :=base64.StdEncoding.DecodeString(data.Jiaceng)//解码 得到一个image byte切片

	//ioutil.WriteFile("./output.png", imgdata, 0666) //还原图像
	img, _ := hd.YYS_Capture()                                //得到一个go类型的窗口句柄图像
	im,_ :=gocv.ImageToMatRGBA(img)                           //得到一个Mat类型的图像
	im_search,_ := gocv.IMDecode(imgdata,gocv.IMReadAnyColor) //从内存中读取 读取 image byte切片
	//im_search :=gocv.IMRead(`D:\go_work\yys\img\3.png`, gocv.IMReadAnyColor)
	defer im_search.Close()
	rs := r.Find_all_template(im,im_search,0.9)
	fmt.Println(rs[0])
	for i,_ :=range rs  {
		fmt.Println(rs[i])
	}
}

func (r *Result)Test_find_all_template(t *testing.T){
	im_source :=gocv.IMRead(`img/1.png`, gocv.IMReadAnyColor)
	im_search :=gocv.IMRead(`img/2.png`, gocv.IMReadAnyColor)

	defer im_source.Close()
	defer im_search.Close()

	//fmt.Println(find_template(im_source,im_search,0.85))
	got := r.Find_all_template(im_source,im_search,0.88)
	fmt.Println(got[0])
	//find_all_template(im_source,im_search,0.88)

	//want :=got
	//if !reflect.DeepEqual(want,got){
	//	t.Errorf("excepted:%v,got:%v",want,got)
	//}
}

func (r *Result)Benchmark_find_all_template(b *testing.B){
	im_source :=gocv.IMRead(`img/1.png`, gocv.IMReadAnyColor)
	im_search :=gocv.IMRead(`img/2.png`, gocv.IMReadAnyColor)

	defer im_source.Close()
	defer im_search.Close()

	//fmt.Println(find_template(im_source,im_search,0.85))
	got := r.Find_all_template(im_source,im_search,0.88)
	fmt.Println(got[0])
	//find_all_template(im_source,im_search,0.88)

	//want :=got
	//if !reflect.DeepEqual(want,got){
	//	t.Errorf("excepted:%v,got:%v",want,got)
	//}
}