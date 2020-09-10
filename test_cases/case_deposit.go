package testcases

import (
	"encoding/json"

	"gitlab.silkrode.com.tw/golang/mq"
	"gitlab.silkrode.com.tw/golang/mq/inside/pub"
	model "gitlab.silkrode.com.tw/team_golang/kbc/integration-tests/models"
)

func (tc *TestCase) Deposit() {
	tc.logger.Info().Msg("--------Runing test Deposit--------")

	pubInstance, err := mq.Init(tc.ctx, tc.config.Pubsub.Deposit, tc.config.Pubsub.CredentialsFilePath, tc.config.Pubsub.ProjectID, mq.InitPub())
	if err != nil {
		tc.logger.Error().Msgf("Failed to create pubInstance: %+v", err)
	}

	_depositReq := model.DepositRequest{
		Type: "vgam",
		Payload: model.OrderCreationVGAMReq{
			MerchantOrderNumber: "order123123123123",
			TrackingNumber:      "TrackingNumberrr065",
			Info: &model.OrderCreationVGAMReq_PlayerInfo{
				PlayerUID: "PlayerUID12333",
			},
			MerchantInfo: &model.KbcMerchantInfo{
				MerchantID:             123,
				MerchantName:           "123",
				MerchantProjectID:      123,
				MerchantProjectName:    "MerchantProjectName",
				MerchantProductID:      1,
				MerchantProductName:    "MerchantProductName",
				MerchantProductAPIID:   123,
				MerchantProductAPIName: "12333",
			},
			CallBackURL:       "CallBackURL",
			Amount:            33.33,
			BankCode:          "301",
			AccountUIDRec:     "12345678",
			AccountName:       "jamie~~~",
			BankNameRec:       "中國銀行",
			BankBranchNameRec: "深圳分行",
			TransactionType:   model.TransactionType_DEPOSIT,
		},
	}

	depositReq, err := json.Marshal(_depositReq)
	if err != nil {
		tc.logger.Error().Msgf("json.Marshal Err: %+v", err)
	}

	pubInstance.Publisher().Options(
		pub.SetErrorHook(func(err error, requestID string) {
			tc.logger.Error().Msgf("failed to publish requestID: %v ; message: %v", requestID, err)
		}),
	).Publish([]byte(string(depositReq)), "Integration case test: deposit")
}
