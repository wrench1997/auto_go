package route

import (
	"context"
	"evenkey/mygrpc"
	pb "evenkey/mygrpc" // 导入生成的 protobuf 代码
	"syscall"

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

func (s *MyServiceServer) MouseClick(ctx context.Context, req *pb.MouseRequest) (*emptypb.Empty, error) {
	// 处理 MouseClick 请求逻辑
	// 获取屏幕缩放比例
	scaleFactor := getScreenScaleFactor()

	// 根据缩放比例调整鼠标坐标
	x := int(float64(req.GetX()) / scaleFactor)
	y := int(float64(req.GetY()) / scaleFactor)
	robotgo.Move(x, y)
	robotgo.Click(req.GetClicktype())
	return &emptypb.Empty{}, nil
}

func getScreenScaleFactor() float64 {
	var (
		user32           = syscall.MustLoadDLL("user32.dll")
		GetDC            = user32.MustFindProc("GetDC")
		GetSystemMetrics = user32.MustFindProc("GetSystemMetrics")
		releaseDC        = user32.MustFindProc("ReleaseDC")
		gdi32            = syscall.MustLoadDLL("gdi32.dll")
		GetDeviceCaps    = gdi32.MustFindProc("GetDeviceCaps")
	)
	hdc, _, _ := GetDC.Call(0)
	defer releaseDC.Call(hdc)
	sw, _, _ := GetSystemMetrics.Call(0)
	dw, _, _ := GetDeviceCaps.Call(hdc, uintptr(118))
	return float64(dw) / float64(sw)
}

func (s *MyServiceServer) ReleaseKey(ctx context.Context, req *pb.KeyRequest) (*emptypb.Empty, error) {

	// 处理 ReleaseKey 请求逻辑

	robotgo.KeyUp(req.Key)
	return &emptypb.Empty{}, nil
}

func (s *MyServiceServer) mustEmbedUnimplementedMyServiceServer() {

}
