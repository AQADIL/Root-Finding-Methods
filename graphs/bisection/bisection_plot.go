package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Function: x*e^x - sin(x)
func f(x float64) float64 {
	return x*math.Exp(x) - math.Sin(x)
}

func main() {

	outputDir := `C:\Users\ACER\GolandProjects\compmath-assik1\graphs\bisection`
	fileName := "bisection_graph.html"

	a, b := -4.0, -2.0
	tol := 0.001
	maxIter := 50

	var iterations []string
	var errors []string

	for i := 1; i <= maxIter; i++ {
		c := (a + b) / 2.0
		err := math.Abs(b - a)

		iterations = append(iterations, fmt.Sprintf("%d", i))
		errors = append(errors, fmt.Sprintf("%.6f", err))

		if err < tol || f(c) == 0 {
			break
		}

		if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}
	}

	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Bisection Convergence</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js"></script>
    <style>
        body { font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background: #f4f4f9; margin: 0; }
        .container { width: 80%%; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
        h2 { text-align: center; color: #333; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Bisection Method: Convergence Analysis</h2>
        <canvas id="myChart"></canvas>
    </div>
    <script>
        const ctx = document.getElementById('myChart').getContext('2d');
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: [%s],
                datasets: [{
                    label: 'Error (Interval Width)',
                    data: [%s],
                    borderColor: '#ff6384',
                    backgroundColor: 'rgba(255, 99, 132, 0.2)',
                    borderWidth: 3,
                    pointRadius: 5,
                    fill: true,
                    tension: 0.3
                }]
            },
            options: {
                responsive: true,
                scales: {
                    y: { 
                        title: {display: true, text: 'Error value'}, 
                        beginAtZero: true 
                    },
                    x: { 
                        title: {display: true, text: 'Iteration Number'} 
                    }
                }
            }
        });
    </script>
</body>
</html>`, strings.Join(iterations, ","), strings.Join(errors, ","))

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		panic(err)
	}

	fullPath := filepath.Join(outputDir, fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(htmlContent)

	fmt.Println("Graph saved to:", fullPath)

	cmd := exec.Command("cmd", "/c", "start", "msedge", fullPath)
	cmd.Run()
}
