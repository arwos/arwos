package webui

func ID(arg string) Node {
	return &object{Type: typeAttr, Key: "id", Value: arg}
}

func Link(arg string) Node {
	return &object{Type: typeAttr, Key: "link", Value: arg}
}

func Class(args ...string) Node {
	return &object{Type: typeAttr, Key: "class", Value: args}
}
