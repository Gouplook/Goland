/**
 * @Author: yinjinlin
 * @File:  goHtmltoPdf
 * @Description:
 * @Date: 2021/12/28 下午1:18
 */

package gopdf

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

//  HTML 转化pdf
//  jdx006884011015-1-1

// 开发摘要：
// 先从master 切一个分支出来开发，最后合并到master和dev上 删除分支
// 涉及到的容器有：rpcCard、rpcOder、 apiCrard.
// 发卡业务、卡包业务。
// 注意问题，开发2期，是否兼顾3期。还是先开发2期，后面在开发三期。
// 先把


const (
	OrientationLandscape = "Landscape" // Landscape mode
	OrientationPortrait  = "Portrait"  // Portrait mode
	// 证件类型
	ICARD_1 = "身份证"
	ICARD_2 = "护照"

	// 项目类型
	ITEM_TYPE_single = 1 // 单项目
	ITEM_TYPE_sm     = 2 // 套餐
	ITEM_TYPE_card   = 3 // 综合卡
	ITEM_TYPE_hcard  = 4 // 限时卡
	ITEM_TYPE_ncard  = 5 // 限次卡
	ITEM_TYPE_hncard = 6 // 限时限次卡
	ITEM_TYPE_rcard  = 7 // 充值卡
	ITEM_TYPE_icard  = 8 // 身份卡
	ITEM_TYPE_pcard  = 9 // 预充卡

	// 支付方式 1= 现金 2=银行卡 3=支付宝 4=微信 5=数字人民币 6=其他 7= 渠道支付
	CASH      = 1
	BANKCARD  = 2
	ZHIFUBAO  = 3
	WEIXIN    = 4
	NUMBERREN = 5
	OTHERPAY  = 6
	CHANNELPAY = 7
)

// 涉及问题点
// 分页问题。
type CompactGoodLists struct {
	GoodId    int
	GoodName  string  // 产品名称
	RealPrice string  // 售价
	Discount  float64 // 产品折扣
}

