<div align="center">

# 🚀 GoRoots

### High-Performance Numerical Root-Finding Library in Go

*Computational Mathematics Project implementing 6 classical root-finding algorithms with interactive visualizations*

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen?style=for-the-badge)](https://github.com)
[![Made with Love](https://img.shields.io/badge/Made%20with-❤️-red?style=for-the-badge)](https://github.com)

[Features](#-features) • [Quick Start](#-quick-start) • [The Math](#-the-math) • [Benchmarks](#-benchmarks) • [Methods](#-methods-implemented) • [Visualizations](#-visualizations)

</div>

---

## ✨ Features

- ⚡ **Blazing Fast:** Pure Go implementation with zero dependencies
- 📊 **Interactive Graphs:** Auto-generates HTML files with Chart.js convergence plots
- 🛠️ **6 Algorithms:** From classic Bisection to advanced Muller's Method
- 🎯 **Precision:** Configurable tolerance and iteration limits
- 💻 **CLI-Friendly:** Beautiful formatted output with iteration tables
- 🧪 **Battle-Tested:** Validated on interval [-4, -1] with root ≈ -2.9907

---

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/GoRoots.git
cd GoRoots

# Run all methods
go run main.go

# Run specific method
go run bisection.go
go run newton.go
go run muller.go
```

**Output:** CLI convergence tables + `.html` visualization files in the project directory.

---

## 📐 The Math

We solve the non-linear equation:

```
f(x) = x·eˣ - sin(x) = 0
```

**LaTeX representation:**

$$
f(x) = x \cdot e^x - \sin(x)
$$

**Target Root:** `x ≈ -2.9907` (on interval `[-4, -1]`)

---

## 🏆 Benchmarks

Performance comparison based on iterations to reach tolerance `ε = 0.001`:

| Rank | Method | Iterations | Speed | Notes |
|------|--------|------------|-------|-------|
| 🥇 | **Muller's Method** | **2** | 🚀 Fastest | Quadratic interpolation wins |
| 🥈 | **Newton-Raphson** | **3** | ⚡ Very Fast | Requires derivative |
| 🥈 | **Secant Method** | **3** | ⚡ Very Fast | No derivative needed |
| 🥉 | **Fixed-Point** | **4** | ✅ Fast | Custom formula g(x) |
| 4️⃣ | **False Position** | **4** | ✅ Good | Guaranteed convergence |
| 5️⃣ | **Bisection** | **14** | 🐢 Reliable | Slowest but bulletproof |

> **Verdict:** Muller's Method is the clear winner for this function! 🎯

---

## 🛠️ Methods Implemented

### 1. 📏 Bisection Method
**File:** `bisection.go`
Classic interval-halving technique. Slow but guaranteed convergence for continuous functions.

### 2. 🔄 Fixed-Point Iteration
**File:** `fixedpoint2v.go`
Transforms `f(x) = 0` into `x = g(x)` using:
```
g(x) = -π - arcsin(x·eˣ)
```

### 3. 📈 Newton-Raphson Method
**File:** `newtonraphson.go`
Uses tangent line approximation. Requires derivative `f'(x) = eˣ(x+1) - cos(x)`.

### 4. ➗ Secant Method
**File:** `secant.go`
Approximates derivative numerically. Faster setup than Newton's method.

### 5. 🎯 False Position (Regula Falsi)
**File:** `falseposition.go`
Hybrid approach combining bisection reliability with linear interpolation.

### 6. 🌟 Muller's Method
**File:** `muller.go`
Parabolic interpolation through 3 points. Best performer on this function!

---

## 📊 Visualizations

Each method generates an interactive HTML file:

```
bisection_convergence.html
newton_convergence.html
muller_convergence.html
...
```

**Features:**
- 📉 Convergence trend line
- 🎨 Styled with Chart.js
- 🖱️ Hover tooltips showing exact values
- 📱 Responsive design

**Example:**

<div align="center">
<img src="graph.png" alt="Convergence Plot" width="600"/>
</div>

---

## 📁 Project Structure

```
compmath-assik1/
├── graphs/                        # Visualization files directory
│   ├── bisection/
│   │   ├── bisection_graph.html
│   │   └── bisection_plot.go
│   ├── falseposition/
│   │   ├── falseposition_graph.html
│   │   └── falseposition-graph.go
│   ├── fixedpoint/
│   │   ├── fixedpoint_graph.html
│   │   └── fixedpoint-graph.go
│   ├── muller/
│   │   ├── muller_graph.html
│   │   └── muller-graph.go
│   ├── newtonraphson/
│   │   ├── newton_graph.html
│   │   └── newtonraphson-graph.go
│   ├── secant/
│   │   ├── secant_graph.html
│   │   └── secant-graph.go
│   └── graph.png
├── bisection.go                   # Bisection method
├── fixedpoint.go                  # Fixed-point iteration (v1)
├── fixedpoint2v.go                # Fixed-point iteration (v2)
├── newtonraphson.go               # Newton-Raphson method
├── secant.go                      # Secant method
├── falseposition.go               # False position method
├── muller.go                      # Muller's method
├── go.mod                         # Go module file
└── README.md                      # You are here!
```

---

## 🎓 Theory Background

### Convergence Rates

| Method | Order | Formula |
|--------|-------|---------|
| Bisection | Linear (1) | `e_{n+1} ≈ e_n / 2` |
| Fixed-Point | Linear (1-2) | Depends on `g'(x)` |
| Secant | Superlinear (1.618) | Golden ratio! |
| Newton-Raphson | Quadratic (2) | `e_{n+1} ≈ e_n²` |
| Muller | Near-Cubic (1.84) | Best of both worlds |

---

## 🧪 How to Test

Run all methods and compare:

```bash
# Run main program
go run main.go

# Open generated HTML files
open bisection_convergence.html
open muller_convergence.html
```

---

## 📝 Author

**Alish Akadil**

<div align="center">

### 🌟 Star this repo if you found it helpful!

Made with ❤️ and lots of ☕ using Go

---

*"In mathematics, you don't understand things. You just get used to them." — John von Neumann*

</div>
