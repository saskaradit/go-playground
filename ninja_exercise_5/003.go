package main

import "fmt"

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {

	hino := truck{
		vehicle: vehicle{
			doors: "2",
			color: "black",
		},
		fourWheel: true,
	}
	sls := sedan{
		vehicle: vehicle{
			doors: "2",
			color: "grey",
		},
		luxury: true,
	}
	fmt.Println(hino)
	fmt.Println(sls)
}
