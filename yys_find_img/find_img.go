package yys_find_img

import "C"
import (
	"encoding/base64"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"yys/getyyshwnd"
	"yys/yys_screenshot"
)
//func main() {
//	im_source :=gocv.IMRead(`D:\SynologyDrive\python\yys\imgs\1.png`, gocv.IMReadAnyColor)
//	im_search :=gocv.IMRead(`D:\SynologyDrive\python\yys\imgs\2.png`, gocv.IMReadAnyColor)
//
//	defer im_source.Close()
//	defer im_search.Close()
//
//	fmt.Println(find_template(im_source,im_search,0.85))
//	//fmt.Printf("initial MatProfile count: %v\n", gocv.MatProfile.Count())
//}

type Result struct {
	Result_img_centen  []int//顶点中间坐标
	Result_img_topleft []int//左上角坐标
	Rectangles         [][]int//矩形四角
	Confidence         []float32//相似度
	Clickrangevalue    []int//相似度
}

func (r *Result)NewResult(result_img []int, result_img_topleft []int, rectangles [][]int,confidence []float32,clickrangevalue []int) *Result { // 返回值指向Student结构体的指针
	return &Result{
		Result_img_centen:  result_img,
		Result_img_topleft: result_img_topleft,
		Rectangles:         rectangles,
		Confidence:         confidence,
		Clickrangevalue:    clickrangevalue,}
}
func (r *Result)Show(img gocv.Mat){
	//显示 一个图片
	window :=gocv.NewWindow("显示")
	defer window.Close()
	window.ResizeWindow(1152, 679)
	//imgs :=gocv.IMRead(img,gocv.IMReadGrayScale)
	window.IMShow(img)
	window.WaitKey(0)

}

func (r *Result)Imrad(filename string)gocv.Mat{
	im := gocv.IMRead(filename,gocv.IMReadGrayScale)
	return im
}




//同时匹配多个不同的图案并返回结果集
func (r *Result)RecognitionsBuTongTuAn(im_searchs []string,threshold float32)[]*Result {
	hs :=yys_screenshot.Yys_windows_screenshot{}
	hwnd := getyyshwnd.Get_expvar_hwnd()
	imgMatsearch :=[]gocv.Mat{}
	for i:=0; i<len(im_searchs);i++{
		imgdata,_ :=base64.StdEncoding.DecodeString(im_searchs[i])//解码 得到一个image byte切片
		im_search,_ := gocv.IMDecode(imgdata,gocv.IMReadAnyColor)//从内存中读取 读取 image byte切片
		imgMatsearch =append(imgMatsearch,im_search)
		defer im_search.Close()
	}
	//ioutil.WriteFile("./output.png", imgdata, 0666) //还原图像
	imgh,err := hs.YYS_Capture_HWND(hwnd)//得到一个go类型的窗口句柄图像
	if err!=nil{
		fmt.Println("Recognitions:",err)
		//panic(err)
		return nil
	}
	im,_ :=gocv.ImageToMatRGBA(imgh)//得到一个Mat类型的图像
	defer im.Close()
	//im_search :=gocv.IMRead(`D:\go_work\yys\img\3.png`, gocv.IMReadAnyColor)


	rs :=r.Find_all_templates(im,imgMatsearch,threshold)
	//fmt.Println(rs[0])
	//for i,_ :=range rs  {
	//	fmt.Println(rs[i])
	//}
	return rs
}
//返回多个相同图像集
func (r *Result)Recognitions(im_searchs string,threshold float32)[]*Result {
	hs :=yys_screenshot.Yys_windows_screenshot{}
	hwnd := getyyshwnd.Get_expvar_hwnd()
	//hd :=Hwnd.YYSHWND{}
	//hwnd :=hd.YYSHWND
	//rt :=win.RECT{}
	//rect :=win.GetWindowRect(hwnd,&rt) //获取阴阳师窗口在桌面的位置
	//fmt.Println(rt.Left, rt.Bottom, rt.Right, rt.Top)
	//w := rt.Right- rt.Left  //得到阴阳师窗口的宽带
	//h := rt.Bottom - rt.Top //得到阴阳师窗口的高度
	//fmt.Println("rect:",rect,w,h)

	imgdata,_ :=base64.StdEncoding.DecodeString(im_searchs)//解码 得到一个image byte切片

	//ioutil.WriteFile("./output.png", imgdata, 0666) //还原图像
	imgh,err := hs.YYS_Capture_HWND(hwnd)//得到一个go类型的窗口句柄图像
	if err!=nil{
		fmt.Println("Recognitions:",err)
		//panic(err)
		return nil
	}
	im,_ :=gocv.ImageToMatRGBA(imgh)//得到一个Mat类型的图像
	im_search,_ := gocv.IMDecode(imgdata,gocv.IMReadAnyColor)//从内存中读取 读取 image byte切片
	//im_search :=gocv.IMRead(`D:\go_work\yys\img\3.png`, gocv.IMReadAnyColor)
	defer im_search.Close()
	defer im.Close()
	rs :=r.Find_all_template(im,im_search,threshold)
	//fmt.Println(rs[0])
	//for i,_ :=range rs  {
	//	fmt.Println(rs[i])
	//}
	return rs
}

