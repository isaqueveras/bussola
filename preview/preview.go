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
	gridColor := color.RGBA{200, 200, 200, 255} // Light gray
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
				drawComponent(img, x, y, w, h, getComponentColor(cell.Content), getComponentName(cell.Content))
			}
		}
	}

	// Save the image
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return jpeg.Encode(f, img, &jpeg.Options{Quality: 90})
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
	default:
		return "Component"
	}
}

func drawComponent(img *image.RGBA, x, y, w, h int, c color.Color, name string) {
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

	// Draw component name centered
	label := name
	face := basicfont.Face7x13
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

func getComponentColor(component bussola.Component) color.Color {
	switch component.(type) {
	case *bussola.Indicator:
		return color.RGBA{173, 216, 230, 255} // Light blue
	case *bussola.Chart:
		return color.RGBA{144, 238, 144, 255} // Light green
	case *bussola.Table:
		return color.RGBA{255, 182, 193, 255} // Light pink
	case *bussola.ProgressBar:
		return color.RGBA{255, 218, 185, 255} // Light orange
	default:
		return color.RGBA{240, 240, 240, 255} // Light gray
	}
}
