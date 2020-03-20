.DEFAULT_GOAL := module

.PHONY: all module clean test action newUser callbackURL ping

clean:
	@rm -f message/*

module: test
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message *.proto
	@echo "generated all"

test:
	@go test ./

action:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message action.proto
	@echo "generated: action"

src/action.proto: action
message/action.pb.go: action

newUser:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message newUser.proto
	@echo "generated: newUser"

src/newUser.proto: newUser
message/newUser.pb.go: newUser

userExists:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message userExists.proto
	@echo "generated: userExists"

src/userExists.proto: userExists
message/userExists.pb.go: userExists


callbackURL:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message callbackURL.proto
	@echo "generated: callbackURL"

src/callbackURL.proto: callbackURL
message/callbackURL.pb.go: callbackURL

ping:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message ping.proto
	@echo "generated: ping"

src/ping.proto: ping
message/ping.pb.go: ping
