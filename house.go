package main

type House struct {
}

func (h House) MaxHeight() int {
	return 10
}

func (h House) GetComposition() string {
	return "rooftile, wall, foundation"
}
