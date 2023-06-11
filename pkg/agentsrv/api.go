package agentsrv

//go:generate easyjson

import (
	"fmt"

	"github.com/deweppro/go-sdk/log"

	"github.com/arwos/arwos/pkg/enums"
	"github.com/arwos/arwos/pkg/wsevents"
	"github.com/deweppro/goppy/plugins/web"
)

func (v *AgentSrv) InjectApi(route web.Router) {
	// ws events
	v.ws.Event(v.registrationEvent, wsevents.AgentRegister)
	route.Get("/_/agent", v.ws.Handling)
	// ajax routes
	route.Get("/api/agent/list", v.GetList)

	route.Get("/api/agent/access", v.ListAccess)
	route.Post("/api/agent/access", v.AddAccess)
	route.Delete("/api/agent/access", v.DelAccess)

	route.Get("/api/agent/tags", v.ListTag)
	route.Post("/api/agent/tags", v.AddTag)
	route.Delete("/api/agent/tags", v.DelTag)

	route.Get("/api/agent/tag-link/{id}", v.ListTagLink)
	route.Post("/api/agent/tag-link/{id}", v.AddTagLink)
	route.Delete("/api/agent/tag-link/{id}", v.DelTagLink)
}

type (
	//easyjson:json
	ListModel struct {
		All    []string    `json:"all"`
		Active []string    `json:"active"`
		Info   []InfoModel `json:"info"`
	}
	//easyjson:json
	InfoModel struct {
		ID    uint64 `json:"id"`
		CID   string `json:"cid"`
		Token string `json:"token"`
		OS    string `json:"os"`
	}
)

func (v *AgentSrv) GetList(ctx web.Context) {
	v.mux.RLock()
	defer v.mux.RUnlock()

	all := make([]string, 0, len(v.all))
	for k := range v.all {
		all = append(all, k)
	}

	active := make([]string, 0, len(v.active))
	for k := range v.active {
		active = append(active, k)
	}

	info := make([]InfoModel, 0, len(v.info))
	for kk, vv := range v.info {
		info = append(info, InfoModel{
			ID:    vv.ID,
			CID:   kk,
			Token: vv.Token,
			OS:    vv.OS,
		})
	}

	model := ListModel{
		All:    all,
		Active: active,
		Info:   info,
	}

	ctx.JSON(200, &model)
}

type (
	//easyjson:json
	AccessModel struct {
		ID    uint64 `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Token string `json:"token"`
	}
)

func (v *AgentSrv) ListAccess(ctx web.Context) {
	var out []AccessModel
	list, err := v.selectAccess()
	if err != nil {
		ctx.Error(400, err)
		return
	}

	for _, s := range list {
		out = append(out, AccessModel{
			ID:    s.ID,
			Name:  s.Name,
			Token: s.Token,
		})
	}

	ctx.JSON(200, &out)
}

func (v *AgentSrv) AddAccess(ctx web.Context) {
	model := AccessModel{}
	if err := ctx.BindJSON(&model); err != nil {
		ctx.Error(400, err)
		return
	}

	id, err := v.insertAccess(model.Name, model.Token)
	if err != nil {
		ctx.Error(400, err)
		return
	}

	v.connUpdateID(model.Token, id)
	v.connActive(model.Token, false)
	if c := v.connByToken(model.Token); c != nil {
		c.Encode(wsevents.AgentRegister, &wsevents.AgentRegisterStatusModel{
			Status: enums.ConnStatusAllow,
		})
	}

	ctx.String(200, "ok")
}

func (v *AgentSrv) DelAccess(ctx web.Context) {
	model := AccessModel{}
	if err := ctx.BindJSON(&model); err != nil {
		ctx.Error(400, err)
		return
	}

	if err := v.deleteAccess(model.Token); err != nil {
		ctx.Error(400, err)
		return
	}

	v.connUpdateID(model.Token, 0)
	v.connActive(model.Token, false)
	if c := v.connByToken(model.Token); c != nil {
		c.Encode(wsevents.AgentRegister, &wsevents.AgentRegisterStatusModel{
			Status: enums.ConnStatusDisallow,
		})
	}

	ctx.String(200, "ok")
}

type (
	//easyjson:json
	TagsListModel []TagItemModel
	//easyjson:json
	TagItemModel struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

func (v *AgentSrv) ListTag(ctx web.Context) {
	model := TagsListModel{}
	list, err := v.selectTags()
	if err != nil {
		ctx.Error(400, err)
		return
	}
	for _, s := range list {
		model = append(model, TagItemModel{
			ID:   s.ID,
			Name: s.Name,
		})
	}
	ctx.JSON(200, &model)
}

type (
	//easyjson:json
	TagsModel []string
)

func (v *AgentSrv) AddTag(ctx web.Context) {
	model := TagsModel{}
	if err := ctx.BindJSON(&model); err != nil {
		ctx.Error(400, err)
		return
	}

	var errTag []string
	for _, tag := range model {
		if _, e := v.insertTag(tag); e != nil {
			log.WithField("tag", tag).Errorf("add agent tags")
			errTag = append(errTag, tag)
		}
	}
	if len(errTag) > 0 {
		ctx.Error(400, fmt.Errorf("tags alredy exist: %v", errTag))
		return
	}

	ctx.String(200, "ok")
}

func (v *AgentSrv) DelTag(ctx web.Context) {
	model := TagsModel{}
	if err := ctx.BindJSON(&model); err != nil {
		ctx.Error(400, err)
		return
	}
	for _, tag := range model {
		if e := v.deleteTag(tag); e != nil {
			log.WithFields(log.Fields{
				"tag": tag,
				"err": e.Error(),
			}).Errorf("delete agent tags")
		}
	}
	ctx.String(200, "ok")
}

type (
	//easyjson:json
	TagLinkModel []uint64
)

func (v *AgentSrv) ListTagLink(ctx web.Context) {
	id, err := ctx.Param("id").Int()
	if err != nil {
		ctx.Error(400, err)
		return
	}
	model := TagsListModel{}
	list, err := v.selectTagLink(uint64(id))
	if err != nil {
		ctx.Error(400, err)
		return
	}
	for _, s := range list {
		model = append(model, TagItemModel{
			ID:   s.ID,
			Name: s.Name,
		})
	}
	ctx.JSON(200, &model)
}

func (v *AgentSrv) AddTagLink(ctx web.Context) {
	id, err := ctx.Param("id").Int()
	if err != nil {
		ctx.Error(400, err)
		return
	}
	model := TagLinkModel{}
	if err = ctx.BindJSON(&model); err != nil {
		ctx.Error(400, err)
		return
	}
	for _, tid := range model {
		if err = v.insertTagLink(uint64(id), tid); err != nil {
			ctx.Error(400, err)
			return
		}
	}
	ctx.String(200, "ok")
}

func (v *AgentSrv) DelTagLink(ctx web.Context) {
	id, err := ctx.Param("id").Int()
	if err != nil {
		ctx.Error(400, err)
		return
	}
	model := TagLinkModel{}
	if err = ctx.BindJSON(&model); err != nil {
		ctx.Error(400, err)
		return
	}
	for _, tid := range model {
		if err = v.deleteTagLink(uint64(id), tid); err != nil {
			ctx.Error(400, err)
			return
		}
	}
	ctx.String(200, "ok")
}
