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
	jjtpnum9 :=[][]int{//选择进攻点击位置
		{340,136,80,30},//1
		{620,136,80,30},//2
		{900,136,80,30},//3
		{340,246,80,30},//4
		{620,246,80,30},//5
		{900,246,80,30},//6
		{340,356,80,30},//7
		{620,356,80,30},//8
		{900,356,80,30},//9
	}
	jjtpnum9_FuZhu :=[][]int{//判断是否已经攻击
		{440,136,12898778},
		{720,136,12898778},
		{1000,136,12898778},
		{440,246,12898778},
		{720,246,12898778},
		{1000,246,12898778},
		{440,356,12898778},
		{720,356,12898778},
		{1000,356,12898778},
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
					f.DJ_Click_Range(904,538,1,1,"结界突破->上锁0")
				}else{
					f.DJ_Click_Range(930,537,1,1,"结界突破->上锁1")
				}
			}
			Jiejietupo_1_end_flag :=r.Recognition(data.Jiejietupo_1_end_flag,0.95)
			if Jiejietupo_1_end_flag!=nil {
				f.Stops()
				break
			}
			for i,_ :=range jjtpnum9{
				if f.StopFlag==false {
					break
				}
				index :=i
				x :=jjtpnum9[index][0]
				y :=jjtpnum9[index][1]
				xrange :=jjtpnum9[index][2]
				yrange :=jjtpnum9[index][3]

				x_FuZhu :=jjtpnum9_FuZhu[index][0]
				y_FuZhu :=jjtpnum9_FuZhu[index][1]
				coloerrfe :=jjtpnum9_FuZhu[index][2]

				if r.Find_Pixels_jjtp9num(x_FuZhu,y_FuZhu, coloerrfe){
					f.DJ_Click_Range(x,y,xrange,yrange,"结界突破->选择")
					time.Sleep(time.Millisecond*600)
					Jiejietupo_2_jingong_click :=r.Recognition(data.Jiejietupo_2_jingong_click,0.9)
					if Jiejietupo_2_jingong_click!=nil {
						f.Dj_click(Jiejietupo_2_jingong_click,"结界突破->进攻")
						time.Sleep(time.Second*2)
						//fmt.Println("True:",jjtpnum9,i)
						break
					}
				}else {
					fmt.Println("跳过无效的",jjtpnum9[index])
				}
				if i ==8{
					//fmt.Println(fp.FlagJieJieTuPoLenQue())
					if fp.FlagJieJieTuPoLenQue() ==true{ //如果没有冷却执行

						f.DJ_Click_Range(1057,169,30,25,"结界突破->刷新")
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
			f.DJ_Click_Range(918,473,70,40,"斗技挑战")
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
	LiaoTuPo_num8 :=[][]int{//选择进攻点击位置
		{585,151,80,30},
		{876,151,80,30},
		{585,251,80,30},
		{876,251,80,30},
		{585,351,80,30},
		{876,351,80,30},
		{585,451,80,30},
		{876,451,80,30},
	}
	//寮突破选择位置
	LiaoTuPo_FuZhu :=[][]int{//判断是否已经攻击
		{685,151,12898778},
		{976,151,12898778},
		{685,251,12898778},
		{976,251,12898778},
		{685,351,12898778},
		{976,351,12898778},
		{685,451,12898778},
		{976,451,12898778},
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
			f.DJ_Click_Range(94,370,26,100,"结界突破->寮突破")
		}
		//结界突破->寮突破->记录锚点
		//Jiejietupo_2_liaotupo_ji_flag:=r.Recognition(data.Jiejietupo_2_liaotupo_ji_flag,0.9)
		//if Jiejietupo_2_liaotupo_ji_flag!=nil {
		if fp.Flag_LiaoTuPo_JieMian(){
			if fp.Flag_LiaoTuPo_Po(){
				//探索->结界突破->寮突破->选择->进攻->如果没有机会等待.
				Liaotupo_flag :=r.Recognition(data.Liaotupo_flag2,0.9)
				if Liaotupo_flag!=nil {
					//if fp.Flag_LiaoTuPo_JinGongCiShu(){
					if fp.Flag_LiaoTuPo_JieMian(){
						time.Sleep(time.Millisecond*300)
						//f.DJ_Click_Range(44,24,1,30,"寮突破->探索5分钟")
						f.DJ_Click_Range(32,40,10,4,"寮突破->探索15破")
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
						f.DJ_Click_Range(32,40,10,4,"寮突破->探索15分")
						time.Sleep(time.Second*900)
						fmt.Println("等待恢复中.....")
					}
					continue
				}
			}
			//自动上锁
			if fp.FlagLiaoTuPoOnLock()==true {
				rd :=rand.Intn(1)
				if rd==0{
					f.DJ_Click_Range(238,540,1,1,"寮突破->上锁0")
				}else{
					f.DJ_Click_Range(264,540,1,1,"寮突破->上锁1")
				}
			}
			for i,_ :=range LiaoTuPo_num8{
				if f.StopFlag==false {
					break
				}
				index :=i
				x :=LiaoTuPo_num8[index][0]
				y :=LiaoTuPo_num8[index][1]
				xrange :=LiaoTuPo_num8[index][2]
				yrange :=LiaoTuPo_num8[index][3]

				x_FuZhu :=LiaoTuPo_FuZhu[index][0]
				y_FuZhu :=LiaoTuPo_FuZhu[index][1]
				coloerrfe :=LiaoTuPo_FuZhu[index][2]

				if r.Find_Pixels_jjtp9num(x_FuZhu,y_FuZhu, coloerrfe){
					f.DJ_Click_Range(x,y,xrange,yrange,"寮突破->选择")
					time.Sleep(time.Millisecond*1000)
					Jiejietupo_2_jingong_click :=r.Recognition(data.Jiejietupo_2_jingong_click,0.9)
					if Jiejietupo_2_jingong_click!=nil {
						f.Dj_click(Jiejietupo_2_jingong_click,"寮突破->进攻")
						time.Sleep(time.Second*2)
					}
					break
				}else {
					fmt.Println("跳过无效的",LiaoTuPo_num8[index])
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
		//竞速秘闻挑战
		if fp.FlagJingSuMiWenTiaoZhan(){
			f.DJ_Click_Range(990,481,60,60,"竞速秘闻->挑战")
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

