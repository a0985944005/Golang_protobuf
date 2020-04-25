# Golang_protobuf
## 1. 安裝 Protocol Buffers 生成工具
```https://github.com/protocolbuffers/protobuf/releases/tag/v3.0.0```
找到自己系統相對應的，下載之後解壓縮，然後將解壓縮的後的檔案放置任意目錄中，之後將bin/檔案丟至系統路徑 $PATH，這樣我們就能夠在終端機執行protoc.exe的功能。
## 2.套件安裝
```go get -u github.com/golang/protobuf/{proto,protoc-gen-go}```

## 3.Protobuf文件宣告 (.proto)
```
// 撰寫格式是 Proto v3。
syntax = "proto3";  
// 生成的程式在 Golang 中將會屬於 `protobuf` 套件。
package protobuf;

// User 帶有使用者資料，如帳號、密碼。
message User {  
    int64  id       = 1;
    string username = 2;
    string password = 3;
}
```

## 4.執行命令
```Cmd -> protoc --go_out=. *.proto```
會得到一個*.pb.go檔可以當pkg引入

可以安裝
```go get -u github.com/gogo/protobuf/proto```
gogo/protobuf 是基於官方 protobuf 的衍生分歧，當初選用的原因是比起官方的 protobuf 還要多一些功能，效能上聽說也有些提升。

## 5.與 JSON 比較，Protocol Buffers 有這些優點

```message User {  
    string username = 1;
    string password = 2;
}
```
* 資料輕量化：資料非常輕量，省去了不必要的 { 或 : 累贅。
* 混淆性：在一般人眼中無法輕易地猜測出資料結構為何，因為有經過編碼。
* 效能高：處理速度很快。
* 極具方便性：結構就是你的資料模型，你能夠直接在程式中使用這些結構，而不用建立新的物件來接納、映射（Mapping）這些資料。
* 清晰明瞭、無需文件：.proto 檔案本身就是你的文件，不需要額外撰寫 API 或結構文件來告訴別人你接受怎樣的資料。
