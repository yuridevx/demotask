module yuridev.com/googdemo/game

go 1.21.5

require (
	github.com/shopspring/decimal v1.3.1
	google.golang.org/grpc v1.59.0
	yuridev.com/googdemo/domain v0.0.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace yuridev.com/googdemo/domain => ../domain
