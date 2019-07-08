https://blog.csdn.net/u010471121/article/details/52605365
https://www.bookstack.cn/read/learning-proto3/guide-overview.md

protoc --go_out=. usermsg.proto


php
/usr/local/bin/protoc

protoc --plugin=grpc_php_plugin --proto_path= /Users/itech8/Desktop/test --php_out= /Users/itech8/Desktop/center_user.proto


/usr/local/protobuf/bin/protoc /Users/itech8/Desktop/test/center_user.proto --php_out=/Users/itech8/Desktop/test


protoc --plugin=grpc_php_plugin --proto_path= /Users/itech8/Desktop/test/  --php_out= /Users/itech8/Desktop/test/center_user.proto


export PATH=$PATH:/Users/itech8/Desktop/test


protoc --plugin=grpc_php_plugin  --php_out= ./center_user.proto

将center_user.proto文件移到test文件夹下面,然后进入test文件夹下面
mkdir Centeruser

protoc --plugin=grpc_php_plugin --php_out=./Centeruser ./center_user.proto


protoc --plugin=grpc_php_plugin --php_out=. ./user.proto



go get -u github.com/golang/protobuf/protoc-gen-go

使用protoc-gen-go内置的grpc插件生成GRPC代码:
protoc --go_out=plugins=grpc:. hello.proto


protoc --plugin=grpc_php_plugin --php_out=./CenterMember ./centermember.proto




