package baseProtocol

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

type BaseRequest struct {
	BaseUrl      string `json:"baseUrl"`
	BaseSendTime int64  `json:"baseSendTime"`
	BaseScgToken string `json:"baseScgToken"`
	MemberId     string `json:"memberId"`
}

type RetStatus struct {
	StatusCode int    `json:"StatusCode" qs:"StatusCode"`
	StatusMsg  string `json:"StatusMsg" qs:"StatusMsg"`
	SystemTime int64  `json:"SystemTime" qs:"SystemTime"`
	CheckCode  int    `json:"CheckCode" qs:"CheckCode"`
}

type marshal RetStatus

type GroupRetStatusCode struct {
	CheckCode int    `json:"StatusCode,omitempty"`
	StatusMsg string `json:"StatusMsg"`
}

// BaseResponse : 基礎回應資料結構
type BaseResponse struct {
	RetStatus RetStatus `json:"retStatus"`
}

const Success = 10000 // Success : 正常返回

// UnknownStatusCode : 未知代碼(可能未宣告const或map)
const UnknownStatusCode = "Unknown Status Code"

const ( // Notification
	Notification      = 3000 // Notification : 基礎號段
	Notification_3001 = 3001 // JSON Parse Error
	Notification_3002 = 3002 // Validate Request Error
	Notification_3003 = 3003 // Unknown Error
	Notification_3004 = 3004 // Rewards Already Exchange
)

const ( // RewardDistribution
	RewardDistribution                             = 3200 // RewardDistribution基礎號段
	RewardDistribution_InvalidInput_3201           = 3201 // 無效參數
	RewardDistribution_ResourceNotFound_3202       = 3202 // 資源不存在
	RewardDistribution_RewardHasClaimed_3203       = 3203 // 獎勵已領取
	RewardDistribution_RewardHasClaimedInFull_3204 = 3204 // 獎勵已被領完
	RewardDistribution_Unknown_3205                = 3205 // 未知錯誤
	RewardDistribution_ExternalServiceError_3206   = 3206 // 外部服務錯誤
)

const ( // SortingHat
	SortingHat                           = 3300 // SortingHat基礎號段
	SortingHat_InvalidInput_3301         = 3301 // 無效參數
	SortingHat_ResourceNotFound_3302     = 3302 // 資源不存在
	SortingHat_Unknown_3303              = 3303 // 未知錯誤
	SortingHat_InternalServiceError_3304 = 3304 // 內部服務錯誤
	SortingHat_ExternalServiceError_3305 = 3305 // 外部服務錯誤
)

const ( // KickCenter
	KickCenter       = 10200 // TokenGuard : 基礎號段
	KickCenter_10201 = 10201 // KickCenter_10201 : 設定 Issuer 發生錯誤
	KickCenter_10202 = 10202 // KickCenter_10202 : 取得 tokenguard 位置發生問題
	KickCenter_10203 = 10203 // KickCenter_10203 : 傳送資料至 tokenguard 發生錯誤

	KickCenter_10204 = 10204 // KickCenter_10204 : 傳送資料至 tokenguard 發生錯誤
	KickCenter_10205 = 10205 // KickCenter_10205 : 傳送資料至 tokenguard 發生錯誤

	KickCenter_10206 = 10206 // KickCenter_10206 : 註冊 tokenguard hostname 至 redis 發生錯誤
	KickCenter_10207 = 10207 // KickCenter_10207 : 取得 Issuer 發生錯誤
	KickCenter_10208 = 10208 // KickCenter_10208 : 從 redis 刪除 restokenguard hostname 發生錯誤
	KickCenter_10209 = 10209 // KickCenter_10209 : 從 redis 取得 KickMemberList 發生錯誤
	KickCenter_10210 = 10210 // KickCenter_10210 : 取得 Kick info 發生錯誤
	KickCenter_10211 = 10211 // KickCenter_10211 : 踢除玩家發生問題
	KickCenter_10212 = 10212 // KickCenter_10212 : 排隊中
)

