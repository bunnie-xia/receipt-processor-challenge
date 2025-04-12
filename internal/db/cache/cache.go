package cache

import (
	"project/internal/db/receipts"
	"reflect"
	// "encoding/json"
	"fmt"
	// "string"
	"strings"
  "strconv"
  "math"
	// "errors"
	"github.com/google/uuid"
)

// type Cache struct {
// 	id string
// 	objData interface {} receipts.processReceipt
// 	Points float64 receipts.calPoints
// }

// type idGenFunc func(data *interface{}) string
// type receipt.convertReceipt func(JSONreceipt []byte(jsonString)) struct

// func generateId(data *interface{}) string {
//   newUUID := uuid.new()
//   data.id = newUUID
// }

// var CacheMap = make(map[string]receipts.Receipt)
var CacheMap = make(map[interface{}]receipts.Receipt)
var CachePoints = make(map[string]int)

var TotalPrice float64 = 0.00

func Set(receipt receipts.Receipt) string {
	for key, cached := range CacheMap {
		if reflect.DeepEqual(receipt, cached) {
			fmt.Println("Receipt is aleady posted with id: %s\n", key)
		}
	}

	id := uuid.New().String()
	CacheMap[id] = receipt
	// CachePoints[id] = CalculatePoints(id)

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
// func getItemInfo(items [][]string) [][]string {
	res := [][]string{}

	for _, item := range items {
		descrip := strings.TrimSpace(item.ShortDescription)
		price := item.Price

		res = append(res, []string {descrip, price})
		// res[0] = append(res[0], descrip)
		// res[1] = append(res[1], price)

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Error parsing price:", err)
			continue
		}
		TotalPrice += priceFloat
	}

	return res
}

func getTotalPrice(id string) float64 {
	receipt := getCachedReceipt(id)
	total := TotalPrice
	if int(receipt.Total) == 0 {
		TotalPrice = float64(0)
		return math.Round(total * 100)/ 100
		// receipt.Total = total

	}

	return total
}
// func getPoints(id string) {
// 	if id == "" {
// 		fmt.Println("Receipt id not found.")
// 		return
// 	}

// 	points, exist := Cache[id]
// 	if !exists {
// 		return errors.New("Item not found.")
// 	}

// 	return object.id.Points.sum
// }
