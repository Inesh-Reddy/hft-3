package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Inesh-Reddy/hft-3/packages/golib/redis"
	"github.com/Inesh-Reddy/hft-3/packages/golib/ws"
)
func main(){
	fmt.Println(`Ticker Service running......`)

	ctx := context.Background()
	log.Println(`Context:`, ctx);
	redis := redis.ConnectToRedis()
	redis.Ping(ctx)
	wes:=ws.ConnectToWs()
	wes.PingHandler();
	_,data,err:=wes.ReadMessage()
	if err != nil {
		log.Fatal("erroredading data from ws:", err)
	}
	log.Println(string(data))


}