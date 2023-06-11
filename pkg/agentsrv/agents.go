package agentsrv

import (
	"fmt"
	"sync"

	"github.com/arwos/arwos/pkg/adapters"
	"github.com/deweppro/goppy/plugins/web"
)

type (
	AgentSrv struct {
		ws web.WebsocketServer
		db adapters.DB

		info   map[string]ConnInfo
		active map[string]struct{}
		all    map[string]web.WebsocketServerProcessor
		tokens map[string]string

		mux sync.RWMutex
	}

	ConnInfo struct {
		ID    uint64
		Token string
		OS    string
	}
)

func NewAgentSrv(ws web.WebsocketServerPool, db adapters.DB) *AgentSrv {
	return &AgentSrv{
		ws:     ws.Create("agent"),
		db:     db,
		info:   make(map[string]ConnInfo, 100),
		active: make(map[string]struct{}, 100),
		tokens: make(map[string]string, 100),
		all:    make(map[string]web.WebsocketServerProcessor, 100),
		mux:    sync.RWMutex{},
	}
}

func (v *AgentSrv) hasConnAccess(cid string) bool {
	v.mux.RLock()
	defer v.mux.RUnlock()

	_, ok := v.active[cid]
	return ok
}

func (v *AgentSrv) connAdd(info ConnInfo, conn web.WebsocketServerProcessor) error {
	v.mux.Lock()
	defer v.mux.Unlock()

	if _, ok := v.tokens[info.Token]; ok {
		return fmt.Errorf("connection alredy exist")
	}

	v.all[conn.ConnectID()] = conn
	v.info[conn.ConnectID()] = info
	v.tokens[info.Token] = conn.ConnectID()

	conn.OnClose(func(cid string) {
		v.connDel(cid)
	})
	return nil
}

func (v *AgentSrv) connDel(cid string) {
	v.mux.Lock()
	defer v.mux.Unlock()

	info, ok := v.info[cid]
	if !ok {
		return
	}

	delete(v.all, cid)
	delete(v.info, cid)
	delete(v.active, cid)
	delete(v.tokens, info.Token)
}

func (v *AgentSrv) connActive(token string, status bool) {
	v.mux.Lock()
	defer v.mux.Unlock()

	cid, ok := v.tokens[token]
	if !ok {
		return
	}

	if _, ok = v.all[cid]; ok {
		if status {
			v.active[cid] = struct{}{}
		} else {
			delete(v.active, cid)
		}
	}
}

func (v *AgentSrv) connByToken(token string) web.WebsocketServerProcessor {
	v.mux.RLock()
	defer v.mux.RUnlock()

	cid, ok := v.tokens[token]
	if !ok {
		return nil
	}
	if c, ok := v.all[cid]; ok {
		return c
	}
	return nil
}

func (v *AgentSrv) connUpdateID(token string, id uint64) {
	v.mux.Lock()
	defer v.mux.Unlock()

	cid, ok := v.tokens[token]
	if !ok {
		return
	}

	info, ok := v.info[cid]
	if !ok {
		return
	}

	info.ID = id
	v.info[cid] = info
}
