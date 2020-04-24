# Golang_protobuf
## 1. 安裝 Protocol Buffers 生成工具
https://github.com/protocolbuffers/protobuf/releases/tag/v3.0.0
找到自己系統相對應的，下載之後解壓縮，然後將解壓縮的後的檔案放置任意目錄中，之後將bin/檔案丟至系統路徑 $PATH，這樣我們就能夠在終端機執行protoc.exe的功能。
## 2.套件安裝
```go get -u github.com/golang/protobuf/{proto,protoc-gen-go}```

## 3.Protobuf文件宣告 (.proto)
code (
// 撰寫格式是 Proto v3。
syntax = "proto3";  
// 生成的程式在 Golang 中將會屬於 `protobuf` 套件。
package protobuf;

// User 帶有使用者資料，如帳號、密碼。
message User {  
    int64  id       = 1;
    string username = 2;
    string password = 3;
})

## 4.執行命令
code(Cmd -> protoc --go_out=. *.proto)
會得到一個*.pb.go檔可以當pkg引入

可以安裝
code(go get -u github.com/gogo/protobuf/proto)
gogo/protobuf 是基於官方 protobuf 的衍生分歧，當初選用的原因是比起官方的 protobuf 還要多一些功能，效能上聽說也有些提升。
