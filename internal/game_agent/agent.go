package gameagent

type Agent interface {
	Start() error
	Ask(string, bool) (*GameResponse, error)
}
