// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
    "expvar"
    "fmt"
    "github.com/lxn/win"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
    "github.com/ying32/govcl/vcl/types/keys"
    "github.com/ying32/govcl/vcl/types/messages"
    win2 "github.com/ying32/govcl/vcl/win"
    "math/rand"
    _ "net/http/pprof"
    "strconv"
    "time"
    "yys/getyyshwnd"
    "yys/data"
    "yys/flagpiex"
    "yys/yys_find_img"
)

//::private::
type TFMainFields struct {
    StopFlag            bool //暂停
    YuHunJueXingOnClock bool //御魂觉醒房间是否上锁
    ClickDaJiuMaFlag    bool //点大舅妈
    ClickDaoCaoRenFlag  bool //点稻草人
    JuXingBuffFlag bool//觉醒buff状态
    YuHunBuffFlag bool //御魂buff状态
    OffBuff int//倒计时 如果没有动作关闭buff计数
    OffNumGame int//记录副本次数如果是0 停止辅助
    FlagNum bool//判断计数是否有效
    TiaoZhanJiShuoff int//当挑战次数满了以后计数,达到计数后自动停止挑战
    HWND win.HWND//窗口句柄
    hotKeyId types.ATOM//热键
    //GLZB bool//狗粮准备
}


func NewTFMainFields( stopflag bool,yuhunjuexingonclock bool,clickdajiuma bool,clickdaocaoren bool)TFMainFields{
    return TFMainFields{StopFlag:stopflag,YuHunJueXingOnClock:yuhunjuexingonclock,ClickDaJiuMaFlag:clickdajiuma,ClickDaoCaoRenFlag:clickdaocaoren,}
}

func init(){
   YYSHWND := getyyshwnd.YYSHWND{}
   hwnd:=YYSHWND.Get_yys_hwnd()
   e:=expvar.NewInt("erhwnd")
   e.Set(int64(hwnd))
}



