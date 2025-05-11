package webui

//go:generate easyjson

type Node interface {
	setNodes(args []Node) Node
	getType() objType
}

type objType int

const (
	typeTag  objType = 1
	typeAttr objType = 2
)

//easyjson:json
type object struct {
	Type  objType `json:"type"`
	Key   string  `json:"key"`
	Value any     `json:"value,omitempty"`
	Attr  []any   `json:"attr,omitempty"`
	Tags  []any   `json:"tags,omitempty"`
}

func (o *object) getType() objType {
	return o.Type
}

func (o *object) setNodes(args []Node) Node {
	for _, arg := range args {
		switch arg.getType() {
		case typeTag:
			o.Tags = append(o.Tags, arg)
		case typeAttr:
			o.Attr = append(o.Attr, arg)
		}
	}

	return o
}
