package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// The g(x) function: x = -Pi - asin(x * e^x)
func g(x float64) float64 {
	return -math.Pi - math.Asin(x*math.Exp(x))
}

func main() {
	// --- SETUP ---
	outputDir := `C:\Users\ACER\GolandProjects\compmath-assik1\graphs\fixedpoint`
	fileName := "fixedpoint_graph.html"

	xPrev := -2.5 // Start guess
	tol := 0.001
	maxIter := 50

	var iterations []string
	var errors []string

	// --- LOGIC ---
	for i := 1; i <= maxIter; i++ {
		xNew := g(xPrev)

		if math.IsNaN(xNew) {
			break // Stop if math breaks
		}

		err := math.Abs(xNew - xPrev)

		// Save Data
		iterations = append(iterations, fmt.Sprintf("%d", i))
		errors = append(errors, fmt.Sprintf("%.6f", err))

		if err < tol {
			break
		}
		xPrev = xNew
	}

	// --- HTML GENERATION ---
	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Fixed-Point Convergence</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js"></script>
    <style>
        body { font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background: #f4f4f9; margin: 0; }
        .container { width: 80%%; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
        h2 { text-align: center; color: #333; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Fixed-Point Method: Error vs Iteration</h2>
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
                    borderColor: '#36a2eb',
                    backgroundColor: 'rgba(54, 162, 235, 0.2)',
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

	// --- SAVE & OPEN ---
	os.MkdirAll(outputDir, os.ModePerm)
	fullPath := filepath.Join(outputDir, fileName)
	file, _ := os.Create(fullPath)
	file.WriteString(htmlContent)
	file.Close()
	exec.Command("cmd", "/c", "start", "msedge", fullPath).Run()
}
