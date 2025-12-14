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
	outputDir := `C:\Users\ACER\GolandProjects\compmath-assik1\graphs\falseposition`
	fileName := "falseposition_graph.html"

	a, b := -4.0, -1.0
	tol := 0.001
	maxIter := 50
	cPrev := 0.0

	var iterations []string
	var errors []string

	for i := 1; i <= maxIter; i++ {
		fa := f(a)
		fb := f(b)

		c := (a*fb - b*fa) / (fb - fa)
		err := math.Abs(c - cPrev)

		iterations = append(iterations, fmt.Sprintf("%d", i))
		errors = append(errors, fmt.Sprintf("%.6f", err))

		if i > 1 && err < tol {
			break
		}

		if fa*f(c) < 0 {
			b = c
		} else {
			a = c
		}
		cPrev = c
	}

	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>False Position Convergence</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js"></script>
    <style>
        body { font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background: #f4f4f9; margin: 0; }
        .container { width: 80%%; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
        h2 { text-align: center; color: #333; }
    </style>
</head>
<body>
    <div class="container">
        <h2>False Position Method Convergence</h2>
        <canvas id="myChart"></canvas>
    </div>
    <script>
        const ctx = document.getElementById('myChart').getContext('2d');
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: [%s],
                datasets: [{
                    label: 'Error (Change in Root)',
                    data: [%s],
                    borderColor: '#9966ff',
                    backgroundColor: 'rgba(153, 102, 255, 0.2)',
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
