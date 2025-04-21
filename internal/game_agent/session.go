package gameagent

import (
	"fmt"
	"log"
	"sync"
	"time"
	"uooobarry/soup/internal/service"
)

type SessionInfo struct {
	Agent      Agent
	LastActive time.Time
}

var (
	sessions   = make(map[string]*SessionInfo)
	sessionMux sync.Mutex
	timeout    = 30 * time.Minute
)

func NewSession(soupID uint, s *service.SoupService) (*SessionInfo, error) {
	sessionMux.Lock()
	defer sessionMux.Unlock()

	agent, err := InitDS(soupID, s)
	if err != nil {
		return nil, err
	}

	uuid := agent.UUID
	session := &SessionInfo{
		Agent:      agent,
		LastActive: time.Now(),
	}
	sessions[uuid] = session

	return session, nil
}

func init() {
	go cleanupSessions()
}

func cleanupSessions() {
	for {
		time.Sleep(5 * time.Minute)
		sessionMux.Lock()
		now := time.Now()
		for uuid, info := range sessions {
			log.Println(info.LastActive)
			if now.Sub(info.LastActive) > timeout {
				delete(sessions, uuid)
				log.Println(fmt.Sprintf("Deleted timeout session %s", uuid))
			}
		}
		sessionMux.Unlock()
	}
}

func GetSession(uuid string) (*SessionInfo, bool) {
	sessionMux.Lock()
	defer sessionMux.Unlock()

	session, exists := sessions[uuid]
	if exists {
		session.UpdateLastActive()
	}
	return session, exists
}

func EndSession(uuid string) {
	sessionMux.Lock()
	defer sessionMux.Unlock()

	delete(sessions, uuid)
}

func (s *SessionInfo) UpdateLastActive() {
	s.LastActive = time.Now()
}
