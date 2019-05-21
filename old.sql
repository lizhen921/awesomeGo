select biz_id,biz_type,chezhan_id,chezhan_title,jump_url,start_time, end_time from rcm_pool.zhiding where biz_type in (1,3,12,13,14,15,35,59,69) and is_delete = 0 and chezhan_id = 104 and chezhan_title is not null and start_time <= current_timestamp and end_time > current_timestamp order by start_time desc
select biz_id,biz_type,title,skis_first,author_icon,author_id,author,author_role,cms_series_ids,nlp_tags_choose,cities,jump_url,img_url2,cms_tags,modify_time,recommend_time ,create_at from rcm_pool.youji_jike where is_delete = 0
select biz_id,biz_type,title,img_url,img_url2,jump_url,cms_series_ids,car_brand_ids,car_brand_names,view_count,modify_time,recommend_time,publish_time,create_at,popularity,refine,media_type,merchant_id from rcm_pool.pinpai_zhanguan where is_delete = 0
select biz_id,biz_type,title,img_url,summary,cms_series_ids,city,city_id,popularity,circle_member,modify_time,publish_time,recommend_time,create_at from rcm_pool.cheyouquan where is_delete = 0 and biz_type = 48
select biz_id,biz_type,img_url,jump_url,modify_time,recommend_time,cms_series_ids,city,min_price,create_at,city_id,title from rcm_pool.2sc_jiage where is_delete = 0
select biz_id,biz_type,img_url,jump_url,modify_time,recommend_time,cms_series_ids,city,min_price,shoufu_price ,create_at from rcm_pool.2sc_fenqi where is_delete = 0
select biz_id,biz_type,img_url,modify_time,recommend_time,publish_time,cms_series_ids,info_obj ,create_at from rcm_pool.xinche_ecom where is_delete = 0
select biz_id,biz_type,img_url,img_url2,modify_time,recommend_time,cms_series_ids,info_obj,rn ,create_at,title from rcm_pool.jinrong_fenqi where is_delete = 0
select biz_id,biz_type,img_url,img_url2,modify_time,recommend_time,cms_series_ids,info_obj ,create_at,is_new_energy,jump_url2,view_count,title_template,price_template,button_text,displaytype,series_score,summary from rcm_pool.dealer_xunjia where is_delete = 0
select biz_id,biz_type,img_url,img_url2,modify_time,recommend_time,cms_series_ids,info_obj,update_at ,create_at from rcm_pool.repair_dianping where is_delete = 0
select biz_id,biz_type,title,cities,modify_time,recommend_time,jump_url,img_url,platform,img_url2 ,create_at,content from rcm_pool.shouyou_lianyun where is_delete = 0
select biz_id,biz_type,sub_biz_id,sub_biz_name,title,jump_url,img_url,modify_time,end_time,cities,cms_series_ids,car_brand_ids,graphic_img_list,update_at ,create_at from rcm_pool.jinrong where is_delete = 0
select biz_id,biz_type,title,img_url,img_url2,jump_url,index_detail,page_index,cms_content_class,reply_count,show_big_img,refine,cms_series_ids,cms_tags,nlp_tags_choose,modify_time,recommend_time,publish_time,v_id,author,author_icon,play_count,duration from rcm_pool.luntan_video where is_delete=0 and date_add(current_timestamp,interval -35 day)<recommend_time
select biz_id, biz_type, title, img_url, img_url2, jump_url, index_detail, page_index, cms_content_class, play_count, show_big_img, refine, cms_series_ids, cms_tags, nlp_tags_choose, modify_time, publish_time, recommend_time, create_at, duration, v_id, reply_count, author, author_icon, nlp_tags_choose2 from rcm_pool.qingshaonian_luntanvideo where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete= 0
select biz_id,biz_type,jump_url,title,cms_tags,reply_count,page_index,index_detail,cms_content_class,content,view_count,cms_series_ids,refine,tuijianforapp,img_url,graphic_img_list,publish_time,modify_time from rcm_pool.zhidao where is_delete=0 and date_add(current_timestamp,interval -45 day)<recommend_time
select biz_id, biz_type, title, img_url, view_count, modify_time, publish_time, recommend_time, create_at, update_at from rcm_pool.7step_buy_car where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete= 0
select biz_id,biz_type,title,img_url,img_url2, is_new_energy, info_obj, cms_series_ids, cms_spec_ids, publish_time, recommend_time, update_at, create_at from rcm_pool.dealer_xunjia_private where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete= 0
select biz_id,biz_type,title,recommend_time,reply_count,view_count,cms_tags,img_url,modify_time,update_at, picitems,create_at,author_id,city_ids from rcm_pool.youji where recommend_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete =0
select biz_id,biz_type,title,author,publish_time,img_url2 as img_url,modify_time,update_time,recommend_time,summary,reply_count,view_count ,img_url2,pk_id,pk_obj,vote_id,vote_obj, create_at from rcm_pool.topic where recommend_time<= current_timestamp and date_add(current_timestamp,interval -30 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete= 0 and pk_obj is not NULL
select biz_id,biz_type,img_url2,jump_url,modify_time,recommend_time,title,city,cms_tags,platform,create_at from rcm_pool.chefuwu where recommend_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete =0
select biz_id,biz_type,title,img_url, img_url2,cms_series_ids,cms_spec_ids,city_id,city,modify_time,view_count,publish_time,update_at,recommend_time,jump_url ,create_at from rcm_pool.chezhu_jiage where recommend_time<= current_timestamp and date_add(current_timestamp,interval -30 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete= 0 and biz_type = 36
select biz_id,biz_type,title,img_url, img_url2,car_brand_ids,car_brand_names,city,modify_time,publish_time,update_at,recommend_time,jump_url ,create_at,cms_series_ids from rcm_pool.tuangou where recommend_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete= 0 and biz_type = 30
select biz_id,biz_type,img_url,img_url2,jump_url,title,city,publish_time,recommend_time,modify_time,start_time,end_time,cms_car_price  ,create_at,cms_series_ids,cms_spec_ids from rcm_pool.dealer where recommend_time<= current_timestamp and start_time <= current_timestamp and end_time >= current_timestamp and date_add(current_timestamp,interval -30 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete= 0
select biz_id,biz_type,img_url,img_url2,title,city,publish_time,recommend_time,modify_time,reply_count,jump_url,create_at from rcm_pool.2sc_article where recommend_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete= 0
select biz_id,biz_type,title,recommend_time,reply_count,view_count,cms_series_ids,cms_tags,cms_content_class,img_url,index_detail,modify_time,graphic_img_list,subject_id,is_close_comment,skis_first,img_url2,duration,author,show_big_img,  content, publish_time, author_id, v_id, youku_id ,create_at from rcm_pool.video_zhanwai where recommend_time<= current_timestamp and recommend_time is not null and recommend_time!='' and is_delete= 0
select biz_id,biz_type,title,img_url,img_url2,cms_tags,spec_kb_score,spec_feature,view_count,reply_count,cms_series_ids,cms_series_names,cms_spec_ids, cms_spec_names,jump_url, modify_time, recommend_time,update_at  ,create_at from rcm_pool.koubei_pinglun where recommend_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete= 0 and biz_type = 43
select biz_id,biz_type,title,jump_url,img_url,modify_time,publish_time,cms_series_ids,cms_spec_ids,view_count,popularity,recommend_time,update_at ,create_at from rcm_pool.vr_pano_ext where is_delete = 0 and date_add(current_timestamp,interval -35 day) < recommend_time
select biz_id,biz_type,title,jump_url,cms_series_ids,cms_spec_ids,view_count,publish_time,update_at,create_at,graphic_img_list from rcm_pool.dealer_zhanting where publish_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < publish_time and is_delete= 0 and biz_type = 61
select biz_id,biz_type,title,content,img_url,img_url2,jump_url,graphic_img_list,author_id,view_count,reply_count,refine,is_video, modify_time,recommend_time,cms_tags,page_index,index_detail,cms_content_class,create_at from rcm_pool.qingshaonian_jinghua where date_add(current_timestamp,interval -25 day) < recommend_time and is_delete = 0
select biz_id, biz_type, title, img_url, img_url2, jump_url, index_detail, page_index, cms_content_class, play_count, show_big_img, refine, cms_series_ids, cms_tags, nlp_tags_choose, modify_time, publish_time, recommend_time, create_at, duration, v_id, reply_count, author, author_icon from rcm_pool.qingshaonian_luntanvideo where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete= 0
select biz_id,biz_type,title,author_icon,author,author_id,cms_series_ids,modify_time,recommend_time,update_at,jump_url,img_url3,img_url2,cms_tags,show_big_img,popularity,reply_count,view_count,content,live_line,subject_id,summary,share_url,geo_point,img_url,duration,video_width,video_height,img_url4,create_at,live_albumname,author_role from rcm_pool.video_small where is_delete = 0 and title != "" and ((img_url2 is not null and img_url2 != "") or (img_url4 is not null and img_url4 != "")) and date_add(current_timestamp,interval -35 day) < recommend_time and (refine=1 or date_add(current_timestamp,interval -24 hour) < recommend_time)
select biz_id,biz_type,title,jump_url,img_url,modify_time,cms_series_ids,graphic_img_list,reply_count,recommend_time,update_at ,create_at from rcm_pool.pic_chezhan where is_delete = 0 and date_add(current_timestamp,interval -35 day) < recommend_time
select biz_id,biz_type,live_line,jump_url,jump_url2,img_url,title,slogan,cms_series_ids,cms_spec_ids,dealer_id,cities,sale_price,suggest_price,dikouquan_price,dikouquanyi_price,yigoumai_num,is_fenqi,modify_time,publish_time,recommend_time,create_at,update_at,shoufu_price,source,is_mengzhu,media_type from rcm_pool.xinche_ecom_channel where is_delete=0 and date_add(current_timestamp,interval -35 day)<recommend_time and (suggest_price-sale_price) > 12000 and sale_price > 10000
select biz_id,biz_type,jump_url,img_url,title,cms_series_ids,cms_spec_ids,cms_car_price,sale_price,meter_mileage,city_ids,modify_time,publish_time,create_at,update_at,cheyuan_score,price_template,title_template from rcm_pool.2sc_cheyuan WHERE is_delete = 0 AND jump_listurl2 IS NOT NULL AND jump_listurl2 != '' AND title_template IS NOT NULL AND price_template IS NOT NULL
select biz_id,biz_type,jump_url,title,author_icon,summary,cms_series_ids,city_id,cms_tags,circle_member,refine,modify_time,publish_time,recommend_time,img_url,img_url2,graphic_img_list, create_at from rcm_pool.cheyouquan_reliao where is_delete=0
select biz_id,biz_type,title,img_url2,jump_url,publish_time,update_at,create_at,img_url from rcm_pool.qichecheng where is_delete = 0
select biz_id,biz_type,title,img_url2,jump_url,publish_time,update_at,create_at,img_url from rcm_pool.xiaoyouxi where is_delete = 0
select biz_id,biz_type,jump_url,title,view_count,img_url,img_url2,modify_time,recommend_time, create_at from rcm_pool.travel_tips where is_delete=0 and date_add(current_timestamp,interval -45 day)<recommend_time
select biz_id,biz_type,title,img_url,img_url2,jump_url,cms_series_ids,cms_series_names,spec_kb_score,reply_count,refine,min_price,max_price,cms_level_ids,cms_level_names,popularity,publish_time,update_at,create_at from rcm_pool.koubei_rank where publish_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < publish_time and is_delete= 0 and biz_type = 80
select biz_id,biz_type,title,img_url,img_url2,jump_url,cms_series_ids,refine,reply_count,content,publish_time,update_at,create_at from rcm_pool.koubei_danxiang where publish_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < publish_time and is_delete= 0 and biz_type = 81 and (refine=1 or refine=2)
select biz_id,biz_type,title,jump_url,img_url2,graphic_img_list,refine,cms_series_ids,author_id,author,view_count,info_obj,publish_time,create_at,update_at,img_url from rcm_pool.guanzhu_topic where publish_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < publish_time and is_delete= 0 and biz_type = 86
select biz_id,biz_type,title,jump_url,img_url,cms_tags,cms_series_ids,view_count,reply_count,publish_time,create_at,update_at from rcm_pool.guanzhu_topic_list where publish_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < publish_time and is_delete= 0 and biz_type = 87
select biz_type, biz_id, title, img_url2, cms_series_ids, cms_tags, graphic_img_list, show_big_img, refine, media_type, jump_url, publish_time, recommend_time, update_at, create_at from rcm_pool.chejiahao_zhuanti where date_add(current_timestamp,interval -35 day)<recommend_time and refine=1 and is_delete=0
select biz_id, biz_type, title, jump_url, img_url, img_url2, cms_series_ids, cities, view_count, publish_time, recommend_time, create_at, update_at from rcm_pool.chezhu_jiage_ai where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete= 0
select biz_id, biz_type, title, img_url, view_count, jump_url2, publish_time, recommend_time, modify_time, create_at, update_at, cms_series_ids, refine from rcm_pool.series_guide where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete= 0
select biz_id, biz_type, title, img_url, jump_url, content, index_detail, cms_tags, cms_series_ids, recommend_time, nlp_tags, nlp_series, nlp_keyword from rcm_pool.jiajia_xiaomi  where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete= 0
select biz_id, biz_type, popularity, graphic_img_list, jump_url, img_url, title, cms_tags, media_type, publish_time, recommend_time, create_at, update_at,cms_series_ids,refine from rcm_pool.car_channel where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete = 0 and media_type = 1
select biz_id, biz_type, graphic_img_list, jump_url, title, cms_tags, reply_count, media_type, publish_time, recommend_time, create_at, update_at,cms_series_ids,refine,rn from rcm_pool.wenda_zonghe where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete = 0
select biz_type,biz_id,title,jump_url,jump_url2,publish_time,view_count,img_url2,author_icon,show_big_img,recommend_time,update_at,create_at from rcm_pool.heicar_tuiguang where is_delete=0
select biz_type, biz_id, title, author, author_id , img_url, summary, reply_count, view_count, content , cms_series_ids, cms_tags, category_tags, graphic_img_list, author_icon , cms_series_names, car_brand_ids, car_brand_names, jump_url, cities , jump_url2, nlp_tags_choose2, publish_time, recommend_time, create_at, update_at from rcm_pool.youchuang WHERE date_add(current_timestamp, INTERVAL -35 DAY) < recommend_time AND is_delete = 0 AND (refine = 1 OR refine = 3) AND biz_type = 206
select biz_type,biz_id,title,jump_url,jump_url2,publish_time,view_count,img_url2,author_icon,show_big_img,recommend_time,update_at ,create_at from rcm_pool.car_dianping where is_delete=0
select biz_type, biz_id, author_icon, author, author_id, editor_id, live_line, summary, jump_url, publish_time, modify_time, recommend_time from rcm_pool.guanzhu_card_hongren where is_delete = 0
select biz_id, biz_type, graphic_img_list, jump_url, img_url, title,reply_count,  publish_time, recommend_time, create_at, update_at,cms_seires_ids from rcm_pool.chefuwu_good_comments where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete = 0
select biz_type,biz_id,title,jump_url,img_url2,graphic_img_list,cms_series_ids,cms_spec_ids,city_ids,merchant_id,view_count,cms_tags,show_big_img,publish_time,create_at,update_at from rcm_pool.dealer_pinpai where is_delete = 0
select biz_type, biz_id, title, jump_url, img_url2 , abstract, author_icon, publish_time, recommend_time, create_at, update_at from rcm_pool.pinpaizhou where is_delete = 0 AND date_add(current_timestamp, INTERVAL -35 DAY)
select biz_id,biz_type,title,recommend_time,reply_count,view_count,cms_tags,img_url,modify_time,update_at, picitems, create_at, author_id, city_ids, source, jump_url2 from rcm_pool.youji_ogc where recommend_time<= current_timestamp and date_add(current_timestamp,interval -35 day) < recommend_time and recommend_time is not null and recommend_time!='' and is_delete =0 and biz_type = 212
select biz_id, biz_type, title, img_url, cms_series_ids, view_count, jump_url, index_detail, publish_time, recommend_time, create_at, update_at from rcm_pool.canjia_zhushou where is_delete = 0 and date_add(current_timestamp,interval -35 day) < recommend_time and biz_type = 213
select biz_id, biz_type, title, modify_time, recommend_time, publish_time, img_url, graphic_img_list, cms_series_ids, cities, view_count, jump_url,index_detail, cms_tags,nlp_tags_choose2, create_at, update_at, reply_count from rcm_pool.jinrong_article where date_add(current_timestamp,interval -35 day) < recommend_time and is_delete = 0
select biz_id, biz_type, title, img_url, index_detail, abstract, graphic_img_list, cms_series_ids, cms_tags, jump_url, jump_url2, publish_time, recommend_time, update_at, create_at from rcm_pool.chejiahao_fantuan where biz_type=215 and is_delete = 0 and date_add(current_timestamp,interval -35 day) < recommend_time
select biz_id, biz_type, img_url2, title, jump_url, view_count, publish_time, create_at, update_at from rcm_pool.cms_operation_card where is_delete = 0 and biz_type = 10001
