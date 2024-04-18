protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --openapiv2_out . --validate_out="lang=go:." --validate_opt=paths=source_relative apu\auth\*.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --openapiv2_out . --validate_out="lang=go:." --validate_opt=paths=source_relative apu\user\*.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --openapiv2_out . --validate_out="lang=go:." --validate_opt=paths=source_relative  apu\server\*.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --openapiv2_out . --validate_out="lang=go:." --validate_opt=paths=source_relative  apu\server_file\*.proto
