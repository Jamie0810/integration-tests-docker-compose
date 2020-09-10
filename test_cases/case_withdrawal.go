package testcases

import (
	pb "gitlab.silkrode.com.tw/golang/kbc_pb/order"
)

func (tc *TestCase) Withdrawal() {
	tc.logger.Info().Msg("--------Runing test Withdrawal--------")

	c := pb.NewOrderCreateServiceClient(tc.grpcConn)
	r, err := c.OrderCreation(tc.ctx, &pb.OrderCreationReq{
		APIKey:              "APIKey_test",
		MerchantOrderNumber: "123",
		MerchantInfo: &pb.KbcMerchantInfo{
			MerchantID:             218,
			MerchantName:           "bkhandsome",
			MerchantProjectID:      10,
			MerchantProjectName:    "bktest5566@@!",
			MerchantProductID:      27,
			MerchantProductName:    "產品測試",
			MerchantProductAPIID:   2,
			MerchantProductAPIName: "ProductAPISetting_1",
		},
		TransactionType: 2,
		PaymentType:     1,
		// PaymentInfo: ,
		Amount: 100,
		PlayerInfo: &pb.PlayerInfo{
			PlayerID:   "123",
			DeviceID:   "123",
			DeviceIP:   "127.0.0.2",
			DeviceType: "1",
			Telephone:  "0933123123",
			Name:       "jamie",
			PayAccount: "12333",
			AlipayName: "jamie",
			Metadata:   "333",
		},
		// DiscountCode: "1",
		CallbackURL: "CallbackURL",
	})
	if err != nil {
		tc.logger.Error().Msgf("Withdrawal Err %+v", err)
		return
	}

	tc.logger.Info().Msgf("GetPaymentURL: %+v", r.GetPaymentURL())
	// tc.logger.Info().Msgf("GetOrderUID: %+v", r.GetOrderUID())
	// tc.logger.Info().Msgf("GetExpiredAt: %+v", r.GetExpiredAt())
	// tc.logger.Info().Msgf("GetBankBranchNameRec: %+v", r.GetBankBranchNameRec())
	// tc.logger.Info().Msgf("GetBankCodeRec: %+v", r.GetBankCodeRec())
	// tc.logger.Info().Msgf("GetAccountRec: %+v", r.GetAccountRec())
	// tc.logger.Info().Msgf("GetAccountUIDRec: %+v", r.GetAccountUIDRec())
	// tc.logger.Info().Msgf("GetExpectedAmount: %+v", r.GetExpectedAmount())
}