const ( // TokenGuard
	TokenGuard       = 10100 // TokenGuard : 基礎號段
	TokenGuard_10101 = 10101 // TokenGuard_10101 : Login Failed
	TokenGuard_10102 = 10102 // TokenGuard_10102 : Member Not Exist
	TokenGuard_10103 = 10103 // TokenGuard_10103 : Password is invalid
	TokenGuard_10104 = 10104 // TokenGuard_10104 : Token Invalid
	TokenGuard_10105 = 10105 // TokenGuard_10105 : Token Error
	TokenGuard_10106 = 10106 // TokenGuard_10106 : Token Create Error
	TokenGuard_10107 = 10107 // TokenGuard_10107 : FB Token Expired
	TokenGuard_10108 = 10108 // TokenGuard_10108 : FB Token Invalid
	TokenGuard_10109 = 10109 // TokenGuard_10109 : Create Account Error
	TokenGuard_10110 = 10110 // TokenGuard_10110 : Account Not Active
	TokenGuard_10111 = 10111 // TokenGuard_10111 : Verification Time Out
	TokenGuard_10112 = 10112 // TokenGuard_10112 : Change Password Fail
	TokenGuard_10113 = 10113 // TokenGuard_10113 : Verification Code Error
	TokenGuard_10114 = 10114 // TokenGuard_10114 : CellPhone Bind Fail
	TokenGuard_10115 = 10115 // TokenGuard_10115 : Member Not Bind CellPhone
	TokenGuard_10116 = 10116 // TokenGuard_10116 : No CellPhone Bind Authority
	TokenGuard_10117 = 10117 // TokenGuard_10117 : No Forget Password Authorit
	TokenGuard_10118 = 10118 // TokenGuard_10118 : Invalid Base Url
	TokenGuard_10119 = 10119 // TokenGuard_10119 : Change Password Has Been
	TokenGuard_10120 = 10120 // TokenGuard_10120 : No LogIn Authority
	TokenGuard_10121 = 10121 // TokenGuard_10121 : Line Token Invalid
	TokenGuard_10122 = 10122 // TokenGuard_10122 : LoginDevice Not Trusted
	TokenGuard_10123 = 10123 // TokenGuard_10123 : Device Bind Fail
	TokenGuard_10124 = 10124 // TokenGuard_10124 : Device Already Exists
	TokenGuard_10125 = 10125 // TokenGuard_10125 : Check Sign Error
	TokenGuard_10126 = 10126 // TokenGuard_10126 : JWT Expired
	TokenGuard_10127 = 10127 // TokenGuard_10127 : JWT Unknown Error
	TokenGuard_10128 = 10128 // TokenGuard_10128 : Base Request Compare Error
	TokenGuard_10129 = 10129 // TokenGuard_10129 : Login from another Device
	TokenGuard_10130 = 10130 // TokenGuard_10130 : Server is Maintenance
	TokenGuard_10131 = 10131 // TokenGuard_10131 : Set Maintenance Error
	TokenGuard_10132 = 10132 // TokenGuard_10132 : Set kick member error
	TokenGuard_10133 = 10133 // TokenGuard_10133 : Bind Account BodyParser Error
	TokenGuard_10134 = 10134 // TokenGuard_10134 : Bind Account CheckDeviceInfo Error
	TokenGuard_10135 = 10135 // TokenGuard_10135 : Only Support Account Bind
	TokenGuard_10136 = 10136 // TokenGuard_10136 : Connect To bind Service Error
	TokenGuard_10137 = 10137 // TokenGuard_10137 : Gen RefreshJWT Error
	TokenGuard_10138 = 10138 // TokenGuard_10138 : kick by backend
	TokenGuard_10139 = 10139 // TokenGuard_10139 : 帳號註銷
	TokenGuard_10140 = 10140 // TokenGuard_10140 : No SafeBox Authority
	TokenGuard_10141 = 10141 // TokenGuard_10141 : Request Time Check Fail
	TokenGuard_10142 = 10142 // get retry error
	TokenGuard_10143 = 10143 // set retry error
	TokenGuard_10144 = 10144 // lock error
	TokenGuard_10145 = 10145 // unlock error
)

