package cache

import (
	"project/internal/db/receipts"
	"reflect"
	"fmt"
	"github.com/google/uuid"
)


var CacheMap = make(map[string]receipts.Receipt)
var CachePoints = make(map[string]int)

var calculatedTotal float64 = 0.00

func Set(receipt receipts.Receipt) string {
	for id, cached := range CacheMap {
		if reflect.DeepEqual(receipt, cached) {
			fmt.Println("Receipt is aleady posted with total points received: ", CachePoints[id])
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

