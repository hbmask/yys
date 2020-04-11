package flagpiex

import (
	//"github.com/lxn/win"
	"yys/yys_find_img"
)

type FLagPiex struct {
	Flagxyp [][]int
}


//战斗界面第一个 96,34 第二个157,34
//战斗中 好友->10601686 世界->10601686
//通用战斗界面判断
func (f FLagPiex)FlagZhanDouJieMian()bool{
	r :=yys_find_img.Result{}
	zdjmxyp:=[][]int{{42,34,10601686},{96,39,10601686},{157,34,10601686},}
	return r.Find_Pixels(zdjmxyp)
}
//通用战斗界面准备判断
func (f FLagPiex)FlagZhanDouJieMianZhunBei()bool{
	r :=yys_find_img.Result{}
	zdjmxyp:=[][]int{{1038,469,9485274},{42,34,10601686},{96,39,10601686},{157,34,10601686},}
	return r.Find_Pixels(zdjmxyp)
}


//通用胜利取色三组
func (f FLagPiex)FlagShengLi()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{408,136,1054842},{408,184,1121186},{450,184,1121187}}
	return r.Find_Pixels(xyp)
}

//通用失败取色三组
func (f FLagPiex)FlagShiBai()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{408,136,5588301},{408,184,1910582},{450,184,6837342}}
	return r.Find_Pixels(xyp)
}


//{461,534,13272126},
//62 132 202 0
//{493,534,657763},
//99 9 10 0
//{527,534,525904},
//80 6 8 0
//{627,534,920959},
//127 13 14 0
//通用胜利宝箱 四组
func (f FLagPiex)FlagShengLiBaoXiang()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{461,534,13272126},{493,534,657763},{527,534,525904},{627,534,920959},}
	return r.Find_Pixels(xyp)
}
//通用庭院标记
func (f FLagPiex)FlagTingYuan()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{960,37,8892630},{1025,37,8761302},{1094,37,8761300},}
	return r.Find_Pixels(xyp)
}

//通用探索标记
func (f FLagPiex)FlagTanSuo()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1000,35,8892630},{105,586,2911137},{621,11,5619679},}
	return r.Find_Pixels(xyp)
}
//御魂->业原火->上锁
func (f FLagPiex)FlagYeYuanHuoOnClock()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{496,516,13211553},{520,516,13867168}}
	return r.Find_Pixels(xyp)
}
//御魂->业原火->选择三层
func (f FLagPiex)FlagYeYuanHuoXuanZeSanCeng()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{240,420,9287630},{300,420,9287630},{330,420,9287630},}
	return r.Find_Pixels(xyp)
}
//{1001,42,9089752},
//216 178 138 0
//{1094,42,8695508},
//212 174 132 0
//{321,50,12967418},
//250 221 197 0
//{800,400,9481005},
//45 171 144 0
//御魂->业原火界面
func (f FLagPiex)FlagYeYuanHuoJiemian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1001,42,9089752},{1094,42,8695508},{321,50,12967418},{800,400,9481005},}
	return r.Find_Pixels(xyp)
}


//探索界面狗粮标记
func (f FLagPiex)FlagTanSuo_GouLiang()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1000,35,8892630},{621,11,5619679},}
	return r.Find_Pixels(xyp)
}
//{541,481,6206195},
//243 178 94 0
//{617,481,6206195},
//243 178 94 0
//{802,481,6206195},
//243 178 94 0
//{882,481,6206195},
//243 178 94 0
//探索界面与狗粮组队界面标记
func (f FLagPiex)FlagTanSuo_GouLiangZuDuiJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1000,35,8892630},{541,481,6206195},{617,481,6206195},{802,481,6206195},{882,481,6206195},}
	return r.Find_Pixels(xyp)
}
//{241,241,13097191},
//231 216 199 0
//{270,241,10136017},
//209 169 154 0
//{291,241,5595823},
//175 98 85 0
//{339,241,8556228},
//196 142 130 0
//{368,241,7438012},
//188 126 113 0
//探索界面 显示邀请是否是困难28
func (f FLagPiex)FlagTanSuo_KunNan28()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{241,241,13097191},{270,241,10136017},{291,241,5595823},{339,241,8556228},{368,241,7438012},}
	return r.Find_Pixels(xyp)
}
//{614,32,11057883},
//219 186 168 0
//{786,32,2845414},
//230 106 43 0
//{999,32,8958679},
//215 178 136 0
//{1089,32,8629972},
//212 174 131 0
//狗粮->进入副本界面(可以看到怪物选择)
func (f FLagPiex)FlagGouliangFuBenJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{614,32,11057883},{786,32,2845414},{999,32,8958679},{1089,32,8629972},}
	return r.Find_Pixels(xyp)
}
//结界突破->界面
func (f FLagPiex)FlagJieJieTuPoJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{270,550,6206195},{282,480,469365},{744,554,2891670},}
	return r.Find_Pixels(xyp)
}
//{380,110,11912916},
//212 198 181 0
//{380,230,11912916},
//212 198 181 0
//{380,350,11912916},
//212 198 181 0
//{690,110,11715794},
//210 196 178 0
//{690,230,11715794},
//210 196 178 0
//{690,350,11715794},
//210 196 178 0
//{990,110,11912916},
//212 198 181 0
//{990,230,11912916},
//212 198 181 0
//{990,350,11912916},
//212 198 181 0
//结界突破->不重复攻击
func (f FLagPiex)FlagJieJieTuPo_ShuaXin()bool{
	r :=yys_find_img.Result{}
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
	return r.Find_Pixels(jjtpnum9_FuZhu)
}

