.PHONY : auth all user

auth :
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
	--openapiv2_out . \
	--validate_out="lang=go:." \
	--validate_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
    apu/auth/*.proto

user :
	protoc  \
	--go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
	--openapiv2_out . \
	--validate_out="lang=go:." \
	--validate_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
    apu/user/*.proto

server :
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
	--openapiv2_out . \
	--validate_out="lang=go:." \
	--validate_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
    apu/server/*.proto

server_file :
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
	--openapiv2_out . \
	--validate_out="lang=go:." \
	--validate_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
    apu/server_file/*.proto

mail_sender :
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
	--openapiv2_out . \
	--validate_out="lang=go:." \
	--validate_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
    apu/mail_sender/*.proto

merge-openapiv2 : 
	go-swagger-merger -o .\apu\vcs_msm.swagger2.json  .\apu\user\user.swagger.json .\apu\auth\auth.swagger.json

all : auth user server
