package main

type SkyScrapper struct {
}

func (s SkyScrapper) MaxHeight() int {
	return 40
}

func (s SkyScrapper) GetComposition() string {
	return "rooftile, wall, foundation, glass"
}
