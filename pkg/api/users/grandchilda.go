package users

type GrandchildAPI struct {
	childa ChildA
}

func (g *GrandchildAPI) Call(value string) {
	g.childa.parent.Save(value)
}
