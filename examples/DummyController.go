package examples

import (
	"context"
	"errors"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-commons-gox/run"
	"github.com/pip-services3-gox/pip-services3-components-gox/log"
	"sync/atomic"
)

type DummyController struct {
	timer    *run.FixedRateTimer
	logger   *log.CompositeLogger
	message  string
	counter1 int
	counter2 int64
}

func NewDummyController() *DummyController {
	c := &DummyController{
		logger:   log.NewCompositeLogger(),
		message:  "Hello World!",
		counter1: 0,
	}

	c.timer = run.NewFixedRateTimerFromTask(c, 1000, 1000, 5)

	return c
}

func (c *DummyController) Message() string {
	return c.message
}

func (c *DummyController) SetMessage(value string) {
	c.message = value
}

func (c *DummyController) Counter() int {
	return c.counter1
}

func (c *DummyController) SetCounter(value int) {
	c.counter1 = value
}

func (c *DummyController) Configure(ctx context.Context, config *config.ConfigParams) {
	c.message = config.GetAsStringWithDefault("message", c.message)
}

func (c *DummyController) SetReferences(ctx context.Context, references refer.IReferences) {
	c.logger.SetReferences(ctx, references)
}

func (c *DummyController) IsOpen() bool {
	return c.timer.IsStarted()
}

func (c *DummyController) Open(ctx context.Context, correlationId string) error {
	c.timer.Start(ctx)
	c.logger.Trace(ctx, correlationId, "Dummy controller opened")
	return nil
}

func (c *DummyController) Close(ctx context.Context, correlationId string) error {
	c.timer.Stop(ctx)
	c.logger.Trace(ctx, correlationId, "Dummy controller closed")
	return nil
}

func (c *DummyController) Notify(ctx context.Context, correlationId string, args *run.Parameters) {
	go func(c *DummyController) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					msg := convert.StringConverter.ToString(r)
					err = errors.New(msg)
				}
				// Send shutdown signal with err to container
				// and close all components
				run.SendShutdownSignalWithErr(ctx, err)
			}
		}()
		atomic.AddInt64(&c.counter2, 1)
	}(c)
	c.logger.Info(ctx, correlationId, "%d - %s", c.counter1, c.message)
	c.counter1++
}
