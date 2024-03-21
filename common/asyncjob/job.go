package asyncjob

import (
	"context"
	"time"
)

type Job interface {
	Exec(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(durations []time.Duration)
}
type JobState int
type JobHandler func(ctx context.Context) error

var defaultRetryDurations = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
var defaultMaxRetry = 3

const (
	StatePending JobState = iota
	StateRunning
	StateSuccess
	StateFailed
	StateRetry
	StateTimeout
	StateCanceled
	StateUnknown
)

func (job JobState) String() string {
	return [8]string{"Pending", "Running", "Success", "Failed", "Retry", "Timeout", "Canceled", "Unknown"}[job]
}

type jobConfig struct {
	Name           string
	MaxRetry       int
	RetryDurations []time.Duration
}
type job struct {
	config     jobConfig
	hander     JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler, options ...OptionHdl) *job {
	j := job{
		config: jobConfig{
			MaxRetry:       defaultMaxRetry,
			RetryDurations: defaultRetryDurations,
		},
		hander:     handler,
		state:      StatePending,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
	for i := range options {
		options[i](&j.config)
	}
	return &j
}

type OptionHdl func(config *jobConfig)

func WithRetryDurations(durations []time.Duration) OptionHdl {
	return func(config *jobConfig) {
		config.RetryDurations = durations
	}
}

func (j *job) Exec(ctx context.Context) error {
	j.state = StateRunning
	err := j.hander(ctx)
	if err != nil {
		j.state = StateFailed
		return err
	}
	j.state = StateSuccess
	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.state = StateRetry
	j.retryIndex++
	if j.retryIndex >= j.config.MaxRetry {
		j.state = StateFailed
		return nil
	}
	select {
	case <-ctx.Done():
		j.state = StateCanceled
		return nil
	case <-time.After(j.config.RetryDurations[j.retryIndex]):
		return j.Exec(ctx)
	}
}
