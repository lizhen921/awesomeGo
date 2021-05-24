package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/**
跳转中心
*/
/*func main() {
	//加载配置
	conf.LoadConfig()
	grpc-gw-server := web.SetupRouter()
	//Listen and Server in 0.0.0.0:8080
	grpc-gw-server.Run(":8080")
}
*/

func main() {
	//routerJson, err := ioutil.ReadFile("router.json")
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	return
	//}
	//sj, err := simplejson.NewJson(routerJson)
	//routers, err := sj.Get("routers").Array()
	//for _, v := range routers {
	//	vJson := v.(map[string]interface{})
	//	if _, ok := vJson["pids"]; ok {
	//		fmt.Printf("%s:%s\n", vJson["name"], vJson["pids"])
	//	}
	//}
	//bRouter, err := routers.MarshalJSON()
	//br1 := bufio.NewReader(routerJson)
	//br1.

	c := &http.Client{
		Timeout: time.Duration(1000) * time.Millisecond,
	}

	URL := "http://nlpaip.corpautohome.com/dm_nlp/v1/get_event_by_bs"

	param := simplejson.New()
	param.Set("uuid", "s122")
	param.Set("series", "途观X")
	paramByte, err := param.MarshalJSON()
	//fmt.Println("---", string(paramByte), err)

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(paramByte))

	input := "cupid-engineBlnKH8ox:DaSDJeCtLqCQ"

	auth := base64.StdEncoding.EncodeToString([]byte(input))
	//fmt.Println(auth)

	req.Header.Add("Authorization", "Basic "+auth)
	res, err := c.Do(req)
	fmt.Println(err)
	//fmt.Println(res)
	body, err := ioutil.ReadAll(res.Body)
	bodyJson, err := simplejson.NewJson(body)
	returnCode, err := bodyJson.Get("returncode").Int64()

	result := bodyJson.Get("result").MustArray()
	if returnCode != 0 {

	}

	for k, v := range result {
		fmt.Printf("== %d, %v\n", k, v)
	}

	fmt.Println(string(body))

	//brand, series := h.getBrandsAndSeries(qc)
	//param := simplejson.New()
	//param.Set("uuid", qc.SrtInfoIn.SrtReqInfoSearch.StrLogUuid)

	//if series != nil || len(series) > 0 {
	//	param.Set("series", series)
	//} else if brand != nil || len(brand) > 0 {
	//	param.Set("brand", brand)
	//} else {
	//	return nil
	//}
	//paramByte, err := param.MarshalJSON()
	//if err != nil {
	//	h.srtLog4sys.Warn(" MarshalJSON err:[%v]", err)
	//	return err
	//}
	req, err = http.NewRequest("POST", URL, bytes.NewBuffer(paramByte))
	auth = base64.StdEncoding.EncodeToString([]byte("cupid-engineBlnKH8ox:DaSDJeCtLqCQ"))
	req.Header.Add("Authorization", "Basic "+auth)
	res, err = c.Do(req)
	if err != nil {
		//h.srtLog4sys.Warn("POST err:[%v], res:[%v], url:[%s]", err, res, URL)
		//return err
	}
	body, err = ioutil.ReadAll(res.Body)

	tianShuRes := &TianShuRes{}
	err = json.Unmarshal(body, &tianShuRes)
	if err != nil || tianShuRes.ReturnCode != 0 {

	}

	//取第一个物料
	if len(tianShuRes.Result) > 0 {
		for _, v := range tianShuRes.Result[0] {
			if len(v) > 0 {
				bizType, err := strconv.Atoi(v[0][0])
				bizId, err := strconv.Atoi(v[0][1])
				fmt.Println(bizType, bizId, err)
			}
		}
	}

	//s := "如今国内的中级车市场的竞争异常激烈，而雅阁、帕萨特、凯美瑞间的夺冠争夺尤甚。在去年中保研碰撞测试后，帕萨特已“翻车”，今年以来销量一直萎靡不振。而一向比较低调的凯美瑞却在2、3月份连续两个月夺得中级车销量冠军。雅阁4月份重新发力，销量大幅上涨，目前来看夺冠已成定局。  随着国内疫情得以有效控制，国内车市在4月份逐渐恢复正常。压抑了两个多月，本田雅阁销量在4月份迎来大爆发。4月份雅阁销量达1.89万辆，环比增长78.1%。如果不出意外，4月份中级车销量冠军或仍将是本田雅阁。虽然帕萨特在3月份通过中汽研碰撞测试又“补考”了一次，而且成绩非常不错，夺得五星的好成绩，但仍无法抹平帕萨特中保研碰撞测试A柱断裂在消费者心中的“阴影”。而且貌似帕萨特的这波操作，越抹越黑，不仅使得消费者对帕萨特的安全性仍存在很大的质疑，而且更坚实了中汽研“五星批发部”的称号。因此帕萨特短期内销量难以出现反弹，夺冠也成为了小概率事件。  而4月份广汽丰田销量已公布，其中凯美瑞4月份销量为1.54万辆，虽同比环比均实现正增长，但凯美瑞4月份销量仍落后于雅阁。除此之外，其他中级车对雅阁不会造成很大的威胁。因此4月份中级车销量冠军十有八九为本田雅阁。  本田雅阁销量的大幅上涨，除了客观环境因素之外，强大的产品力是关键。十代雅阁年轻运动的外观设计，迷倒了很多年轻消费者。尤其低矮的车身姿态以及动感尾部线条，营造出更加动感的效果，非常符合年轻人的口味。而4893mm的车身长度及2830mm的轴距，保证了雅阁具备比较大的内部空间。  此外，发动机一直是本田的强项。雅阁燃油版搭载了1.5T（高低功率版）地球梦发动机，最大功率分别为130kw、143kw。同时配备CVT变速箱，保证了动力输出的平顺性。而雅阁在动力充沛性及燃油经济性方面表现很优秀，这也是消费者在众多中级车中选择雅阁的一个重要因素。除此之外，雅阁混动车型也因较低的油耗，深受消费者的喜爱。4月份雅阁混动车型销量达4586辆，环比增长139.9%，占雅阁品牌整体销量的24.2%。  而在绝大部分品牌都在降价保市场的情况下，本田雅阁也顺势给出了不小的优惠。据了解，在一些地区终端市场雅阁终端优惠达2万元。这在一定程度上也刺激了消费者的购买。写在最后随着国内疫情的好转，4月份本田雅阁销量强势回归，重新夺回中级轿车销量冠军的概率很大。有强大的产品力做支撑，雅阁夺冠显得更为轻松。与之形成对比的是帕萨特在中保研碰撞测试中的“翻船”，安全性备受质疑，产品力大打折扣，销量大幅下滑也就不难理解了。（本文由文武车道新媒体工作室原创出品，转载请注明出处：文武车道，本文作者：夏沐）"
	//nameRune := []rune(s)
	//nameRune[0:4]
}

