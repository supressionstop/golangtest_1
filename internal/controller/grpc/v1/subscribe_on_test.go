package v1_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	v1 "softpro6/internal/controller/grpc/v1"
	"softpro6/internal/controller/grpc/v1/pb"
	"softpro6/internal/entity/sport"
	"softpro6/internal/usecase"
	"softpro6/internal/valueobject"
	"softpro6/pkg/logger"
	"testing"
	"time"
)

func TestGrpcServer_SubscribeOn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		ctxTst := context.Background()
		// server
		ourGrpcServer, getRecentMock := makeMocks(t)
		lis, grpcServer := runGrpcServer(t, ourGrpcServer)
		t.Cleanup(func() {
			grpcServer.Stop()
		})
		// client
		client, err := newGrpcClient(lis)
		require.NoError(t, err)
		t.Cleanup(func() {
			_ = client.conn.Close()
		})
		sports := []usecase.Sport{
			sport.NewBaseball(valueobject.NewRate("foo", 1.0001), time.Now()),
		}
		getRecentMock.
			On("Execute", mock.Anything, mock.Anything).
			Return(sports, nil).
			Once()

		// act
		sub, err := client.subscribe(
			ctxTst,
			&pb.Subscribe{
				Sports:       []string{"baseball"},
				Microseconds: 1,
			},
			"testSub",
		)

		// assert
		require.NoError(t, err)
		got := <-sub.receive.data
		want := &pb.SportsData{
			Results: []*pb.Result{
				{
					SportName: "baseball",
					Rate:      "1.0001",
				},
			},
		}
		assert.Equal(t, want.Results, got.Results)
	})
	t.Run("diff", func(t *testing.T) {
		// arrange
		ctxTst := context.Background()
		// server
		ourGrpcServer, getRecentMock := makeMocks(t)
		lis, grpcServer := runGrpcServer(t, ourGrpcServer)
		t.Cleanup(func() {
			grpcServer.Stop()
		})
		// client
		client, err := newGrpcClient(lis)
		require.NoError(t, err)
		t.Cleanup(func() {
			_ = client.conn.Close()
		})
		// 1st data
		sports1 := []usecase.Sport{
			sport.NewBaseball(valueobject.NewRate("foo", 1.0001), time.Now()),
		}
		getRecentMock.
			On("Execute", mock.Anything, mock.Anything).
			Return(sports1, nil).
			Once()
		// 2nd data
		sports2 := []usecase.Sport{
			sport.NewBaseball(valueobject.NewRate("foo", 1.5001), time.Now()),
		}
		getRecentMock.
			On("Execute", mock.Anything, mock.Anything).
			Return(sports2, nil).
			Once()
		// 3rd data
		sports3 := []usecase.Sport{
			sport.NewBaseball(valueobject.NewRate("foo", 0.0001), time.Now()),
		}
		getRecentMock.
			On("Execute", mock.Anything, mock.Anything).
			Return(sports3, nil).
			Once()

		// act
		sub, err := client.subscribe(
			ctxTst,
			&pb.Subscribe{
				Sports:       []string{"baseball"},
				Microseconds: 1,
			},
			"testSub",
		)
		require.NoError(t, err)

		// assert
		// 1
		got := <-sub.receive.data
		want := &pb.SportsData{
			Results: []*pb.Result{
				{
					SportName: "baseball",
					Rate:      "1.0001",
				},
			},
		}
		assert.Equal(t, want.Results, got.Results)
		// 2
		got = <-sub.receive.data
		want = &pb.SportsData{
			Results: []*pb.Result{
				{
					SportName: "baseball",
					Rate:      "0.5",
				},
			},
		}
		assert.Equal(t, want.Results, got.Results)
		// 3
		got = <-sub.receive.data
		want = &pb.SportsData{
			Results: []*pb.Result{
				{
					SportName: "baseball",
					Rate:      "-1.5",
				},
			},
		}
		assert.Equal(t, want.Results, got.Results)
	})
}

type getRecentSportMock struct {
	mock.Mock
}

func (uc *getRecentSportMock) Execute(ctx context.Context, sports ...valueobject.Sport) ([]usecase.Sport, error) {
	args := uc.Called()
	return args.Get(0).([]usecase.Sport), args.Error(1)
}

func runGrpcServer(t *testing.T, grpcServer *v1.GrpcServer) (*bufconn.Listener, *grpc.Server) {
	t.Helper()
	listener := bufconn.Listen(1024 * 1024)
	srv := grpc.NewServer()
	pb.RegisterProcessorServiceServer(srv, grpcServer)
	go func() {
		err := srv.Serve(listener)
		if err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	return listener, srv
}

func makeMocks(t *testing.T) (*v1.GrpcServer, *getRecentSportMock) {
	t.Helper()
	l := logger.New("debug", "test", "test")
	ucMock := new(getRecentSportMock)
	grpcServer := v1.NewGrpcServer(ucMock, l)

	return grpcServer, ucMock
}

// grpc client
type grpcClient struct {
	conn     *grpc.ClientConn
	pbClient pb.ProcessorServiceClient
	subs     subscriptions
}

func newGrpcClient(listener *bufconn.Listener) (grpcClient, error) {
	dialer := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return grpcClient{}, err
	}

	pbClient := pb.NewProcessorServiceClient(conn)

	return grpcClient{
		conn:     conn,
		pbClient: pbClient,
		subs:     make(map[subscriptionId]subscription),
	}, nil
}

type subscriptions map[subscriptionId]subscription

type subscriptionId string

type subscription struct {
	id      subscriptionId
	client  pb.ProcessorService_SubscribeOnClient
	receive *communicator
}

type communicator struct {
	data chan *pb.SportsData
	err  chan error
}

func (s *subscription) startReceiving() {
	receive := &communicator{
		data: make(chan *pb.SportsData),
		err:  make(chan error),
	}

	go func(comm *communicator) {
		for {
			response, err := s.client.Recv()
			if err != nil {
				comm.err <- err
				return
			}
			comm.data <- response
		}
	}(receive)

	s.receive = receive
}

func (s *subscription) updateSubscription(subscribe *pb.Subscribe) error {
	return s.client.Send(subscribe)
}

func (c grpcClient) subscribe(ctx context.Context, subscribe *pb.Subscribe, subName string) (subscription, error) {
	subId := subscriptionId(subName)
	_, isAlreadySubscribed := c.subs[subId]
	if isAlreadySubscribed {
		return subscription{}, fmt.Errorf("%s already subscribed", subId)
	}

	subClient, err := c.pbClient.SubscribeOn(ctx)
	if err != nil {
		return subscription{}, err
	}

	sub := subscription{
		id:     subId,
		client: subClient,
	}

	err = sub.client.Send(subscribe)
	if err != nil {
		return subscription{}, err
	}

	sub.startReceiving()

	c.subs[subId] = sub

	return sub, nil
}
