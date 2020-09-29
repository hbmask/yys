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
//战斗界面好友->{96,39,10601686},
//战斗界面世界->{157,34,10601686},
//战斗界面加层->{121,608,15006718},
//通用战斗界面判断
func (f FLagPiex)FlagZhanDouJieMianJiaCeng()bool{
	r :=yys_find_img.Result{}
	zdjmxyp:=[][]int{{121,608,15006718},}
	return r.Find_Pixels(zdjmxyp)
}
//通用战斗界面准备判断
func (f FLagPiex)FlagZhanDouJieMianZhunBei()bool{
	r :=yys_find_img.Result{}
	zdjmxyp:=[][]int{{1038,469,9485274},{42,34,10601686},{96,39,10601686},{157,34,10601686},}
	return r.Find_Pixels(zdjmxyp)
}
//战斗界面准备前 判断返回是否存在
//{31,38,10601686},
//214 196 161 0
func (f FLagPiex)FlagZhanDouJieMianZhunBeiFanHui()bool{
	r :=yys_find_img.Result{}
	zdjmxyp:=[][]int{{42,34,10601686},}
	return r.Find_Pixels(zdjmxyp)
}
//战斗界面判断 准备是否存在
//{1045,500,4608092},*
//92 80 70 0
//{1006,502,3355195},*
//59 50 51 0
//{1035,578,8107485},
func (f FLagPiex)FlagZhanDouJieMianZhunBeiFanHui_ZhunBei()bool{
	r :=yys_find_img.Result{}
	zdjmxyp:=[][]int{{1035,578,8107485},}
	return r.Find_Pixels(zdjmxyp)
}
//战斗界面 准备 鼓棒
//{1044,575,330056},*
//72 9 5 0
//{1024,575,528746},*
//106 17 8 0
//{1035,578,8107485},
//221 181 123 0
//func (f FLagPiex)FlagZhanDouJieMianZhunBeiFanHui_GuDiBu()bool{
//	r :=yys_find_img.Result{}
//	zdjmxyp:=[][]int{{1035,578,8107485},}
//	return r.Find_Pixels(zdjmxyp)
//}

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
	xyp:=[][]int{{1025,37,8761302},{1086,45,7643852},}
	return r.Find_Pixels(xyp)
}

//通用探索标记
func (f FLagPiex)FlagTanSuo()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1000,35,8892630},{105,586,2911137},{930,26,10601686},}
	return r.Find_Pixels(xyp)
}
//御魂->业原火->上锁
func (f FLagPiex)FlagYeYuanHuoOnClock()bool{
	r :=yys_find_img.Result{}
	//xyp:=[][]int{{496,516,13211553},{520,516,13867168}}
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
	//xyp:=[][]int{{1094,42,8695508},{321,50,12967418},{800,400,9481005},}
	xyp:=[][]int{{1094,42,8695508},{321,50,12967419},{800,400,9481005},}
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
//{38,119,2969200},
//112 78 45 0
//狗粮队长状态标记
func (f FLagPiex)FlagTanSuo_GouLiangFuBenDuiZhang()bool{
	r :=yys_find_img.Result{}
	//xyp:=[][]int{{29,127,1586518},}
	xyp:=[][]int{{29,127,1586517},}
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
//{614,32,11057883},叶儿粑
//219 186 168 0
//{786,32,2845414},体力
//230 106 43 0
//{999,32,8958679},信封
//215 178 136 0
//{1089,32,8629972},世界聊天
//212 174 131 0
//狗粮->进入副本界面(可以看到怪物选择)
func (f FLagPiex)FlagGouliangFuBenJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{614,32,11057883},{999,32,8958679}}
	return r.Find_Pixels(xyp)
}
//结界突破->界面
func (f FLagPiex)FlagJieJieTuPoJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{114,574,6659779}}
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
	xyp:=[][]int{{869,519,6206195}}
	return r.Find_Pixels(xyp)
}

