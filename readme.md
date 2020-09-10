# KBC Order Integration Tests
### Services

1. mock-service
1. order-grpc
1. order-worker
1. ams-grpc
1. db

### Test Scenario
*Test Scenario 1*: Check the **Deposit** Functionality
* Start → pub to `topic_kbc_integration_test_deposit`  → **order-worker** → pub to `topic_kbc_integration_test_doctor_strange` → **mock-service** (mock Risk Management Process in AMS grpc) → **order-worker** → **ams-grpc** → **order-worker**  → create sub order

*Test Scenario 2*: Check the **Withdrawal** Functionality
* Start → **order-grpc** → **mock-service** (mock BBService in AMS grpc) → **order-grpc** 

*Test Scenario 3*: Check the **Query** Functionality
* Start → **order-grpc** 
### Quick start
1. Run all services:
```
docker-compose up -d
make run-migration
docker-compose up
```

2. Run the tests:
Run `main.go` file and choose a case test to run.