//结界突破->界面->刷新冷却
func (f FLagPiex)FlagJieJieTuPoLenQue()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{881,497,6206195},{890,497,6206195},{900,497,6206195},}
	return r.Find_Pixels(xyp)
}

//结界突破->界面->上锁
func (f FLagPiex)FlagJieJieTuPoOnLock()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{908,551,14197164},{938,552,16101303},}
	return r.Find_Pixels(xyp)
}



//{572,244,16318207},
//255 254 248 0
//{966,244,16515071},
//255 255 251 0
//御魂_觉醒组队->房间标记->占位
func (f FLagPiex)FlagYuhunJueXingFangJianWeiZhi2()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{571,245,16711679},}
	return r.Find_Pixels(xyp)
}
func (f FLagPiex)FlagYuhunJueXingFangJianWeiZhi3()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{966,244,16515071},}
	return r.Find_Pixels(xyp)
}
//{25,585,12158609},
//145 134 185 0
//{51,588,11898529},
//161 142 181 0
//御魂_觉醒组队->房间标记->锁
func (f FLagPiex)FlagYuhunJueXingFangJianOnLock()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{25,585,12158609},{51,588,11898529},}
	return r.Find_Pixels(xyp)
}
//{1095,35,10601686},
//214 196 161 0
//{736,49,13296379},
//251 226 202 0
//{1075,615,5289983},
//255 183 80 0
//御魂觉醒->房间标记
func (f FLagPiex) FlagYuHunJueXingFangJian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1095,35,10601686},{736,49,13296379},{1075,615,5289983},}
	return r.Find_Pixels(xyp)
}
//{1095,35,10601686},
//214 196 161 0
//{736,49,13296379},
//251 226 202 0
//{1075,585,3177214},
//254 122 48 0
//御魂觉醒->打手房间标记
func (f FLagPiex) FlagYuHunJueXingFangJian_DaShou()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1095,35,10601686},{736,49,13296379},{1075,585,3177214},}
	return r.Find_Pixels(xyp)
}
//{498,298,9878744},
//216 188 150 0
//{661,320,8762591},
//223 180 133 0
//{568,282,9091550},
//222 185 138 0
//{568,369,9751273},
//233 202 148 0
//一回目鼓的标记
func (f FLagPiex)FlagYuhunJueXingYiHuiMu()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{498,298,9878744},{661,320,8762591},{568,282,9091550},{568,369,9751273},}
	return r.Find_Pixels(xyp)
}
//{52,50,4867367},
//39 69 74 0
//{83,52,4801573},
//37 68 73 0
//御魂退出战斗数据标记点
func (f FLagPiex)FlagTuiChuZhanDouShuJu()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{52,50,4867367},{83,52,4801573},}
	return r.Find_Pixels(xyp)
}
//{1028,411,14935527},
//231 229 227 0
//{1055,415,3224419},
//99 51 49 0
//{1055,458,3158626},
//98 50 48 0
//{992,458,5206678},
//150 114 79 0
//{992,500,7958119},
//103 110 121 0
//御魂退出贪吃鬼标记
func (f FLagPiex)FlagTuiChuTanChiGui()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1028,411,14935527},{1055,415,3224419},{1055,458,3158626},{992,500,7958119},}
	return r.Find_Pixels(xyp)
}
//{499,322,5070706},
//114 95 77 0
//{616,373,6206195},
//243 178 94 0
//{627,400,6206195},
//243 178 94 0
//{715,389,6206195},
//243 178 94 0
//御魂战斗结束是否组队继续邀请选项
func (f FLagPiex)FlagTuiChuYaoQingJiXu()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{499,322,5070706},{616,373,6206195},{627,400,6206195},{715,389,6206195},}
	return r.Find_Pixels(xyp)
}

