package gameagent

type Agent interface {
	Start() error
	Ask(string) (*GameResponse, error)
}
