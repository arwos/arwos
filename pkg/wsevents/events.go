package wsevents

import "github.com/arwos/arwos/pkg/enums"

//go:generate easyjson

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const AgentRegister uint = 1001

//easyjson:json
type AgentRegisterModel struct {
	Token string `json:"token"`
	OS    string `json:"os"`
}

//easyjson:json
type AgentRegisterStatusModel struct {
	Status string `json:"status"`
}

func (v *AgentRegisterStatusModel) IsValid() bool {
	return enums.IsValidConnStatus(v.Status)
}
