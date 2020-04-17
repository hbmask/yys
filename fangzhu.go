package main

import (
	"fmt"
	"time"
	"yys/flagpiex"
)

func (f TFMain)YuHunOrJueXingFangZHu(i int){
	fp :=flagpiex.FLagPiex{}
	f.StopFlag=true
	for {
		if f.StopFlag == false {
			break
		}
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
		//战斗界面
		if fp.FlagZhanDouJieMian(){
			//显示一回木
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
			continue
		}
		//第一次战斗结束邀请队友继续
		if fp.FlagTuiChuYaoQingJiXu(){
			f.DJ_Click_Range(487,313,21,15,"我继续邀请队友")
			time.Sleep(time.Millisecond*700)
			f.DJ_Click_Range(603,366,140,36,"我确定")
		}
		//在 庭院 探索 房间 //60秒没动作关闭御魂buff
		if fp.FlagTingYuan()||fp.FlagTanSuo()||fp.FlagYuHunJueXingFangJian(){
			if  f.OffBuff>=90||f.OffNumGame==0{//记录副本次
				f.YuHunTingYuanOffBuffJianCha()
				f.YuHunOffBuffJianCha()
				break
			}
			time.Sleep(time.Millisecond *100)
			f.OffBuff =f.OffBuff+1
			fmt.Println(f.OffBuff)

			if fp.FlagYuhunJueXingFangJianOnLock(){
				f.YuHunJueXingOnClock =true
			}else{
				f.YuHunJueXingOnClock =false
			}
			if  f.YuHunBuffFlag ==false{//御魂buff状态
				f.YuHunOnBuffJianCha() //选择御魂是否打开御魂buff
			}
			if i==2{
				if fp.FlagYuhunJueXingFangJianWeiZhi2()==false&&fp.FlagYuHunJueXingFangJian(){ //是不是2人满了
					f.DJ_Click_Range(1065,564,50,25,"挑战")}
					time.Sleep(time.Millisecond*1000)
			}else if i==3{
				if fp.FlagYuhunJueXingFangJianWeiZhi3()==false&&fp.FlagYuHunJueXingFangJian(){ //是不是3人满了
					f.DJ_Click_Range(1065,564,50,25,"挑战")}
					time.Sleep(time.Millisecond*1000)
			}

		}
		//time.Sleep(time.Millisecond*100)
	}
}
