package gameagent

import (
	"sync"
)

var (
	sessions   = make(map[string]Agent)
	sessionMux sync.Mutex
)

func NewSession(soupID uint) (Agent, error) {
	sessionMux.Lock()
	defer sessionMux.Unlock()

	agent, err := InitDS(soupID)
	if err != nil {
		return nil, err
	}

	uuid := agent.UUID
	sessions[uuid] = agent

	return agent, nil
}

func GetSession(uuid string) (Agent, bool) {
	sessionMux.Lock()
	defer sessionMux.Unlock()

	agent, exists := sessions[uuid]
	return agent, exists
}

func EndSession(uuid string) {
	sessionMux.Lock()
	defer sessionMux.Unlock()

	delete(sessions, uuid)
}
