package newsarticles

import "ncu-main-recruitment/internal"

func main() {
	a := internal.App{}
	a.Initialize()
	a.Run(":8080")
}
