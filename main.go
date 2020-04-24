package main

import (
	"fmt"

	"./protobuf"
	"google.golang.org/protobuf/proto"
)

func main() {
	// 建立一個 User 格式，並在其中放入資料。
	data := protobuf.User{
		Id:       12345,
		Account:  "Rosco",
		Password: "pass1",
		Phone:    "0985944005",
	}

	// // 將資料編碼成 Protocol Buffer 格式（請注意是傳入 Pointer）。
	dataBuffer, _ := proto.Marshal(&data)

	// // 將已經編碼的資料解碼成 protobuf.User 格式。
	var user protobuf.User
	proto.Unmarshal(dataBuffer, &user)

	// 輸出解碼結果。
	fmt.Println(user.Id, " ", user.Account, " ", user.Password, " ", user.Phone)
}
