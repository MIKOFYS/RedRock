package main

import "fmt"

type DigitalProductInfo1 struct {
	Edition   string
	UserName  string
}

type BasicPhoneInfo1 struct {
	CameraName  string
	Soc         string
	Battery     int
	State    string
}

type PhoneInfo1 struct {
	DigitalProductInfo1
	BasicPhoneInfo1
}

func (p *BasicPhoneInfo1)Open(){
	p.State = "open"
}

type DigitalProduct interface {
	Open()
	Close()
	Sell()
}
type Goods interface {
	Sell()
}

func main(){
	a := PhoneInfo1{
		DigitalProductInfo1{
			Edition: "mi10",
			UserName: "yzw",
		},
		BasicPhoneInfo1{
			CameraName: "pix10000000",
			Soc: "kirin980",
			Battery: 40000,
			State : "90%",
		},
	}

	a.Open()
	fmt.Println("STATE",a.State)

}
