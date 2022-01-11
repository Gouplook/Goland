/**
 * @Author: yinjinlin
 * @File:  goHtmltoPdf
 * @Description:
 * @Date: 2021/12/28 下午1:18
 */

package gopdf

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io/ioutil"
	"log"
	"os"
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
)

// 涉及问题点
// 分页问题。

func HtmlToPdf() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		log.Fatal(err)
	}

	//
	textStr, err := ioutil.ReadFile("./Contract.html")
	if err != nil {
		return
	}
	htmlStr := string(textStr)
	str := strings.Replace(htmlStr,"[UserName]","张三",-1)

	// 生成一个临时
	fileCach := "cachfile" +time.Now().Format("2006-01-02 13:03-04")+".html"
	f,err :=os.Create(fileCach)
	if err!=nil{
		fmt.Println(err)
		return
	}
	n,err:=f.WriteString(str)
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(n)

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA5)
	pdfg.Orientation.Set(OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage(fileCach)



	// Set options for this page
	// page.FooterRight.Set("[page]")
	str2 := "合同编号：HT000001"
	// page.HeaderCenter.Set(str)
	page.HeaderLeft.Set(str2)
	page.HeaderLine.Set(true)
	page.HeaderFontSize.Set(7)
	page.HeaderSpacing.Set(3)

	page.FooterCenter.Set("第 [page]页")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.9)


	// Add to document  // 4
	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err  = pdfg.Create()
	if err != nil {

	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./12R.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done

}

//
// BusId        int    // 商户ID
// ItemId       int    // 卡项id
// ItemType     int    // 卡项类型
// CardProperty int    // 卡项性质 1= 记名卡 2= 不记名卡
// CardForm     int    // 卡项形式 1= 实体卡 2 = 虚拟卡 3=其他
// HeTongSn     string // 合同编号 *
//
// UserName        string // 用户真实姓名（合同的甲方）*
// UserContactCall string // 用户联系方式 *
// ICardType       int    // 证件类型 1= 身份证 2= 护照 *
// ICardId         string // 证件号 *
// ShopName        string // 分店门店名称（BusCompanyName-ShopName 为合同的乙方）
// BrandName       string // 商户品牌名称
// RegisterAddress string // 注册地址
// Address         string // 经营地址
// CreditNo        string // 统一社会信用代码
// Contact         string // 店铺联系人/负责人
// ContactCall     string // 店铺联系电话（手机号或固定电话）
//
// IsAllSingle                 bool                // 包含全部单项目和产品
// HasGiveSingle               bool                // 是否含有赠送的项目
// SingleLists                 []HeTongSingleLists // 包含的单项目
// GoodLists                   []CompactGoodLists  // 包含的产品
// GiveSingleLists             []HeTongSingleLists // 赠送的单项目
// CardPrice                   float64             // 卡项面值
// PayRealPrice                float64             // 购卡支付的金额
// Discounts                   float64             // 折扣率，没有填 --
// AllSingleAndProductDiscount float64             // 包含全部项目和产品时的折扣
// GiveAmount                  float64             // 乙方赠送的服务项目、产品等值金额 *
// BankName                    string              // 开户银行
// AcctNo                      string              // 卡项资金监管专用账户
// BankCardNo                  string              // 卡项资金经营结算账户
// PayForm                     int                 // 支付方式 1= 现金 2=银行卡 3=支付宝 4=微信 5=数字人民币 6=其他 *
// PayTime                     string              // 付款时间 *
// IsForeverPeriod             bool                // 是否永久有效
// CompactStartTime            string              // 合同起始时间
// CompactEndTime              string              // 合同结束时间 ServicePeriod
// CardPeriod                  int                 // 卡项使用有效期 与合同有效期相同 自合同签订之日三年内有效 无固定期限
// CardRange                   int                 // 卡项使用范围,1=仅限指定门店使用,2=直营连锁门店通用,3=加盟门店通用,4=其它合作门店通用
// CardLimits                  []int               // 卡项使用权限 1=本人使用 2=家人使用 3= 朋友使用 4= 同事使用，5=任何人使用 6=可转让 7=可共享
// CardLimitsArrange           int                 // 卡项使用权限 特别约定 1= 是 2= 否
// PerformanceGuarantee        int                 // 履约保障 1=银行专用存管 2=履约保证保险 3=银行保函 4=信托 5=担保 6=其它 7=无
// HistoricalTrad              int                 // 历史交易次数 1=第一次  2=第二次 3=第三次及以上  *
// Trad                        int                 // 交易原因 1=乙方推荐  2=甲方自主  3=朋友推荐  4=其它
// SpecialAgreements           string              // 特别约定



// 柳州乾锦智能装备股份有限公司
// 税号：9145 0200 MA5K C280 6U
// 开户行：华夏银行股份有限公司柳州支行
// 账号： 1305 4000 0001 76361
// 开票地址：柳州市柳东新区C区 标准厂房4栋1层
// 电话：0772-2990820
// 收票地址：柳州市鱼峰区雒容镇柳东新区花岭片区粤桂黔产业合作园3栋1层
// 收票人 ：韦良月  181 7672 6005

// busId=74
// cardPrice=2
// realPrice=1
// itemType=8
// servicePeriod=5
// before=true
// cardRange=2
// cardLimits=3,4,2
// specialAgreements=11
// specialAgreements=11
// isAllSingle=false
// singleListsStr=%5B%7B"SingleName":"test","SingleDiscount":2%7D,%7B"SingleName":"addName","SingleDiscount":3
//
//
// realPrice=1&itemType=8&servicePeriod=5&before=true&cardRange=2&cardLimits=3,4,2&specialAgreements=11&cardLimitsArrange=2&isAllSingle=false&singleListsStr=%5B%7B%22SingleName%22:%22test%22,%22SingleDiscount%22:2%7D,