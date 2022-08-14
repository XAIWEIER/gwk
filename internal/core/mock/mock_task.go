package mock

import (
	"encoding/json"
	"fmt"

	"github.com/XAIWEIER/gwk/internal/gwkerr"
	"github.com/XAIWEIER/gwk/tool"
)

type reqParam struct {
	Model int `json:"model,omitempty"`
}

type engin struct {
	name string
}

func (e *engin) Exec(param string) error {
	req := reqParam{}
	if err := json.Unmarshal([]byte(param), &req); err != nil {
		return gwkerr.ErrInvalidTaskParam(err)
	}
	rsp := fmt.Sprintf("%v exec success, mode=%v", e.name, req.Model)
	tool.PrintGreen(rsp)
	return nil
}

func NewMockEngin() *engin {
	return &engin{
		name: "mock_engin",
	}
}
