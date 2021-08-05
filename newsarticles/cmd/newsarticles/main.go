package main

import (
	"ncu-main-recruitment/newsarticles/internal"
)

func main() {
	a := internal.App{}
	a.Initialize()
	a.Run(":8080")
}
