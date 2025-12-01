package main

import (
	"context"
	"log"

	"MuXiFresh-Be-2.0/app/task/cmd/rpc/submission/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:30200",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 禁用 TLS（测试环境）
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("连接失败:%v", err)
	}
	defer conn.Close()
	client := pb.NewSubmissionClientClient(conn)

	TestGetSubmissionInfo(client)

}

func TestGetSubmissionInfo(client pb.SubmissionClientClient) {
	req := &pb.GetSubmissionInfoReq{
		AssignmentID: "692c68c4f2e8336c6e3927de",
		UserId:       "66de7b10e860586498d3c904",
	}

	resp, err := client.GetSubmissionInfo(context.Background(), req)
	if err != nil {
		log.Printf("调用 GetSubmissionInfo 失败: %v", err)
		return
	}

	log.Printf("GetSubmissionInfo 响应: 共 %d 条提交信息", len(resp.SubmissionInfos))
	for _, info := range resp.SubmissionInfos {
		log.Printf("  - SubmissionID: %s, Urls: %v, Time: %s", info.SubmissionID, info.Urls, info.Time)
	}
}
