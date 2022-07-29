package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "peers"
	app.Usage = "a file-coin tools"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "url",
			Value: "lotus url",
			Usage: "url",
		},
	}

	app.Action = func(c *cli.Context) error {
		urlstr := c.String("url")
		if urlstr == "" {
			fmt.Printf("config url is nil:url:%s", urlstr)
			return nil
		}
		err := Setup(urlstr)
		if err != nil {
			return err
		}
		PeerInfo()
		//r := gin.Default()
		//r.GET("/ping", func(c *gin.Context) {
		//	c.JSON(200, gin.H{
		//		"message": "pong",
		//	})
		//})
		//r.Run(":13836")

		return nil
	}

}

//
//func InitInterval() {
//	var timeout string = "0/30 * * * * *" // 定时器时间区间，默认精度为30s/次：0,30 * * * * *
//	var intervalId int                    // 定时器id
//
//	go func() {
//		num := 0 // 运行次数
//		// 设置时区
//		local, _ := time.LoadLocation("Local")
//		interval := cron.New(cron.WithLocation(local), cron.WithSeconds()) // 设置时区并且精度按秒。
//		_timeout := timeout
//		_intervalId, err := interval.AddFunc(_timeout, func() {
//			num++
//			// 下面调用其他函数
//			TimeInterval(intervalId, num, _timeout)
//
//		})
//		if err != nil {
//			os.Exit(200)
//		}
//		intervalId = int(_intervalId)
//		interval.Start()
//
//		//关闭着计划任务, 但是不能关闭已经在执行中的任务.
//		defer interval.Stop()
//		select {} // 阻塞主线程而不退出
//
//	}()
//}
//
//func TimeInterval(intervalId int, num int, timeout string) {
//	var maxLog int = 5
//	if num < maxLog { // 不必全部打印，只打印前几个即可
//		if intervalId == 0 && num == 0 {
//
//		} else {
//		}
//	} else if num == maxLog {
//	}
//
//	// 其他任务
//	PeerInfo()
//}

func PeerInfo() {
	var ctx = context.TODO()
	scores, err := lotusNode.node.NetPubsubScores(ctx)
	if err != nil {
		return
	}
	for _, peer := range scores {
		if peer.Score.Score > 1 {
			fmt.Printf("%s, %f\n", peer.ID, peer.Score.Score)
			addrs, err := lotusNode.node.NetFindPeer(ctx, peer.ID)
			if err != nil {
				return
			}
			res := peer.ID.String() + "/p2p/" + addrs.Addrs[0].String()
			fmt.Println(res)
		}
	}
}
