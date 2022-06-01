package main

import (
	"context"
	"flag"
	"grpc-test/pb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:52022", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	sumTwo(conn)
	sumNums(conn)
	getUserNotification(conn)
	getUserInfo(conn)
}

// 对两个数求和
func sumTwo(conn *grpc.ClientConn) {
	c := pb.NewSumNumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SumTwoNums(ctx, &pb.SumTwoRequest{Num1: 10, Num2: 21})
	if err != nil {
		log.Fatalf("could not get sum: %v", err)
	}
	log.Printf("result: %d", r.Sumvalue)
}

//使用客户端流模式对任意个数求和
func sumNums(conn *grpc.ClientConn) {
	c := pb.NewSumNumClient(conn)
	var testNums = []int64{1, 2, 3, 4, 5, 6}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.SumNums(ctx)
	if err != nil {
		log.Print("sendcreate error", err)
	}
	for _, num := range testNums {
		if err := stream.Send(&pb.NumsRequest{Num: num}); err != nil {
			log.Print("send error", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Print("receive error", err)
	}
	log.Printf("receive sum:%d", reply.Sumvalue)
}

//使用服务器端流模式根据用户id获取用户的所有通知信息
func getUserNotification(conn *grpc.ClientConn) {
	c := pb.NewUserClient(conn)
	stream, err := c.UserNotification(context.Background(), &pb.UserRequest{Uid: 1})
	if err != nil {
		log.Print("get stream fail", err)
	}
	log.Print("start stream")
	for {
		notice, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print("recv err", err.Error())
			break
		}
		log.Printf("get notice : %s", notice.Notification)
	}
	log.Print("end stream")
}
func getUserInfo(conn *grpc.ClientConn) {
	uids := []int64{1, 2}
	c := pb.NewUserClient(conn)
	stream, err := c.GetUserInfo(context.Background())
	if err != nil {
		log.Print("sendcreate error", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Print("recv err", err)
			}
			log.Printf("get user name : %s,age:%d,brief:%s", in.Name, in.Age, in.Brief)
		}
	}()
	for _, uid := range uids {
		err = stream.Send(&pb.UserRequest{Uid: uid})
		if err != nil {
			log.Print("send uid error", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
