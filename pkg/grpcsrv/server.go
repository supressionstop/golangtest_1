package grpcsrv

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"softpro6/pkg/logger"
)

type GrpcServer struct {
	logger logger.Interface
	srv    *grpc.Server
	lis    net.Listener
	notify chan error
}

func newGrpcServer(srv *grpc.Server, l logger.Interface) *GrpcServer {
	return &GrpcServer{
		srv:    srv,
		notify: make(chan error, 1),
		logger: l,
	}
}

func NewServer(addr string, logger logger.Interface, registerServer func(server *grpc.Server)) (*GrpcServer, error) {
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)
	gSrv := newGrpcServer(grpcServer, logger)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal("grpc listen", zap.Error(err))
	}
	gSrv.lis = lis

	return gSrv, nil
}

func (s *GrpcServer) start() {
	go func() {
		s.notify <- s.srv.Serve(s.lis)
		close(s.notify)
	}()
}

func (s *GrpcServer) StartLater(waitChan <-chan struct{}) {
	go func() {
		<-waitChan
		s.logger.Info("grpc server started", zap.String("address", s.lis.Addr().String()))
		s.start()
	}()
}

func (s *GrpcServer) Notify() <-chan error {
	return s.notify
}

func (s *GrpcServer) Shutdown() {
	s.srv.GracefulStop()
}