//狗粮->返回多个相同图像集
func (r *Result)RecognitionsGouLiang_2Man(im_searchs string,x int,y int,threshold float32)[]*Result {
	hs :=yys_screenshot.Yys_windows_screenshot{}
	hwnd := getyyshwnd.Get_expvar_hwnd()

	imgdata,_ :=base64.StdEncoding.DecodeString(im_searchs)//解码 得到一个image byte切片
	imgh,err := hs.YYS_Capture_HWND(hwnd)//得到一个go类型的窗口句柄图像
	if err!=nil{
		fmt.Println("Recognitions:",err)
		//panic(err)
		return nil
	}
	im,_ :=gocv.ImageToMatRGBA(imgh)//得到一个Mat类型的图像
	im_search,_ := gocv.IMDecode(imgdata,gocv.IMReadAnyColor)//从内存中读取 读取 image byte切片
	defer im_search.Close()
	defer im.Close()
	rs :=r.Find_all_template(im,im_search,threshold)
	rsman :=rs
	for i,_ :=range rs{
		if rs[i].Result_img_centen[0]>x{//过滤第三个满级
			if rs[i].Result_img_centen[1]<y{
				rsman =append(rs[:i],rs[i+1:]...)
			}
		}
	}

	return rsman
}


//识别图像
func (r *Result)Recognition(im_searchs string,threshold float32) *Result {
	hwnd := getyyshwnd.Get_expvar_hwnd()
	hs :=yys_screenshot.Yys_windows_screenshot{}
	//hwnd :=hd.Get_yys_hwnd()
	//hd :=Hwnd.YYSHWND{}
	//hwnd :=hd.Get_yys_hwnd()
	//rt :=win.RECT{}
	//rt2 :=win.RECT{}
	//win.GetWindowRect(hwnd,&rt) //获取阴阳师窗口在桌面的位置
	//win.GetClientRect(hwnd,&rt2) //获取阴阳师窗口在桌面的位置
	//fmt.Println("GetWindowRect:",rt.Left, rt.Bottom, rt.Right, rt.Top)
	//fmt.Println("GetClientRect:",rt2.Left, rt2.Bottom, rt2.Right, rt2.Top)
	//fmt.Println(rt.Left, rt.Bottom, rt.Right, rt.Top)
	//w := rt.Right- rt.Left  //得到阴阳师窗口的宽带
	//h := rt.Bottom - rt.Top //得到阴阳师窗口的高度
	//fmt.Println("rect:",rect,w,h)

	//buf := bytes.NewBufferString(im_searchs)
	img_byte,_ :=base64.StdEncoding.DecodeString(im_searchs)//解码 得到一个image byte切片
	//ioutil.WriteFile("./output.png", imgdata, 0666) //还原图像
	imgh, err := hs.YYS_Capture_HWND(hwnd)//得到一个go类型的窗口句柄图像
	if err!=nil{
		fmt.Println("Recognition",err)
		//panic(err)
		return nil
	}
	im,_ :=gocv.ImageToMatRGBA(imgh)//得到一个Mat类型的图像
	defer im.Close()
	im_search,_ := gocv.IMDecode(img_byte,gocv.IMReadAnyColor)//从内存中读取 读取 image byte切片
	defer im_search.Close()
	//defer im_search.Close()
	//im_search :=gocv.IMRead(`D:\go_work\yys\img\3.png`, gocv.IMReadAnyColor)
	rs :=r.Find_template(im,im_search,threshold)
	if rs ==nil{
		return nil
	}
	return rs
}

