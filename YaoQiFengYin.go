package main

import (
	"time"
	"yys/data"
	"yys/yys_find_img"
)

//妖气封印选择式神
//Result_img_centen  []int//顶点中间坐标
//Result_img_topleft []int//左上角坐标
//Rectangles         [][]int//矩形四角
//Confidence         []float32//相似度
//Clickrangevalue    []int//相似度
func (f *TFMain) YaoQiFengYing_XuanZeShiShen(r yys_find_img.Result,imgsdata string,ShiShenName string){
	//返回一个 包含 坐标 相似度
	rt :=r.Recognition(imgsdata,0.9)
	if rt!=nil{
		f.Dj_click(rt,ShiShenName)
		time.Sleep(time.Millisecond*800)
		ZiDongPiPei :=r.Recognition(data.YaoQiZiDongPiPeiClick,0.9)
		if ZiDongPiPei!=nil{
			f.Dj_click(ZiDongPiPei,"自动匹配")
			time.Sleep(time.Millisecond*2000)
		}

	}else {
		f.mv_mouse_Range(433,267,1,-300,"")
	}
}
