package bussola

// Grid represents a grid layout for organizing widgets in a dashboard.
type Grid struct {
	BaseWidget
	Title   string        `json:"title"`
	Rows    int           `json:"rows"`
	Columns int           `json:"columns"`
	Cells   [][]*GridCell `json:"cells"`
	Spacing float64       `json:"spacing"`
	Padding float64       `json:"padding"`
}

// GridCell represents a cell in the grid
type GridCell struct {
	Row     int       `json:"row"`
	Column  int       `json:"column"`
	RowSpan int       `json:"rowSpan"`
	ColSpan int       `json:"colSpan"`
	Content Component `json:"content"`
}

// NewGrid creates a new grid with the specified number of rows and columns
func NewGrid(title string, rows, columns int) *Grid {
	cells := make([][]*GridCell, rows)
	for i := range cells {
		cells[i] = make([]*GridCell, columns)
	}

	return &Grid{
		Title:   title,
		Rows:    rows,
		Columns: columns,
		Cells:   cells,
		Spacing: 10,
		Padding: 15,
	}
}

// AddItem adds a component to the grid at the specified position
func (g *Grid) AddItem(component Component, row, col, rowSpan, colSpan int) {
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Columns {
		return
	}

	cell := &GridCell{
		Row:     row,
		Column:  col,
		RowSpan: rowSpan,
		ColSpan: colSpan,
		Content: component,
	}

	g.Cells[row][col] = cell
}

// Render generates a JSON representation of the grid
func (g *Grid) Render() map[string]interface{} {
	result := make(map[string]interface{})
	result["title"] = g.Title
	result["rows"] = g.Rows
	result["columns"] = g.Columns
	result["spacing"] = g.Spacing
	result["padding"] = g.Padding

	cells := []map[string]interface{}{}
	for i := range g.Cells {
		for j := range g.Cells[i] {
			if cell := g.Cells[i][j]; cell != nil {
				cellData := map[string]interface{}{
					"row":     cell.Row,
					"column":  cell.Column,
					"rowSpan": cell.RowSpan,
					"colSpan": cell.ColSpan,
					"content": cell.Content.Render(),
				}
				cells = append(cells, cellData)
			}
		}
	}
	result["cells"] = cells

	return result
}
