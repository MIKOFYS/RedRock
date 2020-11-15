package main

import "fmt"

type UpInformation struct {
	UpName string
	Subscribers string
	Signature string
	VIP bool
}

type VideoInformation struct {
	VideoTittle string
	SubmitTime string
	Like int
	Coin int
	Star int
	Share int
}

type BiliBiliWebsite struct {
	UpInformation
	VideoInformation
}

func InfoPrint(x BiliBiliWebsite){
	fmt.Println("UpName:",x.UpInformation.UpName)
	fmt.Println("Signature:",x.UpInformation.Signature)
	fmt.Println("Subscribers:",x.UpInformation.Subscribers)
	fmt.Println("UpInformation:",x.UpInformation.VIP)
	fmt.Println("VideoTittle:",x.VideoInformation.VideoTittle)
	fmt.Println("SubmitTime:",x.VideoInformation.SubmitTime)
	fmt.Println("Like:",x.VideoInformation.Like)
	fmt.Println("Coin:",x.VideoInformation.Coin)
	fmt.Println("Star:",x.VideoInformation.Star)
	fmt.Println("Share:",x.VideoInformation.Share)
}

func (i *VideoInformation)LikePlus() {
	i.Like += 1
}

func (i *VideoInformation)CoinPlus01(){
	i.Coin += 1
}

func (i *VideoInformation)CoinPlus02(){
	i.Coin += 2
}

func (i *VideoInformation)StarPlus(){
	i.Star += 1
}

func (i *VideoInformation)Three(){
	i.Like += 1
	i.Coin += 1
	i.Star += 1
}



func SubmitVideo(name string, videoTittle string)(Output BiliBiliWebsite){
	var NewVideo BiliBiliWebsite
	NewVideo.UpName = name
	NewVideo.SubmitTime = videoTittle
	Output = NewVideo
	return
}

func main(){
	var NewVideo BiliBiliWebsite
	Hiiro := BiliBiliWebsite{
		UpInformation{
			UpName: "HiiroVTuber",
			Subscribers: "16.2万",
			Signature: "901 IQ, 不是 debu, cat！米娜我是你的爸爸",
			VIP: true,
		},
		VideoInformation{
			VideoTittle: "【Hiiro】我的悲伤是水做的",
			SubmitTime: "2020-10-27 20:28:47",
			Like: 2.3e+4,
			Coin: 2.0e+4,
			Star: 1.1e+4,
			Share: 1473,
		},

	}
	Hiiro.Three()
    InfoPrint(Hiiro)
    fmt.Println("")

	NewVideo = SubmitVideo("Hirro","yu")
	InfoPrint(NewVideo)
}
