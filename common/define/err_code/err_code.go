package ec

const (
	Success = 0
	Unknown = 1
	InvalidParam = 2
	ConfigError = 3

	InvalidAccountName = 10
	InvalidAccountLength = 11
	InvalidAccountPassword = 12
	DuplicateAccountName = 13
	AccountNotFound = 14
	WrongPassword = 15
	InvalidToken = 16
	RequestTimeout = 17
	InvalidSign = 18
	AccountBanned = 19

	DuplicatePlayerName = 201
	CreateCharacterFail = 202

	ErrLogin            = 202 // 登录失败
	ErrAuthInvalid      = 203 // 身份认证失败
	LoginOnOtherTerminal    = 204 // 账号已在其他地方登录
	DuplicateCreateCharacter          = 205 // 角色已创建
	ErrRoleNotExists    = 206 // 角色不存在
	ErrIntranetRecharge = 207 // 内部充值失败
	ErrServ             = 208 // 区服未开启
	ErrInvalidName		= 209 // 名字包含无效字符
	ErrEmptyName		= 210 // 名字不能为空

	ErrActivityGetData			= 300
	ErrActivityInvalidStage		= 301
	ErrActivityNotFinish			= 302
	ErrActivityHasGotReward		= 303

	ErrSetRoleNameIdx              = 1003 // 设置角色昵称索引失败
	ErrGetGlobalRole               = 1010 // 获取角色数据失败
	ErrSetGlobalRole               = 1011 // 设置角色数据失败
	ErrGetUserData                 = 1020 // 获取角色关联数据失败
	ErrSetUserData                 = 1021 // 设置角色关联数据失败
	ErrGetUserQuest                = 1028 // 获取任务数据失败
	ErrGetUserHisV                 = 1034 // 获取历史数据失败
	ErrSetUserHisV                 = 1035 // 设置历史数据失败
	ErrGetGlobalFriend             = 1042 // 获取好友数据失败
	ErrSetGlobalFriend             = 1043 // 设置好友数据失败
	ErrGetUserItem                 = 1044 // 获取道具数据失败
	ErrSetUserItem                 = 1045 // 设置道具数据失败
	ErrGetUserMail                 = 1046 // 获取邮件数据失败
	ErrSetUserMail                 = 1047 // 设置邮件数据失败
	ErrGetUserHisKV                = 1060 // 获取历史键值类数据失败
	ErrSetUserHisKV                = 1061 // 设置历史键值类数据失败
	ErrGetUserGuide                = 1068 // 获取引导数据失败
	ErrSetUserGuide                = 1069 // 设置引导数据失败
	ErrGetUserTitle                = 1086 // 获取玩家称号信息失败

	ErrNotEnough                    = 2001 // 不满足所需条件
	ErrNotEnoughItem                = 2002 // 道具不足
	ErrNotEnoughGold                = 2003 // 金币不足
	ErrRepeatGetReward              = 2006 // 不能重复领取奖励
	ErrReduceCurrency               = 2008 // 扣除货币失败
	ErrReduceItem                   = 2009 // 扣除物品失败
	ErrAddItem                      = 2010 // 添加物品失败
	ErrUseItem                      = 2011 // 物品使用失败
	ErrQuestProgressNotEnough       = 2023 // 任务所需进度不足
	ErrFriendNumLimit               = 2043 // 好友数量已达上限
	ErrTargetFriendNumLimit         = 2044 // 对方好友数量已达上限
	ErrTargetFriendApplyNumLimit    = 2045 // 对方好友申请数量已达上限
	ErrNameAlreadyExists            = 2050 // 昵称已存在
	ErrSensitiveCharacters          = 2051 // 存在敏感字符
	ErrFriendAlreadyInApply         = 2062 // 已在好友申请列表中
	ErrFriendAlready                = 2063 // 已是对方好友
	ErrNotFriend                    = 2064 // 已是对方好友
	ErrNotMakeFriendsWithYourself   = 2066 // 不能加自己为好友
	ErrCgpAlreadyUsed               = 2109 // 礼包码已使用
	ErrCgpExpire                    = 2110 // 礼包码已过期
	ErrCgpPackageTypeNotEnough      = 2111 // 当前渠道不符合使用条件
	ErrCgpServNotEnough             = 2112 // 当前区服不符合使用条件
	ErrCgpUseFail                   = 2113 // 礼包码使用失败
	ErrInvalidCgp					= 2114 // 礼包码无效
	ErrRechargeAmountNotEnough		= 2115

	// 芝麻官
	ErrHeroIsExist					= 3000	// 门客已经存在
	ErrHeroMaxLevel					= 3001	// 门客等级达到上限
	ErrHeroNotEnoughMoney			= 3002 	// 银两不足
	ErrNothingToCost					= 3003	// 没物品消耗
	ErrHeroBookIndex					= 3004	// 书籍索引错误
	ErrNotEnoughBookExp				= 3005	// 书籍经验不足
	ErrCostItemData					= 3006	// 花费的道具数据错误
	ErrHeroSkillIndex					= 3007	// 技能索引错误
	ErrNotEnoughHeroSkillExp		= 3008	// 技能经验不足
	ErrHeroSkillMaxLv					= 3009	// 英雄技能达到最大值
	ErrHeroPeerageMaxLv				= 3010	// 英雄爵位达到最大
	ErrHeroNotEnoughRoleExp			= 3011	// 角色经验不足
	ErrRoleMaxLv						= 3012	// 角色达到最大等级
	ErrBossNotPass					= 3013	// boss没通过
	ErrPrelevelNotPass				= 3014	// 前置关卡未通过
	ErrHeroPlayed						= 3015	// 门客已经参加过
	ErrNotHadHero						= 3016 	// 门客不存在
	ErrMiracleNotEnough				= 3017	// 神迹次数不足
	ErrMiracleDataError				= 3018	// 神迹数据错误
	ErrHeroLvupLogic					= 3019	// 英雄升级逻辑错误
	ErrNotPrisoner					= 3020	// 没有囚犯
	ErrSeatNotOpen					= 3021	// 座位未开启
	ErrHeroSited						= 3022	// 座位上有人
	ErrNotHeroSited					= 3023	// 座位上没有人
	ErrNotEnoughTime					= 3024	// 时间不够
	ErrAlreadyHadReadyHero			= 3025	// 已经有英雄在等待
	ErrAlreadyInFight					= 3026	// 已经在战斗
	ErrNotEnoughFreeTimes			= 3027	// 免费次数不足
	ErrFreeCDNotFinish				= 3028	// 免费冷却没结束
	ErrYamenNoHero					= 3029	// 衙门没有英雄
	ErrHeroDead						= 3030	// 门客死亡
	ErrYamenHadEnemy					= 3031	// 衙门已经有敌人
	ErrYamenFightIdxErr				= 3032	// 衙门战斗索引错误
	ErrYamenWinComboRewardExist		= 3033	// 衙门连胜奖励未领取
	ErrYamenNotComboReward			= 3034	// 没有连胜奖励
	ErrYamenNotRandAttribute		= 3035	// 没有额外属性可以选择
	ErrYamenLvTypeError				= 3036	// 等级类型错误
	ErrNotEnoughGreenPill			= 3037	// 绿色药丸不足
	ErrYamenFoeIndex					= 3038	// 衙门的仇人索引错误
	ErrHeroIdErr						= 3039	// 门客id错误
	ErrYamenHeroPlayed				= 3040	// 门客已经出战
	ErrRoleIdErr						= 3041	// rid错误
	ErrYamenFreeTimeNotOver			= 3042	// 免费次数没用完
	ErrNotEnoughHero					= 3043	// 门客数量不足
	ErrNotEnoughItemCount			= 3044	// 道具次数不足
	ErrRankTypeErr					= 3045	// 排行类型错误
	ErrMoreThanLimitNumber			= 3046	// 超过数量上限
	ErrNotEnoughFindEnergy			= 3047	// 寻访体力不足
	ErrNotEnoughFindFate				= 3048	// 寻访运势不足
	ErrHeroNotPlayed					= 3049	// 门客未出战
	ErrHeroUsedItem					= 3050	// 门客已经使用过道具
	ErrYamenExtraAttrGot				= 3051	// 属性已经购买
	ErrChatContentTooLong			= 3052	// 聊天内容太长
	ErrChatRoomIdErr					= 3053	// 聊天房间id错误
	ErrBlacklistRoleExist			= 3054	// 玩家已经存在黑名单
	ErrBlacklistRoleNotExist		= 3055	// 玩家不存在黑名单
	ErrBlacklistTooManyRole			= 3056	// 黑名单人数太多
	ErrPalaceIsReceived				= 3057	// 皇宫俸禄已经领取
	ErrNotEnoughLv					= 3058	// 等级不足
	ErrYamenNotEnemy					= 3059	// 衙门没有敌人
	ErrChatNotInRoom					= 3060	// 不在房间
	ErrHeroFetterId					= 3061	// 光环id错误
	ErrHeroFetterUnlock1				= 3062	// 前置1未解锁
	ErrHeroFetterUnlock2				= 3063	// 前置2未解锁
	ErrForbidChat						= 3064
	ErrTargetCloseChat					= 3065

	ErrBusinessNotEnoughTimes		= 3100	// 收取次数不足
	ErrBusinessGetData				= 3101	// 获取经营数据失败

	ErrLeagueInside					= 3150	// 已经在家族
	ErrLeagueIdErr					= 3151	// id错误
	ErrLeagueNotInside				= 3152	// 没有进入家族
	ErrLeagueNotLeader				= 3153	// 不是盟主
	ErrLeagueAlreadyReqJoin			= 3154	// 已经申请
	ErrLeagueTooMack					= 3155	// 超过数量上限
	ErrLeagueNotReqJoin				= 3156	// 没有申请
	ErrLeagueNotManager				= 3157	// 不是管理者
	ErrLeagueIsLeader					= 3158	// 盟主
	ErrLeagueNotInSame				= 3158	// 不在同一个联盟
	ErrLeaguePermissionDenied		= 3159	// 没有权限
	ErrLeagueInCD						= 3160	// 冷却中
	ErrLeagueJobError					= 3161	// 职位错误
	ErrLeagueCannotSameJob			= 3162	// 不能设置相同职位
	ErrLeagueNotEnoughContribution	= 3163	// 贡献不足
	ErrLeagueAlradyBuild				= 3164	// 已经建造
	ErrLeagueSetSameValue			= 3165	// 设置了相同的值
	ErrLeagueDonotHaveAutoLeague	= 3166	// 没有自动加入的联盟
	ErrLeagueLogoutTooLong			= 3167	// 登出时间太久
	ErrLeagueBossNotInTime			= 3168	// 不在boss时间
	ErrLeagueBossAlreadyOpened		= 3169	// boss已经开启
	ErrLeagueBossCostTypeErr		= 3170	// boss消耗类型错误
	ErrLeagueNotEnoughLevel			= 3171	// 联盟等级不足
	ErrLeagueNotEnoughMoney			= 3172	// 联盟财富不足
	ErrLeagueBossIdError				= 3173	// BOSSid错误
	ErrLeagueBossIsDead				= 3174	// BOSS死亡
	ErrLeagueNotEnoughBuildCount	= 3175	// 建设次数不足
	ErrLeagueNameUsed					= 3176	// 名字已经被使用
	ErrLeagueLeagueTooMack			= 3177	// 联盟申请超过数量上限

	ErrMiracleGetData					= 3200	// 获取神迹数据失败
	ErrMiracleNotEnoughTimes		= 3201	// 神迹次数不足

	ErrRankCountIsZero				= 3250	// 排名人数为0
	ErrRankNotEnoughCount			= 3251	// 次数不足

	ErrAffairsGetData					= 3300	// 获取政务数据失败
	ErrAffairsNotEnoughTimes		= 3301	// 处理政务次数不足
	ErrAffairsCfgErr					= 3302	// 政务配置错误

	ErrBeautyGetData					= 3400	// 获取红颜数据失败
	ErrBeautyGreetingNotEnoughTimes = 3401	// 传唤次数不足
	ErrBeautyCfgErr					= 3402	// 红颜配置错误
	ErrBeautySkillMaxLev				= 3403	// 技能等级上限
	ErrBeautyExpNotEnough			= 3404	// 红颜经验不足
	ErrBeautyGradeMaxNum				= 3405	// 红颜位份满员

	ErrChildrenTrainGetData			= 3500	// 获取培养数据失败
	ErrChildrenTrainMaxSlot			= 3501
	ErrChildrenSetEmptyName			= 3502
	ErrChildrenTrainNotEnoughStamina = 3503
	ErrChildrenExamNotEnoughLev		= 3504

	ErrChildrenAdultGetData			= 3600
	ErrMarriageGetData				= 3601
	ErrMarriageStatusInvalid		= 3602
	ErrMarriageProposeExpire		= 3603

	ErrActivityClosed					= 3701	// 活动关闭
	ErrActivityBossKilled			= 3702	// BOSS死亡
	ErrActivityRewardNotReceived	= 3703	// 奖励未领取
	ErrActivityInAlive				= 3704	// 活动正在进行
	ErrActivityNotReward				= 3705	// 没有奖励领取

	ErrNeigeInvalidStatus			= 3800
	ErrNeigeGetData					= 3801
	ErrNeigeAlreadyDispatch			= 3802
	ErrNeigeInvalidSeat				= 3803
	ErrNeigeNoHeroSelected			= 3804
	ErrNeigeHeroHasBeenDispatched	= 3805
	ErrNeigeAlreadyGetReward		= 3806

	ErrShopTypeErr					= 3900	// 商店类型错误
	ErrShopGoodsIndexErr				= 3901	// 物品索引错误
	ErrShopNotEnoughBuyCount		= 3902	// 购买次数不足

	ErrHanlinInTeam					= 3950	// 已经在队伍中
	ErrHanlinSeatIdxErr					= 3951	// 座位索引错误
	ErrHanlinTeamIdErr				= 3952	// 队伍id错误
	ErrHanlinKilledLimitTime		= 3953	// 打压的限制时间
	ErrHanlinFailed					= 3954	// 输了
	ErrHanlinNotEnoughExp			= 3955	// 经验不够
	ErrHanlinNotInTeam			= 3956	// 没加入队伍
	ErrHanlinAlreadyInShiled			= 3957	// 已经有黄马褂
	ErrHanlinInShiled					= 3958	// 被黄马褂保护

	ErrMongoliaGetData				= 4000
	ErrMongoliaAlreadyAttack		= 4001
	ErrMongoliaNotStart				= 4002
	ErrMongoliaClear					= 4003
	ErrMongoliaAlreadyUseToken		= 4004

	ErrVipLevNotEnough				= 4100
	ErrVipGiftAlreadyGot				= 4101

	ErrNoRecharge						= 4200
	ErrFirstRechargeGiftAlreadyGot	= 4201

	ErrDailySignGetData				= 4300
	ErrDailySignAlreadyGotReward	= 4301

	ErrGetCombatGiftData				= 4400

	ErrDinnerGetData					= 4500
	ErrDinnerNotFinish				= 4501
	ErrDinnerNotExist					= 4502
	ErrDinnerInvalidSeat				= 4503
	ErrDinnerHeroHasBeenDispatched	= 4504
	ErrDinnerOver						= 4505

	ErrComposeLimitTimes				= 4600

	ErrHasBeenInvited					= 4700
	ErrInvalidInviteCode				= 4701
	ErrInviteNumNotEnough				= 4702
	ErrHasNotBeenInvited				= 4703
	ErrInvitePointNotEnough				= 4704
	ErrCanNotInviteSelf					= 4705

	ErrDemonQueenSearchTimesNotEnough	= 4800
	ErrDemonQueenAllMonsterKilled		= 4801
	ErrDemonQueenMonsterDead			= 4802
	ErrDemonQueenMonsterNotFound		= 4803
	ErrDemonQueenHeroAlreadyFight		= 4804
	ErrDemonQueenInspireMaxTimes		= 4805
	ErrDemonQueenNoNeedBattleToken		= 4806
	ErrDemonQueenAlreadyUseBattleToken	= 4807
)
