package main

import (
	"log"

	cnf "gitlab.silkrode.com.tw/team_golang/kbc/integration-tests/configs"
	internal "gitlab.silkrode.com.tw/team_golang/kbc/integration-tests/pkg/log"
	testcases "gitlab.silkrode.com.tw/team_golang/kbc/integration-tests/test_cases"
	"google.golang.org/grpc"
)

func main() {
	config, err := cnf.InitConfig("./configs")
	if err != nil {
		log.Fatal("Failed to init config.", err)
	}

	logger, err := internal.InitLogger(config)
	if err != nil {
		log.Fatal("Failed to init logger.", err)
	}

	grpcConn, err := grpc.Dial(config.Grpc.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to grpc server.", err)
	}

	tc := testcases.InitTestCase(config, logger, grpcConn)
	tc.TestCases()
}
