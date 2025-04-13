package cache

import (
	"project/internal/db/receipts"
	"reflect"
	"fmt"
	"strings"
  "strconv"
  "math"
	"github.com/google/uuid"
)


var CacheMap = make(map[interface{}]receipts.Receipt)
var CachePoints = make(map[string]int)

var calculatedTotal float64 = 0.00

func Set(receipt receipts.Receipt) string {
	for key, cached := range CacheMap {
		if reflect.DeepEqual(receipt, cached) {
			fmt.Println("Receipt is aleady posted with id: %s\n", key)
		}
	}

	id := uuid.New().String()
	CacheMap[id] = receipt

	return id
}


func getCachedReceipt(id string) receipts.Receipt {
	if id == "" {
		return receipts.Receipt{}
	}

	receipt, exist := CacheMap[id]
	if !exist {
		return receipts.Receipt{}
	}

	return receipt
}


func getItemInfo(items []receipts.Items) [][]string {
	res := [][]string{}

	for _, item := range items {
		descrip := strings.TrimSpace(item.ShortDescription)
		price := item.Price

		res = append(res, []string {descrip, price})

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Error parsing price:", err)
			continue
		}
		calculatedTotal += priceFloat
	}

	return res
}


func roundedRcptTotal(id string) string {
	receipt := getCachedReceipt(id)
	if receipt.Total == "" {
		return "0.01"
	}

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
			fmt.Println("Error parsing total to float:", err)
			return "0.01"
	}

	roundedFloat := math.Round(total * 100) / 100
	floatStr := strconv.FormatFloat(roundedFloat, 'f', 2, 64)
	return floatStr
}
