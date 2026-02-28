gomock:

bin install => go install go.uber.org/mock/mockgen@latest
add path => export PATH=$PATH:$(go env GOPATH)/bin
command => ex: mockgen -source=internal/core/port/user.go -destination=internal/core/port/mock/user.go

air:

bin install => go install github.com/air-verse/air@latest

sqlc: 

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest