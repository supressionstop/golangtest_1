package v1

import (
	"errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"io"
	"softpro6/internal/controller/grpc/v1/pb"
	"softpro6/internal/usecase"
	"softpro6/internal/valueobject"
	"softpro6/pkg/logger"
	"time"
)

func (s GrpcServer) SubscribeOn(stream pb.ProcessorService_SubscribeOnServer) error {
	for {
		// basic credentials
		peerInfo, ok := peer.FromContext(stream.Context())
		if !ok {
			return status.Error(codes.Unauthenticated, "no peer info provided")
		}
		peerAddr := PeerNetAddress(peerInfo.Addr.String())

		clientRequest, err := stream.Recv()
		st, _ := status.FromError(err)

		if err == io.EOF || st.Code() == codes.Canceled {
			s.logger.Info(
				"grpc client disconnected",
				zap.String("client_addr", peerInfo.Addr.String()),
				zap.String("code", st.Code().String()),
			)
			delete(s.peers, peerAddr)
			return nil
		}
		if err != nil {
			s.logger.Error("grpc SubscribeOn", zap.Error(err))
			return err
		}
		s.logger.Debug("grpc server SubscribeOn", zap.Any("request", clientRequest))

		responder, alreadyConnected := s.peers[peerAddr]
		s.logger.Debug("alreadyConnected", alreadyConnected)
		if alreadyConnected {
			responder.Update(clientRequest)
		} else {
			sportList := valueobject.SportsFromArray(clientRequest.Sports)
			newRateResponder := NewResponder(sportList, time.Duration(clientRequest.Microseconds)*time.Microsecond, s.logger, s.getRecentSport)
			newRateResponder.Upstream(stream)
			s.peers[peerAddr] = newRateResponder
		}
	}
}

type Responder struct {
	// Dependencies
	getRecentSport usecase.GetRecentSportsUseCase
	logger         logger.Interface

	// Params
	sports   []valueobject.Sport
	interval time.Duration

	// Vars
	ticker    *time.Ticker
	prevValue RecentSports
}

type RecentSports map[valueobject.Sport]usecase.Sport

func RecentSportsFromSports(sports []usecase.Sport) RecentSports {
	result := RecentSports{}
	for i := range sports {
		result[valueobject.NewSport(sports[i].Name())] = sports[i]
	}
	return result
}

func NewResponder(sports []valueobject.Sport, interval time.Duration, logger logger.Interface, useCase usecase.GetRecentSportsUseCase) *Responder {
	ticker := time.NewTicker(interval)
	ticker.Stop()

	return &Responder{
		sports:         sports,
		ticker:         ticker,
		interval:       interval,
		logger:         logger,
		getRecentSport: useCase,
	}
}

func (r *Responder) Update(request *pb.Subscribe) {
	r.prevValue = nil
	r.ticker.Stop()
	r.sports = valueobject.SportsFromArray(request.Sports)
	r.ticker.Reset(time.Duration(request.Microseconds) * time.Microsecond)
}

func (r *Responder) Upstream(stream pb.ProcessorService_SubscribeOnServer) {
	r.ticker.Reset(r.interval * time.Microsecond)
	go r.upstream(stream)
}

func (r *Responder) upstream(stream pb.ProcessorService_SubscribeOnServer) {
	for {
		select {
		case <-r.ticker.C:
			r.logger.Info("u")
			sports, err := r.getRecentSport.Execute(stream.Context(), r.sports...)
			r.logger.Info("upstream", sports[0].Name(), sports[0].Rate())
			if err != nil {
				r.logger.Error("grpc - subscribeOn - upstream - getRecentSport", zap.Error(err))
				continue
			}

			recentSports := RecentSportsFromSports(sports)
			err = r.upstreamHandle(recentSports, stream)
			r.prevValue = recentSports

			if err == io.EOF {
				r.logger.Info("grpc - subscribeOn - upstream - client terminated connection")
				return
			}
			if err != nil {
				r.logger.Error("grpc send to client", zap.Error(err))
				return
			}
		case <-stream.Context().Done():
			return
		}
	}
}

func (r *Responder) upstreamHandle(recentSports RecentSports, stream pb.ProcessorService_SubscribeOnServer) error {
	isFirstOrAfterUpdate := r.prevValue == nil
	if isFirstOrAfterUpdate {
		sportsData := r.recentToSportData(recentSports)
		return stream.Send(sportsData)
	} else {
		sportsDataDiff, err := r.diffRecent(recentSports, r.prevValue)
		if err != nil {
			return err
		}
		return stream.Send(sportsDataDiff)
	}
}

func (r *Responder) diffRecent(prev, next RecentSports) (*pb.SportsData, error) {
	results := make([]*pb.Result, 0, len(prev))
	for vo, entity := range prev {
		nextEntity, ok := next[vo]
		if !ok {
			return nil, errors.New("new recent does not have some sport")
		}
		diffRate := entity.Rate().Value().Sub(nextEntity.Rate().Value())
		result := &pb.Result{
			SportName: entity.Name(),
			Rate:      diffRate.String(),
		}
		results = append(results, result)
	}

	return &pb.SportsData{Results: results}, nil
}

func (r *Responder) recentToSportData(recentSports RecentSports) *pb.SportsData {
	sportsData := &pb.SportsData{}
	for _, sport := range recentSports {
		sportResult := &pb.Result{
			SportName: sport.Name(),
			Rate:      sport.Rate().Value().String(),
		}
		sportsData.Results = append(sportsData.Results, sportResult)
	}
	return sportsData
}
