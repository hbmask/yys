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
    "net"
    _ "net/http/pprof"
    "os"
    "sort"
    "strconv"
    "time"
    "yys/data"
    "yys/flagpiex"
    "yys/getyyshwnd"
    "yys/yys_find_img"
)

//::private::
type TFMainFields struct {
    StopFlag            bool //暂停
    YuHunJueXingOnClock bool //御魂觉醒房间是否上锁
    ClickDaJiuMaFlag    bool //点大舅妈
    ClickDaoCaoRenFlag  bool //点稻草人
    JuXingBuffFlag bool//觉醒buff状态 启动还是未启动
    YuHunBuffFlag bool //御魂buff状态 启动还是为启动
    FlagNum bool//每次对点怪只进行一次操作.
    FlagDouJiSZ bool//每次对点怪只进行一次操作.
    GuanYuHunNext bool//关闭御魂条件传递到下一个参数
    OffBuff int//计数多少次以后关闭buff.
    OffNumGame int//记录副本次数如果是0 停止辅助..
    StopYuHunNum int //记录已经刷了多少次御魂.
    TiaoZhanJiShuoff int//当挑战次数达到上线时,点击后没有进入副本,停止
    HWND win.HWND//窗口句柄
    hotKeyId types.ATOM//热键
}


func NewTFMainFields( stopflag bool,yuhunjuexingonclock bool,clickdajiuma bool,clickdaocaoren bool)TFMainFields{
    return TFMainFields{StopFlag:stopflag,YuHunJueXingOnClock:yuhunjuexingonclock,ClickDaJiuMaFlag:clickdajiuma,ClickDaoCaoRenFlag:clickdaocaoren}
}
var e=expvar.NewInt("erhwnd")
func init(){
   YYSHWND := getyyshwnd.YYSHWND{}
   hwnd:=YYSHWND.Get_yys_hwnd()
   e.Set(int64(hwnd))
   rand.Seed(time.Now().UnixNano())
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
                if fp.FlagZhanDouJieMianJiaCeng()&&f.CheckBoxGuanYuHun.Checked(){//战斗界面->点击加层
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
                //被邀请进组
                if fp.FlagYuHunZuDuiYaoQing(){ //被邀请进组
                    H10 :=r.Recognition(data.H10,0.85)
                    if H10!=nil {
                        YuHunChiLunClick :=r.Recognition(data.YuHunChiLun_Click,0.85)
                        if YuHunChiLunClick !=nil{ //被邀请进组选择齿轮
                            f.Dj_click(YuHunChiLunClick,"齿轮进入")
                            //f.DJ_Click_Range(198,212,30,30,"从此轮进组")
                            time.Sleep(time.Millisecond*200)
                            f.YuHunJueXingOnClock =false
                            continue
                        }
                        f.DJ_Click_Range(125,233,5,5,"接受魂10邀请")
                        f.YuHunJueXingOnClock =false
                        time.Sleep(time.Millisecond*200)
                        continue
                    }
                    H11 :=r.Recognition(data.H11,0.85)
                    if H11!=nil {
                        YuHunChiLunClick :=r.Recognition(data.YuHunChiLun_Click,0.85)
                        if YuHunChiLunClick !=nil{ //被邀请进组选择齿轮
                            f.Dj_click(YuHunChiLunClick,"齿轮进入")
                            //f.DJ_Click_Range(198,212,30,30,"从此轮进组")
                            f.YuHunJueXingOnClock =false
                            time.Sleep(time.Millisecond*200)
                            continue
                        }
                        f.DJ_Click_Range(125,233,5,5,"接受魂11邀请")
                        f.YuHunJueXingOnClock =false
                        time.Sleep(time.Millisecond*200)
                        continue
                    }
                }
                //在庭院,探索,房间
                if fp.FlagYuHunJueXingFangJian_DaShou()||f.OffNumGame==0||fp.FlagTingYuan()||fp.FlagTanSuo(){


                    if fp.FlagYuhunJueXingFangJianOnLock(){//房间上锁
                        f.YuHunJueXingOnClock =true //房间上锁=自动准备
                        f.ClickDaJiuMaFlag=false//组队房间重置
                        f.ClickDaoCaoRenFlag=false//组队房间重置
                        f.FlagNum=false//计数判定
                    }
                    time.Sleep(time.Millisecond*100)
                    if f.OffBuff>=300{//满足条件关闭御魂
                        f.YuHunTingYuanOffBuffJianCha()
                        f.YuHunOffBuffJianCha()
                    }
                    f.OffBuff++
                    fmt.Println(f.OffBuff)
                }
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
                f.YuHunOrJueXingFangZhu(2,fp,f.CheckBoxGuanYuHun.Checked())
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
                f.YuHunOrJueXingFangZhu(3,fp,f.CheckBoxGuanYuHun.Checked())
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
    f.Off_All_Buttone()
    f.Zhuangtai_3()
    //mbgouliangxy :=make([][]int,0,0)
    go func() {
        f.StopFlag=true
        for {
            if f.StopFlag == false {
                    break
                }
            //有困难28标志和邀请勾选
            if fp.FlagYuHunZuDuiYaoQing(){
                KunNan28Flag :=r.Recognition(data.GouLiangKunNan28_Flag,0.85) //少女与面具
                if KunNan28Flag !=nil{
                    f.DJ_Click_Range(125,233,5,5,"接受狗粮28邀请")
                }
            }
            //探索界面与狗粮组队界面
            if fp.FlagTanSuo_GouLiang()||fp.FlagTanSuo_GouLiangZuDuiJieMian(){//&&fp.FlagTanSuo_KunNan28(){
                if f.OffBuff>=120{//满足条件关闭御魂
                    f.GouLiangOffBuffJianCha()
                    f.OffBuff=0
                    f.YuHunBuffFlag =false
                }
                f.OffBuff++
                time.Sleep(time.Millisecond*100)
            }
            //K28房间开宝箱
            if fp.FlagK28GouLiangFangJian(){
                K28BaoXiangClick :=r.Recognition(data.K28BaoXiang_Click,0.9) //K28房间开宝箱
                if K28BaoXiangClick !=nil {
                    f.Dj_click(K28BaoXiangClick,"开宝箱啦")
                }
            }
            //战斗界面->点击加层
            if fp.FlagZhanDouJieMianJiaCeng(){//战斗界面->点击加层
             if  f.YuHunBuffFlag ==false{//狗粮buff状态
                f.DJ_Click_Range(106,595,26,25,"狗粮经验加层")
                for  {
                    if fp.FlagGouLiangBuffRed50(){//红色状态
                        if fp.FlagGouLiangBuffRed100() { //100红色状态
                            f.DJ_Click_Range(697,319,60,6,"开启100%经验")
                        }
                        time.Sleep(time.Millisecond*500)
                        f.DJ_Click_Range(697,380,60,6,"开启50%经验")
                        f.YuHunBuffFlag =true
                        f.DJ_Click_Range(0,489,600,30,"")
                        //time.Sleep(time.Millisecond*500)
                        f.StopYuHunNum=0
                        break
                    }
                    if fp.FlagGouLiangBuffGold50(){//金色状态
                        //if fp.FlagGouLiangBuffGold100() { //100金色状态
                        //    f.YuHunBuffFlag =true
                        //}
                        //f.DJ_Click_Range(317,489,600,61,"御魂buff已打开")
                        f.YuHunBuffFlag =true
                        f.DJ_Click_Range(0,489,600,30,"buff已经开启")
                        //time.Sleep(time.Millisecond*500)
                        //f.DJ_Click_Range(0,489,600,30,"")
                        f.StopYuHunNum=0
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
            if fp.FlagZhanDouJieMian(){//战斗界面
                if fp.FlagZhanDouJieMianZhunBei(){//战斗准备界面
                    zbGouliangManJiFlag :=r.RecognitionsGouLiang_2Man(data.GouliangManJi_Flag,1100,420,0.85) //获取更换满级的目标
                    if len(zbGouliangManJiFlag)<3&&len(zbGouliangManJiFlag)>0{
                     f.ZhanDouZhunBei()
                     time.Sleep(time.Second)
                    }
                   switch f.ComboBoxGouLiang.ItemIndex() {
                   //1级N
                   case 0:
                       f.GouLiangGengHuan(r,data.GouLiangNKa_Click,"N",data.GouliangManJi_Flag,data.GouLiang1JiN_Click,"1级N")
                       //GouLiangQuanBu_Click:=r.Recognition(data.GouLiangQuanBu_Click,0.9)//狗粮->全部
                       //if GouLiangQuanBu_Click!=nil{
                       //    f.Dj_click(GouLiangQuanBu_Click,"全部")
                       //    time.Sleep(time.Millisecond*300)
                       //    GouLiangNKaClick:=r.Recognition(data.GouLiangNKaClick,0.9)//狗粮N
                       //    if GouLiangNKaClick!=nil{
                       //        f.Dj_click(GouLiangNKaClick,"选择->N")
                       //        time.Sleep(time.Millisecond*100)
                       //    }
                       //}
                       //GouLiangNKaFlag:=r.Recognition(data.GouLiangNKaFlag,0.9)//狗粮N
                       //if GouLiangNKaFlag!=nil{
                       //    mb:=r.RecognitionsGouLiang_2Man(data.GouliangManJiFlag,790,420,0.85)//获取更换满级的目标
                       //    GouLiang1JiN_Click := r.Recognitions(data.GouLiang1JiN_Click, 0.9) //从N卡中找到1级N卡
                       //    if len(GouLiang1JiN_Click)!=0{
                       //        for i,_ :=range mb{
                       //            if i==1{
                       //                time.Sleep(time.Millisecond*500)
                       //                GouLiang1JiN_Click = r.Recognitions(data.GouLiang1JiN_Click, 0.9)//获取第二次1级红坐标
                       //            }
                       //            f.move_click(mb[i].Result_img_centen, GouLiang1JiN_Click, 0, 90, "更换1级N")
                       //            //time.Sleep(time.Millisecond*200)
                       //        }
                       //    }else {
                       //        f.YYSLos("没有找到1级N")
                       //    }
                       //}
                   //1级白
                   case 1:
                       f.GouLiangGengHuan(r,data.GouLiangSuCai_Click,"素材",data.GouLiangSuCaiFlag,data.GouLiang1JiBai_Click,"1级白")
                       //GouLiangQuanBu_Click:=r.Recognition(data.GouLiangQuanBu_Click,0.9)//狗粮->全部
                       //if GouLiangQuanBu_Click!=nil{
                       //    f.Dj_click(GouLiangQuanBu_Click,"全部")
                       //    time.Sleep(time.Millisecond*300)
                       //    GouLiangSuCai_Click:=r.Recognition(data.GouLiangSuCai_Click,0.9)//素材
                       //    if GouLiangSuCai_Click!=nil{
                       //        f.Dj_click(GouLiangSuCai_Click,"选择->素材")
                       //        time.Sleep(time.Millisecond*100)
                       //    }
                       //}
                       //GouLiangSuCaiFlag:=r.Recognition(data.GouLiangSuCaiFlag,0.9)//素材
                       //if GouLiangSuCaiFlag!=nil{
                       //    mb:=r.RecognitionsGouLiang_2Man(data.GouliangManJiFlag,750,360,0.85)//获取更换满级的目标坐标
                       //    GouLiang1JiBai_Click := r.Recognitions(data.GouLiang1JiBai_Click, 0.9) //从素材中找到1级白
                       //    if len(GouLiang1JiBai_Click)!=0{
                       //        for i,_ :=range mb{
                       //            if i>0{
                       //                time.Sleep(time.Millisecond*600)
                       //               GouLiang1JiBai_Click = r.Recognitions(data.GouLiang1JiBai_Click, 0.9)//获取第二次1级坐标 刷新
                       //            }
                       //                f.move_click(mb[i].Result_img_centen, GouLiang1JiBai_Click, 0, 120, strconv.Itoa(len(mb))+"更换1级白0|"+strconv.Itoa(i))
                       //
                       //
                       //            //time.Sleep(time.Second)
                       //        }
                       //    }else {
                       //        f.YYSLos("没有找到1级白")
                       //    }
                       //}
                   //1级红
                   case 2:
                       f.GouLiangGengHuan(r,data.GouLiangSuCai_Click,"素材",data.GouLiangSuCaiFlag,data.GouLiang1JiHong_Click,"1级红")
                       //GouLiangQuanBu_Click:=r.Recognition(data.GouLiangQuanBu_Click,0.9)//狗粮->全部
                       //if GouLiangQuanBu_Click!=nil{
                       //    f.Dj_click(GouLiangQuanBu_Click,"全部")
                       //    time.Sleep(time.Millisecond*300)
                       //    GouLiangSuCai_Click:=r.Recognition(data.GouLiangSuCai_Click,0.9)//狗粮素材
                       //    if GouLiangSuCai_Click!=nil{
                       //        f.Dj_click(GouLiangSuCai_Click,"选择->素材")
                       //        time.Sleep(time.Millisecond*100)
                       //    }
                       //}
                       //GouLiangSuCaiFlag:=r.Recognition(data.GouLiangSuCaiFlag,0.9)//狗粮
                       //if GouLiangSuCaiFlag!=nil{
                       //    mb:=r.RecognitionsGouLiang_2Man(data.GouliangManJiFlag,790,420,0.85)//获取更换满级的目标
                       //    GouLiang1JiHong_Click := r.Recognitions(data.GouLiang1JiHong_Click, 0.9) //从素材中找到1级红
                       //    if len(GouLiang1JiHong_Click)!=0{
                       //        for i,_ :=range mb{
                       //            if i==1{
                       //                time.Sleep(time.Millisecond*500)
                       //                GouLiang1JiHong_Click = r.Recognitions(data.GouLiang1JiHong_Click, 0.9)//获取第二次1级红坐标
                       //            }
                       //            f.move_click(mb[i].Result_img_centen, GouLiang1JiHong_Click, 0, 120, "更换1级红")
                       //            //time.Sleep(time.Millisecond*200)
                       //        }
                       //    }else {
                       //        f.YYSLos("没有找到1级红")
                       //    }
                       //}
                   //20级白
                   case 3:
                       f.GouLiangGengHuan(r,data.GouLiangSuCai_Click,"素材",data.GouLiangSuCaiFlag,data.GouLiang20Ji_Click,"20级白")
                       //GouLiangQuanBu_Click:=r.Recognition(data.GouLiangQuanBu_Click,0.9)//狗粮->全部
                       //if GouLiangQuanBu_Click!=nil{
                       //    f.Dj_click(GouLiangQuanBu_Click,"全部")
                       //    time.Sleep(time.Millisecond*300)
                       //    GouLiangSuCai_Click:=r.Recognition(data.GouLiangSuCai_Click,0.9)//狗粮素材
                       //    if GouLiangSuCai_Click!=nil{
                       //        f.Dj_click(GouLiangSuCai_Click,"选择->素材")
                       //        time.Sleep(time.Millisecond*100)
                       //    }
                       //}
                       //GouLiangSuCaiFlag:=r.Recognition(data.GouLiangSuCaiFlag,0.9)//狗粮
                       //if GouLiangSuCaiFlag!=nil{
                       //    mb:=r.RecognitionsGouLiang_2Man(data.GouliangManJiFlag,790,420,0.85)//获取更换满级的目标
                       //    GouLiang20Ji_Click := r.Recognitions(data.GouLiang20Ji_Click, 0.9) //从素材中找到20级白
                       //    if len(GouLiang20Ji_Click)!=0{
                       //        for i,_ :=range mb{
                       //            if i==1{
                       //                time.Sleep(time.Millisecond*500)
                       //                GouLiang20Ji_Click = r.Recognitions(data.GouLiang20Ji_Click, 0.9)//获取第二次1级红坐标
                       //            }
                       //            f.move_click(mb[i].Result_img_centen, GouLiang20Ji_Click, 0, 120, "更换20级白")
                       //            //time.Sleep(time.Millisecond*200)
                       //        }
                       //    }else {
                       //        f.YYSLos("没有找到1级红")
                       //    }
                       //
                       //}
                   //20级N
                   case 4:
                       GouLiangNKaClick :=r.Recognition(data.GouLiangNKa_Click,0.9) //狗粮N
                       f.Dj_click(GouLiangNKaClick,"选择->N")
                       time.Sleep(time.Millisecond*500)
                       f.YYSLos("此选项暂时无效")
                       if GouLiangNKaClick !=nil{

                       }
                   }
                    GouliangManJiFlag :=r.Recognitions(data.GouliangManJi_Flag,0.85) //获取满级图像
                    if len(GouliangManJiFlag)==3&&fp.FlagGouLiangDiBan()==false{     //3个满级后更换狗粮
                            f.SJ_Click_Range(150,450,10,10,"进入更换狗粮界面.")
                            //time.Sleep(time.Millisecond*600)
                        for  {//直到到达指定界面退出循环
                            if fp.FlagHuanGouLiangShiSHenKuan()==true{
                                break
                            }
                            time.Sleep(time.Millisecond*50)
                        }
                    }
                }
                //time.Sleep(time.Millisecond *300)
            }
            //队长离开副本后 退出
            if  fp.FlagGouliangFuBenJieMian(){
                time.Sleep(time.Millisecond*500)
                fmt.Println(fp.FlagGouliangFuBenJieMian(),fp.FlagTanSuo_GouLiangFuBenDuiZhang())
                if fp.FlagGouliangFuBenJieMian()&&fp.FlagTanSuo_GouLiangFuBenDuiZhang()==false{//狗粮副本界面
                    f.DJ_Click_Range(32,51,12,14,"队长已经退出")
                    time.Sleep(time.Millisecond*400)
                    f.DJ_Click_Range(650,350,100,25,"立刻退出")
                }
            }
            f.XuanShang()
            f.ZhanDouTuiChu()
            time.Sleep(time.Millisecond*100)
        }
    }()
}
//更换狗粮流程
//GouLiang_LeiXing ->全部->选择狗粮类型
//LeiXingName ->(log 显示替换的什么狗粮) N卡
//LEiXingFlag ->选择狗粮类型后 确定 flg常驻 N字样
//Find_GLImg ->查找常驻狗粮的匹配图像 比如 白蛋 红蛋 N卡
//GL_Name ->log显示更换狗粮等级和类型 1级N
func (f *TFMain) GouLiangGengHuan(r yys_find_img.Result, GouLiangLeiXing string,LeiXingName string, LeiXingFlag string, FindGLImg string, GLName string){

    GouLiangQuanBuClick :=r.Recognition(data.GouLiangQuanBu_Click,0.9) //狗粮->全部
    if GouLiangQuanBuClick !=nil{
        f.Dj_click(GouLiangQuanBuClick,"全部")
        time.Sleep(time.Millisecond*300)
        GouLiangLeiXingClick :=r.Recognition(GouLiangLeiXing,0.9) //狗粮N 素材 标记要更换的类型
        if GouLiangLeiXingClick !=nil{
            f.Dj_click(GouLiangLeiXingClick,"选择->"+LeiXingName)
            time.Sleep(time.Millisecond*100)
        }
    }
    LEiXingFlag :=r.Recognition(LeiXingFlag,0.9) //狗粮N 看是否已经是所选择的狗粮类型 比如说 素材
    if LEiXingFlag !=nil{
        mb:=r.RecognitionsGouLiang_2Man(data.GouliangManJi_Flag,790,420,0.85)//获取更换满级的目标
        mb =f.SortResultR(mb)
        FindGLImgs := r.Recognitions(FindGLImg, 0.9) //从N卡中找到1级N卡 在类型中查找是否符合匹配图像
        FindGLImgs = f.SortResultL(FindGLImgs)
        if len(FindGLImgs)!=0{ //检查是否有 匹配的图像个数
            for i :=range mb{
                if i==1{
                    time.Sleep(time.Second)
                    FindGLImgs = r.Recognitions(FindGLImg, 0.9) //获取第二次1级红坐标
                    FindGLImgs = f.SortResultR(FindGLImgs)
                }
                f.move_click(mb[i].Result_img_centen, FindGLImgs, 0, 120, "更换"+GLName)
                //time.Sleep(time.Millisecond*200)
            }
        }else {
            f.YYSLos("没有找到"+ GLName)
        }
    }
}
//从小到大排序
func (f *TFMain) SortResultL(rs []*yys_find_img.Result)[]*yys_find_img.Result  {

    var SortRs []*yys_find_img.Result
    var rsint []int
    for i :=0;i<len(rs);i++{//取x横坐标 到 切片中
        rsint=append(rsint,rs[i].Result_img_centen[0])
    }
    sort.Ints(rsint[:])//正常排序
    fmt.Println(rsint,len(rs))
    for i:=0;i<len(rs);i++{//通过活得的切片对Result进行从小到大排序
        for j:=0;j<len(rs);j++{
            if rsint[i] == rs[j].Result_img_centen[0]{
                SortRs =append(SortRs,rs[j])
                fmt.Println(SortRs[i],"从小到大")
                continue
        }
       }
    }
    return SortRs
}
//倒序 从大到小
func (f *TFMain) SortResultR(rs []*yys_find_img.Result)[]*yys_find_img.Result  {
    var SortRs []*yys_find_img.Result
    var rsint []int
    for i :=0;i<len(rs);i++{
        rsint=append(rsint,rs[i].Result_img_centen[0])
    }
    //sort.Ints(rsint[:])
    //sort.Reverse()
    sort.Sort(sort.Reverse(sort.IntSlice(rsint)))//倒序
    fmt.Println(rsint,len(rs))
    for i:=0;i<len(rs);i++{
        for j:=0;j<len(rs);j++{
            if rsint[i] == rs[j].Result_img_centen[0]{
                SortRs =append(SortRs,rs[j])
                fmt.Println(SortRs[i],"从大到小")
                continue
            }
        }
    }
    return SortRs
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
    switch f.ComboBoxQiTa.ItemIndex() {
    //结界突破 0
    case 0:
        f.StopFlag=true
        go f.JieJieTuPo(r,fp)
    //业原火痴 1
    case 1:
        f.StopFlag=true
        go f.YeYuanHuoChi(r,fp)
    //自动斗技 2
    case 2:
        f.StopFlag=true
        go f.ZiDongDouJi(r,fp)
    //自动御灵 3
    case 3:
        f.StopFlag=true
        go f.ZiDongYuLin(r,fp)
    //寮突破 4
    case 4:
        f.StopFlag=true
        go f.LiaoTuPo(r,fp)
    //全自动挂机5
    case 5:
        f.Zhuangtai_all()
        fmt.Println("全自动 5")
    //召唤厕纸6
    case 6:
        f.StopFlag=true
        go f.ZhaoHuanCeZhi(r,fp)
    //竞速秘闻挑战
    case 7:
        f.StopFlag=true
        go f.JinSuMiWenTiaoZhan(r,fp)
    case 8://结界卡合成
        f.StopFlag=true
        f.JieJieKaHeCheng(r,fp)
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
            if !f.StopFlag {
                break
            }
            f.XuanShang()
            //战斗界面
            if fp.FlagZhanDouJieMian() {
                time.Sleep(time.Millisecond * 500)
                continue
            }
            //战斗退出
            f.ZhanDouTuiChu()
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
            //庭院进组
            if fp.FlagTingYuan(){
                f.DJ_Click_Range(318,558,35,30,"庭院->组队")
                time.Sleep(time.Millisecond * 500)
                continue
            }
            //如果房主 直接挑战
            if fp.Flag_FangJian_TiaoZhan()&&fp.FlagYuHunJueXingFangJian_DaShou(){
                f.DJ_Click_Range(1058,563,50,30,"房间挑战")
                time.Sleep(time.Millisecond*800)
            }
            //判断是否能找到红色妖气
            //fmt.Println(fp.FlagALLZuDuiJieMian())
            if fp.FlagALLZuDuiJieMian(){
                YaoQiFengYinFalg :=r.Recognition(data.YaoQiFengYin_Falg,0.9)
                if YaoQiFengYinFalg !=nil{
                    YaoQiFengYinQuXiaoPiPeiFlag :=r.Recognition(data.YaoQiFengYinQuXiaoPiPeiFlag,0.9)
                    //取消匹配存在返回
                    if YaoQiFengYinQuXiaoPiPeiFlag !=nil{
                        time.Sleep(time.Millisecond*500)
                        continue
                    }

                    switch f.ComboBoxYaoQi.ItemIndex() {
                    //日和坊
                    case 0:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiRiHeFang_Click,"选择日和坊")
                    //鬼使黑
                    case 1:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiGuiShiHei_Click,"选择鬼使黑")
                    //淑图
                    case 2:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiShuTu_Click,"选择淑图")
                    //小松丸
                    case 3:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiXiaoSongWan_Click,"选择小松丸")
                    //二口女
                    case 4:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiErKouNv_Click,"选择二口女")
                    //骨女
                    case 5:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiGuNv_Click,"选择骨女")
                    //饿鬼
                    case 6:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiEGui_Click,"选择饿鬼")
                    //海坊主
                    case 7:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiHaiFangZhu_Click,"选择海坊主")
                    //跳跳哥哥
                    case 8:
                        f.YaoQiFengYing_XuanZeShiShen(r,data.YaoQiTiaoTiaoGeGe_Click,"选择跳跳哥")
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

//获取mac地址
func getMacAddrs() (macAddrs []string) {
    netInterfaces, err := net.Interfaces()
    if err != nil {
        fmt.Printf("fail to get net interfaces: %v", err)
        return macAddrs
    }

    for _, netInterface := range netInterfaces {
        macAddr := netInterface.HardwareAddr.String()
        if len(macAddr) == 0 {
            continue
        }

        macAddrs = append(macAddrs, macAddr)
    }
    return macAddrs
}

func (f *TFMain) OnFormCreate(sender vcl.IObject) {
    hname,_ :=os.Hostname()
    fmt.Println(getMacAddrs(),hname)
    f.ScreenCenter()
    f.hotKeyId = win2.GlobalAddAtom("HotKeyId") - 0xC000
    // rtl.ShiftStateToWord(shift) 这个只是更容易理解，也可以使用 MOD_CONTROL | MOD_ALT 方法
    if !win2.RegisterHotKey(f.Handle(), int32(f.hotKeyId),win2.MOD_NOREPEAT, keys.VkHome) {
        vcl.ShowMessage("注册热键失败。")
    }
    f.YYSLos("本辅助永久免费")
    f.YYSLos("获取更新请加入")
    f.YYSLos("Q群:646105028")
    hwnd := getyyshwnd.Get_expvar_hwnd()
    hd :=strconv.Itoa(int(hwnd))
    if hd=="0"{
        fmt.Println("游戏没有启动....")
        f.YYSLos("游戏没有启动....")
    }
    rt :=win.RECT{}
    win.GetClientRect(hwnd,&rt)
    fmt.Println(rt.Bottom,rt.Left,rt.Right,rt.Top)
    if rt.Bottom!=640&&rt.Right!=1136{
        f.YYSLos("***************")
        f.YYSLos("游戏分辨率有问题")
        f.YYSLos("正确是:1136*640")
        f.YYSLos("***************")
    }

    f.OffNumGame,_ = strconv.Atoi(f.EditCiShu.Text())//初始化御魂次数
    f.ComboBoxBangDing.SetText(hd)
    f.ComboBoxBangDing.SetItemIndex(0)
    f.CheckBoxGuanJueXing.SetEnabled(false)
    f.CheckBoxCaoRen.SetEnabled(false)
    f.ButtonBangDing.SetEnabled(false)
    f.ButtonBangDing.SetTextBuf("没做")
    f.SetCaption(strconv.Itoa(int(time.Now().UnixNano())))
    if time.Now().Year()!=2020&&int(time.Now().Month())<11{
       f.Close()
    }

}
//type Month int
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
    //f.OffNumGame=0//记录副本次数
    f.YuHunBuffFlag =false//停止后重置 buff检查
    f.OffBuff=0//关闭buff计数
    f.On_All_Buttone()
    fmt.Println("暂停")
    f.YYSLos("->暂停<-")
}

func (f *TFMain) YYSLos(s string){
    if s !=""{
        t:=time.Now().Format("15:04:05"                                                                                                                                                                   )
        f.ListBoxLog.Items().Add(t+":"+s)
        f.ListBoxLog.SetItemIndex(f.ListBoxLog.Items().Count()-1)
    }

}





func (f *TFMain) OnComboBoxQiTaChange(sender vcl.IObject) {

}

