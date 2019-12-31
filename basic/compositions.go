package basic

import "time"

// Composition an Object definition
type Composition struct {
	val  int32
	Text string
}

// ExplicitComposing an composition explicitly
type ExplicitComposing struct {
	// date and composition is private,
	// only can use/assign in this file
	date        time.Time
	composition Composition
	Compose     Composition
}

// ImplicitComposing an composition implicit without field name
type ImplicitComposing struct {
	date time.Time
	Composition
}

var (
	// ImplicitCompose ...
	// you can either use "field:value" or "value". But not mix both
	// implicitComposing = ImplicitComposing{
	//   date: time.Now(), // feild:value
	//   Composition{32, "ImplicitComposing"} // value
	// }
	// implicit composition the field is implicitly
	// this force only use "value" initalizer bind
	// is as ImplicitComposing value instead
	ImplicitCompose = ImplicitComposing{time.Now(), Composition{val: 32, Text: "ImplicitComposing"}}

	// ExplicitCompose init with default values
	// is as pointer
	ExplicitCompose = new(ExplicitComposing)
)
