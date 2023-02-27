package main

import "github.com/wanghuiyt/ding"

func main() {
	d := ding.Webhook{
		AccessToken: "f9f7e29c56a678facf1cafcc7bdd5263a4bef01e4335dc97436cb131c33bd439",    // 上面获取的 access_token
		Secret:      "SEC6e004dca28307acaa0dbf36e601a9e3436b78c6d200f8ddd2b553ef9a409e854", // 上面获取的加签的值
		//艾特所有人,必须同时开启 EnableAt 和 AtAll
		EnableAt: true,
		//AtAll:       true,
	}
	_ = d.SendMessage("完犊子了,数据库被删了,准备跑路吧!!!", "18326147303")
	//d.SendMessage("大哥们好。。。")

}
