package gameagent

import (
	"fmt"
	"log"
	"sync"
	"time"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/service"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SessionInfo struct {
	Agent      Agent
	LastActive time.Time
	Users      []*model.User
}

var (
	sessions   = make(map[string]*SessionInfo)
	sessionMux sync.Mutex
	timeout    = 30 * time.Minute
)

func NewSession(soupID uint, user *model.User, s *service.SoupService, l *i18n.Localizer) (*SessionInfo, error) {
	sessionMux.Lock()
	defer sessionMux.Unlock()

	agent, err := InitDS(soupID, s, l)
	if err != nil {
		return nil, err
	}

	uuid := agent.UUID
	session := &SessionInfo{
		Agent:      agent,
		LastActive: time.Now(),
	}
	sessions[uuid] = session
	session.Users = append(session.Users, user)

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
