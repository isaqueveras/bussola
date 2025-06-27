package bussola

// BaseWidget provides common widget functionality
type BaseWidget struct {
	size     Size
	position Position
	hidden   bool
}

func (w *BaseWidget) MinSize() Size      { return w.size }
func (w *BaseWidget) Resize(size Size)   { w.size = size }
func (w *BaseWidget) Position() Position { return w.position }
func (w *BaseWidget) Move(pos Position)  { w.position = pos }
func (w *BaseWidget) Visible() bool      { return !w.hidden }
func (w *BaseWidget) Show()              { w.hidden = false }
func (w *BaseWidget) Hide()              { w.hidden = true }

// Chart represents a chart widget
type Chart struct {
	BaseWidget
	Type     string      `json:"type"` // line, bar, pie, etc
	Data     interface{} `json:"data"`
	Options  interface{} `json:"options"`
	Title    string      `json:"title"`
	Subtitle string      `json:"subtitle"`
}

func NewChart(title string, chartType string) *Chart {
	return &Chart{
		Title: title,
		Type:  chartType,
	}
}

func (c *Chart) Render() map[string]interface{} {
	return map[string]interface{}{
		"type":      "chart",
		"title":     c.Title,
		"subtitle":  c.Subtitle,
		"chartType": c.Type,
		"data":      c.Data,
		"options":   c.Options,
	}
}

// Table represents a table widget with pagination
type Table struct {
	BaseWidget
	Headers     []string                 `json:"headers"`
	Data        []map[string]interface{} `json:"data"`
	PageSize    int                      `json:"pageSize"`
	CurrentPage int                      `json:"currentPage"`
	Title       string                   `json:"title"`
}

func NewTable(title string, headers []string) *Table {
	return &Table{
		Title:       title,
		Headers:     headers,
		PageSize:    10,
		CurrentPage: 1,
	}
}

func (t *Table) Render() map[string]interface{} {
	return map[string]interface{}{
		"type":        "table",
		"title":       t.Title,
		"headers":     t.Headers,
		"data":        t.Data,
		"pageSize":    t.PageSize,
		"currentPage": t.CurrentPage,
	}
}

// Indicator represents a numeric indicator widget
type Indicator struct {
	BaseWidget
	Title       string      `json:"title"`
	Value       interface{} `json:"value"`
	Target      interface{} `json:"target,omitempty"`
	Unit        string      `json:"unit,omitempty"`
	Trend       float64     `json:"trend,omitempty"`
	Description string      `json:"description"`
}

func NewIndicator(title string, value interface{}) *Indicator {
	return &Indicator{
		Title: title,
		Value: value,
	}
}

func (i *Indicator) Render() map[string]interface{} {
	return map[string]interface{}{
		"type":        "indicator",
		"title":       i.Title,
		"value":       i.Value,
		"target":      i.Target,
		"unit":        i.Unit,
		"trend":       i.Trend,
		"description": i.Description,
	}
}

// ProgressBar represents a progress bar widget
type ProgressBar struct {
	BaseWidget
	Title       string  `json:"title"`
	Value       float64 `json:"value"`
	MaxValue    float64 `json:"maxValue"`
	ShowPercent bool    `json:"showPercent"`
}

func NewProgressBar(title string) *ProgressBar {
	return &ProgressBar{
		Title:       title,
		MaxValue:    100,
		ShowPercent: true,
	}
}

func (p *ProgressBar) Render() map[string]interface{} {
	return map[string]interface{}{
		"type":        "progressBar",
		"title":       p.Title,
		"value":       p.Value,
		"maxValue":    p.MaxValue,
		"showPercent": p.ShowPercent,
		"percent":     (p.Value / p.MaxValue) * 100,
	}
}