//御魂觉醒 执行
//打手 0
//房主两人队 1
//房主三人队 2
func (f *TFMain) OnButtonYuhunZhixingClick(sender vcl.IObject) {
    f.ButtonYuhunZhixing.SetCaption("执行中.")
    f.Off_All_Buttone()
    f.CheckBoxGuanYuHun.SetChecked(true)
    r := yys_find_img.Result{}
    fp :=flagpiex.FLagPiex{}
    //fmt.Println(f.ComboBoxYuhun.Text(), f.ComboBoxYuhun.ItemIndex())
    switch f.ComboBoxYuHun.ItemIndex(){
    case 0:
        f.Zhuangtai_3()
        fmt.Println("打手 0")
        go func() {
            f.StopFlag=true
            for {
                if f.StopFlag == false {
                    break
                }
                f.XuanShang()
                //战斗界面
                if fp.FlagZhanDouJieMian(){
                    //fmt.Println("战斗界面")
                    //如果没有上锁 手动点击准备
                    if fp.FlagZhanDouJieMianZhunBei(){
                        if f.YuHunJueXingOnClock ==false{
                            f.ZhanDouZhunBei()
                            f.YuHunJueXingOnClock =true
                        }//点击准备
                        //time.Sleep(time.Millisecond*500)
                    }
                    //在回目一标记大舅妈
                    if fp.FlagYuhunJueXingYiHuiMu()&&f.ClickDaJiuMaFlag ==false  {
                        //fmt.Println("点击->大舅妈")
                        f.DianJiDaJiuMa()//标记大舅妈
                        time.Sleep(time.Millisecond*300)
                    }
                    //在回目一记录执行副本次数
                    if fp.FlagYuhunJueXingYiHuiMu()&&f.FlagNum==false{
                        f.OffNumGame=f.YuHunJueXingShiShiCiShu()
                        f.OffBuff =0//重置关闭御魂时间
                        f.FlagNum =true//已经识别
                    }
                    time.Sleep(time.Millisecond*100)
                    continue
                }
                //在庭院,探索,房间
                if f.OffNumGame==0||fp.FlagTingYuan()||fp.FlagTanSuo()||fp.FlagYuHunJueXingFangJian_DaShou(){
                    //被邀请进组
                    if fp.FlagYuHunZuDuiYaoQing(){
                        H10 :=r.Recognition(data.H10,0.85)
                        if H10!=nil {
                            YuHunChiLun_Click :=r.Recognition(data.YuHunChiLun_Click,0.85)
                            if YuHunChiLun_Click!=nil{ //被邀请进组选择齿轮
                                f.Dj_click(YuHunChiLun_Click,"齿轮进入")
                                //f.DJ_Click_Range(198,212,30,30,"从此轮进组")
                                time.Sleep(time.Millisecond*200)
                                continue
                            }
                            f.DJ_Click_Range(125,233,5,5,"接受魂10邀请")
                            time.Sleep(time.Millisecond*200)
                            continue
                        }
                        H11 :=r.Recognition(data.H11,0.85)
                        if H11!=nil {
                            YuHunChiLun_Click :=r.Recognition(data.YuHunChiLun_Click,0.85)
                            if YuHunChiLun_Click!=nil{ //被邀请进组选择齿轮
                                f.Dj_click(YuHunChiLun_Click,"齿轮进入")
                                //f.DJ_Click_Range(198,212,30,30,"从此轮进组")
                                time.Sleep(time.Millisecond*200)
                                continue
                            }
                            f.DJ_Click_Range(125,233,5,5,"接受魂11邀请")
                            time.Sleep(time.Millisecond*200)
                            continue
                        }
                    }
                    if  f.YuHunBuffFlag ==false{//御魂buff状态
                        f.YuHunOnBuffJianCha() //选择御魂是否打开御魂buff
                    }
                    if fp.FlagYuhunJueXingFangJianOnLock(){//房间上锁
                        f.YuHunJueXingOnClock =true //房间上锁=自动准备
                        f.ClickDaJiuMaFlag=false//组队房间重置
                        f.ClickDaoCaoRenFlag=false//组队房间重置
                        f.FlagNum=false//计数判定
                    }
                    time.Sleep(time.Millisecond*100)
                    if f.OffBuff>=60{//满足条件关闭御魂
                        f.YuHunTingYuanOffBuffJianCha()
                        f.YuHunOffBuffJianCha()
                    }
                    time.Sleep(time.Millisecond *500)
                    f.OffBuff =f.OffBuff+1
                }

                //判断是否在房间
                //if fp.FlagYuHunJueXingFangJian_DaShou(){
                //    //fmt.Println("房间")
                //    if  f.YuHunBuffFlag ==false{//御魂buff状态
                //        f.YuHunOnBuffJianCha() //选择御魂是否打开御魂buff
                //    }
                //    if fp.FlagYuhunJueXingFangJianOnLock(){
                //        f.YuHunJueXingOnClock =true
                //        f.ClickDaJiuMaFlag=false//组队房间重置
                //        f.ClickDaoCaoRenFlag=false//组队房间重置
                //        f.FlagNum=false//计数判定
                //    }
                //    time.Sleep(time.Millisecond*100)
                //}

               f.ZhanDouTuiChu()
               time.Sleep(time.Millisecond*100)
            }
        }()
    case 1:
        f.Zhuangtai_all()
        fmt.Println("房主两人队 1")
        go func() {
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
                    time.Sleep(time.Millisecond*100)
                }
                //战斗界面
                if fp.FlagZhanDouJieMian(){
                    if fp.FlagYuhunJueXingYiHuiMu()&&f.ClickDaJiuMaFlag ==false { //显示一回木
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

                //第一次战斗是否邀请队友
                if fp.FlagTuiChuYaoQingJiXu(){
                    f.DJ_Click_Range(487,313,21,15,"我继续邀请队友")
                    time.Sleep(time.Millisecond*700)
                    f.DJ_Click_Range(603,366,140,36,"我确定")
                }

                //在房间
                if fp.FlagYuHunJueXingFangJian(){
                    if  f.YuHunBuffFlag ==false{//御魂buff状态
                        f.YuHunOnBuffJianCha() //选择御魂打开御魂buff
                    }
                    if fp.FlagYuhunJueXingFangJianOnLock(){
                        f.YuHunJueXingOnClock =true
                    }else{
                        f.YuHunJueXingOnClock =false
                    }
                    if f.OffNumGame==0{//记录副本次
                        f.YuHunTingYuanOffBuffJianCha()
                        f.YuHunOffBuffJianCha()
                        continue
                    }
                    if fp.FlagYuhunJueXingFangJianWeiZhi2()==false{ //是不是2人满了
                        f.DJ_Click_Range(1065,564,50,25,"挑战开始")} //点击挑战
                    time.Sleep(time.Second)
                }
                //在庭院 探索 房间 60秒没动作关闭御魂buff
                if fp.FlagTingYuan()||fp.FlagTanSuo()||fp.FlagYuHunJueXingFangJian(){
                    if  f.OffBuff>=180{
                        f.YuHunTingYuanOffBuffJianCha()
                        f.YuHunOffBuffJianCha()
                    }
                    time.Sleep(time.Millisecond *100)
                    f.OffBuff =f.OffBuff+1
                    fmt.Println(f.OffBuff)

                }
                time.Sleep(time.Millisecond*100)
            }
        }()
    case 2:
        f.Zhuangtai_all()
        fmt.Println("房主三人队 2")
        go func() {
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

                //在不在房间
                if fp.FlagYuHunJueXingFangJian(){
                    if  f.YuHunBuffFlag ==false{//御魂buff状态
                        f.YuHunOnBuffJianCha() //选择御魂是否打开御魂buff
                    }
                    if fp.FlagYuhunJueXingFangJianOnLock(){
                        f.YuHunJueXingOnClock =true
                    }else{
                        f.YuHunJueXingOnClock =false
                    }
                    if f.OffNumGame==0{//记录副本次
                        f.YuHunTingYuanOffBuffJianCha()
                        f.YuHunOffBuffJianCha()
                        continue
                    }
                    if fp.FlagYuhunJueXingFangJianWeiZhi3()==false{ //是不是3人满了
                        f.DJ_Click_Range(1065,564,50,25,"挑战")} //点击挑战
                    time.Sleep(time.Millisecond*1000)
                }
                //在 庭院 探索 房间 //60秒没动作关闭御魂buff
                if f.OffNumGame==0||fp.FlagTingYuan()||fp.FlagTanSuo()||fp.FlagYuHunJueXingFangJian(){
                    if  f.OffBuff>=180{
                        f.YuHunTingYuanOffBuffJianCha()
                        f.YuHunOffBuffJianCha()
                    }
                    time.Sleep(time.Millisecond *100)
                    f.OffBuff =f.OffBuff+1
                    fmt.Println(f.OffBuff)
                }
                time.Sleep(time.Millisecond*100)
            }
        }()
    }
}
//狗粮
func (f *TFMain) OnButtonGouLiangZhiXingClick(sender vcl.IObject) {
    //f.CheckBoxGuanYuHun.SetChecked(true)
    r := yys_find_img.Result{}
    fp :=flagpiex.FLagPiex{}
    f.ButtonGouLiangZhiXing.SetCaption("执行中.")
    f.XuanShang()
    f.Off_All_Buttone()
    //mbgouliangxy :=make([][]int,0,0)
    go func() {
        f.StopFlag=true
        for {
            if f.StopFlag == false {
                    break
                }
            f.XuanShang()
            if fp.FlagTanSuo_GouLiang()||fp.FlagTanSuo_GouLiangZuDuiJieMian(){//&&fp.FlagTanSuo_KunNan28(){//探索界面与狗粮组队界面
                if fp.FlagYuHunZuDuiYaoQing(){//有困难28标志和邀请勾选
                    KunNan28_Flag:=r.Recognition(data.GouLiangKunNan28_Flag,0.85)//少女与面具
                    if KunNan28_Flag!=nil{
                        f.DJ_Click_Range(125,233,5,5,"接受狗粮28邀请")
                    }
                }
                time.Sleep(time.Millisecond*100)
            }
            if fp.FlagZhanDouJieMian(){//战斗界面
                if fp.FlagZhanDouJieMianZhunBei(){//战斗准备界面
                   switch f.ComboBoxGouLiang.ItemIndex() {
                   case 0: //1级N
                       GouLiangQuanBu_Click:=r.Recognition(data.GouLiangQuanBu_Click,0.9)//狗粮->全部
                       if GouLiangQuanBu_Click!=nil{
                           f.Dj_click(GouLiangQuanBu_Click,"全部")
                           time.Sleep(time.Millisecond*500)
                           GouLiangNKa_Click:=r.Recognition(data.GouLiangNKa_Click,0.9)//狗粮N
                           if GouLiangNKa_Click!=nil{
                               f.Dj_click(GouLiangNKa_Click,"选择->N")
                               time.Sleep(time.Millisecond*600)}
                       }
                       GouLiangNKaFlag:=r.Recognition(data.GouLiangNKaFlag,0.9)//狗粮N
                       if GouLiangNKaFlag!=nil{
                           mb:=r.RecognitionsGouLiang_2Man(data.GouliangManJi_Flag,790,420,0.85)//获取更换满级的目标
                           GouLiang1JiN_Click := r.Recognitions(data.GouLiang1JiN_Click, 0.9) //从N卡中找到1级N卡
                           for i,_ :=range mb{
                               f.move_click(mb[i].Result_img_centen, GouLiang1JiN_Click, 0, 90, "更换1级N")
                           }
                       }
                   case 1: //1级白
                       GouLiangQuanBu_Click:=r.Recognition(data.GouLiangQuanBu_Click,0.9)//狗粮->全部
                       if GouLiangQuanBu_Click!=nil{
                           f.Dj_click(GouLiangQuanBu_Click,"全部")
                           time.Sleep(time.Millisecond*500)
                           GouLiangSuCai_Click:=r.Recognition(data.GouLiangSuCai_Click,0.9)//素材
                           if GouLiangSuCai_Click!=nil{
                               f.Dj_click(GouLiangSuCai_Click,"选择->素材")
                               time.Sleep(time.Millisecond*600)}
                       }
                       GouLiangSuCaiFlag:=r.Recognition(data.GouLiangSuCaiFlag,0.9)//狗粮N
                       if GouLiangSuCaiFlag!=nil{
                           mb:=r.RecognitionsGouLiang_2Man(data.GouliangManJi_Flag,790,420,0.85)//获取更换满级的目标
                           GouLiang1JiBai_Click := r.Recognitions(data.GouLiang1JiBai_Click, 0.9) //从素材中找到1级白
                           for i,_ :=range mb{
                               f.move_click(mb[i].Result_img_centen, GouLiang1JiBai_Click, 0, 90, "更换1级白")
                           }
                       }
                   case 2: //1级红
                       GouLiangQuanBu_Click:=r.Recognition(data.GouLiangQuanBu_Click,0.9)//狗粮->全部
                       if GouLiangQuanBu_Click!=nil{
                           f.Dj_click(GouLiangQuanBu_Click,"全部")
                           time.Sleep(time.Millisecond*500)
                           GouLiangSuCai_Click:=r.Recognition(data.GouLiangSuCai_Click,0.9)//狗粮素材
                           if GouLiangSuCai_Click!=nil{
                               f.Dj_click(GouLiangSuCai_Click,"选择->素材")
                               time.Sleep(time.Millisecond*600)}
                       }
                       GouLiangSuCaiFlag:=r.Recognition(data.GouLiangSuCaiFlag,0.9)//狗粮
                       if GouLiangSuCaiFlag!=nil{
                           mb:=r.RecognitionsGouLiang_2Man(data.GouliangManJi_Flag,790,420,0.85)//获取更换满级的目标
                           GouLiang1JiHong_Click := r.Recognitions(data.GouLiang1JiHong_Click, 0.9) //从素材中找到1级白
                           for i,_ :=range mb{
                               f.move_click(mb[i].Result_img_centen, GouLiang1JiHong_Click, 0, 90, "更换1级红")
                           }
                       }
                   case 3: //20级白
                       GouLiangSuCai_Click:=r.Recognition(data.GouLiangSuCai_Click,0.9)//狗粮素材
                       f.Dj_click(GouLiangSuCai_Click,"选择->素材")
                       time.Sleep(time.Second*1)
                       if GouLiangSuCai_Click!=nil{

                       }
                   case 4: //20级N
                       GouLiangNKa_Click:=r.Recognition(data.GouLiangNKa_Click,0.9)//狗粮N
                       f.Dj_click(GouLiangNKa_Click,"选择->N")
                       time.Sleep(time.Second*1)
                       if GouLiangNKa_Click!=nil{

                       }
                   }
                    GouliangManJi_Flag:=r.Recognitions(data.GouliangManJi_Flag,0.85)//获取满级图像
                    if len(GouliangManJi_Flag)==3&&fp.FlagGouLiangDiBan()==false{
                            f.SJ_Click_Range(530,490,10,10,"狗粮满级更换..")
                            time.Sleep(time.Second*2)
                    }
                    zbGouliangManJi_Flag:=r.RecognitionsGouLiang_2Man(data.GouliangManJi_Flag,1100,420,0.85)//获取更换满级的目标
                    if len(zbGouliangManJi_Flag)==1{
                      f.ZhanDouZhunBei()
                      time.Sleep(time.Second*2)
                    }

                }
                time.Sleep(time.Millisecond *100)
            }
            if fp.FlagGouliangFuBenJieMian(){//狗粮副本界面
                GouLiangDuiZhang_Flag:=r.Recognition(data.GouLiangDuiZhang_Flag,0.85)
                if GouLiangDuiZhang_Flag == nil {//标记副本里面的队长是否还在
                    f.DJ_Click_Range(32,51,12,14,"队长已经退出")
                    time.Sleep(time.Millisecond*500)
                    f.DJ_Click_Range(650,350,100,25,"立刻退出")
                }
                time.Sleep(time.Millisecond*500)
            }
            f.ZhanDouTuiChu()
        }
    }()
}
//结界突破
//业原火痴
//自动斗技
//自动御灵
//寮突破
//全自动
//御灵
//厕纸
func (f *TFMain) OnButtonQiTaZhiXingClick(sender vcl.IObject) {
    f.ButtonQiTaZhiXing.SetCaption("执行中.")
    f.Off_All_Buttone()
    r := yys_find_img.Result{}
    fp :=flagpiex.FLagPiex{}

    jjtpnum9 :=[][]int{//点击位置
       {221,141,140,40},//1
       {521,141,140,40},//2
       {830,141,140,40},//3
       {212,270,140,30},//4
       {523,270,140,30},//5
       {833,270,140,30},//6
       {222,391,140,20},//7
       {525,390,140,20},//8
       {830,390,140,20},//9
   }
   jjtpnum9_FuZhu :=[][]int{//判断是否已经攻击
       {380,110,11912916},
       {690,110,11715794},
       {990,110,11912916},
       {380,230,11912916},
       {690,230,11715794},
       {990,230,11912916},
       {380,350,11912916},
       {690,350,11715794},
       {990,350,11912916},
   }
    switch f.ComboBoxQiTa.ItemIndex() {
    //结界突破 0
    case 0:
        f.Zhuangtai_3()

        fmt.Println("结界突破 0")
        go func() {

            f.StopFlag=true
            for{
                if f.StopFlag==false {
                    break
                }
                f.XuanShang()
                //战斗界面
                if fp.FlagZhanDouJieMian(){
                    time.Sleep(time.Millisecond*100)
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
                    if fp.FlagJieJieTuPoOnLock()==false {
                        rd :=rand.Intn(1)
                        if rd==0{
                            f.DJ_Click_Range(908,551,1,1,"结界突破->上锁0")
                        }else{
                            f.DJ_Click_Range(938,552,1,1,"结界突破->上锁1")
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

                                f.DJ_Click_Range(865,465,130,30,"结界突破->刷新")
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
        }()
    //业原火痴 1
    case 1:
        f.Zhuangtai_3()
        fp:=flagpiex.FLagPiex{}
        fmt.Println("业原火痴 1")
        go func() {
            f.StopFlag=true
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
                        }
                    }
                    //御魂->业原火->上锁
                    Yuhun_3_meisuo_click:=r.Recognition(data.Yuhun_3_meisuo_click,0.9)
                    if Yuhun_3_meisuo_click!=nil {//御魂->业原火->上锁
                        f.Dj_click(Yuhun_3_meisuo_click,"上锁")
                        time.Sleep(time.Second*1)
                    }

                }
                //御魂->业原火
                Yuhun_1_yeyuanhuo_clik:=r.Recognition(data.Yuhun_1_yeyuanhuo_clik,0.9)
                if Yuhun_1_yeyuanhuo_clik!=nil {//御魂->业原火
                    f.Dj_click(Yuhun_1_yeyuanhuo_clik,"御魂->业原火")
                    time.Sleep(time.Second*1)
                }
                //探索->御魂
                Yuhun_0_click :=r.Recognition(data.Yuhun_0_click,0.9)
                if Yuhun_0_click!=nil { //探索->御魂
                    f.Dj_click(Yuhun_0_click,"探索->御魂")
                    time.Sleep(time.Second*1)
                }
            }
        }()
    //自动斗技 2
    case 2:
        f.Zhuangtai_all()
        fmt.Println("自动斗技 2")
        f.XuanShang()
    //自动御灵 3
    case 3:
        f.Zhuangtai_3()
        fmt.Println("自动御灵 3")
        go func() {
            f.StopFlag = true
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
                        time.Sleep(time.Millisecond*300)
                    }
                }
                //战斗退出
                f.ZhanDouTuiChu()
                }
            }()
    //寮突破 4
    case 4:
        f.Zhuangtai_3()
        fmt.Println("寮突破 4")
        go func() {
            f.StopFlag=true
            for{
                if f.StopFlag==false {
                    break
                }
                f.XuanShang()
                if fp.FlagZhanDouJieMianZhunBei(){ //如果没有上锁 手动点击准备
                        f.ZhanDouZhunBei()
                    time.Sleep(time.Second)
                }
                f.ZhanDouTuiChu()
                time.Sleep(time.Millisecond*200)
                //探索->结界突破->寮突破->选择->进攻->如果没有机会等待.
                Liaotupo_flag :=r.Recognition(data.Liaotupo_flag,0.9)
                if Liaotupo_flag!=nil {
                    time.Sleep(time.Second*30)
                    continue
                }
                //结界突破->寮突破->记录锚点
                Jiejietupo_2_liaotupo_ji_flag:=r.Recognition(data.Jiejietupo_2_liaotupo_ji_flag,0.9)
                if Jiejietupo_2_liaotupo_ji_flag!=nil {
                    //fmt.Println("请挑战")
                    //结界突破->寮突破->选择
                    Jiejietupo_1_xunzhang_click:=r.Recognition(data.Jiejietupo_1_xunzhang_click,0.7)
                    if Jiejietupo_1_xunzhang_click!=nil {
                        f.Dj_click(Jiejietupo_1_xunzhang_click,">寮突破->选择")
                        time.Sleep(time.Second*1)
                        //探索->结界突破->寮突破->选择->进攻
                        Jiejietupo_2_jingong_click :=r.Recognition(data.Jiejietupo_2_jingong_click,0.85)
                        if Jiejietupo_2_jingong_click!=nil {
                            if  f.TiaoZhanJiShuoff >=3{//次数达到上限退出
                                f.YYSLos("次数达到上限退出")
                                f.Stops()
                            }
                            f.Dj_click(Jiejietupo_2_jingong_click,"寮突破->进攻")
                            f.TiaoZhanJiShuoff +=1
                            time.Sleep(time.Second*2)
                        }
                    }
                    continue
                }else {
                    //结界突破->寮突破
                    Jiejietupo_1_liaotupo_click:=r.Recognition(data.Jiejietupo_1_liaotupo_click,0.9)
                    if Jiejietupo_1_liaotupo_click!=nil {
                        f.Dj_click(Jiejietupo_1_liaotupo_click,"结界突破->寮突破")
                        time.Sleep(time.Second*2)
                    }
                }
                //探索->结界突破
                Jiejietupo_0 :=r.Recognition(data.Jiejietupo_0,0.9)
                if Jiejietupo_0!=nil {
                    f.Dj_click(Jiejietupo_0,"探索->结界突破")
                    time.Sleep(time.Second*2)
                }
            }
        }()
    //全自动挂机5
    case 5:
        f.Zhuangtai_all()
        fmt.Println("全自动 5")
    //召唤厕纸6
    case 6:
        f.Zhuangtai_all()
        fmt.Println("召唤厕纸 6")
        go func() {
            f.StopFlag=true
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
                Cezhi_click :=r.Recognitions(data.Cezhi_click,0.9)
                if Cezhi_click!=nil {
                    f.Dj_clicks(Cezhi_click,"召唤祖安")
                    //time.Sleep(time.Second*1)
                }

            }
        }()
    //竞速秘闻挑战
    case 7:
        f.Zhuangtai_3()

        fmt.Println("竞速秘闻挑战 7")
        go func() {

            f.StopFlag=true
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

        }()
    case 8:
        go func() {
            for{
                if fp.FlagHuoDongWanShiWu(){
                    f.XuanShang()
                    f.ZhanDouZhunBei()
                    f.ZhanDouTuiChu()
                    //领取
                    HuoDongWanShiWu2 :=r.Recognition(data.HuoDongWanShiWu2,0.85)
                    if HuoDongWanShiWu2!=nil{
                        f.Dj_click(HuoDongWanShiWu2,"领取")
                        //f.DJ_Click_Range(910,569,39,41,"万事屋收取")
                        time.Sleep(time.Second*1)
                        f.DJ_Click_Range(46,103,50,50,"")
                        continue
                    }
                    //提交
                    HuoDongTijiao :=r.Recognition(data.HuoDongTijiao,0.85)
                    if HuoDongTijiao!=nil{
                        f.Dj_click(HuoDongTijiao,"提交")
                        time.Sleep(time.Second*1)
                        f.DJ_Click_Range(46,103,50,50,"")
                        continue
                    }
                    //突发情况
                    TuFaZhuangKuang :=r.Recognition(data.TuFaZhuangKuang,0.85)
                    if TuFaZhuangKuang!=nil{
                        f.DJ_Click_Range(46,103,50,50,"")
                        //f.DJ_Click_Range(910,569,39,41,"万事屋收取")
                        time.Sleep(time.Second*1)
                        continue
                    }
                    LingQIJiFa :=r.Recognition(data.LingQIJiFa,0.85)
                    if LingQIJiFa!=nil{
                        QianWang :=r.Recognition(data.QianWang,0.85)
                        if QianWang!=nil{
                            f.Dj_click(QianWang,"前往")
                        }
                        //f.DJ_Click_Range(910,569,39,41,"万事屋收取")
                        time.Sleep(time.Second*2)
                        TiaoZhan :=r.Recognition(data.TiaoZhan,0.85)
                        if QianWang!=nil{
                            f.Dj_click(TiaoZhan,"挑战")
                        }
                        continue
                    }
                    time.Sleep(time.Second*10)
                }
            }
        }()
    }

}
//妖气封印
func (f *TFMain) OnButtonYaoQiZhiXingClick(sender vcl.IObject) {
    f.ButtonYaoQiZhiXing.SetCaption("执行中.")
    f.Off_All_Buttone()
    r := yys_find_img.Result{}
    fp :=flagpiex.FLagPiex{}
    go func() {
        f.StopFlag=true
        for {
            if f.StopFlag == false {
                break
            }
            f.XuanShang()
            //庭院->妖气封印排队等待
            if fp.FlagYaoQiFengYinPaiDui(){
                time.Sleep(time.Millisecond*500)
                continue
            }
            //战斗主备手动点击准备
            if fp.FlagZhanDouJieMianZhunBei(){
                f.ZhanDouZhunBei()
                time.Sleep(time.Second)
                continue
            }
            //战斗界面
            if fp.FlagZhanDouJieMian() {
                time.Sleep(time.Millisecond * 100)
                continue
            }
            //战斗退出
            f.ZhanDouTuiChu()

            //庭院进组
            if fp.FlagTingYuan(){
                f.DJ_Click_Range(318,558,35,30,"庭院->组队")
                continue
            }
            //判断是否能找到红色妖气
            if fp.FlagALLZuDuiJieMian(){
                YaoQiFengYin_Falg :=r.Recognition(data.YaoQiFengYin_Falg,0.9)
                if YaoQiFengYin_Falg!=nil{
                    YaoQiFengYinQuXiaoPiPeiFlag :=r.Recognition(data.YaoQiFengYinQuXiaoPiPeiFlag,0.9)
                    //取消匹配存在返回
                    if YaoQiFengYinQuXiaoPiPeiFlag !=nil{
                        time.Sleep(time.Millisecond*500)
                        continue
                    }
                    switch f.ComboBoxYaoQi.ItemIndex() {
                    //日和坊
                    case 0:
                        YaoQiRiHeFang_Click :=r.Recognition(data.YaoQiRiHeFang_Click,0.9)
                        if YaoQiRiHeFang_Click!=nil{
                            f.Dj_click(YaoQiRiHeFang_Click,"选择日和坊")
                            time.Sleep(time.Millisecond*500)
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,267,1,-300,"")
                        }
                    //鬼使黑
                    case 1:
                        YaoQiGuiShiHei_Click :=r.Recognition(data.YaoQiGuiShiHei_Click,0.9)
                        if YaoQiGuiShiHei_Click!=nil{
                            f.Dj_click(YaoQiGuiShiHei_Click,"选择鬼使黑")
                            time.Sleep(time.Millisecond*500)
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,267,1,-300,"")
                        }
                    //淑图
                    case 2:
                        YaoQiShuTu_Click :=r.Recognition(data.YaoQiShuTu_Click,0.9)
                        if YaoQiShuTu_Click!=nil{
                            f.Dj_click(YaoQiShuTu_Click,"选择淑图")
                            time.Sleep(time.Millisecond*500)
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,267,1,300,"")
                            time.Sleep(time.Millisecond*200)
                        }
                    //小松丸
                    case 3:
                        YaoQiXiaoSongWan_Click :=r.Recognition(data.YaoQiXiaoSongWan_Click,0.9)
                        if YaoQiXiaoSongWan_Click!=nil{
                            f.Dj_click(YaoQiXiaoSongWan_Click,"选择小松丸")
                            time.Sleep(time.Millisecond*500)
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,267,1,-300,"")
                            time.Sleep(time.Millisecond*200)
                        }
                    //二口女
                    case 4:
                        YaoQiErKouNv_Click :=r.Recognition(data.YaoQiErKouNv_Click,0.9)
                        if YaoQiErKouNv_Click!=nil{
                            f.Dj_click(YaoQiErKouNv_Click,"选择二口女")
                            time.Sleep(time.Millisecond*500)
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,267,1,-300,"")
                            time.Sleep(time.Millisecond*200)
                        }
                    //骨女
                    case 5:
                        YaoQiGuNv_Click :=r.Recognition(data.YaoQiGuNv_Click,0.9)
                        if YaoQiGuNv_Click!=nil{
                            f.Dj_click(YaoQiGuNv_Click,"选择骨女")
                            time.Sleep(time.Millisecond*500)
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,267,1,300,"")
                            time.Sleep(time.Millisecond*200)
                        }
                    //饿鬼
                    case 6:
                        YaoQiEGui_Click :=r.Recognition(data.YaoQiEGui_Click,0.9)
                        if YaoQiEGui_Click!=nil{
                            f.Dj_click(YaoQiEGui_Click,"选择饿鬼")
                            time.Sleep(time.Millisecond*500)
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,467,1,300,"")
                            time.Sleep(time.Millisecond*200)
                        }
                    //海坊主
                    case 7:
                        YaoQiHaiFangZhu_Click :=r.Recognition(data.YaoQiHaiFangZhu_Click,0.9)
                        if YaoQiHaiFangZhu_Click!=nil{
                            f.Dj_click(YaoQiHaiFangZhu_Click,"选择海坊主")
                            time.Sleep(time.Millisecond*500)
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,267,1,-300,"")
                            time.Sleep(time.Millisecond*200)
                        }
                    //跳跳哥哥
                    case 8:
                        YaoQiTiaoTiaoGeGe_Click :=r.Recognition(data.YaoQiTiaoTiaoGeGe_Click,0.9)
                        if YaoQiTiaoTiaoGeGe_Click!=nil{
                            f.Dj_click(YaoQiTiaoTiaoGeGe_Click,"选择跳跳哥")
                            time.Sleep(time.Millisecond*500)
                            if fp.FlagYaoQiFengYinPaiDui(){
                                continue
                            }
                            f.Dj_click(r.Recognition(data.YaoQiZiDongPiPeiClick,0.9),"自动匹配")
                            time.Sleep(time.Millisecond*2000)
                        }else {
                            f.mv_mouse_Range(433,467,1,600,"")
                            time.Sleep(time.Millisecond*200)
                        }
                    }
                }else {
                    YaoQiFengYinZuDui :=r.Recognition(data.YaoQiFengYinZuDui,0.9)
                    if YaoQiFengYinZuDui !=nil{
                        f.Dj_click(YaoQiFengYinZuDui,"妖气封印")
                        time.Sleep(time.Millisecond*100)
                        continue
                    }else {
                        f.mv_mouse_Range(131,146,1,-300,"")
                        time.Sleep(time.Millisecond*100)
                        continue
                    }

                }

            }
        }
    }()
}

