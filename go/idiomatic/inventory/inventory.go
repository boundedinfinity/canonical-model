package inventory

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ItemKind string

type itemKinds struct {
	Unknown        ItemKind
	PhysicalItem   ItemKind
	PaperItem      ItemKind
	ElectronicItem ItemKind
}

var ItemKinds = itemKinds{
	Unknown:        "unknown",
	PhysicalItem:   "physical-item",
	PaperItem:      "paper-item",
	ElectronicItem: "electronic-item",
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Item interface {
	Kind() ItemKind
	GetName() string
}
