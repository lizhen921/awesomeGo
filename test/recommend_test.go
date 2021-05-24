package test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

//内容中台summary接口
func TestSummaryGid(t *testing.T) {
	//url := "http://power-summary-prod.autohome.com.cn/api/v1/summary/resourceProbuf?version=1.0&itemKeys=755510764302475264,679385066593660928"
	//var summary = &powerpbsummary.SummaryResponse{}
	//c := &http.Client{
	//	Timeout: time.Duration(1000) * time.Millisecond,
	//}
	//resp, err := c.Get(url)
	//body, err := ioutil.ReadAll(resp.Body)
	//err = proto.Unmarshal(body, summary)
	//b, err := json.Marshal(summary)
	//fmt.Println(string(b))
	//fmt.Println(err)
}

func TestUnmarsal(t *testing.T) {
	var mm map[string]map[string]bool
	var dd map[string]map[string]bool
	str := "{\"autohmpg2\":{\"IsBoostDefaultRankModel_ALL\":true,\"IsBoostDefaultRankModel_PRE\":true,\"banBusiness\":false,\"is7StepBuyCar\":true,\"isAgeEERecallRatio\":true,\"isAgeEEV23\":true,\"isAntiTypeRatio\":true,\"isAttentionRecallAdjustV2\":true,\"isAutoDianpingApp\":true,\"isBSNumberV1\":true,\"isBoostLRRankModel_V1\":true,\"isBoostLRRankModel_V2\":true,\"isBoostLSTMRankModel_V1\":true,\"isBoostLSTMRankModel_V2\":true,\"isBoostXGBoostLRRankModel_V2\":true,\"isBusiness\":true,\"isBussinessAiYingXiao\":true,\"isCFMergeV3\":true,\"isCallSummary\":true,\"isCandiByAuthorV4\":true,\"isCardVideoSmallProductLib\":true,\"isCardVideoSmallSeries\":true,\"isCardVideoSmallbiz2rank\":true,\"isCarsoShuoChe\":true,\"isCfMissNewV4\":true,\"isChangAnV1\":true,\"isChejiahaoVrVedio\":true,\"isCloseRemedyCityFilterSeries\":true,\"isCookBookV2ATest\":true,\"isDeepPredictorMonitor\":true,\"isDefaultRankModel_ALL\":true,\"isDianshangRecallBizrank\":true,\"isDmsAndEmbRecallNum\":true,\"isDocVector\":true,\"isFMRecallV3\":true,\"isFastConfAndCompeteSeries\":true,\"isFastsysCandidate\":true,\"isFeedbackOpt\":true,\"isForceType\":true,\"isFpModelStrategyV2\":true,\"isGBDTTagRelevance\":true,\"isGrpcModel\":true,\"isHeyCarApp\":true,\"isHighQualityVideo\":true,\"isHistory7DPvFilter\":true,\"isHotItemRecall\":true,\"isIntentParallel\":true,\"isInterventionOpt\":true,\"isItemAgeEEScore\":true,\"isItemClusterEmbedRecallV3Vearch\":true,\"isItemGroupEE\":true,\"isJustFastsys\":true,\"isLocalModel\":true,\"isLstmItemEmbeddingRecallVearch\":true,\"isMabSeriesRecallV2\":true,\"isMergeSerType\":true,\"isModifyRatioSKYCCW\":true,\"isMultiSeriToufangV4\":true,\"isNegativeFeedbackV5\":true,\"isNewDeviceProfile\":true,\"isNewItemBehaveCluster\":true,\"isNewSmartMarketing\":true,\"isNewVideo\":true,\"isNoBertItemEmbeddingRecall\":true,\"isNoCardVideoSmallChejiahao\":true,\"isNoGraphSeriesEmbRecall\":true,\"isNoLabelTest\":true,\"isNoOPUTypeOptimize2\":true,\"isNoRecallSize\":true,\"isNoSeriNumAdjust\":true,\"isNoSortService\":true,\"isNoStgVisualType\":true,\"isOPUTypeOptimize4\":true,\"isOnlyHourAndEEScore\":true,\"isOriginalTop\":true,\"isPinPaiZhou\":true,\"isProblemSeriesPvCtrl\":true,\"isQuantityControlAndSeparateV3\":true,\"isQuantityControlAndSeparateV4\":true,\"isQuantityControlByVideo\":true,\"isRandomPictureIdea\":true,\"isRandomShow\":true,\"isRankAllEnv\":true,\"isRankBizrankMerge\":true,\"isRankModelFree\":true,\"isRealTimeItemRecall\":true,\"isRealtimeEeFusion\":true,\"isRealtimeProfileRcm\":true,\"isRealtimeSeparate\":true,\"isRealtimeSeriesSeparate\":true,\"isRecallAdjust\":true,\"isRecallMerge\":true,\"isRecallOffline1\":true,\"isRecallSomebyBizSerTag\":true,\"isRecommendReason\":true,\"isRelationMoreContentV2\":true,\"isRelationMoreContentV3\":true,\"isRequestPack\":true,\"isRtICFRecall\":true,\"isScoreWeightByAuthor\":true,\"isScoreWeightByNewitem\":true,\"isScoreWeightBySeries\":true,\"isScoreWeightByTag\":true,\"isScoreWeightByType\":true,\"isSearchWordOpt\":true,\"isSeriDSSMSort\":true,\"isSeriRecallOpt\":true,\"isSeriesArticleRecall\":true,\"isSeriesCard52\":true,\"isSeriesCardGroupV2\":true,\"isSeriesSortWayB\":true,\"isShortWillSeriMerge\":true,\"isSimilarClick\":true,\"isSimilarClickTraveller\":true,\"isSmQuantityControl\":true,\"isSmallvideoCardRecallsometag\":true,\"isSmallvideoPortrayalsModel\":true,\"isSmallvideocardsAdjust\":true,\"isSmallvideouserSvidoSeri\":true,\"isSomeauthorSeriesRecallV2\":true,\"isSpecPK\":true,\"isSpecpkCompSeriV2\":true,\"isSpecpkRankV2\":true,\"isStgPure\":true,\"isTextSeries\":true,\"isTouFangNewIndexV2\":true,\"isTransactionUser\":true,\"isTwoCarCompare\":true,\"isU2IRecall\":true,\"isU2URecall\":true,\"isUCFRecall\":true,\"isUserInterestRecallNumAdjustV2\":true,\"isVideoSmallGroup\":true,\"isVipTouFang\":true,\"isWicfV1\":true,\"isXGBoostExp2\":true,\"isYuanChuangIntervention\":true,\"isYuanChuangTouTiaoInterest\":true,\"isZhongtai700Serve\":true,\"isCallAlgoSummary\":true,\"isUseRecallEngine\":true,\"isRoutineRecengine\":true}}"

	err := json.Unmarshal([]byte(str), &mm)

	expInfoDecoded, _ := base64.StdEncoding.DecodeString(str)

	err = json.Unmarshal(expInfoDecoded, &dd)

	fmt.Println(err)

}
