package engine

type ITaskEngine interface {
	Exec(param string) error
}
