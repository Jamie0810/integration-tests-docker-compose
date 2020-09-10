package model

type AccountOrderCheckReq struct {
	ProductAPIID               int64    `protobuf:"varint,1,opt,name=ProductAPIID,proto3" json:"ProductAPIID,omitempty"`                               // 產品街口主鍵號
	MerchantName               string   `protobuf:"bytes,2,opt,name=MerchantName,proto3" json:"MerchantName,omitempty"`                                // 商戶姓名
	TrackingNumber             string   `protobuf:"bytes,3,opt,name=TrackingNumber,proto3" json:"TrackingNumber,omitempty"`                            // 訂單編號
	PlayerID                   string   `protobuf:"bytes,4,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`                                        // 玩家ID
	DepositDailyLimit          float64  `protobuf:"fixed64,5,opt,name=DepositDailyLimit,proto3" json:"DepositDailyLimit,omitempty"`                    // 玩家ID單日打款總額
	DepositAccountAmount       float64  `protobuf:"fixed64,6,opt,name=DepositAccountAmount,proto3" json:"DepositAccountAmount,omitempty"`              // 打款額度
	WithdrawAccountAmount      float64  `protobuf:"fixed64,7,opt,name=WithdrawAccountAmount,proto3" json:"WithdrawAccountAmount,omitempty"`            // 收款額度
	PaymentType                string   `protobuf:"bytes,8,opt,name=PaymentType,proto3" json:"PaymentType,omitempty"`                                  // 打收款方式
	AccountNumber              string   `protobuf:"bytes,9,opt,name=AccountNumber,proto3" json:"AccountNumber,omitempty"`                              // 卡號或收付款帳號
	AccountDepositDailyTotal   float64  `protobuf:"fixed64,10,opt,name=AccountDepositDailyTotal,proto3" json:"AccountDepositDailyTotal,omitempty"`     // 玩家帳號當日打款總額
	IsNewPlayer                bool     `protobuf:"varint,11,opt,name=IsNewPlayer,proto3" json:"IsNewPlayer,omitempty"`                                // 是否為新玩家
	IsNewAccount               bool     `protobuf:"varint,12,opt,name=IsNewAccount,proto3" json:"IsNewAccount,omitempty"`                              // 是否為舊玩家新收款帳號
	AllowStep                  []string `protobuf:"bytes,13,rep,name=AllowStep,proto3" json:"AllowStep,omitempty"`                                     // 要跳過的風控邏輯代碼
	MerchantDepositDailyAmount float64  `protobuf:"fixed64,14,opt,name=MerchantDepositDailyAmount,proto3" json:"MerchantDepositDailyAmount,omitempty"` // 商戶當日已打款金額
	CallbackURL                string   `protobuf:"bytes,15,opt,name=CallbackURL,proto3" json:"CallbackURL,omitempty"`                                 // 訂單回調地址
}

type AccountOrderCheckResp struct {
	IsAllow                  bool               `protobuf:"varint,1,opt,name=IsAllow,proto3" json:"IsAllow,omitempty"`                                                               // 是否通過風控
	PaymentErrorType         PaymentErrorType   `protobuf:"varint,2,opt,name=PaymentErrorType,proto3,enum=ams.PaymentErrorType" json:"PaymentErrorType,omitempty"`                   // 錯誤類型
	AllowChannel             []*Channel         `protobuf:"bytes,3,rep,name=AllowChannel,proto3" json:"AllowChannel,omitempty"`                                                      // 允許的渠道權重從大至小排序
	DepositToPlayerOption    TransferOptionType `protobuf:"varint,4,opt,name=DepositToPlayerOption,proto3,enum=ams.TransferOptionType" json:"DepositToPlayerOption,omitempty"`       // 打款手續費選項
	WithdrawFromPlayerOption TransferOptionType `protobuf:"varint,5,opt,name=WithdrawFromPlayerOption,proto3,enum=ams.TransferOptionType" json:"WithdrawFromPlayerOption,omitempty"` // 收款手續費選項
	DepositFeeToPlayer       float64            `protobuf:"fixed64,6,opt,name=DepositFeeToPlayer,proto3" json:"DepositFeeToPlayer,omitempty"`                                        // 打款手續費
	WithdrawFeeFromPlayer    float64            `protobuf:"fixed64,7,opt,name=WithdrawFeeFromPlayer,proto3" json:"WithdrawFeeFromPlayer,omitempty"`                                  // 收款手續費
	AllowNegative            bool               `protobuf:"varint,8,opt,name=AllowNegative,proto3" json:"AllowNegative,omitempty"`                                                   // 允許錢包負數
	BBCode                   []string           `protobuf:"bytes,9,rep,name=BBCode,proto3" json:"BBCode,omitempty"`                                                                  // 風控暫停原因
	TrackingNumber           string             `protobuf:"bytes,10,opt,name=TrackingNumber,proto3" json:"TrackingNumber,omitempty"`                                                 // 訂單編號
}