type TianShuRes struct {
	ReturnCode int
	Message    string
	Result     []map[string][][]string
}

//

/**
处理字符串sql
*/
/*func main() {
	file1, err := os.Open("new.sql")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer file1.Close()

	br1 := bufio.NewReader(file1)
	for {
		a, _, c := br1.ReadLine()
		sqlStr := string(a)
		index := strings.Index(sqlStr, "from")
		if index != -1 {
			//替换
			subSql := sqlStr[index:]
			//fmt.Println(subSql)
			allSql := "select count(1) " + subSql
			fmt.Println(allSql)
			//查询旧sql
			//查询新sql
		}
		if c == io.EOF {
			break
		}
	}
}*/

/**
处理sql
*/
/*func main() {
	var tb_richmedia_pool_ref = []string{"2sc_article", "ah100", "article", "car_channel", "car_dianping", "blog", "cheyouquan", "cheyouquan_reliao", "chezhu_jiage", "chezhu_jiage_ai", "dealer", "dealer_ai", "jiajia_xiaomi", "jinrong_article", "koubei", "koubei_danxiang", "koubei_pinglun", "koubei_rank", "weixin", "wenda_zonghe", "xiaoyouxi", "youchuang", "youji", "youji_jike", "youji_ogc", "zhibo", "zhidao", "dealer_shijia", "product_word", "chejiahao_fantuan", "cms_operation_card", "content", "expertblog", "series_guide", "kuaibao", "luntan_video", "qichecheng", "qingshaonian", "qingshaonian_jinghua", "qingshaonian_luntanvideo", "shouyou_lianyun", "tiezi", "tiezi_jingxuan", "tiezi_external", "topic", "toutiao", "travel_tips", "video", "video_small", "video_zhanwai", "waimei"}
	var tb_car_pool_ref = []string{"2sc_cheyuan", "2sc_fenqi", "2sc_jiage", "chefuwu", "chefuwu_good_comments", "dealer_price", "dealer_xunjia", "dealer_xunjia_private", "jiayouka", "new_energy", "pic", "pic_chezhan", "repair_dianping", "vr", "vr_ext", "vr_pano", "vr_pano_ext", "xinche_ecom", "xinche_ecom_channel", "xinche_ecom_normal", "dealer_zhanting_shop", "com_xinche_card", "canjia_zhushou", "series", "series_base", "serues_info"}
	var tb_active_pool_ref = []string{"chejiahao_hongren", "guanzhu_card_hongren", "guanzhu_topic", "guanzhu_topic_list", "show_user"}
	var tb_user_pool_ref = []string{"7step_buy_car", "7step_buycar_msite", "chejiahao_zhuanti", "dealer_pinpai", "dealer_zhanting", "jinrong", "jinrong_fenqi", "marketing_ideas", "pinpai_zhanguan", "pinpaizhou", "tuangou", "zhiding", "recommend_nlp_info", "heicar_tuiguang", "car_dianping"}

	b, err := ioutil.ReadFile("old.sql")
	if err != nil {
		fmt.Print(err)
	}
	oldSql := string(b)
	//fmt.Println(oldSql)
	for _, v := range tb_richmedia_pool_ref {
		oldSql = strings.Replace(oldSql, "from rcm_pool."+v+" ", "from rcm_pool.tb_richmedia_pool ", -1)
	}
	for _, v := range tb_car_pool_ref {
		oldSql = strings.Replace(oldSql, "from rcm_pool."+v+" ", "from rcm_pool.tb_car_pool ", -1)
	}
	for _, v := range tb_active_pool_ref {
		oldSql = strings.Replace(oldSql, "from rcm_pool."+v+" ", "from rcm_pool.tb_active_pool ", -1)
	}
	for _, v := range tb_user_pool_ref {
		oldSql = strings.Replace(oldSql, "from rcm_pool."+v+" ", "from rcm_pool.tb_user_pool ", -1)
	}

	d1 := []byte(oldSql)
	err = ioutil.WriteFile("new.sql", d1, 0644)
	if err != nil {
		fmt.Errorf("写异常")
	}

	fi, err := os.Open("old.sql")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a) + " limit 10")
	}

	fmt.Println("/n/n")

	file1, err := os.Open("new.sql")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer file1.Close()

	br1 := bufio.NewReader(file1)
	for {
		a, _, c := br1.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a) + " limit 10")
	}

}
*/