type ReplyPrepaidCardsCompact struct {
	BusId        int    // 商户ID
	ItemId       int    // 卡项id
	ItemType     int    // 卡项类型
	CardProperty int    // 卡项性质 1= 记名卡 2= 不记名卡
	CardForm     int    // 卡项形式 1= 实体卡 2 = 虚拟卡 3=其他
	HeTongSn     string // 合同编号
	Cid          int
	MainBindId   int
	Did          int

	UserName        string // 用户真实姓名（合同的甲方）
	UserContactCall string // 用户联系方式
	ICardType       int    // 证件类型 1= 身份证 2= 护照
	ICardId         string // 证件号
	ShopName        string // 分店门店名称（BusCompanyName-ShopName 为合同的乙方）
	BrandName       string // 商户品牌名称
	RegisterAddress string // 注册地址
	Address         string // 经营地址
	CreditNo        string // 统一社会信用代码
	Contact         string // 店铺联系人/负责人
	ContactCall     string // 店铺联系电话（手机号或固定电话）

	IsAllSingle                 bool                // 包含全部单项目和产品
	HasGiveSingle               bool                // 是否含有赠送的项目
	SingleLists                 []HeTongSingleLists // 包含的单项目
	GoodLists                   []CompactGoodLists  // 包含的产品
	GiveSingleLists             []HeTongSingleLists // 赠送的单项目
	CardPrice                   float64             // 卡项面值
	PayRealPrice                float64             // 购卡支付的金额
	Discounts                   float64             // 折扣率，没有填 --
	AllSingleAndProductDiscount float64             // 包含全部项目和产品时的折扣
	GiveAmount                  float64             // 乙方赠送的服务项目、产品等值金额
	BankName                    string              // 开户银行
	AcctNo                      string              // 卡项资金监管专用账户
	BankCardNo                  string              // 卡项资金经营结算账户
	PayForm                     int                 // 支付方式 1= 现金 2=银行卡 3=支付宝 4=微信 5=数字人民币 6=其他
	PayTime                     string              // 付款时间
	IsForeverPeriod             bool                // 是否永久有效
	CompactStartTime            string              // 合同起始时间
	CompactEndTime              string              // 合同结束时间 ServicePeriod
	CardPeriod                  int                 // 卡项使用有效期 与合同有效期相同 自合同签订之日三年内有效 无固定期限
	CardRange                   int                 // 卡项使用范围,1=仅限指定门店使用,2=直营连锁门店通用,3=加盟门店通用,4=其它合作门店通用
	CardLimits                  []int               // 卡项使用权限 1=本人使用 2=家人使用 3= 朋友使用 4= 同事使用，5=任何人使用 6=可转让 7=可共享
	CardLimitsArrange           int                 // 卡项使用权限 特别约定 1= 是 2= 否
	PerformanceGuarantee        int                 // 履约保障 1=银行专用存管 2=履约保证保险 3=银行保函 4=信托 5=担保 6=其它 7=无
	HistoricalTrad              int                 // 历史交易次数 1=第一次  2=第二次 3=第三次及以上
	Trad                        int                 // 交易原因 1=乙方推荐  2=甲方自主  3=朋友推荐  4=其它
	SpecialAgreements           string              // 特别约定

	// 与旧版本兼容字段
	BuyCardDiscountDesc string // 购卡优惠详细
	IndustryId          int    // 分店经营领域id
	BusCompanyName      string // 企业/商户工商营业执照名称
	ServicePeriod       int    // 有效期,单位月（0-永久有效）

}
// 文档替换
func htmlReplace(args *ReplyPrepaidCardsCompact)(replaceFileName string) {
	// 先读取HTML文件
	htmlBytes,err := ioutil.ReadFile("./Contract.html")
	if err != nil {
		return
	}
	htmlStr := string(htmlBytes)
	// 证件类型
	var icardStr string
	if args.ICardType == 1 {
		icardStr = ICARD_1
	} else {
		icardStr = ICARD_2
	}
	// 卡项种类
	single := "❑"
	sm := "❑"
	card := "❑"
	hcard := "❑"
	ncard := "❑"
	hncard := "❑"
	rcard := "❑"
	icard := "❑"
	otherCard := "❑"

	switch args.ItemType {
	case ITEM_TYPE_single:
		single = "✓"
	case ITEM_TYPE_sm:
		sm = "✓"
	case ITEM_TYPE_card:
		card = "✓"
	case ITEM_TYPE_hcard:
		hcard = "✓"
	case ITEM_TYPE_ncard:
		ncard = "✓"
	case ITEM_TYPE_hncard:
		hncard = "✓"
	case ITEM_TYPE_rcard:
		rcard = "✓"
	case ITEM_TYPE_icard:
		icard = "✓"
	default:
		otherCard = "✓"
	}
	// 卡项性质
	registerdCard := "✓"
	unRegisterdCard := "❑"
	// 卡项形式
	physicalCard := "❑"
	virtualCard := "✓"
	other := "❑"

	// 产品
	serverLists := ""
	productLists := ""

	if !args.IsAllSingle {
		var serverTime string
		var discountStr string
		var productStr string
		for _, v := range args.SingleLists {
			numStr := strconv.Itoa(v.Num)
			if v.ServerTime != 0 {
				serverTime = strconv.Itoa(v.ServerTime)
			} else {
				serverTime = "--"
			}
			if v.SingleDiscount == 0.0 {
				discountStr = "--"
			} else {
				discountStr = strconv.FormatFloat(v.SingleDiscount, 'G', 10, 64)
			}

			serverLists += "<li style=\"list-style:none;\"> 服务名称：" + v.SingleName +
				" 标准：" + serverTime + "分钟" +
				" 次数：" + numStr + "次" +
				"&nbsp 折扣：" + discountStr +
				"&nbsp 金额：" + v.RealPrice + "元"+
				"</li>"

		}

		for _, p := range args.GoodLists {
			numStr := "1"
			goodSsp := "--"
			if p.Discount == 0 {
				productStr = "--"
			} else {
				productStr = strconv.FormatFloat(p.Discount, 'G', 10, 64)
			}
			productLists += "<li style=\"list-style:none;\"> 产品名称：" + p.GoodName +
				" 规格：" + goodSsp +
				" 数量：" + numStr +
				"&nbsp折扣：" + productStr +
				"&nbsp金额：" + p.RealPrice + "元" +
				"</li>"
		}

	} else {
		var allDiscount string
		allDiscount = strconv.FormatFloat(args.AllSingleAndProductDiscount, 'G', 10, 64)
		if args.ItemType == ITEM_TYPE_rcard || args.ItemType == ITEM_TYPE_icard {
			serverLists = "<li style=\"list-sytle:none;\"> 适用全部产品 " + "&nbsp折扣: " + allDiscount + "</li>"
		} else {
			serverLists = "适用全部产品"
		}

	}

	// 赠送项目
	giveProductLists := ""
	giveServerLists := ""
	if args.HasGiveSingle == true {
		var serverTime string
		for _, v := range args.GiveSingleLists {
			numStr := strconv.Itoa(v.Num)
			if v.ServerTime != 0 {
				serverTime = strconv.Itoa(v.ServerTime)
			} else {
				serverTime = "--"
			}
			// giveServerLists += "<div style=\"width:100%; display: flex;\">" +
			// 	"<div style=\"min-width:200px;\">服务名称："+ v.SingleName +"</div>" +
			// 	"<div style=\"min-width:150px;\">标准："+ serverTime + "</div>" +
			// 	"<div style=\"min-width:100px;\">次数："+ numStr + "</div>" +
			// 	"<div style=\"min-width:100px;\">金额："+ "1-100"+"</div>"+
			// 	"</div>"
			giveServerLists += "<li style=\"list-style:none;\"> 服务名称：" + v.SingleName +
				" 标准：" + serverTime + "分钟" +
				" 次数：" + numStr + "次" +
				"&nbsp金额：" + v.RealPrice + "元"+
				"</li>"
		}

	} else {
		giveServerLists = "无赠送"
	}

	// 产品与服务 补充说明
	productAdditionalNotes := "无"
	giveAdditionalNotes := "无"

	// 卡项金额
	cardPrice := strconv.FormatFloat(args.CardPrice, 'G', 10, 64)
	discount := strconv.FormatFloat(args.Discounts, 'G', 10, 64)
	payRealPrice := strconv.FormatFloat(args.PayRealPrice, 'G', 10, 64)
	giveAmount := strconv.FormatFloat(args.GiveAmount, 'G', 10, 64)

	// 支付方式
	cash := "❑"
	bankCard := "❑"
	zhiFuBao := "❑"
	weiXin := "❑"
	numberRen := "❑"
	otherPay := "❑"
	channelPay := "❑"

	switch args.PayForm {
	case  CASH:
		cash = "✓"
	case  BANKCARD:
		bankCard = "✓"
	case  ZHIFUBAO:
		zhiFuBao = "✓"
	case  WEIXIN:
		weiXin = "✓"
	case NUMBERREN:
		numberRen = "✓"
	case CHANNELPAY:
		channelPay = "✓"
	default:
		otherPay = "✓"
	}

	// 付款时间
	payTimesp, _ := strconv.ParseInt(args.PayTime, 10, 64)
	tm := time.Unix(payTimesp, 0)
	payYear := strconv.Itoa(tm.Year())
	payMonth := tm.Format("01")
	payDate := strconv.Itoa(tm.Day())
	payTime := strconv.Itoa(tm.Hour())
	paySecond := strconv.Itoa(tm.Second())

	// 有效期间
	startTimesp, _ := strconv.ParseInt(args.CompactStartTime, 10, 64)
	endTimesp, _ := strconv.ParseInt(args.CompactEndTime, 10, 64)
	sTm := time.Unix(startTimesp, 0)
	eTm := time.Unix(endTimesp, 0)
	contractPeriodStartYear := strconv.Itoa(sTm.Year())
	contractPeriodStartMonth := sTm.Format("01")
	contractPeriodStartDate := strconv.Itoa(tm.Day())

	contractPeriodEndYear := strconv.Itoa(eTm.Year())
	contractPeriodEndMonth := eTm.Format("01")
	contractPeriodEndDate := strconv.Itoa(eTm.Day())

	// 卡项使用有效期
	contractPeriodAlike := "✓"
	contractPeriodThree := "❑"
	contractPeriodLimit := "❑"

	// 卡项使用范围
	// 1=仅限指定门店使用,2=直营连锁门店通用,3=加盟门店通用,4=其它合作门店通用
	ats := "❑"
	atls := "❑"
	atjs := "❑"
	atos := "❑"

	// 卡项使用权限 (多选）
	myselfUse := "❑"
	homeUse := "❑"
	friedUse := "❑"
	colleaguesUse := "❑"
	anyoneUse := "❑"
	transferable := "❑"
	shareable := "❑"

	// 卡项使用权限 1=本人使用 2=家人使用 3= 朋友使用 4= 同事使用，5=任何人使用 6=可转让 7=可共享

	// 特别约定：非本人使用时，是否需要本人授权同意
	spaYes := "❑"
	spaNo := "❑"

	// 履约保障
	bankSpecificDepository := "❑"
	performanceGuaranteeInsurance := "❑"
	bankGuarantee := "❑"
	affiance := "❑"
	surety := "❑"
	insuranceOther := "❑"
	insuranceNo := "❑"

	// 签订时间
	signingTime := time.Now().Local().Format("2006-01-02")

	// 冷静期
	calmDate := "7"
	written := "❑"
	emial := "❑"
	Note := "✓"
	calmWeixin := "❑"
	calmOther := "❑"
	paymentDate := "5"
	paymentAmount := "30"

	// 争议解决方式
	arbitrate := "上海市宝山区"
	judgement := "上海市宝山区"

	// 文本替换
	textStr := strings.Replace(htmlStr, "[UserName]", args.UserName, -1)
	textStr = strings.Replace(textStr, "[UserContactCall]", args.UserContactCall, -1)
	textStr = strings.Replace(textStr, "[ICardType]", icardStr, -1)
	textStr = strings.Replace(textStr, "[ICardId]", args.ICardId, -1)
	textStr = strings.Replace(textStr, "[ShopName]", args.ShopName, -1)
	textStr = strings.Replace(textStr, "[BrandName]", args.BrandName, -1)
	textStr = strings.Replace(textStr, "[RegisterAddress]", args.RegisterAddress, -1)
	textStr = strings.Replace(textStr, "[Address]", args.Address, -1)
	textStr = strings.Replace(textStr, "[CreditNo]", args.CreditNo, -1)

	textStr = strings.Replace(textStr, "[Single]", single, -1)
	textStr = strings.Replace(textStr, "[Sm]", sm, -1)
	textStr = strings.Replace(textStr, "[Card]", card, -1)
	textStr = strings.Replace(textStr, "[Hcard]", hcard, -1)
	textStr = strings.Replace(textStr, "[Ncard]", ncard, -1)
	textStr = strings.Replace(textStr, "[HNcard]", hncard, -1)
	textStr = strings.Replace(textStr, "[Rcard]", rcard, -1)
	textStr = strings.Replace(textStr, "[Icard]", icard, -1)
	textStr = strings.Replace(textStr, "[OtherCard]", otherCard, -1)

	textStr = strings.Replace(textStr, "[RegisterdCard]", registerdCard, -1)
	textStr = strings.Replace(textStr, "[UnRegisterdCard]", unRegisterdCard, -1)
	textStr = strings.Replace(textStr, "[PhysicalCard]", physicalCard, -1)
	textStr = strings.Replace(textStr, "[VirtualCard]", virtualCard, -1)
	textStr = strings.Replace(textStr, "[Other]", other, -1)

	// 产品与服务
	textStr = strings.Replace(textStr, "[ProductLists]", productLists, -1)
	textStr = strings.Replace(textStr, "[ServerLists]", serverLists, -1)

	textStr = strings.Replace(textStr, "[GiveProductLists]", giveProductLists, -1)
	textStr = strings.Replace(textStr, "[GiveServerLists]", giveServerLists, -1)

	textStr = strings.Replace(textStr, "[ProductAdditionalNotes]", productAdditionalNotes, -1)
	textStr = strings.Replace(textStr, "[GiveAdditionalNotes]", giveAdditionalNotes, -1)

	textStr = strings.Replace(textStr, "[CardPrice]", cardPrice, -1)
	textStr = strings.Replace(textStr, "[Discount]", discount, -1)
	textStr = strings.Replace(textStr, "[PayRealPrice]", payRealPrice, -1)
	textStr = strings.Replace(textStr, "[GiveAmount]", giveAmount, -1)

	textStr = strings.Replace(textStr, "[BankName]", args.BankName, -1)
	textStr = strings.Replace(textStr, "[AcctNo]", args.AcctNo, -1)
	textStr = strings.Replace(textStr, "[BankCardNo]", args.BankCardNo, -1)

	textStr = strings.Replace(textStr, "[Cash]", cash, -1)
	textStr = strings.Replace(textStr, "[BankCard]", bankCard, -1)
	textStr = strings.Replace(textStr, "[ZhiFuBao]", zhiFuBao, -1)
	textStr = strings.Replace(textStr, "[WeiXin]", weiXin, -1)
	textStr = strings.Replace(textStr, "[ChannelPayments]", channelPay, -1)
	textStr = strings.Replace(textStr, "[NumberRen]", numberRen, -1)
	textStr = strings.Replace(textStr, "[OtherPay]", otherPay, -1)

	textStr = strings.Replace(textStr, "[PayYear]", payYear, -1)
	textStr = strings.Replace(textStr, "[PayMonth]", payMonth, -1)
	textStr = strings.Replace(textStr, "[PayDate]", payDate, -1)
	textStr = strings.Replace(textStr, "[PayTime]", payTime, -1)
	textStr = strings.Replace(textStr, "[PaySecond]", paySecond, -1)

	textStr = strings.Replace(textStr, "[ContractPeriodStartYear]", contractPeriodStartYear, -1)
	textStr = strings.Replace(textStr, "[ContractPeriodStartMonth]", contractPeriodStartMonth, -1)
	textStr = strings.Replace(textStr, "[ContractPeriodStartDate]", contractPeriodStartDate, -1)
	textStr = strings.Replace(textStr, "[ContractPeriodEndYear]", contractPeriodEndYear, -1)
	textStr = strings.Replace(textStr, "[ContractPeriodEndMonth]", contractPeriodEndMonth, -1)
	textStr = strings.Replace(textStr, "[ContractPeriodEndDate]", contractPeriodEndDate, -1)

	textStr = strings.Replace(textStr, "[ContractPeriodAlike]", contractPeriodAlike, -1)
	textStr = strings.Replace(textStr, "[ContractPeriodThree]", contractPeriodThree, -1)
	textStr = strings.Replace(textStr, "[ContractPeriodLimit]", contractPeriodLimit, -1)

	textStr = strings.Replace(textStr, "[ApplyToShop]", ats, -1)
	textStr = strings.Replace(textStr, "[ApplyToLinkageShop]", atls, -1)
	textStr = strings.Replace(textStr, "[ApplyToJoinShop]", atjs, -1)
	textStr = strings.Replace(textStr, "[ApplyToOtherShop]", atos, -1)

	textStr = strings.Replace(textStr, "[MyselfUse]", myselfUse, -1)
	textStr = strings.Replace(textStr, "[HomeUse]", homeUse, -1)
	textStr = strings.Replace(textStr, "[FriedUse]", friedUse, -1)
	textStr = strings.Replace(textStr, "[ColleaguesUse]", colleaguesUse, -1)
	textStr = strings.Replace(textStr, "[AnyoneUse]", anyoneUse, -1)
	textStr = strings.Replace(textStr, "[Transferable]", transferable, -1)
	textStr = strings.Replace(textStr, "[Shareable]", shareable, -1)

	textStr = strings.Replace(textStr, "[SpAYes]", spaYes, -1)
	textStr = strings.Replace(textStr, "[SpANo]", spaNo, -1)

	textStr = strings.Replace(textStr, "[BankSpecificDepository]", bankSpecificDepository, -1)
	textStr = strings.Replace(textStr, "[PerformanceGuaranteeInsurance]", performanceGuaranteeInsurance, -1)
	textStr = strings.Replace(textStr, "[BankGuarantee]", bankGuarantee, -1)
	textStr = strings.Replace(textStr, "[Affiance]", affiance, -1)
	textStr = strings.Replace(textStr, "[Surety]", surety, -1)
	textStr = strings.Replace(textStr, "[InsuranceOther]", insuranceOther, -1)
	textStr = strings.Replace(textStr, "[InsuranceNo]", insuranceNo, -1)

	// textStr = strings.Replace(textStr, "[HistoryFirst]", historyFirst, -1)
	// textStr = strings.Replace(textStr, "[HistorySecond]", historySecond, -1)
	// textStr = strings.Replace(textStr, "[HistoryThree]", historyThree, -1)
	//
	// textStr = strings.Replace(textStr, "[RecommendsB]", recommendsB, -1)
	// textStr = strings.Replace(textStr, "[RecommendsA]", recommendsA, -1)
	// textStr = strings.Replace(textStr, "[FriendRecommend]", friendRecommend, -1)
	// textStr = strings.Replace(textStr, "[RecommendsOther]", recommendsOther, -1)

	textStr = strings.Replace(textStr, "[SpecialAgreements]", args.SpecialAgreements, -1)

	textStr = strings.Replace(textStr, "[SigningTime]", signingTime, -1)
	textStr = strings.Replace(textStr, "[Contact]", args.Contact, -1)
	textStr = strings.Replace(textStr, "[ContactCall]", args.ContactCall, -1)

	textStr = strings.Replace(textStr, "[CalmDate]", calmDate, -1)
	textStr = strings.Replace(textStr, "[Written]", written, -1)
	textStr = strings.Replace(textStr, "[Email]", emial, -1)
	textStr = strings.Replace(textStr, "[Note]", Note, -1)
	textStr = strings.Replace(textStr, "[CalmWeiXin]", calmWeixin, -1)
	textStr = strings.Replace(textStr, "[CalmOther]", calmOther, -1)
	textStr = strings.Replace(textStr, "[PaymentDate]", paymentDate, -1)
	textStr = strings.Replace(textStr, "[PaymentAmount]", paymentAmount, -1)

	textStr = strings.Replace(textStr, "[Arbitrate]", arbitrate, -1)
	textStr = strings.Replace(textStr, "[Judgement]", judgement, -1)

	replaceFileName = time.Now().Format("2006-01-02-13-04-05")+".html"
	// 创建文件
	f, err := os.Create("./" + replaceFileName)
	if err != nil {
		return
	}


	_,err = f.WriteString(textStr)

	if err != nil {
		return
	}
	return
}


func HtmlToPdf() {

	// 初始化pdf
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return
	}
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA5)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)  // pdf 灰度生成

	var args ReplyPrepaidCardsCompact
	//
	args.ItemType = 1
	args.HeTongSn = "YF-0001"
	// 模板替换（HTML替换)
	replaceFileName := htmlReplace(&args)
	page := wkhtmltopdf.NewPage("./"+ replaceFileName)
	// 设置页眉和页脚
	page.HeaderLeft.Set("合同编号：" + args.HeTongSn)
	page.HeaderLine.Set(true)
	page.HeaderFontSize.Set(7)
	page.HeaderSpacing.Set(3)
	page.FooterCenter.Set("第[page]页")
	page.FooterFontSize.Set(7)
	page.FooterSpacing.Set(3)
	page.Zoom.Set(0.9)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return
	}
	if args.HeTongSn == "" {
		args.HeTongSn ="YF-120001-00001"
	}
	fileName := args.HeTongSn + ".pdf"
	err = pdfg.WriteFile("./"+fileName)
	if err != nil {
		return
	}

	defer func() {
		_ = os.Remove("./" + replaceFileName)
	}()



}
