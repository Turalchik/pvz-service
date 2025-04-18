module github.com/Turalchik/pvz-service

go 1.24.2

require (
	github.com/Turalchik/pvz-service/pkg/pvz_service v0.0.0-00010101000000-000000000000
	github.com/jmoiron/sqlx v1.4.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/Masterminds/squirrel v1.5.4 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250414145226-207652e42e2e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250409194420-de1ac958c67a // indirect
	google.golang.org/grpc v1.71.1 // indirect
)

replace github.com/Turalchik/pvz-service/pkg/pvz_service => ./pkg/pvz_service
