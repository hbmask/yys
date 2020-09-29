package main

import (
	"fmt"
	"math/rand"
	"time"
	"yys/data"
	"yys/flagpiex"
	"yys/yys_find_img"
)

//结界突破
func (f *TFMain) JieJieTuPo(r yys_find_img.Result,fp flagpiex.FLagPiex){
	f.Zhuangtai_3()
	fmt.Println("结界突破 0")
	//选择点击位置
	jjtpnum9_XuanZe :=[][]int{ //判断是否已经攻击
		//x  y   -x ,xr,yr, color
		{403,148,303,80,30,12372953},
		{694,148,594,80,30,12439002},
		{992,148,892,80,30,12439002},
		{393,265,293,80,30,12439002},
		{694,265,594,80,30,12439002},
		{986,265,886,80,30,12439002},
		{395,389,295,80,30,12439002},
		{688,389,588,80,30,12439002},
		{990,389,890,80,30,12439002},
	}
	for{
		if f.StopFlag==false {
			break
		}
		f.XuanShang()
		//战斗界面
		if fp.FlagZhanDouJieMian(){
			time.Sleep(time.Millisecond*400)
			continue
		}
		//战斗退出
		f.ZhanDouTuiChu()
		//探索场景
		if fp.FlagTanSuo(){
			f.DJ_Click_Range(254,572,46,30,"探索->结界突破")
			time.Sleep(time.Millisecond*100)
		}
		//如果在突破界面,继续下面操作
		if fp.FlagJieJieTuPoJieMian(){
			//自动上锁
			if fp.FlagJieJieTuPoOnLock()==true {
				rd :=rand.Intn(1)
				if rd==0{
					f.DJ_Click_Range(736,533,1,1,"结界突破->上锁0")
				}else{
					f.DJ_Click_Range(746,533,1,1,"结界突破->上锁1")
				}
			}
			//当没有票时候停止
			Jiejietupo_1_end_flag :=r.Recognition(data.Jiejietupo_1_end_flag,0.95)
			if Jiejietupo_1_end_flag!=nil {
				f.Stops()
				break
			}
			for i,_ :=range jjtpnum9_XuanZe{
				if f.StopFlag==false {
					break
				}
				index :=i
				//x  y   -x ,xr,yr, color
				x :=jjtpnum9_XuanZe[index][0]
				y :=jjtpnum9_XuanZe[index][1]
				x_XuanZe := jjtpnum9_XuanZe[index][2]
				xrange :=jjtpnum9_XuanZe[index][3]
				yrange :=jjtpnum9_XuanZe[index][4]
				coloerrfe := jjtpnum9_XuanZe[index][5]

				if r.Find_Pixels_jjtp9num(x,y, coloerrfe){
					f.DJ_Click_Range(x_XuanZe,y,xrange,yrange,"结界突破->选择")
					time.Sleep(time.Millisecond*750)
					Jiejietupo_2_jingong_click :=r.Recognition(data.Jiejietupo_2_jingong_click,0.9)
					if Jiejietupo_2_jingong_click!=nil {
						f.Dj_click(Jiejietupo_2_jingong_click,"结界突破->进攻")
						time.Sleep(time.Second*2)
						break
					}
				}else {
					fmt.Println("跳过无效的",jjtpnum9_XuanZe[index])
				}
				if i ==8{
					//fmt.Println(fp.FlagJieJieTuPoLenQue())
					if fp.FlagJieJieTuPoLenQue() ==true{ //如果没有冷却执行

						f.DJ_Click_Range(869,519,30,25,"结界突破->刷新")
						time.Sleep(time.Second)
						f.DJ_Click_Range(603,367,130,30,"结界突破->确定")
						time.Sleep(time.Second)
					}
					continue
				}
			}
			time.Sleep(time.Millisecond*100)
		}
	}
}
//业原火痴
func (f *TFMain) YeYuanHuoChi(r yys_find_img.Result,fp flagpiex.FLagPiex){
	f.Zhuangtai_3()
	fmt.Println("业原火痴 1")
	for {
		if f.StopFlag==false {
			break
		}
		f.XuanShang()
		if fp.FlagZhanDouJieMian(){
			time.Sleep(time.Millisecond*1000)
			continue
		}
		f.ZhanDouTuiChu()//退出战斗
		//业原火界面
		if fp.FlagYeYuanHuoJiemian(){//业原火界面
			//御魂->业原火>选择三层
			if fp.FlagYeYuanHuoXuanZeSanCeng()==false {//御魂->业原火>选择三层
				Yuhun_2_1_chijuan_click := r.Recognition(data.Yuhun_2_1_chijuan_click, 0.9)
				if Yuhun_2_1_chijuan_click != nil {
					f.Dj_click(Yuhun_2_1_chijuan_click,"选择三层")
					time.Sleep(time.Second * 1)
					continue
				}
			}
			//御魂->业原火->上锁->挑战
			if fp.FlagYeYuanHuoOnClock(){//御魂->业原火->上锁->挑战
				Yuhun_4_suo_tiaozhan_click:=r.Recognition(data.Yuhun_4_suo_tiaozhan_click,0.9)
				if Yuhun_4_suo_tiaozhan_click!=nil {
					if f.ShiShiCiShu() ==0||f.TiaoZhanJiShuoff>=3{ //次数达到上限退出
						f.YYSLos("次数达到上限退出")
						f.Stops()
					}
					f.Dj_click(Yuhun_4_suo_tiaozhan_click,"上锁->挑战")
					f.TiaoZhanJiShuoff +=1
					time.Sleep(time.Second*1)
					continue
				}
			}
			//御魂->业原火->上锁
			Yuhun_3_meisuo_click:=r.Recognition(data.Yuhun_3_meisuo_click,0.9)
			if Yuhun_3_meisuo_click!=nil {//御魂->业原火->上锁
				f.Dj_click(Yuhun_3_meisuo_click,"上锁")
				time.Sleep(time.Second*1)
				continue
			}

		}
		//御魂->业原火
		Yuhun_1_yeyuanhuo_clik:=r.Recognition(data.Yuhun_1_yeyuanhuo_clik,0.9)
		if Yuhun_1_yeyuanhuo_clik!=nil {//御魂->业原火
			f.Dj_click(Yuhun_1_yeyuanhuo_clik,"御魂->业原火")
			time.Sleep(time.Second*1)
			continue
		}
		//探索->御魂
		Yuhun_0_click :=r.Recognition(data.Yuhun_0_click,0.9)
		if Yuhun_0_click!=nil { //探索->御魂
			f.Dj_click(Yuhun_0_click,"探索->御魂")
			time.Sleep(time.Second*1)
			continue
		}
	}
}
//自动斗技
func (f *TFMain) ZiDongDouJi(r yys_find_img.Result,fp flagpiex.FLagPiex){
	f.Zhuangtai_3()
	fmt.Println("自动斗技 2")
	f.XuanShang()
	for {
		if f.StopFlag == false {
			break
		}
		f.ZhanDouZhunBei()
		f.ZhanDouTuiChu()
		if fp.FlagDouJiZhanDouZhong()&&f.FlagNum==false{//战斗时选择自动
			//time.Sleep(time.Second*4)
			f.DJ_Click_Range(52,576,6,6,"自动战斗")
			f.FlagNum =true
			f.FlagDouJiSZ=false
		}
		if fp.FlagDouJiJieMian(){//斗技界面
			f.DJ_Click_Range(1049,555,40,40,"斗技挑战")
			f.FlagNum =false
			f.FlagDouJiSZ=false
		}
		if fp.FlagDouJi1700ZiDongShangZHen()&&f.FlagDouJiSZ==false{//斗技1700分 自动上阵
			f.DJ_Click_Range(52,141,5,5,"斗技自动上阵")
			f.FlagNum =false
			f.FlagDouJiSZ=true
		}
		if fp.FlagDouJiBaDeTouChou(){//拔得头筹
			f.FlagNum =false
			f.FlagDouJiSZ=false
			f.DJ_Click_TuiChu()
		}

		if time.Now().Hour()==14{
			f.Stops()
			f.YYSLos("2点咯..")
		}
		if fp.FlagDouJiShenJi(){
			f.DJ_Click_Range(918,473,70,40,"斗技挑战")
		}
	}
}
//自动御灵
func (f *TFMain) ZiDongYuLin(r yys_find_img.Result,fp flagpiex.FLagPiex){
	f.Zhuangtai_3()
	fmt.Println("自动御灵 3")
	for {
		if f.StopFlag == false {
			break
		}
		f.XuanShang()
		//战斗界面
		if fp.FlagZhanDouJieMian() {//战斗界面
			time.Sleep(time.Millisecond * 100)
			continue
		}
		if fp.FlagYuLingTiaoZhanJieMian(){//战斗界面战斗准备
			if fp.FlagYuLingTiaoZhanJieMianSanCeng()!=true {
				f.DJ_Click_Range(240,472,100,50,"选择三层")
				time.Sleep(time.Millisecond*100)
			}
			if fp.FlagYuLingTiaoZhanJieShangSuo()!=true{
				rand.Seed(time.Now().UnixNano())
				i :=rand.Intn(1)
				if i==0{
					f.DJ_Click_Range(495,516,1,1,"上锁1")
					time.Sleep(time.Millisecond*100)
				}else {
					f.DJ_Click_Range(519,516,1,1,"上锁2")
					time.Sleep(time.Millisecond*100)
				}
			}else {

				//在挑战记录执行副本次数
				if f.ShiShiCiShu() ==0 ||f.TiaoZhanJiShuoff >=3{//次数达到上限退出
					f.YYSLos("次数达到上限退出")
					f.Stops()
				}
				f.DJ_Click_Range(995,541,55,47,"挑战")
				f.TiaoZhanJiShuoff +=1
				time.Sleep(time.Millisecond*1000)
			}
		}
		//战斗退出
		f.ZhanDouTuiChu()
	}
}
//寮突破
func (f *TFMain) LiaoTuPo(r yys_find_img.Result,fp flagpiex.FLagPiex){
	//寮突破选择位置
	LiaoTuPo_XuanZe :=[][]int{ //判断是否已经攻击
		//x  y   -x ,xr,yr, color
		{643,137,542,80,30,12439002},
		{947,137,840,80,30,12372954},
		{643,258,542,80,30,12439002},
		{947,258,840,80,30,12372954},
		{643,379,542,80,30,12439002},
		{947,379,840,80,30,12372954},
		{643,499,542,80,30,12439002},
		{947,499,840,80,30,12372954},
	}

	f.Zhuangtai_3()
	fmt.Println("寮突破 4")
	for{
		if f.StopFlag==false {
			break
		}
		f.XuanShang()
		//战斗中..等待
		if fp.FlagZhanDouJieMian(){
			time.Sleep(time.Millisecond*400)
			continue
		}
		f.ZhanDouTuiChu()
		//探索场景
		if fp.FlagTanSuo(){
			f.DJ_Click_Range(254,572,46,30,"探索->结界突破")
			time.Sleep(time.Second*1)
			f.DJ_Click_Range(1078,329,26,50,"结界突破->寮突破")
			time.Sleep(time.Second*1)
		}
		//结界突破->寮突破->记录锚点
		//Jiejietupo_2_liaotupo_ji_flag:=r.Recognition(data.Jiejietupo_2_liaotupo_ji_flag,0.9)
		//if Jiejietupo_2_liaotupo_ji_flag!=nil {
		if fp.Flag_LiaoTuPo_JieMian(){
			if fp.Flag_LiaoTuPo_Po(){
				//探索->结界突破->寮突破->选择->进攻->如果没有机会等待.
				Liaotupo_flag :=r.Recognition(data.Liaotupo_flag,0.9)
				if Liaotupo_flag!=nil {
					//if fp.Flag_LiaoTuPo_JinGongCiShu(){
					if fp.Flag_LiaoTuPo_JieMian(){
						time.Sleep(time.Millisecond*300)
						//f.DJ_Click_Range(44,24,1,30,"寮突破->探索5分钟")
						f.DJ_Click_Range(1066,113,10,4,"寮突破->探索15破")
						time.Sleep(time.Second*900)
						fmt.Println("等待恢复中.....")
					}
					continue
				}
			}else {
				//探索->结界突破->寮突破->选择->进攻->如果没有机会等待.
				Liaotupo_flag :=r.Recognition(data.Liaotupo_flag,0.9)
				if Liaotupo_flag!=nil {
					//if fp.Flag_LiaoTuPo_JinGongCiShu(){
					if fp.Flag_LiaoTuPo_JieMian(){
						time.Sleep(time.Millisecond*300)
						//f.DJ_Click_Range(44,24,1,30,"寮突破->探索5分钟")
						f.DJ_Click_Range(1066,113,10,4,"寮突破->探索15分")
						time.Sleep(time.Second*900)
						fmt.Println("等待恢复中.....")
					}
					continue
				}
			}
			//自动上锁
			if fp.FlagLiaoTuPoOnLock() {
				rd :=rand.Intn(1)
				if rd==0{
					f.DJ_Click_Range(177,550,1,1,"寮突破->上锁0")
				}else{
					f.DJ_Click_Range(200,550,1,1,"寮突破->上锁1")
				}
			}
			for i,_ :=range LiaoTuPo_XuanZe {
				if f.StopFlag==false {
					break
				}
				//x  y   -x ,xr,yr, color
				index :=i
				x := LiaoTuPo_XuanZe[index][0]
				y := LiaoTuPo_XuanZe[index][1]
				xrange := LiaoTuPo_XuanZe[index][3]
				yrange := LiaoTuPo_XuanZe[index][4]
				x_xuanze := LiaoTuPo_XuanZe[index][2]
				coloerrfe := LiaoTuPo_XuanZe[index][5]

				if r.Find_Pixels_jjtp9num(x, y, coloerrfe){
					f.DJ_Click_Range(x_xuanze,y,xrange,yrange,"寮突破->选择")
					time.Sleep(time.Millisecond*1000)
					Jiejietupo_2_jingong_click :=r.Recognition(data.Jiejietupo_2_jingong_click,0.9)
					if Jiejietupo_2_jingong_click!=nil {
						f.Dj_click(Jiejietupo_2_jingong_click,"寮突破->进攻")
						time.Sleep(time.Second*2)
					}
					break
				}else {
					fmt.Println("跳过无效的", LiaoTuPo_XuanZe[index])
					if i ==7{
						f.Stops()
					}
				}

			}
		}
	}
}
//全自动
func (f *TFMain) QuanZiDong(r yys_find_img.Result,fp flagpiex.FLagPiex){}
//召唤厕纸
func (f *TFMain) ZhaoHuanCeZhi(r yys_find_img.Result,fp flagpiex.FLagPiex){
	f.Zhuangtai_all()
	fmt.Println("召唤厕纸 6")
	for {
		if f.StopFlag==false {
			break
		}
		f.XuanShang()
		Cezhi_zaohuan_click :=r.Recognition(data.Cezhi_zaohuan_click,0.9)
		if Cezhi_zaohuan_click!=nil {
			f.Dj_click(Cezhi_zaohuan_click,"再次召唤厕纸")
			time.Sleep(time.Second*1)
		}
		//Cezhi_click :=r.Recognitions(data.Cezhi_click,0.9)
		//if Cezhi_click!=nil {
		//	f.Dj_clicks(Cezhi_click,"召唤祖安")
		//	//time.Sleep(time.Second*1)
		//}

	}
}
//竞速秘闻挑战
func (f *TFMain) JinSuMiWenTiaoZhan(r yys_find_img.Result,fp flagpiex.FLagPiex){
	f.Zhuangtai_3()
	fmt.Println("竞速秘闻挑战 7")
	for{
		if f.StopFlag==false {
			break
		}
		f.XuanShang()
		//战斗准备界面
		if fp.FlagZhanDouJieMianZhunBei(){
			//自动上锁
			f.ZhanDouZhunBei()
			time.Sleep(time.Millisecond*500)
		}
		//战斗界面
		if fp.FlagZhanDouJieMian(){
			time.Sleep(time.Millisecond*100)
			continue
		}

		//活动点击挑战
		if f.StopFlag==false {
			break
		}
		f.XuanShang()
		//if r.Recognition(data.HD,0.99)!=nil {
		//	f.Stops()
		//	break
		//}
		if fp.FlagHuDong_TZ_SUO(){
			f.DJ_Click_Range(828,455,1,1,"活动->上锁")
			time.Sleep(time.Millisecond*500)
		}
		if fp.FlagHuDong_TZ(){
			f.DJ_Click_Range(1040,550,20,30,"活动->挑战")
			time.Sleep(time.Millisecond*500)
		}
		//竞速秘闻挑战
		if fp.FlagJingSuMiWenTiaoZhan(){
			f.DJ_Click_Range(1033,498,10,60,"竞速秘闻->挑战")
			time.Sleep(time.Millisecond*500)
		}
		//战斗退出
		f.ZhanDouTuiChu()
	}

}
//结界卡合成
func (f *TFMain) JieJieKaHeCheng(r yys_find_img.Result,fp flagpiex.FLagPiex){
	for{
		fmt.Println()
		if f.StopFlag==false {
			break
		}
		f.XuanShang()
		if fp.FlagJieJieKa_JiXuTianJia(){
			f.DJ_Click_Range(923,523,30,12,"结界卡->继续添加")
			time.Sleep(time.Millisecond*500)
			f.DJ_Click_Range(694,514,130,30,"结界卡->开始合成")
			time.Sleep(time.Millisecond*700)
			continue
		}
		f.Stops()
	}
}