const ( // GoWebsocket
	GoWebsocket       = 30000 // GoWebsocket : 基礎號段
	GoWebsocket_30086 = 30086 // GoWebsocket_30086 : 取得玩家資料發生錯誤
	GoWebsocket_30087 = 30087 // GoWebsocket_30087 : 玩家不在線上
	GoWebsocket_30088 = 30088 // GoWebsocket_30088 : 玩家被禁言
	GoWebsocket_30089 = 30089 // GoWebsocket_30089 : 沒有權限
	GoWebsocket_30090 = 30090 // GoWebsocket_30090 : 頻道找不到
	GoWebsocket_30091 = 30091 // GoWebsocket_30091 : 頻道已經存在
	GoWebsocket_30092 = 30092 // GoWebsocket_30092 : Action Key 找不到
	GoWebsocket_30093 = 30093 // GoWebsocket_30093 : JSON 反序列化錯誤
	GoWebsocket_30094 = 30094 // GoWebsocket_30094 : JSON 序列化錯誤
	GoWebsocket_30095 = 30095 // GoWebsocket_30095 : 取得本人個人資料發生錯誤
	GoWebsocket_30096 = 30096 // GoWebsocket_30096 : 寫入離線訊息發生問題
	GoWebsocket_30097 = 30097 // GoWebsocket_30097 : 存取 jcggame 服務發生異常
	GoWebsocket_30098 = 30098 // GoWebsocket_30098 : 發送/儲存資料發生異常
	GoWebsocket_30099 = 30099 // GoWebsocket_30099 : 未知錯誤
)

const ( // XPK
	XPK_21099 = 21099 // XPK_21099 : Unknown Error
	XPK_21098 = 21098 // XPK_21098 : Request Body Parse Error
	XPK_21097 = 21097 // XPK_21097 : Get Machine Data Error
	XPK_21096 = 21096 // XPK_21096 : Get Balance Error
	XPK_21095 = 21095 // XPK_21095 : Database Error
	XPK_21094 = 21094 // XPK_21094 : Get Not Taken Deal Error
	XPK_21093 = 21093 // XPK_21093 : Get Draw Data Error
	XPK_21092 = 21092 // XPK_21092 : Get DoubleUp Data Error
	XPK_21091 = 21091 // XPK_21091 : Get Deal Data Error
	XPK_21090 = 21090 // XPK_21090 : Set Extend Error
	XPK_21089 = 21089 // XPK_21089 : Set Pick Error
	XPK_21088 = 21088 // XPK_21088 : Parse DoubleUp Error
	XPK_21087 = 21087 // XPK_21087 : Check Bet
	XPK_21086 = 21086 // XPK_21086 : Generate Deal Error
	XPK_21085 = 21085 // XPK_21085 : Generate Draw Error
	XPK_21084 = 21084 // XPK_21084 : Generate DoubleUp Error
	XPK_21083 = 21083 // XPK_21083 : Insert Deal Error
	XPK_21082 = 21082 // XPK_21082 : Balance Bet Error
	XPK_21081 = 21081 // XPK_21081 : Aready Taken Win
	XPK_21080 = 21080 // XPK_21080 : Get Deal Data is Empty
	XPK_21079 = 21079 // XPK_21079 : Mag Is Zero
	XPK_21078 = 21078 // XPK_21078 : Win Out Of Max
	XPK_21077 = 21077 // XPK_21077 : Games Out Of Max
	XPK_21076 = 21076 // XPK_21076 : DoubleUp Data Check Error
	XPK_21075 = 21075 // XPK_21075 : RandNR Error
	XPK_21074 = 21074 // XPK_21074 : Unknown Cate
	XPK_21073 = 21073 // XPK_21073 : Add Member Games Error
	XPK_21072 = 21072 // XPK_21072 : Not Current Pod Name
	XPK_21071 = 21071 // XPK_21071 : HTTP Client Error
	XPK_21070 = 21070 // XPK_21070 : Get Pods Error
)

