package examples

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/pip-services3-go/pip-services3-components-go/log"
)

type DummyController struct {
	timer   *run.FixedRateTimer
	logger  *log.CompositeLogger
	message string
	counter int
}

func NewDummyController() *DummyController {
	c := &DummyController{
		logger:  log.NewCompositeLogger(),
		message: "Hello World!",
		counter: 0,
	}

	c.timer = run.NewFixedRateTimerFromTask(c, 1000, 1000)

	return c
}

func (c *DummyController) Message() string {
	return c.message
}

func (c *DummyController) SetMessage(value string) {
	c.message = value
}

func (c *DummyController) Counter() int {
	return c.counter
}

func (c *DummyController) SetCounter(value int) {
	c.counter = value
}

func (c *DummyController) Configure(config *config.ConfigParams) {
	c.message = config.GetAsStringWithDefault("message", c.message)
}

func (c *DummyController) SetReferences(references refer.IReferences) {
	c.logger.SetReferences(references)
}

func (c *DummyController) IsOpen() bool {
	return c.timer.IsStarted()
}

func (c *DummyController) Open(correlationId string) error {
	c.timer.Start()
	c.logger.Trace(correlationId, "Dummy controller opened")
	return nil
}

func (c *DummyController) Close(correlationId string) error {
	c.timer.Stop()
	c.logger.Trace(correlationId, "Dummy controller closed")
	return nil
}

func (c *DummyController) Notify(correlationId string, args *run.Parameters) {
	c.logger.Info(correlationId, "%d - %s", c.counter, c.message)
	c.counter++
}