//觉醒buff{700,139,2427830},
//182 11 37 0
//御魂buff{701,199,2427830},
//182 11 37 0
//金币100%{699,259,2427830},
//182 11 37 0
//经验100%{702,320,2362295},
//183 11 36 0
//经验50%{702,380,2362295},
//183 11 36 0
//御魂buff 关闭状态 红色,标记颜色位置
func (f FLagPiex)FlagJueXingBUffRead()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{700,139,2427830},}
	return r.Find_Pixels(xyp)
}
func (f FLagPiex)FlagYuHunBuffRed()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{701,199,2427830},}
	return r.Find_Pixels(xyp)
}
//御魂开启状态
//觉醒buff{700,139,39395},
//227 153 0 0
//御魂buff{701,199,38883},
//227 151 0 0
//金币100%{699,259,39908},
//228 155 0 0
//经验100%{702,320,40164},
//228 156 0 0
//经验50%{702,380,40164},
//228 156 0 0
//御魂buff 开启状态 金色,标志位置
func (f FLagPiex)FlagJueXingBUffGold()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{700,139,39395},}
	return r.Find_Pixels(xyp)
}
func (f FLagPiex)FlagYuHunBuffGold()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{701,199,38883},}
	return r.Find_Pixels(xyp)
}

//
//{124,233,6007889},
//81 172 91 0
//{217,243,6533721},
//89 178 99 0
//{210,224,9488108},
//236 198 144 0
//御魂组队再次邀请->齿轮
func (f FLagPiex)FlagYuHunZuDuiYaoQingChiLun()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{124,233,6007889},{217,243,6533721},{210,224,9488108},}
	return r.Find_Pixels(xyp)
}

//
//{41,226,5794521},
//217 106 88 0
//{125,233,6270805},
//85 175 95 0
//御魂组队普通邀请
func (f FLagPiex)FlagYuHunZuDuiYaoQing()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{41,226,5794521},{125,233,6270805},}
	return r.Find_Pixels(xyp)
}

//
//{485,164,4410607},
//239 76 67 0
//{756,374,6139730},
//82 175 93 0
//{757,462,5597400},
//216 104 85 0
//悬赏三色定位
func (f FLagPiex)FlagXuanShangDingWei()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{485,164,4410607},{756,374,6139730},{757,462,5597400},}
	return r.Find_Pixels(xyp)
}

//{61,555,2698040},
//56 43 41 0
//{59,609,8136495},
//47 39 124 0
//{92,589,5978674},
//50 58 91 0
//换狗粮狗粮->全部

//{819,605,1187128},
//56 29 18 0
//{82,626,2566447},
//47 41 39 0
//{909,625,2180198},
//102 68 33 0
//{903,608,1186354},
//50 26 18 0
//换狗粮狗粮->选择性点击地板
func (f FLagPiex)FlagGouLiangDiBan()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{819,605,1187128},{82,626,2566447},{909,625,2180198},{903,608,1186354},}
	return r.Find_Pixels(xyp)
}

//{998,60,7380681},
//201 158 112 0
//{794,33,6006978},
//194 168 91 0
//{55,36,9457473},
//65 79 144 0
//御灵挑战界面
func (f FLagPiex)FlagYuLingTiaoZhanJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{998,60,7380681},{794,33,6006978},{55,36,9457473},}
	return r.Find_Pixels(xyp)
}
//{240,472,9155014},
//198 177 139 0
//{340,472,9090509},
//205 181 138 0
//御灵挑战界面->三层判断
func (f FLagPiex)FlagYuLingTiaoZhanJieMianSanCeng()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{240,472,9155014},{340,472,9090509},}
	return r.Find_Pixels(xyp)
}
//{495,516,9664640},
//128 120 147 0
//{519,516,13539488},
//160 152 206 0
//御灵挑战界面->上锁状态
func (f FLagPiex)FlagYuLingTiaoZhanJieShangSuo()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{495,516,9664640},{519,516,13539488},}
	return r.Find_Pixels(xyp)
}

//{1060,31,10601686},
//214 196 161 0
//{33,30,10278892},
//236 215 156 0
//{97,99,4340799},
//63 60 66 0
//组队界面
func (f FLagPiex)FlagALLZuDuiJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1060,31,10601686},{33,30,10278892},{97,99,4340799},}
	return r.Find_Pixels(xyp)
}

//{131,146,2249621},
//149 83 34 0
//{244,160,7057114},
//218 174 107 0
//妖气封印选项择
func (f FLagPiex) FlagYaoQiFengYinXuanZe()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{131,146,2249621},{244,160,7057114},}
	return r.Find_Pixels(xyp)
}

//{354,35,1917532},
//92 66 29 0
//{354,57,1984876},
//108 73 30 0
//{361,46,1851223},
//87 63 28 0
//妖气封印排队
func (f FLagPiex)FlagYaoQiFengYinPaiDui()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{354,35,1917532},{361,46,1851223},}
	return r.Find_Pixels(xyp)
}