const ( // GameAdapter (21200~21299)
	GameAdapterInputDataError                 = 21200 + iota // [21200] 輸入資料錯誤
	GameAdapterKickOutError                                  // [21201] 踢出錯誤
	GameAdapterGetPartnerMemberPointsError                   // [21202] 取得合作廠商會員點數錯誤
	GameAdapterGetPartnerMemberPointsNotZero                 // [21203] 取得合作廠商會員點數不為零
	GameAdapterGetMemberBalanceError                         // [21204] 取得會員點數錯誤
	GameAdapterMemberPointsNotEnough                         // [21205] 會員點數不足
	GameAdapterCreateRandomIDError                           // [21206] 產生隨機ID錯誤
	GameAdapterWithdrawFromCommonError                       // [21207] 從 common 提款錯誤
	GameAdapterCreateTransactionError                        // [21208] 建立交易錯誤
	GameAdapterDBError                                       // [21209] 資料庫錯誤
	GameAdapterMemberWalletsNotUnusedAndEmpty                // [21210] 會員錢包不是未使用且為空
	GameAdapterDepositToPartnerError                         // [21211] 存款到合作廠商錯誤
	GameAdapterLoginToPartnerError                           // [21212] 登入合作廠商錯誤
	GameAdapterCheckCommonWalletError                        // [21213] 檢查 common 錢包錯誤
	GameAdapterAddOrMinusCommonWalletError                   // [21214] 加減 common 錢包錯誤
	GameAdapterFinishCommonWalletError                       // [21215] 結束 common 錢包錯誤
	GameAdapterGetCommonWalletNotZero                        // [21216] 取得 common 錢包不為零
	GameAdapterAddNewMediatorWalletError                     // [21217] 新增中介錢包錯誤
	GameAdapterGetMediatorWalletsError                       // [21218] 取得中介錢包錯誤
	GameAdapterUpdateMediatorWalletError                     // [21219] 更新中介錢包錯誤
	GameAdapterMemberWalletsIsUnusedOrEmpty                  // [21220] 會員錢包是未使用或為空
	GameAdapterKickOutAllError                               // [21221] 踢出所有玩家發生錯誤
	GameAdapterGetPodsError                                  // [21222] 取得 Pod 錯誤
	GameAdapterWithdrawPartnerWalletError                    // [21223] 從合作廠商錢包提款到錯誤
	GameAdapterLockCommonWalletError                         // [21224] 鎖定 common 錢包錯誤
	GameAdapterHallIDNotExist                                // [21225] 廳館 ID 不存在
	GameAdapterGetOutSideGameIDError                         // [21226] 取得外部遊戲 ID 錯誤
	GameAdapterGetDemoURLError                               // [21227] 取得 Demo URL 錯誤
)

