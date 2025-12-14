package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func f(x float64) float64 { return x*math.Exp(x) - math.Sin(x) }

func main() {
	outputDir := `C:\Users\ACER\GolandProjects\compmath-assik1\graphs\secant`
	fileName := "secant_graph.html"

	x0, x1 := -2.0, -3.0
	tol := 0.001
	maxIter := 50

	var iterations []string
	var errors []string

	for i := 1; i <= maxIter; i++ {
		fx0 := f(x0)
		fx1 := f(x1)

		if math.Abs(fx1-fx0) < 1e-9 {
			break
		}

		xNew := x1 - fx1*(x1-x0)/(fx1-fx0)
		err := math.Abs(xNew - x1)

		iterations = append(iterations, fmt.Sprintf("%d", i))
		errors = append(errors, fmt.Sprintf("%.6f", err))

		if err < tol {
			break
		}
		x0 = x1
		x1 = xNew
	}

	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Secant Convergence</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js"></script>
    <style>
        body { font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background: #f4f4f9; margin: 0; }
        .container { width: 80%%; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
        h2 { text-align: center; color: #333; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Secant Method Convergence</h2>
        <canvas id="myChart"></canvas>
    </div>
    <script>
        const ctx = document.getElementById('myChart').getContext('2d');
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: [%s],
                datasets: [{
                    label: 'Error',
                    data: [%s],
                    borderColor: '#ff9f40',
                    backgroundColor: 'rgba(255, 159, 64, 0.2)',
                    borderWidth: 3,
                    pointRadius: 5,
                    fill: true,
                    tension: 0.3
                }]
            },
            options: { responsive: true, scales: { y: { beginAtZero: true } } }
        });
    </script>
</body>
</html>`, strings.Join(iterations, ","), strings.Join(errors, ","))

	os.MkdirAll(outputDir, os.ModePerm)
	fullPath := filepath.Join(outputDir, fileName)
	file, _ := os.Create(fullPath)
	file.WriteString(htmlContent)
	file.Close()
	exec.Command("cmd", "/c", "start", "msedge", fullPath).Run()
}
