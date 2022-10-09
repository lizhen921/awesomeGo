package main

import (
	"encoding/json"
	"fmt"
	algosummary "git.corpautohome.com/cupid/base_service/protocol/algosummary/protobuf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	Summary_URL = "http://algo-summary-v2-prod.corpautohome.com/api/v1/algoSummary/resourceProtobuf?version=1.0"
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

/*
var Wait sync.WaitGroup
var Counter int = 0

func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}
	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := Counter
		time.Sleep(1 * time.Microsecond)
		value++
		Counter = value
	}
	Wait.Done()
}
*/

/*var i = 0

func g() int {
	i++
	return i
}

func main() {
	_ = g()
}
*/

/*type Config struct {
	a []int
}

func main() {
	cfg := &Config{}

	go func() {
		i := 0
		for {
			i++
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
		}
	}()
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
*/

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

}*/

func main() {
	/*
		rConf := &RecallConf{
			Rway: "rrr",
			Must: Must{
				Rtype:             []string{"t1", "t2"},
				MustIndex:         []string{"uniq_series_id", "uniq_category_name"},
				NotConsiderNumber: make(map[string]bool, 0),
				ProfileData: map[string][]string{
					"uniq_series_id":     {"series---1", "series---2", "series---3"},
					"uniq_category_name": {"category===1", "category===2"},
					"author":             {"author--1", "author--2"}},
			},
		}

		tempIntents := make([]*Intent, 0)
		for _, rtype := range rConf.Must.Rtype {

			intentTypes := make([]*IntentType, 0)
			for i, index := range rConf.Must.MustIndex {
				//按照索引找画像字段
				profileNames := rConf.Must.ProfileData[index]

				//按照画像字段查用户画像
				profieDatas := make([]interface{}, 0)
				for _, pn := range profileNames {
					tempProfileData := GetProfileByName(pn)
					profieDatas = append(profieDatas, tempProfileData...)
				}

				if len(profieDatas) == 0 {
					continue
				}

				//需要考虑单个索引值召回条数
				for _, pd := range profieDatas {
					intentType := &IntentType{
						Index:  index,
						Values: []interface{}{pd},
					}
					intentTypes = append(intentTypes, intentType)

					if i == len(rConf.Must.MustIndex)-1 {
						intent := &Intent{
							RecallName:  "rrr",
							Rtype:       rtype,
							Size:        rConf.PerNum,
							IntentTypes: intentTypes,
						}
						tempIntents = append(tempIntents, intent)
						intentTypes = intentTypes[:len(intentTypes)-1]
					}
				}
			}
		}
		fmt.Println(tempIntents)
	*/

	/*slice1 := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "j", "k"}
	slice2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice3 := []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "J", "K"}

	sets := CartesianProduct(slice1, slice2, slice3)
	fmt.Println(len(sets))
	fmt.Println(sets)
	//
	for _, set := range sets {
		fmt.Println(set)
	}*/

	algo()
}

func CartesianProduct(sets ...[]interface{}) [][]interface{} {
	lens := func(i int) int { return len(sets[i]) }
	product := [][]interface{}{}
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []interface{}
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}

func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

func algo() {

	fmt.Println(time.Now().Format("2006010215"))

	itemkeys := []string{"030120030310039-10131796"}
	httpClient := &http.Client{
		Timeout: time.Duration(6000) * time.Millisecond,
	}
	var res *http.Response
	request_url := Summary_URL + "&itemKeys=" + strings.Join(itemkeys, ",")
	t1 := time.Now().UnixNano()
	res, err := httpClient.Get(request_url)
	if err != nil {
		res, err = httpClient.Get(request_url)
	}
	t2 := time.Now().UnixNano()
	if err != nil {
		fmt.Printf("recall call summary http get error:url[%+v] err[%+v] time[%d ms]", request_url, err, (t2-t1)/1e6)
		return
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("http response data err[%+v]", err)
		return
	}

	series_cnt := make(map[string]int)
	brand_cnt := make(map[int64]int)
	author_cnt := make(map[string]int)
	summaryResponse := &algosummary.SummaryResponse{}
	err = proto.Unmarshal([]byte(content), summaryResponse)
	if err != nil || summaryResponse == nil {
		fmt.Printf("Error When Parse summary pb[%+v]", err)
		return
	} else {
		for pos, item := range summaryResponse.ItemList {
			fmt.Printf("pos:%v, itemKey:%v, biztype:%v, title:%s, seriesid:%v  brand:%v  author:%v source:%v direction:%v original_author:%v content_type:%v\n", pos, item.AlgoBase.ItemKey, item.AlgoBase.BizType, item.AlgoBase.Title, item.AlgoNlp.UniqSeriesId, item.AlgoNlp.UniqBrandsId, item.AlgoBase.Author, item.AlgoBase.Source, item.AlgoBase.Direction, item.AlgoBase.OriginalAuthor, item.AlgoBase.ContentType)
			for _, s := range strings.Split(item.AlgoNlp.UniqSeriesId, ",") {
				series_cnt[s] += 1
			}

			brand_cnt[item.AlgoBase.GetClubClassId()] += 1

		}
	}
	a, _ := json.Marshal(series_cnt)
	fmt.Printf("车系统计： %+v\n", string(a))

	a, _ = json.Marshal(brand_cnt)
	fmt.Printf("品牌统计：%+v\n", string(a))

	a, _ = json.Marshal(author_cnt)
	fmt.Printf("作者统计：%+v\n", string(a))
}
