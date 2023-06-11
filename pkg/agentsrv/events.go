package agentsrv

import (
	"github.com/arwos/arwos/pkg/enums"
	"github.com/arwos/arwos/pkg/wsevents"
	"github.com/deweppro/goppy/plugins/web"
)

func (v *AgentSrv) registrationEvent(event web.WebsocketEventer, conn web.WebsocketServerProcessor) error {
	model := wsevents.AgentRegisterModel{}
	if err := event.Decode(&model); err != nil {
		return err
	}
	info := ConnInfo{
		ID:    0,
		Token: model.Token,
		OS:    model.OS,
	}
	if id, err := v.selectAccessByToken(info.Token); err == nil {
		info.ID = id
	}
	if err := v.connAdd(info, conn); err != nil {
		return err
	}

	resp := &wsevents.AgentRegisterStatusModel{}
	if info.ID > 0 {
		v.connActive(info.Token, true)
		resp.Status = enums.ConnStatusAllow
	} else {
		resp.Status = enums.ConnStatusDisallow
	}
	conn.EncodeEvent(event, resp)
	return nil
}
