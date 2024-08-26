package main

import (
	filemanager "example/price-calculator/fileManager"
	"example/price-calculator/price"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := filemanager.New("pricbbes.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdM := cmdmanager.New()
		priceJob := price.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
			return
		}
	}
}
