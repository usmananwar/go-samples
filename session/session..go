package session

import (
	"sync"
	"time"
)

type Manager struct {
	cookieName string
	lock sync.Mutex
	provider Provider
	maxLifeTime int64
}

type Provider interface {
	SessionInit(sid string) (SessionInterface, error)
	SessionRead(sid string) (SessionInterface, error)
	SessionDestroy(sid string) (SessionInterface, error)
	SessionGC(maxLifeTime int64)
	//SessionUpdate(sid string) error
}

type SessionInterface interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}


type Session struct {
	sid string
	timeAccessed time.Time
	value map[interface{}] interface{}
}

func (st *Session) Set(key, value interface{}) error {
	st.value[key] = value
	st.updateTime()
	return nil
}

func (st *Session) Get(key interface{}) interface{} {
	st.updateTime()
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

func (st *Session) Delete(key interface{}) error {
	delete(st.value, key)
	//st.updateTime()
	return nil
}

func (st *Session) SessionID() string {
	return st.sid
}

func (st *Session) updateTime() {
	st.timeAccessed = time.Now()
}


