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