//绑定
func (f *TFMain) OnButtonBangDingClick(sender vcl.IObject) {

}
func (f *TFMain) OnButtonStopClick(sender vcl.IObject) {
    f.Stops()
}

func (f *TFMain) OnFormCreate(sender vcl.IObject) {
    f.ScreenCenter()
    f.hotKeyId = win2.GlobalAddAtom("HotKeyId") - 0xC000
    // rtl.ShiftStateToWord(shift) 这个只是更容易理解，也可以使用 MOD_CONTROL | MOD_ALT 方法
    if !win2.RegisterHotKey(f.Handle(), int32(f.hotKeyId),win2.MOD_NOREPEAT, keys.VkHome) {
        vcl.ShowMessage("注册热键失败。")
    }

    hwnd := getyyshwnd.Get_expvar_hwnd()
    hd :=strconv.Itoa(int(hwnd))
    if hd=="0"{
        fmt.Println("游戏没有启动....")
        f.YYSLos("游戏没有启动....")
    }
    f.YYSLos("本辅助永久免费")
    f.YYSLos("获取更新请加入")
    f.YYSLos("Q群:646105028")
    f.ComboBoxBangDing.SetText(hd)
    f.ComboBoxBangDing.SetItemIndex(0)
    f.CheckBoxGuanJueXing.SetEnabled(false)
    f.CheckBoxCaoRen.SetEnabled(false)
    f.ButtonBangDing.SetEnabled(false)
    f.ButtonBangDing.SetTextBuf("没做")
    f.SetCaption(strconv.Itoa(int(time.Now().UnixNano())))
    if time.Now().Year()!=2020&&int(time.Now().Month())<6{
       f.Close()
    }

}
type Month int
func (f *TFMain) OnFormDestroy(sender vcl.IObject) {//解锁热键
    if f.hotKeyId > 0 {
        win2.UnregisterHotKey(f.Handle(), int32(f.hotKeyId))
        win2.GlobalDeleteAtom(f.hotKeyId)
    }
}
func (f *TFMain) OnFormWndProc(msg *types.TMessage) {//响应热键

    f.InheritedWndProc(msg)
    /*
       TWMHotKey = record
         Msg: Cardinal;
         MsgFiller: TDWordFiller;
         HotKey: WPARAM;
         Unused: LPARAM;
         Result: LRESULT;
       end;
    */
    if msg.Msg == messages.WM_HOTKEY {
        if msg.WParam == types.WAPRAM(f.hotKeyId) {
            //vcl.ShowMessage("按下了Ctrl+F1")
            f.Stops()
        }
    }
}

func (f *TFMain) Stops() {
    f.YuHunJueXingOnClock =false//重置御魂房间锁
    f.StopFlag =false//停止重置
    f.ClickDaJiuMaFlag =false//重置点大舅妈
    f.ClickDaoCaoRenFlag =false//重置点草人
    f.FlagNum=false//重置玉环关闭计数判定
    f.OffNumGame=0//记录副本次数
    f.YuHunBuffFlag =false//停止后重置 buff检查
    f.OffBuff=0
    f.On_All_Buttone()
    fmt.Println("暂停")
    f.YYSLos("->暂停<-")
}

func (f *TFMain) YYSLos(s string){
    if s !=""{
        t:=time.Now().Format("15:04:05")
        f.ListBoxLog.Items().Add(t+":"+s)
        f.ListBoxLog.SetItemIndex(f.ListBoxLog.Items().Count()-1)
    }

}




