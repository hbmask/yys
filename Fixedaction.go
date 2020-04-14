package main

import (
	"strconv"
	"time"
	"yys/data"
	"yys/flagpiex"
	"yys/yys_find_img"
)

//次数打满退出
func (f *TFMain) ShiShiCiShu()int{
	zhixing_num,_ :=strconv.Atoi(f.EditCiShu.Text())
	for i:=0;i<zhixing_num;i++{
		zhixing_num=zhixing_num-1
		zhixing_string :=strconv.Itoa(zhixing_num)
		f.EditCiShu.SetText(zhixing_string)
		return zhixing_num
	}
	return 0
}

//御魂觉醒次数打满退出
func (f *TFMain) YuHunJueXingShiShiCiShu()int{
	zhixing_num,_ :=strconv.Atoi(f.EditCiShu.Text())
	for i:=0;i<zhixing_num;i++{
		zhixing_num=zhixing_num-1
		zhixing_string :=strconv.Itoa(zhixing_num)
		f.EditCiShu.SetText(zhixing_string)
		return zhixing_num
	}
	return 0
}

//关闭buff 定时器
func (f *TFMain) OffYuHunTimer( ) {
	fp :=flagpiex.FLagPiex{}
	//新建计时器，120秒以后触发，go触发计时器的方法比较特别，就是在计时器的channel中发送值
	tick :=time.NewTicker( 60 * time.Second)
	for {
		select {
		//此处在等待channel中的信号，因此执行此段代码时会阻塞120秒
		case <-tick.C:
			if fp.FlagTingYuan(){
				f.YuHunOffBuffJianCha() //执行我们想要的操作
			}
		}
	}
}

