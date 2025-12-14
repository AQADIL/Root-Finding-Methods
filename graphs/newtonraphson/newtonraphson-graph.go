package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func f(x float64) float64  { return x*math.Exp(x) - math.Sin(x) }
func df(x float64) float64 { return math.Exp(x)*(1+x) - math.Cos(x) }

func main() {
	outputDir := `C:\Users\ACER\GolandProjects\compmath-assik1\graphs\newtonraphson`
	fileName := "newton_graph.html"

	x := -2.5
	tol := 0.001
	maxIter := 50

	var iterations []string
	var errors []string

	for i := 1; i <= maxIter; i++ {
		fx := f(x)
		dfx := df(x)

		if dfx == 0 {
			break
		}

		xNew := x - (fx / dfx)
		err := math.Abs(xNew - x)

		iterations = append(iterations, fmt.Sprintf("%d", i))
		errors = append(errors, fmt.Sprintf("%.6f", err))

		if err < tol {
			break
		}
		x = xNew
	}

	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Newton Method Convergence</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js"></script>
    <style>
        body { font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background: #f4f4f9; margin: 0; }
        .container { width: 80%%; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
        h2 { text-align: center; color: #333; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Newton-Raphson: Super Fast Convergence</h2>
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
                    borderColor: '#4bc0c0',
                    backgroundColor: 'rgba(75, 192, 192, 0.2)',
                    borderWidth: 3,
                    pointRadius: 5,
                    fill: true,
                    tension: 0.1
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
