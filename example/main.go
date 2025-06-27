package main

import (
	"fmt"

	"github.com/isaqueveras/bussola"
)

func main() {
	// Create a new dashboard
	dashboard := bussola.NewDashboard("Analytics Dashboard", "Real-time performance metrics")

	// Create a main grid with 2 rows and 3 columns
	mainGrid := bussola.NewGrid("Main Grid", 2, 3)

	// Create some indicators
	sales := bussola.NewIndicator("Total Sales", 15234.56)
	sales.Unit = "R$"
	sales.Trend = 5.7 // 5.7% increase
	sales.Target = "http://localhost:4040/api/v1/query/sales/indicator"

	users := bussola.NewIndicator("Active Users", 1234)
	users.Description = "Currently active users"

	conversion := bussola.NewProgressBar("Conversion Rate")
	conversion.Value = 67.5

	// Create a chart
	revenueChart := bussola.NewChart("Revenue Over Time", "line")
	revenueChart.Data = []float64{1200, 1900, 3000, 5000, 4100, 4500}
	revenueChart.Options = map[string]interface{}{
		"xAxis": []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		"color": "#1976D2",
	}

	// Create a table
	userTable := bussola.NewTable("Recent Users", []string{"ID", "Name", "Last Access", "Status"})
	userTable.Data = []map[string]interface{}{
		{"id": 1, "name": "John Doe", "lastAccess": "2025-06-26", "status": "Active"},
		{"id": 2, "name": "Jane Smith", "lastAccess": "2025-06-25", "status": "Inactive"},
	}

	// Add components to the grid
	mainGrid.AddItem(sales, 0, 0, 1, 1)
	mainGrid.AddItem(users, 0, 1, 1, 1)
	mainGrid.AddItem(conversion, 0, 2, 1, 1)
	mainGrid.AddItem(revenueChart, 1, 0, 1, 2) // Spans 2 columns
	mainGrid.AddItem(userTable, 1, 2, 1, 1)

	// Set the main grid as the dashboard layout
	dashboard.SetLayout(mainGrid)

	// Customize the dashboard theme
	dashboard.SetTheme(&bussola.Theme{
		Primary:    "#2196F3",
		Secondary:  "#FFC107",
		Background: "#F5F5F5",
		TextColor:  "#212121",
		FontFamily: "Inter, sans-serif",
	})

	// Generate and print the JSON
	fmt.Printf("Dashboard JSON:\n%v\n", dashboard.GenerateJSON())
}
