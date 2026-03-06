package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle (g geometry) {
	// c 转换后的 circle 值; ok 是否转换成功 (true / false)
	// 尝试把接口 g 转换成 circle 类型
	/**
	* if 初始化语句; 条件 {
	* } 
	*/
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main () {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 尝试把接口 g 转换成 circle 类型s
	measure(r)
	measure(c)

	detectCircle(r)
  detectCircle(c)
}