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
	outputDir := `C:\Users\ACER\GolandProjects\compmath-assik1\graphs\muller`
	fileName := "muller_graph.html"

	x0, x1, x2 := -2.5, -3.0, -3.5
	tol := 0.001
	maxIter := 50

	var iterations []string
	var errors []string

	for i := 1; i <= maxIter; i++ {
		f0, f1, f2 := f(x0), f(x1), f(x2)

		h1 := x1 - x0
		h2 := x2 - x1
		d1 := (f1 - f0) / h1
		d2 := (f2 - f1) / h2
		d := (d2 - d1) / (h2 + h1)

		b := d2 + h2*d
		D := math.Sqrt(b*b - 4*f2*d)

		var E float64
		if math.Abs(b-D) < math.Abs(b+D) {
			E = b + D
		} else {
			E = b - D
		}

		h := -2 * f2 / E
		xNew := x2 + h
		err := math.Abs(h)

		iterations = append(iterations, fmt.Sprintf("%d", i))
		errors = append(errors, fmt.Sprintf("%.6f", err))

		if err < tol {
			break
		}
		x0, x1, x2 = x1, x2, xNew
	}

	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Muller Method Convergence</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js"></script>
    <style>
        body { font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background: #f4f4f9; margin: 0; }
        .container { width: 80%%; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
        h2 { text-align: center; color: #333; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Muller's Method Convergence</h2>
        <canvas id="myChart"></canvas>
    </div>
    <script>
        const ctx = document.getElementById('myChart').getContext('2d');
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: [%s],
                datasets: [{
                    label: 'Error (Step Size)',
                    data: [%s],
                    borderColor: '#ffcd56',
                    backgroundColor: 'rgba(255, 205, 86, 0.2)',
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
