package types

const (
	PermTypeMenu = "menu"
	PermTypePage = "page"
	PermTypeElem = "Elem"
)

var PermType = &permType{
	Menu: PermTypeMenu,
	Page: PermTypePage,
	Elem: PermTypeElem,
}

type permType struct {
	Menu string
	Page string
	Elem string
}
