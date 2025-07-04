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
	sales := bussola.NewIndicator("Total Sales")
	sales.Unit = "R$"
	sales.Trend = 5.7 // 5.7% increase
	sales.Target = "http://localhost:4040/api/v1/query/sales/indicator"

	users := bussola.NewIndicator("Active Users")
	users.Description = "Currently active users"

	tma := bussola.NewIndicator("TMA") // Average time to action in minutes
	tma.Unit = "min"
	tma.Description = "Average time to action (TMA)"
	// Set the target URL for TMA
	tma.Target = "http://localhost:4040/api/v1/query/tma/indicator"

	issues := bussola.NewIndicator("Total Issues")
	issues.Unit = "Issues"

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
	revenueChart.Options = map[string]any{
		"xAxis": []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		"color": "#1976D2",
	}

	userTable := bussola.NewTable("Recent Users", []string{"ID", "Name", "Last Access", "Status"})
	userTable.Data = []map[string]any{
		{"id": 1, "name": "John Doe", "lastAccess": "2025-06-26", "status": "Active"},
		{"id": 2, "name": "Jane Smith", "lastAccess": "2025-06-25", "status": "Inactive"},
	}

	filterBar := bussola.NewFilterBar("Filtros Gerais")
	filterBar.AddFilter(bussola.NewFilterDate("Periodo", "period"))
	filterBar.AddFilter(bussola.NewFilterSelect("Tipo de Problema", "problem_type", []string{"Todos", "Erro", "Aviso", "Info"}))
	filterBar.AddFilter(bussola.NewFilterText("Nome do Cliente", "client_name"))
	filterBar.AddFilter(bussola.NewFilterSearch("Pesquisar", "search", "Search by name or ID"))

	mainGrid.AddItem(filterBar, 0, 0, 1, 3)

	// Adicionando indicadores automaticamente no grid de indicadores
	indicators := bussola.NewGrid("Indicators", 2, 4)
	indicators.AddNext(sales)
	indicators.AddNext(users)
	indicators.AddNext(sales)
	indicators.AddNext(users)
	indicators.AddNext(issues)
	indicators.AddNext(tma)
	indicators.AddNext(issues)
	indicators.AddNext(tma)
	mainGrid.AddItem(indicators, 1, 0, 1, 3)

	mainGrid.AddItem(nestedGrid, 2, 0, 1, 1)
	mainGrid.AddItem(revenueChart, 2, 1, 1, 2)
	mainGrid.AddItem(userTable, 3, 0, 1, 2)

	ranking := bussola.NewRanking("Ranking de Clientes")
	ranking.AddItem(bussola.NewRankingItem(1, "Empresa Alpha", "Maior faturamento", "https://randomuser.me/api/portraits/men/1.jpg"))
	ranking.AddItem(bussola.NewRankingItem(2, "Empresa Beta", "Crescimento rápido", "https://randomuser.me/api/portraits/women/2.jpg"))
	ranking.AddItem(bussola.NewRankingItem(3, "Empresa Gama", "Melhor avaliação", "https://randomuser.me/api/portraits/men/3.jpg"))
	ranking.SetOrder("desc")
	mainGrid.AddItem(ranking, 3, 2, 1, 1)

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

	fmt.Printf("Dashboard JSON:\n%v\n", dashboard.GenerateJSON())

	if err := preview.GeneratePreview(dashboard, "dashboard.jpg"); err != nil {
		log.Fatalf("Error generating preview: %v", err)
	}
}
