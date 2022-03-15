package main

import (
	"fmt"
	"strconv"
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

	s := "012"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	nums := make([]int, 0, 5)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)

	fmt.Println(nums[:0])
	fmt.Println(nums[1:])
	fmt.Printf("%p\n", nums)
	temp := append(nums[:0], nums[1:]...)
	fmt.Printf("%p\n", nums)
	fmt.Printf("%p\n", temp)
	fmt.Println(temp)
	fmt.Println(nums)

	str := "0"
	for i := 1; i < 200; i++ {
		str = fmt.Sprintf("%s-%s-%d", str, ",", i)
	}

}
