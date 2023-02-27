package usecase

import (
	"context"
)

type FirstSyncSubscribe struct {
}

func NewFirstSyncSubscribe() *FirstSyncSubscribe {
	return &FirstSyncSubscribe{}
}

func (uc FirstSyncSubscribe) Execute(ctx context.Context, publishers []FirstSyncPublisher) (FirstSyncSubscription, error) {
	return NewFirstSyncSub(ctx, publishers), nil
}

type FirstSyncSub struct {
	syncsLeft int
	syncChan  chan struct{}
	allSynced chan struct{}
}

func NewFirstSyncSub(ctx context.Context, publishers []FirstSyncPublisher) *FirstSyncSub {
	syncsLeft := len(publishers)
	syncChan := make(chan struct{})

	// wait info from pubs
	for i := range publishers {
		go func(i int) {
			select {
			case s := <-publishers[i].IAmSynced():
				syncChan <- s
			case <-ctx.Done():
				return
			}
		}(i)
	}

	// wait for all pubs synced
	allSynced := make(chan struct{})
	go func() {
		for {
			select {
			case <-syncChan:
				syncsLeft--
				if syncsLeft == 0 {
					allSynced <- struct{}{}
					close(allSynced)
					break
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return &FirstSyncSub{
		syncsLeft: syncsLeft,
		syncChan:  syncChan,
		allSynced: allSynced,
	}
}

func (s *FirstSyncSub) IsSynced() <-chan struct{} {
	return s.allSynced
}
