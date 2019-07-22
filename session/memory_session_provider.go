package session

import (
	"container/list"
	"sync"
	"time"
)

type MemorySessionProvider struct {
	lock sync.Mutex
	sessions map[string] *list.Element
	list *list.List
}


func (provider *MemorySessionProvider) SessionInit(sid string) (SessionInterface, error) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &Session{sid: sid, timeAccessed: time.Now(), value: v}
	element := provider.list.PushBack(newsess)
	provider.sessions[sid] = element
	return newsess, nil
}

func (provider *MemorySessionProvider) SessionRead(sid string) (SessionInterface, error) {
	if element, ok := provider.sessions[sid]; ok {
		return element.Value.(*Session), nil
	} else {
		sess, err := provider.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (provider *MemorySessionProvider) SessionDestroy(sid string) error {
	if element, ok := provider.sessions[sid]; ok {
		delete(provider.sessions, sid)
		provider.list.Remove(element)
		return nil
	}
	return nil
}

func (provider *MemorySessionProvider) SessionGC(maxlifetime int64) {
	provider.lock.Lock()
	defer provider.lock.Unlock()

	for {
		element := provider.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*Session).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			provider.list.Remove(element)
			delete(provider.sessions, element.Value.(*Session).sid)
		} else {
			break
		}
	}
}

func (provider *MemorySessionProvider) SessionUpdate(sid string) error {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	if element, ok := provider.sessions[sid]; ok {
		element.Value.(*Session).timeAccessed = time.Now()
		provider.list.MoveToFront(element)
		return nil
	}
	return nil
}
