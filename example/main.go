package main

import (
	"fmt"
	"log"

	"github.com/isaqueveras/bussola"
	"github.com/isaqueveras/bussola/preview"
)

func main() {
	// Create a new dashboard
	dashboard := bussola.NewDashboard("Analytics Dashboard", "Real-time performance metrics")

	mainGrid := bussola.NewGrid("Main Grid", 4, 3)

	// Create some indicators
	sales := bussola.NewIndicator("Total Sales", 15234.56)
	sales.Unit = "R$"
	sales.Trend = 5.7 // 5.7% increase
	sales.Target = "http://localhost:4040/api/v1/query/sales/indicator"

	users := bussola.NewIndicator("Active Users", 1234)
	users.Description = "Currently active users"

	tma := bussola.NewIndicator("TMA", 2.5) // Average time to action in minutes
	tma.Unit = "min"
	tma.Description = "Average time to action (TMA)"
	// Set the target URL for TMA
	tma.Target = "http://localhost:4040/api/v1/query/tma/indicator"

	conversionRate := bussola.NewProgressBar("Conversion Rate")
	conversionRate.Value = 75.0 // 75% conversion rate
	conversionRate.MaxValue = 100.0
	conversionRate.ShowPercent = true

	nestedGrid := bussola.NewGrid("ProgressBar Grid", 2, 1)
	nestedGrid.AddItem(conversionRate, 0, 0, 1, 1)
	nestedGrid.AddItem(bussola.NewProgressBar("Total Conversion"), 1, 0, 1, 1)

	// Create a chart
	revenueChart := bussola.NewChart("Revenue Over Time", "line")
	revenueChart.Data = []float64{1200, 1900, 3000, 5000, 4100, 4500}
	revenueChart.Options = map[string]interface{}{
		"xAxis": []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		"color": "#1976D2",
	}

	userTable := bussola.NewTable("Recent Users", []string{"ID", "Name", "Last Access", "Status"})
	userTable.Data = []map[string]interface{}{
		{"id": 1, "name": "John Doe", "lastAccess": "2025-06-26", "status": "Active"},
		{"id": 2, "name": "Jane Smith", "lastAccess": "2025-06-25", "status": "Inactive"},
	}

	filterBar := bussola.NewFilterBar("Filtros Gerais")
	filterBar.AddFilter(bussola.NewFilterDate("Periodo", "period"))
	filterBar.AddFilter(bussola.NewFilterSelect("Tipo de Problema", "problem_type", []string{"Todos", "Erro", "Aviso", "Info"}))
	filterBar.AddFilter(bussola.NewFilterText("Nome do Cliente", "client_name"))
	filterBar.AddFilter(bussola.NewFilterSearch("Pesquisar", "search", "Search by name or ID"))

	// Add components to the grid
	mainGrid.AddItem(filterBar, 0, 0, 1, 3)
	mainGrid.AddItem(sales, 1, 0, 1, 1)
	mainGrid.AddItem(users, 1, 1, 1, 1)
	mainGrid.AddItem(tma, 1, 2, 1, 1)
	mainGrid.AddItem(nestedGrid, 2, 0, 1, 1)
	mainGrid.AddItem(revenueChart, 2, 1, 1, 2) // Spans 2 columns
	mainGrid.AddItem(userTable, 3, 0, 1, 3)

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

	// Generate a preview image of the layout
	if err := preview.GeneratePreview(dashboard, "dashboard_layout.jpg"); err != nil {
		log.Fatalf("Error generating preview: %v", err)
	}
	fmt.Println("Layout preview generated as 'dashboard_layout.jpg'")
}
