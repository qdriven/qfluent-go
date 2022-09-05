// Package golink 连接
package golink

import (
	"context"
	pb "github.com/qdriven/qfluent-go/app/stresstests/proto"
	"github.com/qdriven/qfluent-go/app/stresstests/server/client"
	model "github.com/qdriven/qfluent-go/pkg/qcurl"
	"github.com/qdriven/qfluent-go/pkg/qtime"
	"sync"
	"time"
)

// Grpc grpc 接口请求
func Grpc(ctx context.Context, chanID uint64, ch chan<- *model.RequestResults, totalNumber uint64, wg *sync.WaitGroup,
	request *model.Request, ws *client.GrpcSocket) {
	defer func() {
		wg.Done()
	}()
	defer func() {
		_ = ws.Close()
	}()
	for i := uint64(0); i < totalNumber; i++ {
		grpcRequest(chanID, ch, i, request, ws)
	}
	return
}

// grpcRequest 请求
func grpcRequest(chanID uint64, ch chan<- *model.RequestResults, i uint64, request *model.Request,
	ws *client.GrpcSocket) {
	var (
		startTime = time.Now()
		isSucceed = false
		errCode   = model.HTTPOk
	)
	// 需要发送的数据
	conn := ws.GetConn()
	if conn == nil {
		errCode = model.RequestErr
	} else {
		// TODO::请求接口示例
		c := pb.NewApiServerClient(conn)
		var (
			ctx = context.Background()
			req = &pb.Request{
				UserName: request.Body,
			}
		)
		rsp, err := c.HelloWorld(ctx, req)
		// fmt.Printf("rsp:%+v", rsp)
		if err != nil {
			errCode = model.RequestErr
		} else {
			// 200 为成功
			if rsp.Code != 200 {
				errCode = model.RequestErr
			} else {
				isSucceed = true
			}
		}
	}
	requestTime := uint64(qtime.DiffNano(startTime))
	requestResults := &model.RequestResults{
		Time:      requestTime,
		IsSucceed: isSucceed,
		ErrCode:   errCode,
	}
	requestResults.SetID(chanID, i)
	ch <- requestResults
}
