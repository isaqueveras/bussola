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

func (g *Grid) AddNext(component Component, values ...int) {
	rowSpan, colSpan := 1, 1
	if len(values) > 0 {
		rowSpan = values[0]
	}

	if len(values) > 1 {
		colSpan = values[1]
	}

	for row := 0; row < g.Rows; row++ {
		for col := 0; col < g.Columns; col++ {
			occupied := false
			for r := row; r < row+rowSpan && r < g.Rows; r++ {
				for c := col; c < col+colSpan && c < g.Columns; c++ {
					if g.Cells[r][c] != nil {
						occupied = true
						break
					}
				}
				if occupied {
					break
				}
			}
			if !occupied {
				g.AddItem(component, row, col, rowSpan, colSpan)
				return
			}
		}
	}
}

// Render generates a JSON representation of the grid
func (g *Grid) Render() map[string]any {
	result := make(map[string]any)
	result["title"] = g.Title
	result["rows"] = g.Rows
	result["columns"] = g.Columns
	result["spacing"] = g.Spacing
	result["padding"] = g.Padding

	cells := []map[string]any{}
	for i := range g.Cells {
		for j := range g.Cells[i] {
			if cell := g.Cells[i][j]; cell != nil {
				cellData := map[string]any{
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