var retStatusCode = map[int]GroupRetStatusCode{
	Success: {0, "Success"},

	// BaseStatusCode >>
	TokenGuard: {0, "[TokenGuard] BaseStatusCode"},
	KickCenter: {0, "[KickCenter] BaseStatusCode"},
	// BaseStatusCode <<

	// TokenGuard>>
	TokenGuard_10101: {0, "[TokenGuard] Login Failed"},
	TokenGuard_10102: {0, "[TokenGuard] Member Not Exist"},
	TokenGuard_10103: {0, "[TokenGuard] Password is invalid"},
	TokenGuard_10104: {0, "[TokenGuard] JWT Invalid"},
	TokenGuard_10105: {0, "[TokenGuard] JWT Error"},
	TokenGuard_10106: {0, "[TokenGuard] JWT Create Error"},
	TokenGuard_10107: {0, "[TokenGuard] FB Token Expired"},
	TokenGuard_10108: {0, "[TokenGuard] FB Token Invalid"},
	TokenGuard_10109: {0, "[TokenGuard] Create Account Error"},
	TokenGuard_10110: {0, "[TokenGuard] Account Not Active"},
	TokenGuard_10111: {0, "[TokenGuard] Verification Time Out"},
	TokenGuard_10112: {0, "[TokenGuard] Change Password Fail"},
	TokenGuard_10113: {0, "[TokenGuard] Verification Code Error"},
	TokenGuard_10114: {0, "[TokenGuard] CellPhone Bind Fail"},
	TokenGuard_10115: {0, "[TokenGuard] Member Not Bind CellPhone"},
	TokenGuard_10116: {0, "[TokenGuard] No CellPhone Bind Authority"},
	TokenGuard_10117: {0, "[TokenGuard] No Forget Password Authorit"},
	TokenGuard_10118: {0, "[TokenGuard] Invalid Base Url"},
	TokenGuard_10119: {0, "[TokenGuard] Change Password Has Been"},
	TokenGuard_10120: {0, "[TokenGuard] No LogIn Authority"},
	TokenGuard_10121: {0, "[TokenGuard] Line Token Invalid"},
	TokenGuard_10122: {0, "[TokenGuard] LoginDevice Not Trusted"},
	TokenGuard_10123: {0, "[TokenGuard] Device Bind Fail"},
	TokenGuard_10124: {0, "[TokenGuard] Device Already Exists"},
	TokenGuard_10125: {0, "[TokenGuard] Check Sign Error"},
	TokenGuard_10126: {0, "[TokenGuard] JWT Expired"},
	TokenGuard_10127: {0, "[TokenGuard] JWT Unknown Error"},
	TokenGuard_10128: {0, "[TokenGuard] Base Request Compare Error"},
	TokenGuard_10129: {0, "[TokenGuard] Login from another Device"},
	TokenGuard_10130: {0, "[TokenGuard] Server is Maintenance"},
	TokenGuard_10131: {0, "[TokenGuard] Set Maintenance Error"},
	TokenGuard_10132: {0, "[TokenGuard] Set kick member error"},
	TokenGuard_10133: {0, "[TokenGuard] Bind Account BodyParser Error"},
	TokenGuard_10134: {0, "[TokenGuard] Bind Account CheckDeviceInfo Error"},
	TokenGuard_10135: {0, "[TokenGuard] Only Support Account Bind"},
	TokenGuard_10136: {0, "[TokenGuard] Connect To bind Service Error"},
	TokenGuard_10137: {0, "[TokenGuard] Gen RefreshJWT Error"},
	TokenGuard_10138: {0, "[TokenGuard] Kick By Backend"},
	TokenGuard_10139: {1, "[JAVA_Common_10139] Account is cancellation"},
	TokenGuard_10140: {0, "[TokenGuard] No SafeBox Authority"},
	TokenGuard_10141: {0, "[TokenGuard] Request Time Check Fail"},
	TokenGuard_10142: {0, "[TokenGuard] Get ReTry Error"},
	TokenGuard_10143: {0, "[TokenGuard] Set ReTry Error"},
	TokenGuard_10144: {0, "[TokenGuard] Lock Error"},
	TokenGuard_10145: {0, "[TokenGuard] Unlock Error"},
	// TokenGuard <<

	// KickCenter>>
	KickCenter_10201: {0, "[KickCenter_10201] Set Issuer error"},
	KickCenter_10202: {0, "[KickCenter_10202] Get tokenguard host(s) error"},
	KickCenter_10203: {0, "[KickCenter_10203] Send data to tokenguard error"},
	KickCenter_10204: {0, "[KickCenter_10204] Get data from tokenguard error"},
	KickCenter_10205: {0, "[KickCenter_10205] tokenguard retStatus error"},
	KickCenter_10206: {0, "[KickCenter_10206] tokenguard register hostname error"},
	KickCenter_10207: {0, "[KickCenter_10207] get Issuer error"},
	KickCenter_10208: {0, "[KickCenter_10208] del tokenguard hostname error"},
	KickCenter_10209: {0, "[KickCenter_10209] get KickMemberList error"},
	KickCenter_10210: {0, "[KickCenter_10210] get kick member info error"},
	KickCenter_10211: {0, "[KickCenter_10211] kick member error"},
	KickCenter_10212: {0, "[KickCenter_10212] in the line"},
	// KickCenter <<

	// GoWebsocket>>
	GoWebsocket_30086: {CheckCode: 0, StatusMsg: "[GoWebsocket] Get Member Info Error"},
	GoWebsocket_30087: {CheckCode: 0, StatusMsg: "[GoWebsocket] Member Not Online"},
	GoWebsocket_30088: {CheckCode: 0, StatusMsg: "[GoWebsocket] Member was Banned"},
	GoWebsocket_30089: {CheckCode: 0, StatusMsg: "[GoWebsocket] Permission denied"},
	GoWebsocket_30090: {CheckCode: 0, StatusMsg: "[GoWebsocket] Channel NotFound"},
	GoWebsocket_30091: {CheckCode: 0, StatusMsg: "[GoWebsocket] Channel Already Exists"},
	GoWebsocket_30092: {CheckCode: 0, StatusMsg: "[GoWebsocket] Action Key NotFound"},
	GoWebsocket_30093: {CheckCode: 0, StatusMsg: "[GoWebsocket] json decode Error"},
	GoWebsocket_30094: {CheckCode: 0, StatusMsg: "[GoWebsocket] json encode error"},
	GoWebsocket_30095: {CheckCode: 0, StatusMsg: "[GoWebsocket] Get Self Member Info Error"},
	GoWebsocket_30096: {CheckCode: 0, StatusMsg: "[GoWebsocket] write offline message error"},
	GoWebsocket_30097: {CheckCode: 0, StatusMsg: "[GoWebsocket] connect to jcggame service error"},
	GoWebsocket_30098: {CheckCode: 0, StatusMsg: "[GoWebsocket] Send Data Error"},
	GoWebsocket_30099: {CheckCode: 0, StatusMsg: "[GoWebsocket] WS Unknow Error"},
	// GoWebsocket <<

	// XPK>>
	XPK_21099: {CheckCode: 0, StatusMsg: "[XPK Games] Unknown Error"},
	XPK_21098: {CheckCode: 0, StatusMsg: "[XPK Games] Request Body Parse Error"},
	XPK_21097: {CheckCode: 0, StatusMsg: "[XPK Games] Get Machine Data Error"},
	XPK_21096: {CheckCode: 0, StatusMsg: "[XPK Games] Get Balance Error"},
	XPK_21095: {CheckCode: 0, StatusMsg: "[XPK Games] Database Error"},
	XPK_21094: {CheckCode: 0, StatusMsg: "[XPK Games] Get Not Taken Deal Error"},
	XPK_21093: {CheckCode: 0, StatusMsg: "[XPK Games] Get Draw Data Error"},
	XPK_21092: {CheckCode: 0, StatusMsg: "[XPK Games] Get DoubleUp Data Error"},
	XPK_21091: {CheckCode: 0, StatusMsg: "[XPK Games] Get Deal Data Error"},
	XPK_21090: {CheckCode: 0, StatusMsg: "[XPK Games] Set Extend Error"},
	XPK_21089: {CheckCode: 0, StatusMsg: "[XPK Games] Set Pick Error"},
	XPK_21088: {CheckCode: 0, StatusMsg: "[XPK Games] Parse DoubleUp Error"},
	XPK_21087: {CheckCode: 0, StatusMsg: "[XPK Games] Check Bet"},
	XPK_21086: {CheckCode: 0, StatusMsg: "[XPK Games] Generate Deal Error"},
	XPK_21085: {CheckCode: 0, StatusMsg: "[XPK Games] Generate Draw Error"},
	XPK_21084: {CheckCode: 0, StatusMsg: "[XPK Games] Generate DoubleUp Error"},
	XPK_21083: {CheckCode: 0, StatusMsg: "[XPK Games] Insert Deal Error"},
	XPK_21082: {CheckCode: 0, StatusMsg: "[XPK Games] Balance Bet Error"},
	XPK_21081: {CheckCode: 0, StatusMsg: "[XPK Games] Aready Taken Win"},
	XPK_21080: {CheckCode: 0, StatusMsg: "[XPK Games] Get Deal Data is Empty"},
	XPK_21079: {CheckCode: 0, StatusMsg: "[XPK Games] Mag Is Zero"},
	XPK_21078: {CheckCode: 0, StatusMsg: "[XPK Games] Win Out Of Max"},
	XPK_21077: {CheckCode: 0, StatusMsg: "[XPK Games] Games Out Of Max"},
	XPK_21076: {CheckCode: 0, StatusMsg: "[XPK Games] DoubleUp Data Check Error"},
	XPK_21075: {CheckCode: 0, StatusMsg: "[XPK Games] RandNR Error"},
	XPK_21074: {CheckCode: 0, StatusMsg: "[XPK Games] Unknown Cate"},
	XPK_21073: {CheckCode: 0, StatusMsg: "[XPK Games] Add Member Games Error"},
	XPK_21072: {CheckCode: 0, StatusMsg: "[XPK Games] Not Current Pod Name"},
	XPK_21071: {CheckCode: 0, StatusMsg: "[XPK Games] HTTP Client Error"},
	XPK_21070: {CheckCode: 0, StatusMsg: "[XPK Games] Get Pods Error"},
	// XPK <<

	// GameAdapter>>
	GameAdapterInputDataError:                 {CheckCode: 0, StatusMsg: "[GameAdapter] Input Data Error"},
	GameAdapterKickOutError:                   {CheckCode: 0, StatusMsg: "[GameAdapter] Kick Out Error"},
	GameAdapterGetPartnerMemberPointsError:    {CheckCode: 0, StatusMsg: "[GameAdapter] Get Partner Member Points Error"},
	GameAdapterGetPartnerMemberPointsNotZero:  {CheckCode: 0, StatusMsg: "[GameAdapter] Get Partner Member Points Not Zero"},
	GameAdapterGetMemberBalanceError:          {CheckCode: 0, StatusMsg: "[GameAdapter] Get Member Balance Error"},
	GameAdapterMemberPointsNotEnough:          {CheckCode: 0, StatusMsg: "[GameAdapter] Member Points Not Enough"},
	GameAdapterCreateRandomIDError:            {CheckCode: 0, StatusMsg: "[GameAdapter] Create Random ID Error"},
	GameAdapterWithdrawFromCommonError:        {CheckCode: 0, StatusMsg: "[GameAdapter] Withdraw From Common Error"},
	GameAdapterCreateTransactionError:         {CheckCode: 0, StatusMsg: "[GameAdapter] Create Transaction Error"},
	GameAdapterDBError:                        {CheckCode: 0, StatusMsg: "[GameAdapter] DB Error"},
	GameAdapterMemberWalletsNotUnusedAndEmpty: {CheckCode: 0, StatusMsg: "[GameAdapter] Member Wallets Not Unused And Empty"},
	GameAdapterDepositToPartnerError:          {CheckCode: 0, StatusMsg: "[GameAdapter] Deposit To Partner Error"},
	GameAdapterLoginToPartnerError:            {CheckCode: 0, StatusMsg: "[GameAdapter] Login To Partner Error"},
	GameAdapterCheckCommonWalletError:         {CheckCode: 0, StatusMsg: "[GameAdapter] Check Common Wallet Error"},
	GameAdapterAddOrMinusCommonWalletError:    {CheckCode: 0, StatusMsg: "[GameAdapter] Add Or Minus Common Wallet Error"},
	GameAdapterFinishCommonWalletError:        {CheckCode: 0, StatusMsg: "[GameAdapter] Finish Common Wallet Error"},
	GameAdapterGetCommonWalletNotZero:         {CheckCode: 0, StatusMsg: "[GameAdapter] Get Common Wallet Not Zero"},
	GameAdapterAddNewMediatorWalletError:      {CheckCode: 0, StatusMsg: "[GameAdapter] Add New Mediator Wallet Error"},
	GameAdapterGetMediatorWalletsError:        {CheckCode: 0, StatusMsg: "[GameAdapter] Get Mediator Wallets Error"},
	GameAdapterUpdateMediatorWalletError:      {CheckCode: 0, StatusMsg: "[GameAdapter] Update Mediator Wallet Error"},
	GameAdapterMemberWalletsIsUnusedOrEmpty:   {CheckCode: 0, StatusMsg: "[GameAdapter] Member Wallets Is Unused Or Empty"},
	GameAdapterKickOutAllError:                {CheckCode: 0, StatusMsg: "[GameAdapter] Kick Out All Error"},
	GameAdapterGetPodsError:                   {CheckCode: 0, StatusMsg: "[GameAdapter] Get Pods Error"},
	GameAdapterLockCommonWalletError:          {CheckCode: 0, StatusMsg: "[GameAdapter] Lock Common Wallet Error"},
	GameAdapterHallIDNotExist:                 {CheckCode: 0, StatusMsg: "[GameAdapter] Hall ID Not Exist"},
	GameAdapterGetOutSideGameIDError:          {CheckCode: 0, StatusMsg: "[GameAdapter] Get OutSide Game ID Error"},
	GameAdapterGetDemoURLError:                {CheckCode: 0, StatusMsg: "[GameAdapter] Get Demo URL Error"},
	// GameAdapter <<

	// Notification >>
	Notification_3001: {CheckCode: 0, StatusMsg: "[Notification] : JSON Parse Error"},
	Notification_3002: {CheckCode: 0, StatusMsg: "[Notification] : Validate Request Error"},
	Notification_3003: {CheckCode: 0, StatusMsg: "[Notification] : Unknown Error"},
	Notification_3004: {CheckCode: 0, StatusMsg: "[Notification] : Rewards Already Exchange"},

	// Notification <<

	// RewardDistribution >>

	RewardDistribution_InvalidInput_3201:           {CheckCode: 0, StatusMsg: "[RewardDistribution] Invalid Input"},
	RewardDistribution_ResourceNotFound_3202:       {CheckCode: 0, StatusMsg: "[RewardDistribution] Resource Not Found"},
	RewardDistribution_RewardHasClaimed_3203:       {CheckCode: 0, StatusMsg: "[RewardDistribution] Reward Has Claimed"},
	RewardDistribution_RewardHasClaimedInFull_3204: {CheckCode: 0, StatusMsg: "[RewardDistribution] Reward Has Claimed In Full"},
	RewardDistribution_Unknown_3205:                {CheckCode: 0, StatusMsg: "[RewardDistribution] Unknown Error"},
	RewardDistribution_ExternalServiceError_3206:   {CheckCode: 0, StatusMsg: "[RewardDistribution] External Service Error"},

	// RewardDistribution <<

	// SortingHat >>

	SortingHat_InvalidInput_3301:         {CheckCode: 0, StatusMsg: "[SortingHat] Invalid Input"},
	SortingHat_ResourceNotFound_3302:     {CheckCode: 0, StatusMsg: "[SortingHat] Resource Not Found"},
	SortingHat_Unknown_3303:              {CheckCode: 0, StatusMsg: "[SortingHat] Unknown Error"},
	SortingHat_InternalServiceError_3304: {CheckCode: 0, StatusMsg: "[SortingHat] Internal Service Error"},
	SortingHat_ExternalServiceError_3305: {CheckCode: 0, StatusMsg: "[SortingHat] External Service Error"},

	// SortingHat <<
}

