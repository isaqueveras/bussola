package bussola

import "encoding/json"

// Dashboard represents a dashboard in Bussola.
type Dashboard struct {
	BaseWidget
	Title       string `json:"title"`
	Description string `json:"description"`
	Layout      *Grid  `json:"layout"`
	Theme       *Theme `json:"theme"`
}

// Theme represents the visual theme of the dashboard
type Theme struct {
	Primary    string `json:"primary"`
	Secondary  string `json:"secondary"`
	Background string `json:"background"`
	TextColor  string `json:"textColor"`
	FontFamily string `json:"fontFamily"`
}

// NewDashboard creates a new Dashboard instance with default values.
func NewDashboard(title, desc string) *Dashboard {
	return &Dashboard{
		Title:       title,
		Description: desc,
		Theme: &Theme{
			Primary:    "#1976D2",
			Secondary:  "#424242",
			Background: "#FFFFFF",
			TextColor:  "#212121",
			FontFamily: "Roboto, sans-serif",
		},
	}
}

// SetLayout sets the main grid layout for the dashboard
func (d *Dashboard) SetLayout(grid *Grid) {
	d.Layout = grid
}

// SetTheme sets the theme for the dashboard
func (d *Dashboard) SetTheme(theme *Theme) {
	d.Theme = theme
}

// Render generates a JSON representation of the dashboard
func (d *Dashboard) Render() map[string]interface{} {
	result := make(map[string]interface{})
	result["title"] = d.Title
	result["description"] = d.Description
	result["theme"] = d.Theme

	if d.Layout != nil {
		result["layout"] = d.Layout.Render()
	}

	return result
}

// GenerateJSON generates a JSON string representation of the dashboard
func (d *Dashboard) GenerateJSON() string {
	data, _ := json.Marshal(d.Render())
	return string(data)
}
