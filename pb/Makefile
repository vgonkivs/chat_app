export PATH := $(GOPATH)/bin:$(PATH)
PB = $(wildcard *.proto)
GO = $(PB:.proto=.pb.go)

all: $(GO)

%.pb.go: %.proto
		protoc --go_out=plugins=grpc:. *.proto
clean:
		rm -f *.pb.go