func (r *Result)Find_template(im_source gocv.Mat,im_search gocv.Mat,threshold float32)*Result {
	//@return find location
	//if not found; return None
	results := r.Find_all_template(im_source, im_search, threshold)
	if len(results) ==0 {
		return nil
	}
	//for i,_ :=range results{
	//	fmt.Println(results[i])
	//}
	return results[0]
}


func (r *Result)Find_all_template(im_source gocv.Mat,im_search gocv.Mat,threshold float32)[]*Result {
	//,threshold float64,maxcnt int, ,bgremove bool
	//Locate image position with cv2.templateFind
	//
	//Use pixel match to find pictures.
	//
	//	Args:
	//im_source(string): 图像、素材
	//im_search(string): 需要查找的图片
	//threshold: 阈值，当相识度小于该阈值的时候，就忽略掉
	//
	//Returns:
	//	A tuple of found [(point, score), ...]
	//
	//Raises:
	//IOError: when file read error
	res :=gocv.NewMat()
	defer res.Close()

	m :=gocv.NewMat()
	defer m.Close()

	i_gray_img :=gocv.NewMat()
	defer i_gray_img.Close()

	s_gray_img :=gocv.NewMat()
	defer s_gray_img.Close()

	//CvtColor 将图像从一个颜色空间转换为另一个颜色空间。
	///它使用包含所需颜色转换代码颜色空间的代码参数将 src Mat 图像转换为 dst Mat。

	gocv.CvtColor(im_search,&s_gray_img, gocv.ColorRGBToGray)//模板
	gocv.CvtColor(im_source,&i_gray_img, gocv.ColorRGBToGray)//源图
	//fmt.Println("im_search 获取三位:",im_search.Channels())
	//show(i_gray_img)
	//show(s_gray_img)
	//fmt.Println("i_gray_img通道数量:",i_gray_img.Channels())
	//fmt.Println(i_gray_img.GetDoubleAt())
	// 边界提取(来实现背景去除的功能)
	//if bgremove:
	//Canny 使用 Canny 算法查找图像中的边缘。函数查找输入图像图像中的边，
	//并使用 Canny 算法在输出贴图边缘中标记它们。阈值 1 和阈值 2 之间的最小值用于边链接。
	//最大值用于查找强边的初始段。
	//gocv.Canny(im_search,&s_gray_img, 100, 200)
	//gocv.Canny(im_source,&i_gray_img, 100, 200)
	//show(i_gray_img)
	//show(s_gray_img)
	//参数解释：
	//InputArray Image: 待搜索的图像，且图像必须为8-bit或32-bit的浮点型图像
	//InputArray templ: 用于进行模板匹配的模板图像，类型和原图像一致，但是尺寸不能大于原图像
	//OutputArray Result: 模板搜索结果输出图像，必须为单通道32-bit位浮点型图像，如果图像尺寸是WxH而template尺寸是wxh，则此参数result一定是(W-w+1)x(H-h+1)
	//int method: 模板匹配计算类型，在匹配原理中已经介绍过这六种方法了，这里不再赘述
	//InputArray mask=noArray(): 图像匹配时用的掩膜板，必须和模板图像有相同的数据类型和尺寸

	results :=make([]*Result,0)

	for {
		//循环查找相匹配的图像,找到以后记录后屏蔽
		gocv.MatchTemplate(i_gray_img, s_gray_img,&res,gocv.TmCcoeffNormed ,m)
		h,w :=im_search.Rows(),im_search.Cols()

		// MinMaxLoc查找数组中的全局最小值和最大值。
		_, max_val, _, max_loc := gocv.MinMaxLoc(res)//返回图像顶点坐标和相似度
		if max_val < threshold{
			//fmt.Printf("%f的最大置信度太低。MatchTemplate在场景中找不到模板。\n", max_val)
			break
		}

		top_left :=max_loc//重命名取 顶端左边
		//calculator middle point计算器中间点
		middle :=image.Pt(top_left.X+w/2,top_left.Y+h/2)
		//显示四个角的坐标
		rectangle :=[][]int{
			{top_left.X,top_left.Y},//左上角
			{top_left.X,top_left.Y+h},//左下角
			{top_left.X+w,top_left.Y+h},//右下角
			{top_left.X+w,top_left.Y},//右上角
		}
		//获取随机点击范围值
		//rxx :=top_left.X+w-top_left.X
		//rxx =rxx/2
		//rxy :=top_left.Y+h-top_left.Y
		//rxy =rxy/2
		clickrangevalue :=[]int{w,h}

		max_vals :=[]float32{max_val}
		result_img :=[]int{middle.X,middle.Y}//模板图像中间点
		result_img_topleft :=[]int{top_left.X,top_left.Y}//模板图像左下角
		results =append(results,
			&Result{Result_img_centen: result_img, //模板图像中间点
				Result_img_topleft: result_img_topleft, //坐下角度
				Rectangles:         rectangle,          //显示四个角的坐标
				Confidence:         max_vals,           //相似度
				Clickrangevalue:    clickrangevalue,    //随机点击范围
			})
		if len(results) ==0 {
			return nil
		}
		//绘图->原图+模板图坐标进行准确绘图 top_left向右偏移模板图size
		// color for the rect when faces detected检测到脸部时画框的颜色
		blue := color.RGBA{0, 0, 255, 0}
		//fmt.Println("随机点击范围:",clickrangevalue)
		//对像素进行标记点,填满位置point ,对识别后的模板进行涂抹
		pts :=[][]image.Point{{//正确涂抹方式
			image.Pt(rectangle[0][0],rectangle[0][1]),//1.先绘制topleft 左上角
			image.Pt(rectangle[1][0],rectangle[1][1]),//2.在绘制topleft+h 左下角
			image.Pt(rectangle[2][0],rectangle[2][1]),//3.topleft+w+h 右下角
			image.Pt(rectangle[3][0],rectangle[3][1]),//4 topleft+w  右上角
							},}
		//执行匹配到以后做图像填充
		gocv.FillPoly(&i_gray_img,pts,blue)
		//template_rectangle := image.Pt(top_left.X,top_left.Y)
		//gocv.Rectangle(&i_gray_img,image.Rectangle{template_rectangle,max_loc}, blue, 2)
		//win :=gocv.NewWindow("123")
		//win.ResizeWindow(1138, 640)
		//win.IMShow(i_gray_img)
		//gocv.WaitKey(0)
		//defer win.Close()
	}

	return results

}


