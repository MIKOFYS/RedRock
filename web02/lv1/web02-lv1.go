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

func main(){
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

	fmt.Println("UpName:",Hiiro.UpInformation.UpName)
	fmt.Println("Signature:",Hiiro.UpInformation.Signature)
	fmt.Println("Subscribers:",Hiiro.UpInformation.Subscribers)
	fmt.Println("UpInformation:",Hiiro.UpInformation.VIP)
	fmt.Println("VideoTittle:",Hiiro.VideoInformation.VideoTittle)
	fmt.Println("SubmitTime:",Hiiro.VideoInformation.SubmitTime)
	fmt.Println("Like:",Hiiro.VideoInformation.Like)
	fmt.Println("Coin:",Hiiro.VideoInformation.Coin)
	fmt.Println("Star:",Hiiro.VideoInformation.Star)
	fmt.Println("Share:",Hiiro.VideoInformation.Share)
}
