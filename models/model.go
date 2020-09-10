package model

type TransactionType int32

const (
	TransactionType_TRANSACTION_UNKNOW TransactionType = 0
	TransactionType_DEPOSIT            TransactionType = 1
	TransactionType_WITHDRAW           TransactionType = 2
)

type PaymentType int32

const (
	PaymentType_PAYMENT_UNKNOW PaymentType = 0
	PaymentType_ALIPAY         PaymentType = 1
	PaymentType_WECHAT         PaymentType = 2
	PaymentType_BANKCARD       PaymentType = 3
)

type DepositRequest struct {
	Type    string               `json:"type"`
	Payload OrderCreationVGAMReq `json:"payload"`
}
type OrderCreationVGAMReq struct {
	TrackingNumber      string                           `protobuf:"bytes,1,opt,name=TrackingNumber,proto3" json:"TrackingNumber,omitempty"`
	WangGuanMerchantID  string                           `protobuf:"bytes,2,opt,name=WangGuanMerchantID,proto3" json:"WangGuanMerchantID,omitempty"`
	MerchantOrderNumber string                           `protobuf:"bytes,3,opt,name=MerchantOrderNumber,proto3" json:"MerchantOrderNumber,omitempty"`
	MerchantID          int64                            `protobuf:"varint,4,opt,name=MerchantID,proto3" json:"MerchantID,omitempty"`
	MerchantName        string                           `protobuf:"bytes,5,opt,name=MerchantName,proto3" json:"MerchantName,omitempty"`
	OldOrderUID         string                           `protobuf:"bytes,6,opt,name=OldOrderUID,proto3" json:"OldOrderUID,omitempty"`
	Amount              float64                          `protobuf:"fixed64,7,opt,name=Amount,proto3" json:"Amount,omitempty"`
	TransactionType     TransactionType                  `protobuf:"varint,8,opt,name=TransactionType,proto3,enum=order.TransactionType" json:"TransactionType,omitempty"`
	PaymentType         PaymentType                      `protobuf:"varint,9,opt,name=PaymentType,proto3,enum=order.PaymentType" json:"PaymentType,omitempty"`
	CallBackURL         string                           `protobuf:"bytes,10,opt,name=CallBackURL,proto3" json:"CallBackURL,omitempty"`
	MerchantInfo        *KbcMerchantInfo                 `protobuf:"bytes,11,opt,name=MerchantInfo,proto3" json:"MerchantInfo,omitempty"`
	Info                *OrderCreationVGAMReq_PlayerInfo `protobuf:"bytes,12,opt,name=Info,proto3" json:"Info,omitempty"`
	BusinessID          int64                            `protobuf:"varint,13,opt,name=BusinessID,proto3" json:"BusinessID,omitempty"`
	IsSettleByBS        int32                            `protobuf:"varint,14,opt,name=IsSettleByBS,proto3" json:"IsSettleByBS,omitempty"`
	IsAutomatic         int32                            `protobuf:"varint,15,opt,name=IsAutomatic,proto3" json:"IsAutomatic,omitempty"`
	FeeType             int32                            `protobuf:"varint,16,opt,name=FeeType,proto3" json:"FeeType,omitempty"`
	Fee                 float64                          `protobuf:"fixed64,17,opt,name=Fee,proto3" json:"Fee,omitempty"`
	BankCode            string                           `protobuf:"bytes,18,opt,name=BankCode,proto3" json:"BankCode,omitempty"`
	AccountUIDRec       string                           `protobuf:"bytes,19,opt,name=AccountUIDRec,proto3" json:"AccountUIDRec,omitempty"`
	AccountName         string                           `protobuf:"bytes,20,opt,name=AccountName,proto3" json:"AccountName,omitempty"`
	BankNameRec         string                           `protobuf:"bytes,21,opt,name=BankNameRec,proto3" json:"BankNameRec,omitempty"`
	BankBranchNameRec   string                           `protobuf:"bytes,22,opt,name=BankBranchNameRec,proto3" json:"BankBranchNameRec,omitempty"`
	BankShengRec        string                           `protobuf:"bytes,23,opt,name=BankShengRec,proto3" json:"BankShengRec,omitempty"`
	BankShiRec          string                           `protobuf:"bytes,24,opt,name=BankShiRec,proto3" json:"BankShiRec,omitempty"`
	SleepSeconds        int64                            `protobuf:"varint,25,opt,name=SleepSeconds,proto3" json:"SleepSeconds,omitempty"`
	BsCreator           string                           `protobuf:"bytes,26,opt,name=BsCreator,proto3" json:"BsCreator,omitempty"`
	BsIP                string                           `protobuf:"bytes,27,opt,name=BsIP,proto3" json:"BsIP,omitempty"`
	BsMeta              string                           `protobuf:"bytes,28,opt,name=BsMeta,proto3" json:"BsMeta,omitempty"`
}

type KbcMerchantInfo struct {
	// 商戶編號
	MerchantID uint64 `protobuf:"varint,1,opt,name=MerchantID,proto3" json:"MerchantID,omitempty"`
	// 商戶名稱
	MerchantName string `protobuf:"bytes,2,opt,name=MerchantName,proto3" json:"MerchantName,omitempty"`
	// 商戶項目編號
	MerchantProjectID uint64 `protobuf:"varint,3,opt,name=MerchantProjectID,proto3" json:"MerchantProjectID,omitempty"`
	// 商戶項目名稱
	MerchantProjectName string `protobuf:"bytes,4,opt,name=MerchantProjectName,proto3" json:"MerchantProjectName,omitempty"`
	// 商戶產品編號
	MerchantProductID uint64 `protobuf:"varint,5,opt,name=MerchantProductID,proto3" json:"MerchantProductID,omitempty"`
	// 商戶產品名稱
	MerchantProductName string `protobuf:"bytes,6,opt,name=MerchantProductName,proto3" json:"MerchantProductName,omitempty"`
	// 商戶產品API ID
	MerchantProductAPIID uint64 `protobuf:"varint,7,opt,name=MerchantProductAPIID,proto3" json:"MerchantProductAPIID,omitempty"`
	// 商戶產品API 名稱
	MerchantProductAPIName string `protobuf:"bytes,8,opt,name=MerchantProductAPIName,proto3" json:"MerchantProductAPIName,omitempty"`
}

type OrderCreationVGAMReq_PlayerInfo struct {
	PlayerUID  string `protobuf:"bytes,1,opt,name=PlayerUID,proto3" json:"PlayerUID,omitempty"`
	AccountUID string `protobuf:"bytes,2,opt,name=AccountUID,proto3" json:"AccountUID,omitempty"`
}
