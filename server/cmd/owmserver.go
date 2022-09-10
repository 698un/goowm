package main

import (
	"cnfg"
	"fmt"
	"mymath"
	"netcontrol"
)

func main() {
	//fmt.Println(calc.AddInts(1, 2))

	fmt.Println(cnfg.PathCash)

	var testVector mymath.MVector
	fmt.Println(testVector)

	fmt.Println("Start OWM server...")

	netcontrol.NetControlStart()

	fmt.Println("The end")

}
