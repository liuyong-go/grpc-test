package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-test/pb"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 52022, "The server port")
)

type sumServer struct {
	pb.UnimplementedSumNumServer
}
type userServer struct {
	pb.UnimplementedUserServer
}

// 对两个数求和
func (s *sumServer) SumTwoNums(ctx context.Context, in *pb.SumTwoRequest) (*pb.SumReply, error) {
	result := in.Num1 + in.Num2
	return &pb.SumReply{Sumvalue: result}, nil
}

//使用客户端流模式对任意个数求和
func (s *sumServer) SumNums(stream pb.SumNum_SumNumsServer) error {
	var sumValue int64 = 0
	for {
		numRequest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.SumReply{Sumvalue: sumValue})
		}
		sumValue += numRequest.Num
	}
}

//使用服务器端流模式根据用户id获取用户的所有通知信息
func (s *userServer) UserNotification(user *pb.UserRequest, stream pb.User_UserNotificationServer) error {
	var userNotice = map[int64][]string{
		1: {"hello1", "here1"},
		2: {"hello2", "here2"},
	}
	uid := user.Uid
	if _, ok := userNotice[uid]; !ok {
		return nil
	}
	for _, notice := range userNotice[uid] {
		time.Sleep(1 * time.Second)
		if err := stream.Send(&pb.Notification{Notification: notice}); err != nil {
			return err
		}
	}
	fmt.Println("send over")
	return nil
}

//使用双向流模式根据一组用户id获取用户详细信息
func (s *userServer) GetUserInfo(stream pb.User_GetUserInfoServer) error {
	var userInfo = map[int64]*pb.UserInfo{
		1: {
			Name:  "liu",
			Age:   1,
			Brief: "a good man",
		},
		2: {
			Name:  "yong",
			Age:   2,
			Brief: "a good man",
		},
	}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		uid := in.Uid
		if _, ok := userInfo[uid]; !ok {
			return nil
		}
		if err := stream.Send(userInfo[uid]); err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSumNumServer(s, &sumServer{})
	pb.RegisterUserServer(s, &userServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