type PaymentErrorType int32

const (
	PaymentErrorType_NO_ERROR                  PaymentErrorType = 0  // 沒有錯誤
	PaymentErrorType_OVER_DAILY_LIMIT          PaymentErrorType = 1  // 超過單日額度
	PaymentErrorType_OVER_CREATE_FREQUENCY     PaymentErrorType = 2  // 超過M分鐘建立 N筆訂單
	PaymentErrorType_BECLOCKED                 PaymentErrorType = 3  // 被黑名單擋掉
	PaymentErrorType_NO_CHANNEL                PaymentErrorType = 4  // 沒有支援的付款渠道
	PaymentErrorType_ERR_PAYMENT               PaymentErrorType = 5  // 檢查付款方式發生錯誤
	PaymentErrorType_INVALID_PAYMENT           PaymentErrorType = 6  // 不支援的付款方式
	PaymentErrorType_NOFOUND_PRODUCTAPI        PaymentErrorType = 7  // 找不到產品接口
	PaymentErrorType_DISABLE_PRODUCTAPI        PaymentErrorType = 8  // 產品接口未啟用
	PaymentErrorType_NOFOUND_MERCHANT          PaymentErrorType = 9  // 找不到商戶
	PaymentErrorType_DISABLE_MERCHANT          PaymentErrorType = 10 // 商戶未啟用
	PaymentErrorType_DISABLE_MERCHANT_DEPOSIT  PaymentErrorType = 11 // 商戶未啟用打款
	PaymentErrorType_DISABLE_MERCHANT_WITHDRAW PaymentErrorType = 12 // 商戶未啟用收款
	PaymentErrorType_NOFOUND_PROJECT           PaymentErrorType = 13 // 找不到項目
	PaymentErrorType_DISABLE_PROJECT           PaymentErrorType = 14 // 項目未啟用
	PaymentErrorType_NOFOUND_PRODUCT           PaymentErrorType = 15 // 找不到產品
	PaymentErrorType_DISABLE_PRODUCT           PaymentErrorType = 16 // 產品未啟用
	PaymentErrorType_ERR_BLACKLISST            PaymentErrorType = 17 // 取得黑名單錯誤
	PaymentErrorType_ERR_CHECK_FREQUENCY       PaymentErrorType = 18 // 檢查玩家打單頻率錯誤
	PaymentErrorType_ERR_DAILY_LIMIT           PaymentErrorType = 19 // 檢查單日額度錯誤
	PaymentErrorType_ERR_FEE                   PaymentErrorType = 20 // 取得手續費設定錯誤
	PaymentErrorType_ERR_MARK_ORDER            PaymentErrorType = 21 // 標記成功訂單失敗（計算建單頻率用）
	PaymentErrorType_PAUSE                     PaymentErrorType = 22 // 風控暫停
	PaymentErrorType_OTHERS                    PaymentErrorType = 23 // 其他錯誤
)

type Channel struct {
	ID      uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`           // 渠道編號
	Code    string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`        // 渠道代碼
	Name    string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`        // 渠道名稱
	Timeout int64  `protobuf:"varint,4,opt,name=Timeout,proto3" json:"Timeout,omitempty"` // 訂單建立過多久沒收到回調轉為 [待確認] （time.Duration)
}

type TransferOptionType int32

const (
	TransferOptionType_FIXED TransferOptionType = 0 // 固定費率
	TransferOptionType_RATE  TransferOptionType = 1 // 比率
)
