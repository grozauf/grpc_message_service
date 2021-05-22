module github.com/grozauf/message_service

go 1.16

replace github.com/grozauf/message_service/pkg/message_service => ./pkg/message_service

require (
	github.com/grozauf/message_service/pkg/message_service v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210521195947-fe42d452be8f
	google.golang.org/grpc v1.38.0
)
