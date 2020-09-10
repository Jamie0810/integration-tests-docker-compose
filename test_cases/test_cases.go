package testcases

import (
	"context"
	"fmt"

	config "gitlab.silkrode.com.tw/team_golang/kbc/integration-tests/configs"
	"gitlab.silkrode.com.tw/team_golang/kbc/integration-tests/pkg/log"
	"google.golang.org/grpc"
)

type TestCase struct {
	ctx      context.Context
	config   config.Config
	logger   *log.Logger
	grpcConn *grpc.ClientConn
}

func InitTestCase(config config.Config, logger *log.Logger, grpcConn *grpc.ClientConn) *TestCase {
	ctx := context.Context(context.Background())

	tc := &TestCase{
		ctx:      ctx,
		config:   config,
		logger:   logger,
		grpcConn: grpcConn,
	}
	return tc
}

func (tc *TestCase) TestCases() {
	var testCode int

	fmt.Print("Choose a test case (Enter the number):" + "\n " +
		"(1) Deposit\n (2) Withdrawal\n (3) Query\n (4) Quit\n")
	fmt.Scanf("%d", &testCode)

	switch testCode {
	case 1:
		tc.Deposit()
	case 2:
		tc.Withdrawal()
	case 3:
		tc.Query()
	}
}
