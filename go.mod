module main

go 1.13

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/salrashid123/grpc_keylog/echo v0.0.0
	golang.org/x/net v0.0.0-20211118161319-6a13c67c3ce4
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/salrashid123/grpc_keylog/echo => ./src/echo
