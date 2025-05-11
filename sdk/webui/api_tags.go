package webui

// GRID

func Area(args ...Node) Node {
	return (&object{Type: typeTag, Key: "area"}).setNodes(args)
}

func AreaFluid(args ...Node) Node {
	return (&object{Type: typeTag, Key: "areaFluid"}).setNodes(args)
}

func Col(args ...Node) Node {
	return (&object{Type: typeTag, Key: "col"}).setNodes(args)
}

func Row(args ...Node) Node {
	return (&object{Type: typeTag, Key: "row"}).setNodes(args)
}

func IFrame(src string, args ...Node) Node {
	return (&object{Type: typeTag, Key: "iframe", Value: src}).setNodes(args)
}

func Text(args ...string) Node {
	return &object{Type: typeTag, Key: "text", Value: args}
}
