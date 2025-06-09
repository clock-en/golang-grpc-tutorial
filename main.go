package main

import (
	"fmt"
	"github.com/clock-en/golang-grpc-tutorial/pb"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "96xx",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is 96xx",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	// シリアライズ
	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize:", err)
	}

	// ファイル生成
	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("Can't write:", err)
	}

	// ファイルの読み込み
	in, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't read file:", err)
	}

	// デシリアライズ
	readEmployee := &pb.Employee{}
	err = proto.Unmarshal(in, readEmployee)

	if err != nil {
		log.Fatalln("Can't deserialize:", err)
	}

	// デシリアライズ後の構造が一致していることを確認
	fmt.Println("employee:", employee)
	fmt.Println("readEmployee:", readEmployee)
}