//结界突破->界面->上锁
func (f FLagPiex)FlagJieJieTuPoOnLock()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{736,533,1318440},}
	return r.Find_Pixels(xyp)
}
//结界突破->寮突破界面->上锁
func (f FLagPiex)FlagLiaoTuPoOnLock()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{177,550,1581871},}
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
	//xyp:=[][]int{{1095,35,10601686},{736,49,13296379},{1075,615,5355519},}
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
	//xyp:=[][]int{{1095,35,10601686},{736,49,13296379},{1075,585,3177214},}
	xyp:=[][]int{{1095,35,10601686},{736,49,13296379},}
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
//{72,573,1976635},
//59 41 30 0
//{72,599,1646894},
//46 33 25 0
//御魂退出战斗数据标记点 下
func (f FLagPiex)FlagTuiChuZhanDouShuJu1()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{72,573,1976635},{72,599,1646894},}
	return r.Find_Pixels(xyp)
}

//{71,51,2108221},
//61 43 32 0
//{71,60,5603494},
//166 128 85 0
//御魂退出战斗数据标记点 上
func (f FLagPiex)FlagTuiChuZhanDouShuJu2()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{71,51,2108221},{71,60,5603494},}
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
//狗粮经验buff 金色标志 100
func (f FLagPiex)FlagGouLiangBuffGold100()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{697,319,37603},}
	return r.Find_Pixels(xyp)
}
//狗粮经验buff 金色标志 50
func (f FLagPiex)FlagGouLiangBuffGold50()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{697,380,37603},}
	return r.Find_Pixels(xyp)
}
//狗粮经验buff 红色标志 100
func (f FLagPiex)FlagGouLiangBuffRed100()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{697,319,2690226}}
	return r.Find_Pixels(xyp)
}
//狗粮经验buff 红色标志 50
func (f FLagPiex)FlagGouLiangBuffRed50()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{697,380,2690226},}
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
	xyp:=[][]int{{484,162,4282352},{484,123,7770803},}
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
//{33,30,10213100},
//236 215 156 0
//{97,99,4340799},
//63 60 66 0
//组队界面
func (f FLagPiex)FlagALLZuDuiJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{38,34,10872308},{97,99,4340799},}
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

//{677,48,10601686},
//妖气封印排队
func (f FLagPiex)FlagYaoQiFengYinPaiDui()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{677,48,10601686},}
	return r.Find_Pixels(xyp)
}


//竞速秘闻挑战
func (f FLagPiex)FlagJingSuMiWenTiaoZhan()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1033,498,12177372},{1033,566,12966885}}
	return r.Find_Pixels(xyp)
}
//20200924活动
func (f FLagPiex)FlagHuDong_TZ()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{48,46,16314864}}
	return r.Find_Pixels(xyp)
}

func (f FLagPiex)FlagHuDong_TZ_SUO()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{828,455,1581871}}
	return r.Find_Pixels(xyp)
}
//{432,64,1121155},
//131 27 17 0
//{406,101,1186973},
//157 28 18 0
//{455,104,1121186},
//162 27 17 0
//竞速秘闻胜利退出
func (f FLagPiex)FlagJingSuMiWenShengLiTuiChu()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{432,64,1121155},{406,101,1186973},{455,104,1121186},}
	return r.Find_Pixels(xyp)
}

//{110,613,1712692},
//52 34 26 0
//{508,613,2375002},
//90 61 36 0
//{1018,615,1318188},
//44 29 20 0
//活动万事屋
func (f FLagPiex)FlagHuoDongWanShiWu()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{110,613,1712692},{508,613,2375002},{1018,615,1318188},}
	return r.Find_Pixels(xyp)
}







//{892,452,11454428},
//220 199 174 0
//{913,451,9549523},
//211 182 145 0
//{917,503,8898283},
//235 198 135 0
//万事屋挑战
func (f FLagPiex)FlagWanShiWuTiaoZhan()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{892,452,11454428},{913,451,9549523},{917,503,8898283},}
	return r.Find_Pixels(xyp)
}

