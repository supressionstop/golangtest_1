package worker

import "context"

type Worker interface {
	ID() string
	Start(context.Context)
	Stop()
	Notify() <-chan error
}
