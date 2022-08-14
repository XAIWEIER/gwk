package main

import (
	"flag"
	"fmt"

	"github.com/XAIWEIER/gwk/internal/core"
	"github.com/XAIWEIER/gwk/tool"
)

var task, param string

func init() {
	flag.StringVar(&task, "task", "", "task for gwk")
	flag.StringVar(&param, "param", "", "param for task")
}

func main() {
	tool.PrintGreen("====  GWK START  ====")
	flag.Parse()

	// task = "mock"
	// param = `{"model": 1,""}`

	if err := core.Core.Exec(task, param); err != nil {
		tool.PrintRed(err.Error())
	} else {
		tool.PrintGreen(fmt.Sprintf("%v exec success", task))
	}

	tool.PrintGreen("====   GWK END   ====")
}
