package business

import (
	"log"
	"strings"

	"github.com/pkg/errors"
)

// NewConver manages the set of API's for Convert access.
type NewConvert struct {
	log *log.Logger
}

// New constructs a Convert for api access.
func New(log *log.Logger) NewConvert {
	return NewConvert{log}
}

// ConvertToNaira converts the input amount to naira
func (nc NewConvert) ConvertToNaira(cm ConvertMoney) (ConvertMoney, error) {
	// get the rate from the currency table
	getRate, err := nairaCurrencyTable(strings.ToLower(cm.Currency))
	if err != nil {
		return ConvertMoney{}, errors.Wrap(err, "could not get rate")
	}
	// convert the amount to naira
	convert := cm.Amount * getRate

	// return the converted amount
	c := ConvertMoney{
		Currency: "NGN",
		Amount: convert,
	}

	// return the converted amount
	log.Printf("converted %v %s to %v %s", cm.Amount, cm.Currency, convert, c.Currency)
	return c, nil
}

// ConvertToCedis converts the input amount to cedis
func (nc NewConvert) ConvertToCedis(cm ConvertMoney) (ConvertMoney, error) {
	// get the rate from the currency table
	getRate, err := cedisCurrencyTable(strings.ToLower(cm.Currency))
	if err != nil {
		return ConvertMoney{}, errors.Wrap(err, "could not get rate")
	}
	// convert the amount to cedis
	convert := cm.Amount * getRate

	// return the converted amount
	c := ConvertMoney{
		Currency: "GHS",
		Amount: convert,
	}

	log.Printf("converted %v %s to %v %s", cm.Amount, cm.Currency, convert, c.Currency)
	// 	return the converted amount
	return c, nil
}

// ConvertToShilling converts the input amount to shilling
func (nc NewConvert) ConvertToShillings(cm ConvertMoney) (ConvertMoney, error) {
	// get the rate from the currency table
	getRate, err := shillingCurrencyTable(strings.ToLower(cm.Currency))
	if err != nil {
		return ConvertMoney{}, errors.Wrap(err, "could not get rate")
	}
	// convert the amount to shillings
	convert := cm.Amount * getRate

	// return the converted amount
	c := ConvertMoney{
		Currency: "KSH",
		Amount: convert,
	}

	log.Printf("converted %v %s to %v %s", cm.Amount, cm.Currency, convert, c.Currency)
	// return the converted amount
	return c, nil
}

// nairaCurrencyTable returns the rate of the currency convertion to naira
func nairaCurrencyTable (currency string) (float64, error) {
	
	// check if the currency is empty
	if currency == "" {
		return 0, errors.Wrap(nil, "currency is empty")
	}
	// create a map of currency and rate
	nct := map[string]float64 {
		"ghs": 66.47,
		"ksh": 3.66,
	}
	
	// return the rate of the currency
	return nct[currency], nil
}

// cedisCurrencyTable returns the rate of the currency convertion to cedis
func cedisCurrencyTable (currency string) (float64, error) {
	
	// check if the currency is empty
	if currency == "" {
		return 0, errors.Wrap(nil, "currency is empty")
	}
	// create a map of currency and rate
	nct := map[string]float64 {
		"ngn": 0.015,
		"ksh": 0.055,
	}
	
	// return the rate of the currency
	return nct[currency], nil
}

// shillingCurrencyTable returns the rate of the currency convertion to shilling
func shillingCurrencyTable (currency string) (float64, error) {
	
	// check if the currency is empty
	if currency == "" {
		return 0, errors.Wrap(nil, "currency is empty")
	}
	// create a map of currency and rate
	nct := map[string]float64 {
		"ngn": 0.27,
		"ghs": 18.18,
	}
	
	// return the rate of the currency
	return nct[currency], nil
}