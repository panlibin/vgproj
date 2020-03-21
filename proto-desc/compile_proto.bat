protoc.exe --proto_path=./ --plugin=protoc-gen-go=./protoc-gen-go.exe --go_out=paths=source_relative:../proto ./common/*.proto
protoc.exe --proto_path=./ --plugin=protoc-gen-go=./protoc-gen-go.exe --go_out=paths=source_relative:../proto ./globalrpc/*.proto
protoc.exe --proto_path=./ --plugin=protoc-gen-go=./protoc-gen-go.exe --go_out=paths=source_relative:../proto ./msg/*.proto
protoc.exe --proto_path=./ --plugin=protoc-gen-go=./protoc-gen-go.exe --go_out=plugins=grpc,paths=source_relative:../proto ./masterrpc/*.proto
protoc.exe --proto_path=./ --plugin=protoc-gen-go=./protoc-gen-go.exe --go_out=plugins=grpc,paths=source_relative:../proto ./loginrpc/*.proto
protoc.exe --proto_path=./ --plugin=protoc-gen-go=./protoc-gen-go.exe --go_out=plugins=grpc,paths=source_relative:../proto ./gamerpc/*.proto
pause