// ErrKeyConflict : retStatusCode衝突
var ErrKeyConflict = errors.New("Key Conflict")

// ErrKeyOutOfRange : retStatusCode 超出可自訂範圍(自訂範圍為6位數)
var ErrKeyOutOfRange = errors.New("Key Out Of Range")

// New : 新增額外的retCode
func New(retStatusMap map[int]GroupRetStatusCode) error {
	for code, v := range retStatusMap {
		if _, isGet := retStatusCode[code]; isGet {
			return ErrKeyConflict
		}
		if code < 999999 && code > 1000 {
			retStatusCode[code] = v
		} else {
			return ErrKeyOutOfRange
		}
	}
	return nil
}

// GetStatusMsg : 將 statusCode 轉為 error msg
func GetStatusMsg(statusCode int) (s string) {
	data, isGet := retStatusCode[statusCode]
	if !isGet {
		s = UnknownStatusCode
	} else {
		s = data.StatusMsg
	}
	return s
}

func (b BaseResponse) Error() string { return b.RetStatus.Error() }

func (r RetStatus) Error() string { return "[Code:" + strconv.Itoa(r.StatusCode) + "] " + r.StatusMsg }

// TODO: this is for v2 version
// func (r RetStatus) Is(err error) bool {
// 	t, ok := err.(RetStatus)
// 	if !ok {
// 		return false
// 	}
// 	return r.StatusCode == t.StatusCode
// }