//{1038,500,10933230},
//238 211 166 0
//{1038,565,10604519},
//231 207 161 0
//万事屋出发
func (f FLagPiex)FlagWanShiWuChuFa()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1038,500,10933230},{1038,565,10604519},}
	return r.Find_Pixels(xyp)
}
//{1,1,2500671},
//63 40 38 0
//{1,200,2693673},
//41 26 41 0
//{1,400,3751787},
//107 63 57 0
//{1,800,4294967295},
//进入万事屋
func (f FLagPiex)FlagJinWanShiWu()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1,1,2500671},{1,200,2693673},{1,400,3751787},{1,800,4294967295},}
	return r.Find_Pixels(xyp)
}

//{376,376,9020091},
//187 162 137 0
//{503,376,10270155},
//203 181 156 0
//{550,376,10270155},
//203 181 156 0
//{600,376,10270155},
//203 181 156 0
//获得奖励
func (f FLagPiex)FlagWanShiWuHuoDeJiangLi()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{376,376,9020091},{503,376,10270155},{550,376,10270155},{600,376,10270155},}
	return r.Find_Pixels(xyp)
}

//{447,101,4413055},
//127 86 67 0
//{673,101,3886193},
//113 76 59 0
//斗技 挑战界面
func (f FLagPiex)FlagDouJiJieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{266,79,8576242},{183,25,1455961},}
	return r.Find_Pixels(xyp)
}
//斗技 升级后
func (f FLagPiex)FlagDouJiShenJi()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{741,312,3293656},}
	return r.Find_Pixels(xyp)
}
//{52,141,5400962},
//130 105 82 0
func (f FLagPiex)FlagDouJi1700ZiDongShangZHen()bool{//斗技自动上阵
	r :=yys_find_img.Result{}
	xyp:=[][]int{{52,141,5400962},}
	return r.Find_Pixels(xyp)
}

//{1037,575,3550814},
//94 46 54 0
//{973,575,10589007},
//79 147 161 0
//斗技 拔得头筹
func (f FLagPiex)FlagDouJiBaDeTouChou()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1037,575,3550814},{973,575,10589007},}
	return r.Find_Pixels(xyp)
}
//斗技 对战中...
//{215,22,4740210},
//114 84 72 0
func (f FLagPiex)FlagDouJiZhanDouZhong()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{215,24,10601686},}
	return r.Find_Pixels(xyp)
}


//结界卡三个点位置
func (f FLagPiex)FlagJieJieKa()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{772,139,924982},{898,394,924981},{636,399,858931},}
	return r.Find_Pixels(xyp)
}

//结界卡 是否有添加字样
func (f FLagPiex)FlagJieJieKa_JiXuTianJia()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{945,535,1910325}}
	return r.Find_Pixels(xyp)
}
//御魂战斗失败是否继续战斗
func (f FLagPiex)FlagYuHun_JiXu_ZhanDou()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{624,386,6206195},{708,386,6206195}}
	return r.Find_Pixels(xyp)
}
//房间挑战
func (f FLagPiex)Flag_FangJian_TiaoZhan()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{1088,569,8311532}}
	return r.Find_Pixels(xyp)
}

//寮突破 界面识别
func (f FLagPiex) Flag_LiaoTuPo_JieMian()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{87,581,8367052}}
	return r.Find_Pixels(xyp)
}

//寮突破 进攻次数识别
func (f FLagPiex)Flag_LiaoTuPo_JinGongCiShu()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{312,497,11844546}}
	return r.Find_Pixels(xyp)
}

//寮突破 攻破后 从新标记 位置
func (f FLagPiex)Flag_LiaoTuPo_Po()bool{
	r :=yys_find_img.Result{}
	xyp:=[][]int{{395,142,12439002}}
	return r.Find_Pixels(xyp)
}