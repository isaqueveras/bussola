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
	Type     string `json:"type"` // line, bar, pie, etc
	Data     any    `json:"data"`
	Options  any    `json:"options"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

func NewChart(title string, chartType string) *Chart {
	return &Chart{
		Title: title,
		Type:  chartType,
	}
}

func (c *Chart) Render() map[string]any {
	return map[string]any{
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
	Headers     []string         `json:"headers"`
	Data        []map[string]any `json:"data"`
	PageSize    int              `json:"pageSize"`
	CurrentPage int              `json:"currentPage"`
	Title       string           `json:"title"`
}

func NewTable(title string, headers []string) *Table {
	return &Table{
		Title:       title,
		Headers:     headers,
		PageSize:    10,
		CurrentPage: 1,
	}
}

func (t *Table) Render() map[string]any {
	return map[string]any{
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
	Title       string  `json:"title"`
	Value       any     `json:"value"`
	Target      any     `json:"target,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	Trend       float64 `json:"trend,omitempty"`
	Description string  `json:"description"`
}

// NewIndicator create a new indicator
func NewIndicator(title string) *Indicator {
	return &Indicator{Title: title}
}

func (i *Indicator) Render() map[string]any {
	return map[string]any{
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

// NewProgressBar create a new progress bar
func NewProgressBar(title string) *ProgressBar {
	return &ProgressBar{
		Title:       title,
		MaxValue:    100,
		ShowPercent: true,
	}
}

func (p *ProgressBar) Render() map[string]any {
	return map[string]any{
		"type":        "progressBar",
		"title":       p.Title,
		"value":       p.Value,
		"maxValue":    p.MaxValue,
		"showPercent": p.ShowPercent,
		"percent":     (p.Value / p.MaxValue) * 100,
	}
}

type FilterBar struct {
	BaseWidget
	Title   string   `json:"title"`
	Filters []Filter `json:"filters"`
}

type Filter interface {
	Render() map[string]any
}

func NewFilterBar(title string) *FilterBar {
	return &FilterBar{
		Title:   title,
		Filters: []Filter{},
	}
}

func (f *FilterBar) AddFilter(filter Filter) {
	f.Filters = append(f.Filters, filter)
}

func (f *FilterBar) Render() map[string]any {
	filters := []map[string]any{}
	for _, flt := range f.Filters {
		filters = append(filters, flt.Render())
	}
	return map[string]any{
		"type":    "filterBar",
		"title":   f.Title,
		"filters": filters,
	}
}

type FilterDate struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

func NewFilterDate(label, key string) *FilterDate {
	return &FilterDate{Label: label, Key: key}
}

func (f *FilterDate) Render() map[string]any {
	return map[string]any{
		"type":  "date",
		"label": f.Label,
		"key":   f.Key,
	}
}

type FilterSelect struct {
	Label   string   `json:"label"`
	Key     string   `json:"key"`
	Options []string `json:"options"`
}

func NewFilterSelect(label, key string, options []string) *FilterSelect {
	return &FilterSelect{Label: label, Key: key, Options: options}
}

func (f *FilterSelect) Render() map[string]any {
	return map[string]any{
		"type":    "select",
		"label":   f.Label,
		"key":     f.Key,
		"options": f.Options,
	}
}

type FilterText struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

func NewFilterText(label, key string) *FilterText {
	return &FilterText{Label: label, Key: key}
}

func (f *FilterText) Render() map[string]any {
	return map[string]any{
		"type":  "text",
		"label": f.Label,
		"key":   f.Key,
	}
}

type FilterBool struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

func NewFilterBool(label, key string) *FilterBool {
	return &FilterBool{Label: label, Key: key}
}

func (f *FilterBool) Render() map[string]any {
	return map[string]any{
		"type":  "bool",
		"label": f.Label,
		"key":   f.Key,
	}
}

type FilterNumber struct {
	Label string  `json:"label"`
	Key   string  `json:"key"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}

func NewFilterNumber(label, key string, min, max float64) *FilterNumber {
	return &FilterNumber{Label: label, Key: key, Min: min, Max: max}
}

func (f *FilterNumber) Render() map[string]any {
	return map[string]any{
		"type":  "number",
		"label": f.Label,
		"key":   f.Key,
		"min":   f.Min,
		"max":   f.Max,
	}
}

type FilterRange struct {
	Label string  `json:"label"`
	Key   string  `json:"key"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}

func NewFilterRange(label, key string, min, max float64) *FilterRange {
	return &FilterRange{Label: label, Key: key, Min: min, Max: max}
}

func (f *FilterRange) Render() map[string]any {
	return map[string]any{
		"type":  "range",
		"label": f.Label,
		"key":   f.Key,
		"min":   f.Min,
		"max":   f.Max,
	}
}

type FilterCheckbox struct {
	Label   string   `json:"label"`
	Key     string   `json:"key"`
	Options []string `json:"options"`
}

func NewFilterCheckbox(label, key string, options []string) *FilterCheckbox {
	return &FilterCheckbox{Label: label, Key: key, Options: options}
}

func (f *FilterCheckbox) Render() map[string]any {
	return map[string]any{
		"type":    "checkbox",
		"label":   f.Label,
		"key":     f.Key,
		"options": f.Options,
	}
}

type FilterRadio struct {
	Label   string   `json:"label"`
	Key     string   `json:"key"`
	Options []string `json:"options"`
}

func NewFilterRadio(label, key string, options []string) *FilterRadio {
	return &FilterRadio{Label: label, Key: key, Options: options}
}

func (f *FilterRadio) Render() map[string]any {
	return map[string]any{
		"type":    "radio",
		"label":   f.Label,
		"key":     f.Key,
		"options": f.Options,
	}
}

type FilterMultiSelect struct {
	Label   string   `json:"label"`
	Key     string   `json:"key"`
	Options []string `json:"options"`
}

func NewFilterMultiSelect(label, key string, options []string) *FilterMultiSelect {
	return &FilterMultiSelect{Label: label, Key: key, Options: options}
}

func (f *FilterMultiSelect) Render() map[string]any {
	return map[string]any{
		"type":    "multiSelect",
		"label":   f.Label,
		"key":     f.Key,
		"options": f.Options,
	}
}

type FilterSlider struct {
	Label string  `json:"label"`
	Key   string  `json:"key"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Value float64 `json:"value"`
}

func NewFilterSlider(label, key string, min, max, value float64) *FilterSlider {
	return &FilterSlider{Label: label, Key: key, Min: min, Max: max, Value: value}
}

func (f *FilterSlider) Render() map[string]any {
	return map[string]any{
		"type":  "slider",
		"label": f.Label,
		"key":   f.Key,
		"min":   f.Min,
		"max":   f.Max,
		"value": f.Value,
	}
}

type FilterToggle struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

func NewFilterToggle(label, key string) *FilterToggle {
	return &FilterToggle{Label: label, Key: key}
}

func (f *FilterToggle) Render() map[string]any {
	return map[string]any{
		"type":  "toggle",
		"label": f.Label,
		"key":   f.Key,
	}
}

type FilterSearch struct {
	Label       string `json:"label"`
	Key         string `json:"key"`
	Placeholder string `json:"placeholder"`
}

func NewFilterSearch(label, key, placeholder string) *FilterSearch {
	return &FilterSearch{Label: label, Key: key, Placeholder: placeholder}
}

func (f *FilterSearch) Render() map[string]any {
	return map[string]any{
		"type":        "search",
		"label":       f.Label,
		"key":         f.Key,
		"placeholder": f.Placeholder,
	}
}

type FilterColor struct {
	Label string `json:"label"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewFilterColor(label, key, value string) *FilterColor {
	return &FilterColor{Label: label, Key: key, Value: value}
}

func (f *FilterColor) Render() map[string]any {
	return map[string]any{
		"type":  "color",
		"label": f.Label,
		"key":   f.Key,
		"value": f.Value,
	}
}

type RankingItem struct {
	Position    int    `json:"position"`
	ImageURL    string `json:"imageUrl,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

func NewRankingItem(position int, title string, description string, imageURL string) RankingItem {
	return RankingItem{
		Position:    position,
		ImageURL:    imageURL,
		Title:       title,
		Description: description,
	}
}

type Ranking struct {
	BaseWidget
	Title string        `json:"title"`
	Items []RankingItem `json:"items"`
	Order string        `json:"order"` // "asc" ou "desc"
}

func NewRanking(title string) *Ranking {
	return &Ranking{
		Title: title,
		Items: []RankingItem{},
		Order: "desc",
	}
}

func (r *Ranking) AddItem(item RankingItem) {
	r.Items = append(r.Items, item)
}

func (r *Ranking) SetOrder(order string) {
	if order == "asc" || order == "desc" {
		r.Order = order
	}
}

func (r *Ranking) Render() map[string]any {
	items := []map[string]any{}
	for _, it := range r.Items {
		item := map[string]any{
			"position":    it.Position,
			"title":       it.Title,
			"description": it.Description,
		}
		if it.ImageURL != "" {
			item["imageUrl"] = it.ImageURL
		}
		items = append(items, item)
	}
	return map[string]any{
		"type":  "ranking",
		"title": r.Title,
		"order": r.Order,
		"items": items,
	}
}
