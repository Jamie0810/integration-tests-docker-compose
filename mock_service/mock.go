package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	pb "gitlab.silkrode.com.tw/golang/kbc_pb/ams"

	"gitlab.silkrode.com.tw/golang/mq"
	"gitlab.silkrode.com.tw/golang/mq/inside/pub"
	"gitlab.silkrode.com.tw/golang/mq/inside/sub"
	"gitlab.silkrode.com.tw/team_golang/KM/order/integration_test/mock_service/config"
	"gitlab.silkrode.com.tw/team_golang/KM/order/integration_test/mock_service/model"
	"google.golang.org/grpc"
)

func main() {
	cnf, err := config.InitConfig("./config")
	if err != nil {
		log.Fatal("Failed to init config.", err)
	}

	go amsBBService(cnf)             //mock 收款經黑名單那段
	go amsRiskManagementProcess(cnf) //mock 打款經風控那段
	select {}
}

func amsBBService(cnf config.Config) {
	log.Println("===== Running Mock service: amsBBService =====")
	addr := cnf.Grpc.Addr + ":" + cnf.Grpc.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterBBServiceServer(s, &grpcServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}

type grpcServer struct {
	pb.UnimplementedBBServiceServer
}

func (s *grpcServer) AccountOrderCheck(ctx context.Context, req *pb.AccountOrderCheckReq) (*pb.AccountOrderCheckResp, error) {
	log.Printf("TrackingNumber: %v", req.TrackingNumber)

	polypay := pb.Channel{
		ID:      1,
		Code:    "mock-polypay",
		Name:    "聚合支付Mock",
		Timeout: 15,
	}

	resp := pb.AccountOrderCheckResp{
		IsAllow:                  true,
		PaymentErrorType:         0,
		AllowChannel:             []*pb.Channel{&polypay},
		WithdrawFromPlayerOption: 0,
		WithdrawFeeFromPlayer:    100,
		AllowNegative:            true,
		TrackingNumber:           req.TrackingNumber,
	}

	return &resp, nil
}

func amsRiskManagementProcess(cnf config.Config) {
	log.Println("===== Running Mock service: amsRiskManagementProcess =====")

	ctx := context.Background()
	subInstance, err := mq.Init(ctx, cnf.Pubsub.TopicBBInput, cnf.Pubsub.CredentialsFilePath, cnf.Pubsub.ProjectID, mq.InitSub())
	if err != nil {
		log.Printf("Failed to create subInstance: %v", err)
		return
	}

	subInstance.Subscriber().Options(
		sub.AsyncMode(),
		sub.SetErrorHook(func(err error, msg string, msgID string) {
			log.Fatalf("Failed to pull message from pub/sub: %v", err)
			return
		})).
		Subscribe(func(ctx context.Context, msg []byte, msgId string) error {
			log.Println("Received a message from " + cnf.Pubsub.TopicBBInput)
			log.Println("Message ID: " + msgId)
			log.Println("Message: " + string(msg))

			var obj model.AccountOrderCheckReq
			if err := json.Unmarshal(msg, &obj); err != nil {
				log.Fatalf("json unmarshal Err: %v", err)
			}

			// Pub to TopicBBOutput
			pubInstance, err := mq.Init(ctx, cnf.Pubsub.TopicBBOutput, cnf.Pubsub.CredentialsFilePath, cnf.Pubsub.ProjectID, mq.InitPub())
			if err != nil {
				log.Fatalf("Failed to create pubInstance: %v", err)
			}

			bankcard := model.Channel{
				ID:      1,
				Code:    "mock-bankcard",
				Name:    "銀行卡Mock",
				Timeout: 15,
			}

			depositReq := model.AccountOrderCheckResp{
				IsAllow:               true,
				PaymentErrorType:      0,
				AllowChannel:          []*model.Channel{&bankcard},
				DepositToPlayerOption: 0,
				DepositFeeToPlayer:    100,
				AllowNegative:         false,
				// BBCode: ,
				TrackingNumber: obj.TrackingNumber,
			}
			pubInstance.Publisher().Options(
				pub.SetErrorHook(func(err error, requestID string) {
					log.Printf("failed to publish requestID: %v ; message: %v", requestID, err)
				}),
			).PublishJSON(&depositReq, fmt.Sprintf("ams-bb-worker-check-order:%d", time.Now().UTC().Unix()))
			return nil
		})
}
