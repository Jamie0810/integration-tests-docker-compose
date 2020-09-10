package testcases

import (
	pb "gitlab.silkrode.com.tw/golang/kbc_pb/order"
)

func (tc *TestCase) Query() {
	tc.logger.Info().Msg("--------Runing test Query--------")

	c := pb.NewOrderServiceClient(tc.grpcConn)
	r, err := c.OrderQuery(tc.ctx, &pb.OrderQueryReq{
		MerchantID:          123,
		ProjectID:           123,
		MerchantOrderNumber: "order123123123123",
	})

	if err != nil {
		tc.logger.Error().Msgf("Query Err %+v", err)
		return
	}

	tc.logger.Info().Msgf("GetMerchatOrderNumber: %+v", r.GetMerchatOrderNumber())
	tc.logger.Info().Msgf("GetPlayerID: %+v", r.GetPlayerID())
}
