package preview

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/isaqueveras/bussola"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const (
	cellWidth  = 200
	cellHeight = 150
	padding    = 20
	margin     = 10
)

// GeneratePreview creates a preview image of the dashboard layout
func GeneratePreview(dashboard *bussola.Dashboard, outputPath string) error {
	if dashboard.Layout == nil {
		return nil
	}

	grid := dashboard.Layout
	totalWidth := grid.Columns*cellWidth + (grid.Columns+1)*margin
	totalHeight := grid.Rows*cellHeight + (grid.Rows+1)*margin

	// Create a new white image
	img := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Draw grid lines
	gridColor := color.RGBA{255, 255, 255, 255}
	for row := 0; row <= grid.Rows; row++ {
		y := row * (cellHeight + margin)
		drawHorizontalLine(img, y, totalWidth, gridColor)
	}
	for col := 0; col <= grid.Columns; col++ {
		x := col * (cellWidth + margin)
		drawVerticalLine(img, x, totalHeight, gridColor)
	}

	// Draw cells with components
	for row := range grid.Cells {
		for col := range grid.Cells[row] {
			cell := grid.Cells[row][col]
			if cell != nil && cell.Content != nil {
				x := col*(cellWidth+margin) + margin
				y := row*(cellHeight+margin) + margin
				w := cell.ColSpan*cellWidth + (cell.ColSpan-1)*margin
				h := cell.RowSpan*cellHeight + (cell.RowSpan-1)*margin

				// Draw component rectangle
				drawComponent(img, x, y, w, h, getComponentColor(cell.Content), getComponentName(cell.Content), cell.Content)
			}
		}
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return jpeg.Encode(f, img, &jpeg.Options{
		Quality: 100,
	})
}

func drawHorizontalLine(img *image.RGBA, y, width int, c color.Color) {
	for x := 0; x < width; x++ {
		img.Set(x, y, c)
	}
}

func drawVerticalLine(img *image.RGBA, x, height int, c color.Color) {
	for y := 0; y < height; y++ {
		img.Set(x, y, c)
	}
}

func getComponentName(component bussola.Component) string {
	switch component.(type) {
	case *bussola.Indicator:
		return "Indicator"
	case *bussola.Chart:
		return "Chart"
	case *bussola.Table:
		return "Table"
	case *bussola.ProgressBar:
		return "ProgressBar"
	case *bussola.FilterBar:
		return "FilterBar"
	case *bussola.Ranking:
		return "Ranking"
	default:
		return ""
	}
}

func drawComponent(img *image.RGBA, x, y, w, h int, c color.Color, name string, comp ...bussola.Component) {
	var component bussola.Component
	if len(comp) > 0 {
		component = comp[0]
	}

	if grid, ok := component.(*bussola.Grid); ok {
		rows := grid.Rows
		cols := grid.Columns
		if rows == 0 || cols == 0 {
			return
		}
		cellW := w / cols
		cellH := h / rows
		for row := range grid.Cells {
			for col := range grid.Cells[row] {
				cell := grid.Cells[row][col]
				if cell != nil && cell.Content != nil {
					x0 := x + col*cellW
					y0 := y + row*cellH
					cw := cell.ColSpan * cellW
					ch := cell.RowSpan * cellH
					drawComponent(img, x0, y0, cw, ch, getComponentColor(cell.Content), getComponentName(cell.Content), cell.Content)
				}
			}
		}

		face := basicfont.Face7x13
		label := name
		labelWidth := font.MeasureString(face, label).Ceil()
		labelX := x + (w-labelWidth)/2
		labelY := y + 15
		col := color.RGBA{30, 30, 30, 255}
		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(col),
			Face: face,
			Dot:  fixed.P(labelX, labelY),
		}
		d.DrawString(label)
		return
	}

	if filterBar, ok := component.(*bussola.FilterBar); ok {
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				img.Set(i, j, c)
			}
		}

		borderColor := color.RGBA{100, 100, 100, 255}
		for i := x; i < x+w; i++ {
			img.Set(i, y, borderColor)
			img.Set(i, y+h-1, borderColor)
		}

		for j := y; j < y+h; j++ {
			img.Set(x, j, borderColor)
			img.Set(x+w-1, j, borderColor)
		}

		face := basicfont.Face7x13
		label := name
		labelWidth := font.MeasureString(face, label).Ceil()
		labelX := x + (w-labelWidth)/2
		labelY := y + 18
		col := color.RGBA{30, 30, 30, 255}
		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(col),
			Face: face,
			Dot:  fixed.P(labelX, labelY),
		}
		d.DrawString(label)

		filterCount := len(filterBar.Filters)
		if filterCount > 0 {
			filterW := (w - 20) / filterCount
			filterH := h - 30
			for i, f := range filterBar.Filters {
				fx := x + 10 + i*filterW
				fy := y + 25
				var fc color.Color
				switch f.(type) {
				case *bussola.FilterDate:
					fc = color.RGBA{200, 230, 255, 255}
				case *bussola.FilterSelect:
					fc = color.RGBA{220, 255, 200, 255}
				case *bussola.FilterText:
					fc = color.RGBA{255, 255, 200, 255}
				case *bussola.FilterBool:
					fc = color.RGBA{255, 220, 220, 255}
				case *bussola.FilterSearch:
					fc = color.RGBA{220, 200, 255, 255}
				case *bussola.FilterCheckbox:
					fc = color.RGBA{255, 240, 200, 255}
				case *bussola.FilterRadio:
					fc = color.RGBA{255, 200, 240, 255}
				case *bussola.FilterMultiSelect:
					fc = color.RGBA{240, 200, 255, 255}
				case *bussola.FilterBar:
					fc = color.RGBA{220, 220, 220, 255}
				case *bussola.FilterNumber:
					fc = color.RGBA{255, 255, 200, 255}
				case *bussola.FilterRange:
					fc = color.RGBA{200, 255, 200, 255}
				case *bussola.FilterToggle:
					fc = color.RGBA{255, 220, 200, 255}
				case *bussola.FilterSlider:
					fc = color.RGBA{200, 255, 220, 255}
				case *bussola.FilterColor:
					fc = color.RGBA{220, 220, 255, 255}
				default:
					fc = color.RGBA{240, 240, 240, 255}
				}

				for i2 := fx; i2 < fx+filterW-8; i2++ {
					for j2 := fy; j2 < fy+filterH-8; j2++ {
						img.Set(i2, j2, fc)
					}
				}

				for i2 := fx; i2 < fx+filterW-8; i2++ {
					img.Set(i2, fy, borderColor)
					img.Set(i2, fy+filterH-9, borderColor)
				}

				for j2 := fy; j2 < fy+filterH-8; j2++ {
					img.Set(fx, j2, borderColor)
					img.Set(fx+filterW-9, j2, borderColor)
				}

				labelF := f.Render()["label"].(string)
				labelFW := font.MeasureString(face, labelF).Ceil()
				labelFX := fx + ((filterW-8)-labelFW)/2
				labelFY := fy + (filterH-8)/2
				d2 := &font.Drawer{
					Dst:  img,
					Src:  image.NewUniform(col),
					Face: face,
					Dot:  fixed.P(labelFX, labelFY),
				}
				d2.DrawString(labelF)

				var typeF string
				switch f.(type) {
				case *bussola.FilterDate:
					typeF = "date"
				case *bussola.FilterSelect:
					typeF = "select"
				case *bussola.FilterText:
					typeF = "text"
				case *bussola.FilterBool:
					typeF = "bool"
				case *bussola.FilterSearch:
					typeF = "search"
				case *bussola.FilterCheckbox:
					typeF = "checkbox"
				case *bussola.FilterRadio:
					typeF = "radio"
				case *bussola.FilterMultiSelect:
					typeF = "multi-select"
				case *bussola.FilterBar:
					typeF = "filter-bar"
				case *bussola.FilterNumber:
					typeF = "number"
				case *bussola.FilterRange:
					typeF = "range"
				case *bussola.FilterToggle:
					typeF = "toggle"
				case *bussola.FilterSlider:
					typeF = "slider"
				case *bussola.FilterColor:
					typeF = "color"
				default:
					typeF = ""
				}

				typeFW := font.MeasureString(face, typeF).Ceil()
				typeFX := fx + ((filterW-8)-typeFW)/2
				typeFY := labelFY + 13
				d3 := &font.Drawer{
					Dst:  img,
					Src:  image.NewUniform(color.RGBA{120, 120, 120, 255}),
					Face: face,
					Dot:  fixed.P(typeFX, typeFY),
				}
				d3.DrawString(typeF)
			}
		}

		return
	}

	// Draw filled rectangle
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			img.Set(i, j, c)
		}
	}

	// Draw border
	borderColor := color.RGBA{100, 100, 100, 255}
	for i := x; i < x+w; i++ {
		img.Set(i, y, borderColor)
		img.Set(i, y+h-1, borderColor)
	}

	for j := y; j < y+h; j++ {
		img.Set(x, j, borderColor)
		img.Set(x+w-1, j, borderColor)
	}

	face := basicfont.Face7x13
	var title string
	switch c := component.(type) {
	case *bussola.Indicator:
		title = c.Title
	case *bussola.Chart:
		title = c.Title
	case *bussola.Table:
		title = c.Title
	case *bussola.ProgressBar:
		title = c.Title
	case *bussola.FilterBar:
		title = c.Title
	case *bussola.Grid:
		title = c.Title
	case *bussola.Ranking:
		title = c.Title
	}

	if title != "" {
		titleW := font.MeasureString(face, title).Ceil()
		titleX := x + (w-titleW)/2
		titleY := y + (h-13)/2
		dTitle := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(color.RGBA{30, 30, 30, 255}),
			Face: face,
			Dot:  fixed.P(titleX, titleY),
		}
		dTitle.DrawString(title)

		typeStr := name
		if typeStr != "" {
			typeW := font.MeasureString(face, typeStr).Ceil()
			typeX := x + (w-typeW)/2
			typeY := titleY + 13
			dType := &font.Drawer{
				Dst:  img,
				Src:  image.NewUniform(color.RGBA{120, 120, 120, 255}),
				Face: face,
				Dot:  fixed.P(typeX, typeY),
			}
			dType.DrawString(typeStr)
		}
	} else {
		label := name
		labelWidth := font.MeasureString(face, label).Ceil()
		labelHeight := 13 // height of Face7x13
		labelX := x + (w-labelWidth)/2
		labelY := y + (h+labelHeight)/2 - 4
		col := color.RGBA{30, 30, 30, 255}
		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(col),
			Face: face,
			Dot:  fixed.P(labelX, labelY),
		}
		d.DrawString(label)
	}
}

func getComponentColor(component bussola.Component) color.Color {
	switch component.(type) {
	case *bussola.Indicator:
		return color.RGBA{173, 216, 230, 255} // Light blue
	case *bussola.Chart:
		return color.RGBA{144, 238, 144, 255} // Light green
	case *bussola.Table:
		return color.RGBA{255, 182, 193, 255} // Light pink
	case *bussola.ProgressBar:
		return color.RGBA{255, 228, 181, 255} // Light yellowish
	case *bussola.Grid:
		return color.RGBA{255, 255, 224, 255} // Light yellow
	case *bussola.FilterBar:
		return color.RGBA{220, 220, 220, 255} // Light gray
	case *bussola.Ranking:
		return color.RGBA{216, 191, 216, 255} // Light purple
	default:
		return color.RGBA{240, 240, 240, 255} // Light gray
	}
}