// Err : 將 statusCode 轉為 error (自動判斷)
func Err(statusCode int) (e error) {
	if statusCode == Success {
		return nil
	}
	return &BaseResponse{*Create(statusCode)}
}

// Err : 將 statusCode 轉為 error  (自動判斷)
func (retStatus *RetStatus) Err() (e error) {
	if retStatus == nil {
		return &BaseResponse{RetStatus: *Create(0)}
	}
	if retStatus.StatusCode == Success {
		return nil
	}
	return &BaseResponse{*retStatus}
}

// Get : 取得基於 retStatus 的 error / 若為 10000 則輸出 nil
func Get(statusCode int) (int, error) {
	if statusCode == Success {
		return statusCode, nil
	}
	return statusCode, Err(statusCode)
}

// Update : 更新retStatus
func (retStatus *RetStatus) Update(statusCode int) {
	*retStatus = *Create(statusCode)
}

// SetSuccess : 設定為 Success
func (retStatus *RetStatus) SetSuccess() {
	retStatus.Update(Success)
}

// IsSuccess : 判斷是否為 Success
func (retStatus *RetStatus) IsSuccess() bool {
	return retStatus.Is(Success)
}

// IsSuccess : 判斷是否為指定 Code
func (retStatus *RetStatus) Is(retStatusCode int) bool {
	return retStatus.StatusCode == retStatusCode
}

