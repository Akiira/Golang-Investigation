package Normal_Version 

import "cmpsc"

type product struct {}

var buffer Buffer

func Producer() {
	for {
		//Create product
		prdct := product{}
		
		//add product to buffer
		buffer.Add(prdct)
	}
}

func Consumer() {
	for {	
		//Get some product from buffer
		prdct := buffer.Get()
		
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

