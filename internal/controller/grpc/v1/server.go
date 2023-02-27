package v1

import (
	"softpro6/internal/controller/grpc/v1/pb"
	"softpro6/internal/usecase"
	"softpro6/pkg/logger"
)

type GrpcServer struct {
	// Dependencies
	getRecentSport usecase.GetRecentSportsUseCase
	logger         logger.Interface

	// Vars
	peers map[PeerNetAddress]*Responder

	// Others
	pb.UnimplementedProcessorServiceServer
}

type PeerNetAddress string

func NewGrpcServer(getRecentSport usecase.GetRecentSportsUseCase, logger logger.Interface) *GrpcServer {
	return &GrpcServer{
		getRecentSport: getRecentSport,
		logger:         logger,
		peers:          map[PeerNetAddress]*Responder{},
	}
}
