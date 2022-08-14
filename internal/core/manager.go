package core

import (
	"fmt"
	"sync"

	"github.com/XAIWEIER/gwk/internal/core/engine"
	"github.com/XAIWEIER/gwk/internal/core/mock"
	"github.com/XAIWEIER/gwk/tool"

	"github.com/pkg/errors"
)

var (
	_coreInitOnce sync.Once

	Core *gwkCore
)

type gwkCore struct {
	taskEngineMap map[string]engine.ITaskEngine
}

func (c *gwkCore) register(taskName string, engin engine.ITaskEngine) {
	if !tool.IsNil(engin) {
		c.taskEngineMap[taskName] = engin
	}
}

func (c *gwkCore) Exec(taskName, param string) error {
	if e, ok := c.taskEngineMap[taskName]; ok {
		if err := e.Exec(param); err != nil {
			return errors.Wrap(err, fmt.Sprintf("gwk core exec task:%v err", taskName))
		}
		return nil
	}
	return fmt.Errorf("nonsupport task:%v", taskName)
}

func initGlobManager() *gwkCore {
	_coreInitOnce.Do(func() {
		core := &gwkCore{
			taskEngineMap: make(map[string]engine.ITaskEngine),
		}
		core.register("mock", mock.NewMockEngin())
		Core = core
	})
	return Core
}

func init() {
	initGlobManager()
}
