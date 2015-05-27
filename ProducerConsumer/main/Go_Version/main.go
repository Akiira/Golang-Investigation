package main

import "cmpsc"

type product struct{}

var buffer chan product

func Producer() {
	for {
		//Create product
		prdct := product{}

		//add product to buffer
		buffer <- prdct
	}
}

func Consumer() {
	for {
		//Get some product from buffer
		prdct := (<-buffer)

		//Consumer it
		_ = prdct
	}
}

func main() {
	cmpsc.CoBegin(
		Producer,
		Consumer,
		Consumer,
		Consumer,
	)
}
