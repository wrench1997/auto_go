package route

import (
	"context"
	"evenkey/mygrpc"
	pb "evenkey/mygrpc" // 导入生成的 protobuf 代码

	"github.com/go-vgo/robotgo"
	"google.golang.org/protobuf/types/known/emptypb"
)

// MyServiceServer 实现了MyService服务的接口
type MyServiceServer struct {
	Port     int
	Addr     string
	Username string
	Password string
	Keydelay int
	mygrpc.UnimplementedMyServiceServer
}

func (s *MyServiceServer) Authenticate(ctx context.Context, req *pb.AuthRequest) (*emptypb.Empty, error) {
	// 处理 Authenticate 请求逻辑
	robotgo.KeySleep = 100
	return &emptypb.Empty{}, nil
}

func (s *MyServiceServer) PressKey(ctx context.Context, req *pb.KeyRequest) (*emptypb.Empty, error) {
	// 处理 PressKey 请求逻辑
	robotgo.KeyDown(req.Key)
	// time.Sleep(1000 * time.Millisecond)
	// robotgo.KeyUp(req.Key)
	return &emptypb.Empty{}, nil
}

func (s *MyServiceServer) Click(ctx context.Context, req *pb.KeyRequest) (*emptypb.Empty, error) {
	// 处理 Click 请求逻辑
	robotgo.KeyTap(req.Key)
	return &emptypb.Empty{}, nil
}

func (s *MyServiceServer) ReleaseKey(ctx context.Context, req *pb.KeyRequest) (*emptypb.Empty, error) {

	// 处理 ReleaseKey 请求逻辑

	robotgo.KeyUp(req.Key)
	return &emptypb.Empty{}, nil
}

func (s *MyServiceServer) mustEmbedUnimplementedMyServiceServer() {

}
