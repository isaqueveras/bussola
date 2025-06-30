package bussola

// Component represents any visual element that can be added to the dashboard
type Component interface {
	MinSize() Size
	Resize(size Size)
	Position() Position
	Move(pos Position)
	Render() map[string]any
}

// CanvasObject represents a visual component that can be added to a container
type CanvasObject interface {
	Component
	Visible() bool
	Show()
	Hide()
}

// Size represents the dimensions of a component
type Size struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Position represents the position of a component
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Container represents a component that can contain other components
type Container interface {
	Component
	Add(objects ...CanvasObject)
	Remove(obj CanvasObject)
	Objects() []CanvasObject
}