//点击大舅妈
func (f *TFMain) DianJiDaJiuMa(){
	r:=yys_find_img.Result{}
	if f.CheckBoxDaJiuMa.Checked(){
		YuHunDaianJiDaJiuMa_click:=r.Recognition(data.YuHunDaianJiDaJiuMa_click,0.9)
		if YuHunDaianJiDaJiuMa_click!=nil {
			f.Dj_click_imgpy(YuHunDaianJiDaJiuMa_click,1,100,"标记->大舅妈")
			f.ClickDaJiuMaFlag =true//点击成功
			time.Sleep(time.Second)
		}
	}
}
//点击稻草人
func (f *TFMain) DianJiDaoCaoRen(){
	if f.CheckBoxCaoRen.Checked(){

	}
}
//觉醒Buff开检擦
func (f *TFMain) JueXingOnBuffJianCha(){
	r:=yys_find_img.Result{}
	fp:=flagpiex.FLagPiex{}
	if f.CheckBoxGuanJueXing.Checked(){
		JiaCeng:=r.Recognition(data.JiaCeng,0.9)
		if JiaCeng!=nil {
			f.Dj_click(JiaCeng,"打开加成界面")
			time.Sleep(time.Millisecond*500)
			if fp.FlagJueXingBUffGold(){
				f.JuXingBuffFlag =true
				f.DJ_Click_Range(317,489,600,61,"觉醒开启状态")
				return
			}
			if fp.FlagJueXingBUffRead(){
				f.DJ_Click_Range(700,139,20,6,"启用开觉醒buff")
				f.JuXingBuffFlag =true
				f.DJ_Click_Range(317,489,600,61,"退出buff加成界面")
			}
		}
	}
}
//御魂Buff开检擦
func (f *TFMain) YuHunOnBuffJianCha(){
	r:=yys_find_img.Result{}
	fp:=flagpiex.FLagPiex{}
	if f.CheckBoxGuanYuHun.Checked(){
		JiaCeng:=r.Recognition(data.JiaCeng,0.9)
		if JiaCeng!=nil {
			f.Dj_click(JiaCeng,"打开加层界面")
			time.Sleep(time.Millisecond*500)
			if fp.FlagYuHunBuffGold(){
				//f.DJ_Click_Range(317,489,600,61,"御魂buff已打开")
				f.YuHunBuffFlag =true
				f.DJ_Click_Range(317,489,600,61,"退出buff加成界面")
				return
			}
			if fp.FlagYuHunBuffRed(){
				f.DJ_Click_Range(701,199,20,6,"打开御魂buff")
				f.YuHunBuffFlag =true
				f.DJ_Click_Range(317,489,600,61,"退出buff加成界面")
				return
			}
		}
	}
}
//御魂Buff关检擦
func (f *TFMain) YuHunOffBuffJianCha(){
	r:=yys_find_img.Result{}
	fp:=flagpiex.FLagPiex{}
	if f.CheckBoxGuanYuHun.Checked(){
		JiaCeng:=r.Recognition(data.JiaCeng,0.9)
		if JiaCeng!=nil {
			f.Dj_click(JiaCeng,"打开加成界面")
			time.Sleep(time.Millisecond*500)
			if fp.FlagYuHunBuffGold(){
				f.DJ_Click_Range(701,199,20,6,"关闭御魂buff")
				f.YuHunBuffFlag =false
				f.DJ_Click_Range(317,489,600,61,"退出加成界面")
				f.Stops()
				return
			}
		}
	}
}
//庭院御魂Buff关闭
func (f *TFMain) YuHunTingYuanOffBuffJianCha(){
	r:=yys_find_img.Result{}
	fp:=flagpiex.FLagPiex{}
	if f.CheckBoxGuanYuHun.Checked(){
		TingYuanJiaCeng:=r.Recognition(data.TingYuanJiaCeng,0.9)
		if TingYuanJiaCeng!=nil {
			f.Dj_click(TingYuanJiaCeng,"打开加成")
			time.Sleep(time.Millisecond*500)
			if fp.FlagYuHunBuffGold(){
				f.DJ_Click_Range(701,199,20,6,"关闭御魂buff")
				f.YuHunBuffFlag =false
				f.DJ_Click_Range(317,489,600,61,"退出加成界面")
				f.Stops()
				return
			}
		}
	}
}
//状态检查 四个选项队长
func (f *TFMain) Zhuangtai_all(){
	f.CheckBoxTuiChu.SetChecked(true)//胜利退出
	f.CheckBoxZhunBei.SetChecked(true)//开局准备
	f.CheckBoxFangZhu.SetChecked(true)//房主
	f.CheckBoxXuanShang.SetChecked(true)//接悬赏
}
//状态检查 四个选项
func (f *TFMain) Zhuangtai_3(){
	f.CheckBoxTuiChu.SetChecked(true)//胜利退出
	f.CheckBoxZhunBei.SetChecked(true)//开局准备
	f.CheckBoxFangZhu.SetChecked(false)//房主
	f.CheckBoxXuanShang.SetChecked(true)//接悬赏
}
//隐藏所有执行按钮
func (f *TFMain) Off_All_Buttone() {
	f.ButtonGouLiangZhiXing.SetEnabled(false)
	f.ButtonQiTaZhiXing.SetEnabled(false)
	f.ButtonYaoQiZhiXing.SetEnabled(false)
	f.ButtonYuhunZhixing.SetEnabled(false)
}
//显示所有执行按钮
func (f *TFMain) On_All_Buttone() {
	f.ButtonGouLiangZhiXing.SetEnabled(true)
	f.ButtonQiTaZhiXing.SetEnabled(true)
	f.ButtonYaoQiZhiXing.SetEnabled(true)
	f.ButtonYuhunZhixing.SetEnabled(true)
	f.ButtonGouLiangZhiXing.SetCaption("执行")
	f.ButtonQiTaZhiXing.SetCaption("执行")
	f.ButtonYaoQiZhiXing.SetCaption("执行")
	f.ButtonYuhunZhixing.SetCaption("执行")
	//f.Refresh()
}
//悬赏
func (f *TFMain) XuanShang(){
	r:=yys_find_img.Result{}
	fp :=flagpiex.FLagPiex{}
	if fp.FlagXuanShangDingWei(){
		xuanshangdata:= []string{data.XuanShangTiLi,data.XuanShangGouYu}
		rdata :=r.RecognitionsBuTongTuAn(xuanshangdata,0.9)
		if len(rdata)==0{
			f.DJ_Click_Range(820,455,30,12,"拒绝悬赏")
			return
		}else {
			f.DJ_Click_Range(821,368,30,12,"接受悬赏")
			return
		}
		//for i,_:=range rdata{
		//	if rdata[i]!=nil{
		//		f.DJ_Click_Range(821,368,30,12,"接受悬赏")
		//		break
		//	}else {
		//		f.DJ_Click_Range(820,455,30,12,"拒绝悬赏")
		//		break
		//	}
		//}


	}
}
//准备
func (f *TFMain) ZhanDouZhunBei(){
	r:=yys_find_img.Result{}
	//准备->查看标记是否存在
	Kaijuzhunbei_flag:=r.Recognition(data.Kaijuzhunbei_flag,0.85)
	if Kaijuzhunbei_flag!=nil {
		//准备->标记->点击准备
		Kaijuzhunbei_click:=r.Recognition(data.Kaijuzhunbei_click,0.85)
		if Kaijuzhunbei_click!=nil {
			f.Dj_click(Kaijuzhunbei_click,"准备战斗")
			return
		}
	}
}
//战斗退出
func (f *TFMain) ZhanDouTuiChu(){
	r:=yys_find_img.Result{}
	fp:=flagpiex.FLagPiex{}
	if fp.FlagShengLi()||fp.FlagJingSuMiWenShengLiTuiChu(){
		f.DJ_Click_TuiChu()
		f.YYSLos("退出战斗")
		f.ClickDaJiuMaFlag =false//战斗退出重置
		f.ClickDaoCaoRenFlag =false//战斗退出重置
		f.FlagNum=false//计数判定
		//time.Sleep(time.Millisecond*500)
		return
	}
	if fp.FlagShengLiBaoXiang(){
		f.DJ_Click_TuiChu()
		f.ClickDaJiuMaFlag =false//战斗退出重置
		f.ClickDaoCaoRenFlag =false//战斗退出重置
		f.FlagNum=false//计数判定
		//time.Sleep(time.Millisecond*500)
		return
	}
	if fp.FlagTuiChuTanChiGui(){
		f.DJ_Click_TuiChu()
		f.ClickDaJiuMaFlag =false//战斗退出重置
		f.ClickDaoCaoRenFlag =false//战斗退出重置
		f.FlagNum=false//计数判定
		//time.Sleep(time.Millisecond*500)
		return
	}
	if fp.FlagShiBai(){
		//失败->点击鼓面
		End_shibai_gu_click:=r.Recognition(data.End_shibai_gu_click,0.89)
		if End_shibai_gu_click!=nil {
			f.Dj_click(End_shibai_gu_click,"太丢人了")
			time.Sleep(time.Millisecond*500)
			return
		}
		//失败->点击四点
		End_shibai_sidian_click:=r.Recognition(data.End_shibai_sidian_click,0.89)
		if End_shibai_sidian_click!=nil {
			f.Dj_click(End_shibai_sidian_click,"真的丢人")
			time.Sleep(time.Millisecond*500)
			return
		}
		//失败->点击字体->自己
		End_shibai_ziji_click:=r.Recognition(data.End_shibai_ziji_click,0.89)
		if End_shibai_ziji_click!=nil {
			f.Dj_click(End_shibai_ziji_click,"真丢人")
			time.Sleep(time.Millisecond*500)
			return
		}
	}
	//胜利->点击图案
	End_dianjituan_click:=r.Recognition(data.End_dianjituan_click,0.85)
	if End_dianjituan_click!=nil {
		f.DJ_Click_TuiChu()
		time.Sleep(time.Millisecond*500)
		return
	}
	//胜利->点击屏幕
	End_dianjipingmu_click:=r.Recognition(data.End_dianjipingmu_click,0.85)
	if End_dianjipingmu_click!=nil {
		f.DJ_Click_TuiChu()
		time.Sleep(time.Millisecond*500)
		return
	}
}