// CreateSuccess : 建立一個 Success 的 RetStatus
func CreateSuccess() (retStatus *RetStatus) {
	return Create(Success)
}

// Create : 新增retStatus
func Create(statusCode int) (retStatus *RetStatus) {
	retStatus = &RetStatus{
		StatusCode: statusCode,
		SystemTime: time.Now().UnixMilli(),
	}
	if statusCode < 1000 {
		retStatus.StatusMsg = fasthttp.StatusMessage(statusCode)
	} else {
		g, idGet := retStatusCode[statusCode]
		if idGet {
			retStatus.CheckCode = g.CheckCode
			retStatus.StatusMsg = g.StatusMsg
		} else {
			log.Printf("%s : %d", UnknownStatusCode, statusCode)
			retStatus.StatusMsg = UnknownStatusCode
		}
	}
	return retStatus
}

func (r RetStatus) MarshalJSON() ([]byte, error) {
	if r == (RetStatus{}) {
		r.SetSuccess()
	}
	return json.Marshal(marshal(r))
}

// 取代掉預設的 retStatusCode
func Override(retStatusMap map[int]GroupRetStatusCode) {
	clear()

	for code, v := range retStatusMap {
		retStatusCode[code] = v
	}
}

// 清除 retStatusCode
func clear() {
	retStatusCode = map[int]GroupRetStatusCode{}
}