//同时在一张原图上处理多个匹配图像
func (r *Result)Find_all_templates(im_source gocv.Mat,im_search []gocv.Mat,threshold float32)[]*Result {
	//,threshold float64,maxcnt int, ,bgremove bool
	//Locate image position with cv2.templateFind
	//Use pixel match to find pictures.
	//	Args:
	//im_source(string): 图像、素材
	//im_search(string): 需要查找的图片
	//threshold: 阈值，当相识度小于该阈值的时候，就忽略掉
	//
	//Returns:
	//	A tuple of found [(point, score), ...]
	//
	//Raises:
	//IOError: when file read error
	res :=gocv.NewMat()
	defer res.Close()
	m :=gocv.NewMat()
	defer m.Close()
	i_gray_img :=gocv.NewMat()
	defer i_gray_img.Close()
	s_gray_img :=gocv.NewMat()
	defer s_gray_img.Close()
	//CvtColor 将图像从一个颜色空间转换为另一个颜色空间。
	///它使用包含所需颜色转换代码颜色空间的代码参数将 src Mat 图像转换为 dst Mat。
	results :=make([]*Result,0)
	for i:=0;i<len(im_search);i++{
		gocv.CvtColor(im_search[i],&s_gray_img, gocv.ColorRGBToGray)//模板
		gocv.CvtColor(im_source,&i_gray_img, gocv.ColorRGBToGray)//源图
		//fmt.Println("im_search 获取三位:",im_search.Channels())
		//show(i_gray_img)
		//show(s_gray_img)
		//fmt.Println("i_gray_img通道数量:",i_gray_img.Channels())
		//fmt.Println(i_gray_img.GetDoubleAt())
		// 边界提取(来实现背景去除的功能)
		//if bgremove:
		//Canny 使用 Canny 算法查找图像中的边缘。函数查找输入图像图像中的边，
		//并使用 Canny 算法在输出贴图边缘中标记它们。阈值 1 和阈值 2 之间的最小值用于边链接。
		//最大值用于查找强边的初始段。
		//gocv.Canny(im_search,&s_gray_img, 100, 200)
		//gocv.Canny(im_source,&i_gray_img, 100, 200)
		//show(i_gray_img)
		//show(s_gray_img)
		//参数解释：
		//InputArray Image: 待搜索的图像，且图像必须为8-bit或32-bit的浮点型图像
		//InputArray templ: 用于进行模板匹配的模板图像，类型和原图像一致，但是尺寸不能大于原图像
		//OutputArray Result: 模板搜索结果输出图像，必须为单通道32-bit位浮点型图像，如果图像尺寸是WxH而template尺寸是wxh，则此参数result一定是(W-w+1)x(H-h+1)
		//int method: 模板匹配计算类型，在匹配原理中已经介绍过这六种方法了，这里不再赘述
		//InputArray mask=noArray(): 图像匹配时用的掩膜板，必须和模板图像有相同的数据类型和尺寸
		//循环查找相匹配的图像,找到以后记录后屏蔽
		for {
			gocv.MatchTemplate(i_gray_img, s_gray_img,&res,gocv.TmCcoeffNormed ,m)
			h,w :=im_search[i].Rows(),im_search[i].Cols()
			// MinMaxLoc查找数组中的全局最小值和最大值。
			_, max_val, _, max_loc := gocv.MinMaxLoc(res)//返回图像顶点坐标和相似度
			if max_val < threshold{
				//fmt.Printf("%f的最大置信度太低。MatchTemplate在场景中找不到模板。\n", max_val)
				break
			}
			top_left :=max_loc//重命名取 顶端左边
			//calculator middle point计算器中间点
			middle :=image.Pt(top_left.X+w/2,top_left.Y+h/2)
			//显示四个角的坐标
			rectangle :=[][]int{
				{top_left.X,top_left.Y},//左上角
				{top_left.X,top_left.Y+h},//左下角
				{top_left.X+w,top_left.Y+h},//右下角
				{top_left.X+w,top_left.Y},//右上角
			}
			//获取随机点击范围值
			//rxx :=top_left.X+w-top_left.X
			//rxx =rxx/2
			//rxy :=top_left.Y+h-top_left.Y
			//rxy =rxy/2
			clickrangevalue :=[]int{w,h}
			max_vals :=[]float32{max_val}
			result_img :=[]int{middle.X,middle.Y}//模板图像中间点
			result_img_topleft :=[]int{top_left.X,top_left.Y}//模板图像左下角
			results =append(results,
				&Result{Result_img_centen: result_img, //模板图像中间点
					Result_img_topleft: result_img_topleft, //坐下角度
					Rectangles:         rectangle,          //显示四个角的坐标
					Confidence:         max_vals,           //相似度
					Clickrangevalue:    clickrangevalue,    //随机点击范围
				})
			if len(results) ==0 {
				return nil
			}
			//绘图->原图+模板图坐标进行准确绘图 top_left向右偏移模板图size
			// color for the rect when faces detected检测到脸部时画框的颜色
			blue := color.RGBA{0, 0, 255, 0}
			//fmt.Println("随机点击范围:",clickrangevalue)
			//对像素进行标记点,填满位置point ,对识别后的模板进行涂抹
			pts :=[][]image.Point{{//正确涂抹方式
				image.Pt(rectangle[0][0],rectangle[0][1]),//1.先绘制topleft 左上角
				image.Pt(rectangle[1][0],rectangle[1][1]),//2.在绘制topleft+h 左下角
				image.Pt(rectangle[2][0],rectangle[2][1]),//3.topleft+w+h 右下角
				image.Pt(rectangle[3][0],rectangle[3][1]),//4 topleft+w  右上角
			},}
			//执行匹配到以后做图像填充
			gocv.FillPoly(&i_gray_img,pts,blue)
			//template_rectangle := image.Pt(top_left.X,top_left.Y)
			//gocv.Rectangle(&i_gray_img,image.Rectangle{template_rectangle,max_loc}, blue, 2)
			//win :=gocv.NewWindow("123")
			//win.ResizeWindow(1138, 640)
			//win.IMShow(i_gray_img)
			//gocv.WaitKey(0)
			//defer win.Close()
		}
	}
	return results

}