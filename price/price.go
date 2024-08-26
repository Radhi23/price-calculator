package price

import (
	"example/price-calculator/conversion"
	iomanager "example/price-calculator/ioManager"
	"fmt"
)

// A new struct has been created for taxjob
type TaxIncludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	InputPrice       []float64           `json:"input_price"`
	TaxIncludedPrice map[string]string   `json:"tax_included_price"`
}

// A new method which will load data
func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	prices, err := conversion.StringstoFloats(lines)
	if err != nil {
		return err
	}
	job.InputPrice = prices
	return nil
}

// A new method
func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrice {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrice = result
	return job.IOManager.WriteResult(job)
}

// A new contsructor func
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxrate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:  iom,
		TaxRate:    taxrate,
		InputPrice: []float64{10, 20, 30},
	}
}
