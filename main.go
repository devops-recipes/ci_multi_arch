package main

import "fmt"
import "runtime"

func main() {
	fmt.Println("Architecture: ", runtime.GOARCH)
	fmt.Println("Operating system: ", runtime.GOOS)
}
