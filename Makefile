.PHONY : auth all user

auth :
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
	--openapiv2_out . \
    apu/auth/*.proto

user :
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
	--openapiv2_out . \
    apu/user/*.proto

merge-openapiv2 : 
	go-swagger-merger -o .\apu\vcs_msm.swagger2.json  .\apu\user\user.swagger.json .\apu\auth\auth.swagger.json

all : auth user merge-openapiv2
