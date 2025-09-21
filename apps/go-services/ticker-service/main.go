package main

import (
	pb "github.com/Inesh-Reddy/HFT-3/packages/golib/proto"
)
type TickerServer struct {
	pb.UnimplementedTickerSeerviceServer
	
}
func main(){
}

// func main(){
// 	fmt.Println(`Ticker Service running......`)

// 	ctx := context.Background()
// 	log.Println(`Context:`, ctx);
// 	redis := redis.ConnectToRedis()
// 	redis.Ping(ctx)
// 	wes:=ws.ConnectToWs()
// 	wes.PingHandler();
// 	_,data,err:=wes.ReadMessage()
// 	if err != nil {
// 		log.Fatal("erroredading data from ws:", err)
// 	}
// 	log.Println(string(data))


// }