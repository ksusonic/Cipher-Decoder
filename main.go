package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func TakePG() (int, int) {
	var trash string
	var g, p int
	_, err := fmt.Scanf("%s %s %d %s %s %s %d", &trash, &trash, &g, &trash, &trash, &trash, &p)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	return g, p
}

func TakeA() int {
	var trash string
	var a int
	_, err := fmt.Scanf("%s %s %d", &trash, &trash, &a)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(3)
	}
	return a
}

func ComputeB(g *int, b *int, p *int) int {
	return int(math.Pow(float64(*g), float64(*b))) % *p
}

func main() {
	g, p := TakePG()
	fmt.Println("OK")
	b := rand.Intn(p) + 1
	B := ComputeB(&g, &b, &p)
	TakeA()
	fmt.Printf("B is %d\n", B)
}
