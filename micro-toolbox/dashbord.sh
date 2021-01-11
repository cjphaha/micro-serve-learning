cd Services/protos
protoc --micro_out=../ --go_out=../ test.proto
cd .. && cd..