package main

import (
	"time"
	"yys/flagpiex"
)

func (f *TFMain) YuHunOrJueXingFangZhu(i int,fp flagpiex.FLagPiex,GuanYuhunNext_v bool){
		f.XuanShang()
		f.ZhanDouTuiChu()
		//如果没有上锁 手动点击准备
		if fp.FlagZhanDouJieMianZhunBei(){
			if f.YuHunJueXingOnClock ==false{
				f.ZhanDouZhunBei()
				f.YuHunJueXingOnClock =true
				//action.DJ_Click_Range(993,473,70,50)
			}//点击准备
			time.Sleep(time.Millisecond*300)
		}
		if fp.FlagZhanDouJieMianJiaCeng()&&GuanYuhunNext_v==true{//战斗界面->点击加层
			if  f.YuHunBuffFlag ==false{//御魂buff状态
				f.DJ_Click_Range(106,595,26,25,"加层检查")
				for  {
					if fp.FlagYuHunBuffRed(){//红色状态
						f.DJ_Click_Range(701,199,20,6,"开启御魂buff")
						f.YuHunBuffFlag =true
						f.DJ_Click_Range(0,489,600,30,"")
						//time.Sleep(time.Millisecond*500)
						f.StopYuHunNum=0
						f.OffBuff=0
						break
					}
					if fp.FlagYuHunBuffGold(){//金色状态
						//f.DJ_Click_Range(317,489,600,61,"御魂buff已打开")
						f.YuHunBuffFlag =true
						f.DJ_Click_Range(0,489,600,30,"buff已经开启")
						//time.Sleep(time.Millisecond*500)
						//f.DJ_Click_Range(0,489,600,30,"")
						f.StopYuHunNum=0
						f.OffBuff=0
						break
					}
					f.StopYuHunNum++
					if f.StopYuHunNum>=20{
						f.StopYuHunNum=0
						break
					}
					f.StopYuHunNum++
					time.Sleep(time.Millisecond*50)
				}
			}
		}
		//战斗界面
		if fp.FlagZhanDouJieMian(){
			//显示一回木 点大舅妈
			if fp.FlagYuhunJueXingYiHuiMu()&&f.ClickDaJiuMaFlag ==false{
				f.DianJiDaJiuMa()//标记大舅妈
				time.Sleep(time.Millisecond*100)
			}
			//记录副本次数
			if fp.FlagYuhunJueXingYiHuiMu()&&f.FlagNum==false{
				f.OffNumGame=f.YuHunJueXingShiShiCiShu()
				f.OffBuff =0
				f.FlagNum =true
			}
			time.Sleep(time.Millisecond*100)
		}
		//第一次战斗结束邀请队友继续
		if fp.FlagTuiChuYaoQingJiXu(){
			f.DJ_Click_Range(487,313,21,15,"我继续邀请队友")
			time.Sleep(time.Millisecond*700)
			f.DJ_Click_Range(603,366,140,36,"我确定")
		}
		//在 庭院 探索 房间 //60秒没动作关闭御魂buff
		if fp.FlagTingYuan()||fp.FlagTanSuo()||fp.FlagYuHunJueXingFangJian(){
			//if  f.YuHunBuffFlag ==false{//御魂buff状态
			//	f.YuHunOnBuffJianCha() //选择御魂是否打开御魂buff
			//}
			if  f.OffBuff>=90||f.OffNumGame==0{//记录副本次
				f.YuHunTingYuanOffBuffJianCha()
				f.YuHunOffBuffJianCha()
				return
			}
			time.Sleep(time.Millisecond *100)
			f.OffBuff =f.OffBuff+1

			if fp.FlagYuhunJueXingFangJianOnLock(){
				f.YuHunJueXingOnClock =true
			}else{
				f.YuHunJueXingOnClock =false
			}
			if i==2{
				if fp.FlagYuhunJueXingFangJianWeiZhi2()==false&&fp.FlagYuHunJueXingFangJian(){ //是不是2人满了
					f.DJ_Click_Range(1065,564,50,25,"挑战")
				}
					time.Sleep(time.Millisecond*1000)
			}else if i==3{
				if fp.FlagYuhunJueXingFangJianWeiZhi3()==false&&fp.FlagYuHunJueXingFangJian(){ //是不是3人满了
					f.DJ_Click_Range(1065,564,50,25,"挑战")}
					time.Sleep(time.Millisecond*1000)
			}

		}
		//time.Sleep(time.Millisecond*100)

}
