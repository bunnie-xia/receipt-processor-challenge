package cache

import (
  "fmt"
	// "strconv"
  "strings"
	// "math"
	// "sort"
	"project/internal/db/receipts"
)
var percent float64 = 0.65


func isLLMtext(words []string) bool {
	decision := false
	if words == nil|| len(words) == 1 {
		return decision
	}

	conjoint := map[string]bool{"the": true, "a": true, "an": true, "to": true, "of": true, "for": true, "in": true, "at": true, "and": true, "is": true, "are": true, "it": true, "that": true, "that's": true, "it's": true}

	repeatedWord := words[0]
	localCount, finalCount,freq := 0, 0, 0

	for i := 1; i < len(words); i++ {
		word := strings.ToLower(words[i])
		if conjoint[word] {
			continue
		}

		if word == repeatedWord {
			localCount++
		} else {
			if localCount > finalCount {
				finalCount = localCount
				freq++
			} else {
				repeatedWord = word
				localCount = 1
			}
		}
	}

	if finalCount >= 3 && freq > 2 {
		decision = true
	}

	return decision
}


func isLLMnumStr(nums []string) bool {
	cnt := 0
	for _, num := range nums {
		if strings.HasSuffix(num,".99") {
			cnt++
		}
	}

	return (float64(cnt) / float64(len(nums))) > float64(percent)
}

func isLLMGenerated(receipt receipts.Receipt, id string) bool {
	behaviorsCount := 0
	totalTest := 3

	// testing data prep
	itemInfo := getItemInfo(receipt.Items)

	retailer := strings.Split(receipt.Retailer, " ")
	itemDesList := itemInfo[0]
	itemPriceList := itemInfo[1]
	totalStr := fmt.Sprintf("%.2f", getTotalPrice(id))

	// Debugging prints
	fmt.Println("itemInfo:", itemInfo)
	// fmt.Println("Retailer:", retailer)
	// fmt.Println("Item Descriptions:", itemDesList)
	// fmt.Println("Item Prices:", itemPriceList)
	// fmt.Println("Total String:", totalStr)

	/**
		behavior tendency %
	**/

	// test: retailer name
	if isLLMtext(retailer) {
		behaviorsCount++
	}

	// test: each item shortDescription
	cnt := 0.00
	for _, description := range itemDesList {
		totalTest++
		// if descSlice, ok := description.([]string); ok {
		if isLLMtext(strings.Split(description, " ")) {
			cnt++
			behaviorsCount += 2
		}
		// } else {
		// 	fmt.Println("Data type error.")
		// }
	}

	if cnt > float64(len(itemDesList)) / float64(0.75) {
		return true
	}


	// test: item prices
	// if itemPrice, ok := itemPriceList.([]string); ok {
	if isLLMnumStr(itemPriceList) {
		behaviorsCount++
	}
	// } else {
	// 	fmt.Println("Data type error.")
	// }

	// test: total
	if strings.HasSuffix(totalStr,".99") {
		behaviorsCount++
	}

	return float64(behaviorsCount) / float64(totalTest) > float64(percent)
}




