package tts

// Speaker 说话人
type Speaker string

const (
	// 男童-新新（动漫人物配音，中英双语）
	SPEAKER_MALE_KID_XIN_XIN Speaker = "tts.other.BV050_streaming"

	// 男童-林林（活泼欢快，童声朗读）
	SPEAKER_MALE_KID_LIN_LIN Speaker = "tts.other.BV415_streaming"

	// 男童-昆廷
	SPEAKER_MALE_KID_KUN_TING Speaker = "tts.other.BV153_streaming"

	// 男童-懒小羊（童年经典小伙伴）
	SPEAKER_MALE_KID_LAN_XIAO_YANG Speaker = "tts.other.BV426_streaming"

	// 男童-萌宝（特色音色）
	SPEAKER_MALE_MID_MENG_BAO Speaker = "tts.other.BV051_streaming"

	// 男童-云亦
	SPEAKER_MALE_MID_YUN_YI Speaker = "tts.other.BV061_streaming"

	// 男童-小海绵（卡通动漫角色声音）
	SPEAKER_MALE_MID_XIAO_HAI_MIAN Speaker = "tts.other.BV063_streaming"

	// 男青年-嘻哈歌手
	SPEAKER_MALE_YOUNG_RAP Speaker = "zh_male_rap"

	// 男青年-王松（自然亲切，发音流畅）
	SPEAKER_MALE_YOUNG_WANG_SONG Speaker = "tts.other.BV008_streaming"

	// 男青年-阳天（抖音阳光男声）
	SPEAKER_MALE_YOUNG_YANG_TIAN Speaker = "tts.other.BV056_streaming"

	// 男青年-加布里埃尔
	SPEAKER_MALE_YOUNG_GABRIEL Speaker = "tts.other.BV414_streaming"

	// 男青年-辰韦（广告配音）
	SPEAKER_MALE_YOUNG_CHEN_WEI Speaker = "tts.other.BV401_streaming"

	// 男青年-林舟（知识科普、视频旁白）
	SPEAKER_MALE_YOUNG_LIN_ZHOU Speaker = "tts.other.BV410_streaming"

	// 男青年-温言（就是你常听到的男主小帅啊）
	SPEAKER_MALE_YOUNG_WEN_YAN Speaker = "tts.other.BV411_streaming"

	// 男青年-威廉
	SPEAKER_MALE_YOUNG_WILLIAM Speaker = "tts.other.BV156_streaming"

	// 男青年-陆安（纨绔青年，情绪饱满）
	SPEAKER_MALE_YOUNG_LU_AN Speaker = "tts.other.BV159_streaming"

	// 男青年-世文（沉稳风格，剧情片解说）
	SPEAKER_MALE_YOUNG_SHI_WEN Speaker = "tts.other.BV142_streaming"

	// 男青年-千山（阳光男声，舒适自然）
	SPEAKER_MALE_YOUNG_QIAN_SHAN Speaker = "tts.other.BV143_streaming"

	// 男青年-弗兰克
	SPEAKER_MALE_YOUNG_FRANK Speaker = "tts.other.BV130_streaming"

	// 男青年-小陌（反卷青年，一切随缘）
	SPEAKER_MALE_YOUNG_XIAO_MO Speaker = "tts.other.BV120_streaming"

	// 男青年-皓一（影视解说）
	SPEAKER_MALE_YOUNG_HAO_YI Speaker = "tts.other.BV107_streaming"

	// 男青年-说唱小哥（说唱小哥）
	SPEAKER_MALE_YOUNG_RAPPER_BOY Speaker = "tts.other.BR001_streaming"

	// 男青年-方言-乔治（东北话）
	SPEAKER_MALE_YOUNG_DIALECT_QIAO_ZHI Speaker = "tts.other.BV021_streaming"

	// 男青年-方言-凡达（天津话）
	SPEAKER_MALE_YOUNG_DIALECT_FAN_DA Speaker = "tts.other.BV212_streaming"

	// 男青年-方言-广西表哥（广西普通话）
	SPEAKER_MALE_YOUNG_DIALECT_GX_BIAO_GE Speaker = "tts.other.BV213_streaming"

	// 男青年-方言-牛哥（郑州话）
	SPEAKER_MALE_YOUNG_DIALECT_NIU_GE Speaker = "tts.other.BV214_streaming"

	// 男青年-方言-乔纳森（乔纳森）
	SPEAKER_MALE_YOUNG_DIALECT_JONATHAN Speaker = "tts.other.BV215_streaming"

	// 男青年-方言-乔尔（乔尔）
	SPEAKER_MALE_YOUNG_DIALECT_JOEL Speaker = "tts.other.BV219_streaming"

	// 男青年-方言-阿伟（京腔京韵）
	SPEAKER_MALE_YOUNG_DIALECT_A_WEI Speaker = "tts.other.BV222_streaming"

	// 男青年-方言-轮轮（台湾普通话）
	SPEAKER_MALE_YOUNG_DIALECT_LUN_LUN Speaker = "tts.other.BV227_streaming"

	// 男青年-方言-重庆小伙（川渝话）
	SPEAKER_MALE_YOUNG_DIALECT_CHONG_QING_XIAO_HUO Speaker = "tts.other.BV019_streaming"

	// 男中年-莫远（新闻、资讯播报）
	SPEAKER_MALE_MIDDLE_MO_YUAN Speaker = "tts.other.BV002_streaming"

	// 男中年-本杰明（本杰明）
	SPEAKER_MALE_MIDDLE_BENJAMIN Speaker = "tts.other.BV124_streaming"

	// 男中年-白枫（知性男声，沉着冷静）
	SPEAKER_MALE_MIDDLE_BAI_FENG Speaker = "tts.other.BV419_streaming"

	// 男中年-袁华（有声阅读）
	SPEAKER_MALE_MIDDLE_YUAN_HUA Speaker = "tts.other.BV701_streaming"

	// 男中年-尼古拉斯（尼古拉斯）
	SPEAKER_MALE_MIDDLE_NICHOLAS Speaker = "tts.other.BV144_streaming"

	// 男中年-泽维尔（泽维尔）
	SPEAKER_MALE_MIDDLE_XAVIER Speaker = "tts.other.BV182_streaming"

	// 男中年-布莱恩（布莱恩）
	SPEAKER_MALE_MIDDLE_BRIAN Speaker = "tts.other.BV185_streaming"

	// 男中年-乐为（视频配音）
	SPEAKER_MALE_MIDDLE_LE_WEI Speaker = "tts.other.BV408_streaming"

	// 男中年-派派（动漫人物配音，中英双语）
	SPEAKER_MALE_MIDDLE_PAI_PAI Speaker = "tts.other.BV417_streaming"

	// 男中年-理查德（理查德）
	SPEAKER_MALE_MIDDLE_RICHARD Speaker = "tts.other.BV154_streaming"

	// 男中年-解说陈老（历史故事、神话旁白）
	SPEAKER_MALE_MIDDLE_CHEN_LAO Speaker = "tts.other.BV158_streaming"

	// 男中年-奥利弗（奥利弗）
	SPEAKER_MALE_MIDDLE_OLIVER Speaker = "tts.other.BV148_streaming"

	// 男中年-帕特里克（帕特里克）
	SPEAKER_MALE_MIDDLE_PATRICK Speaker = "tts.other.BV150_streaming"

	// 男中年-艾萨克（艾萨克）
	SPEAKER_MALE_MIDDLE_ISAAC Speaker = "tts.other.BV134_streaming"

	// 男中年-杰克（杰克）
	SPEAKER_MALE_MIDDLE_JACK Speaker = "tts.other.BV139_streaming"

	// 男中年-凯文（凯文）
	SPEAKER_MALE_MIDDLE_KEVIN Speaker = "tts.other.BV140_streaming"

	// 男中年-应晖（高清男声，广泛适用）
	SPEAKER_MALE_MIDDLE_YING_HUI Speaker = "tts.other.BV123_streaming"

	// 男中年-格雷戈里（格雷戈里）
	SPEAKER_MALE_MIDDLE_GREGORY Speaker = "tts.other.BV131_streaming"

	// 男中年-爱德华（爱德华）
	SPEAKER_MALE_MIDDLE_EDWARD Speaker = "tts.other.BV129_streaming"

	// 男中年-查尔斯（查尔斯）
	SPEAKER_MALE_MIDDLE_CHARLES Speaker = "tts.other.BV125_streaming"

	// 男中年-神尊（通用赘婿,解说）
	SPEAKER_MALE_MIDDLE_SHEN_ZUN Speaker = "tts.other.BV119_streaming"

	// 男中年-穆柔（温柔男声，双语教学）
	SPEAKER_MALE_MIDDLE_MU_ROU Speaker = "tts.other.BV033_streaming"

	// 男中年-斯蒂芬（斯蒂芬）
	SPEAKER_MALE_MIDDLE_STEPHEN Speaker = "tts.other.BV110_streaming"

	// 男中年-亚瑟（亚瑟）
	SPEAKER_MALE_MIDDLE_ARTHUR Speaker = "tts.other.BV117_streaming"

	// 男中年-沐明（质朴男声，朴素朴实）
	SPEAKER_MALE_MIDDLE_MU_MING Speaker = "tts.other.BV100_streaming"

	// 男中年-文森特（文森特）
	SPEAKER_MALE_MIDDLE_VINCENT Speaker = "tts.other.BV101_streaming"

	// 男中年-闵晨（儒雅青年,韵味解说feel）
	SPEAKER_MALE_MIDDLE_MIN_CHEN Speaker = "tts.other.BV102_streaming"

	// 男中年-季辰（广告、宣传片配音）
	SPEAKER_MALE_MIDDLE_JI_CHEN Speaker = "tts.other.BV006_streaming"

	// 男中年-于辉（新闻资讯，字正腔圆）
	SPEAKER_MALE_MIDDLE_YU_HUI Speaker = "tts.other.BV012_streaming"

	// 男中年-雅各布（雅各布）
	SPEAKER_MALE_MIDDLE_JACOB Speaker = "tts.other.BV003_streaming"

	// 男中年-高远（开朗男声，自信饱满）
	SPEAKER_MALE_MIDDLE_GAO_YUAN Speaker = "tts.other.BV004_streaming"

	// 男中年-安小航（影视配音）
	SPEAKER_MALE_MIDDLE_AN_XIAO_HANG Speaker = "zh_male_xiaoming"

	// 男中年-安小振（男主播）
	SPEAKER_MALE_MIDDLE_AN_XIAO_ZHEN Speaker = "zh_male_zhubo"

	// 男中年-安小峰（醇厚男声，自信饱满，表情符号省略）
	SPEAKER_MALE_MIDDLE_AN_XIAO_FENG Speaker = "zh_male_chunhou"

	// 男中年-方言-吴江（粤语）
	SPEAKER_MALE_MIDDLE_DIALECT_WU_JIANG Speaker = "tts.other.BV026_streaming"

	// 女童-德博拉（德博拉）
	SPEAKER_FEMALE_DEBORAH Speaker = "tts.other.BV207_streaming"

	// 女童-艾琳（艾琳）
	SPEAKER_FEMALE_ERIN Speaker = "tts.other.BV427_streaming"

	// 女童-安小豆（少儿故事）
	SPEAKER_FEMALE_AN_XIAO_DOU Speaker = "zh_female_story"

	// 女童-克里斯蒂娜（Christina）
	SPEAKER_FEMALE_KID_CHRISTINA Speaker = "tts.other.BV118_streaming"

	// 女童-萝拉（Laura）
	SPEAKER_FEMALE_KID_LAURA Speaker = "tts.other.BV057_streaming"

	// 女青年-安小恩（female young - An Xiao'en）
	SPEAKER_FEMALE_YOUNG_AN_XIAO_EN Speaker = "zh_female_zhubo"

	// 女青年-海莉（Hailey）
	SPEAKER_FEMALE_YOUNG_HAILEY Speaker = "tts.other.BV409_streaming"

	// 女青年-微晴（Weiqing）
	SPEAKER_FEMALE_YOUNG_WEIQING Speaker = "tts.other.BV405_streaming"

	// 女青年-梦欣（Mengxin）
	SPEAKER_FEMALE_YOUNG_MENGXIN Speaker = "tts.other.BV406_streaming"

	// 女青年-思月（Siyue）
	SPEAKER_FEMALE_YOUNG_SI_YUE Speaker = "tts.other.BV412_streaming"

	// 女青年-雅薇（Yarvi）
	SPEAKER_FEMALE_YOUNG_YARVI Speaker = "tts.other.BV418_streaming"

	// 女青年-青青（Qingqing）
	SPEAKER_FEMALE_YOUNG_QINGQING Speaker = "tts.other.BV428_streaming"

	// 女青年-黛玉（Daiyu）
	SPEAKER_FEMALE_YOUNG_DAIYU Speaker = "tts.other.BV425_streaming"

	// 女青年-灿灿（Cancan）
	SPEAKER_FEMALE_YOUNG_CANCAN Speaker = "tts.other.BV700_V2_streaming"

	// 女青年-然月（Ranyue）
	SPEAKER_FEMALE_YOUNG_RANYUE Speaker = "tts.other.BV407_streaming"

	// 女青年-海伦娜（Helena）
	SPEAKER_FEMALE_YOUNG_HELENA Speaker = "tts.other.BV132_streaming"

	// 女青年-艾琳（Erin）
	SPEAKER_FEMALE_YOUNG_ERIN Speaker = "tts.other.BV300_streaming"

	// 女青年-菲奥娜（Fiona）
	SPEAKER_FEMALE_YOUNG_FIONA Speaker = "tts.other.BV301_streaming"
)
