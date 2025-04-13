package cache

import (
  "strings"
	"sort"
	"project/internal/db/receipts"
)

var percent float64 = 0.80
var overrideToTrue bool = false


func isLLMtext(words []string) bool {
	if words == nil|| len(words) == 1 {
		return false
	}

	conjoint := map[string]bool{"the": true, "a": true, "an": true, "to": true, "of": true, "for": true, "in": true, "at": true, "and": true, "is": true, "are": true, "it": true, "that": true, "that's": true, "it's": true, "": true}

	sort.Strings(words)
	repeatedWord := words[0]
	localCount, mostFreq, totalRepWords := 0, 0, 0

	for i := 1; i < len(words); i++ {
		word := strings.ToLower(words[i])
		if conjoint[word] {
			continue
		}

		if word == repeatedWord {
			localCount++
		} else {
			if localCount > mostFreq {
				mostFreq = localCount
			}
			localCount = 1
			repeatedWord = word
			totalRepWords++
		}
	}

	if mostFreq > 3 && totalRepWords >= 3 {
		return true
	}

	return false
}


func isLLMnumStr(nums []string) bool {
	if nums == nil {
		return false
	}

	cnt := 0
	for _, num := range nums {
		if strings.HasSuffix(num, ".99") {
			cnt++
		}
	}

	return float64(cnt) / float64(len(nums)) > float64(percent)
}


func itemsPrep(items []receipts.Items) [][]string {
	res := [][]string{{}, {}}

	for _, item := range items {
		if item.ShortDescription == "" || item.Price == "" {
      continue
    }

		preLen := len(item.ShortDescription)
		descrip := strings.TrimSpace(item.ShortDescription)
		aftLen := len(descrip)
		if preLen != aftLen {
			overrideToTrue = true
			return [][]string{{}, {}}
		}

		price := item.Price

		res[0] = append(res[0], descrip)
		res[1] = append(res[1], price)
	}

	return res
}


func isLLMGenerated(receipt receipts.Receipt, id string) bool {
	behaviorsCount := 0
	totalTest := 2

	/** test data prepand "overrideToTrue" checks **/
	// retailer
	preLen := len(receipt.Retailer)
	retailerStr := strings.TrimSpace(receipt.Retailer)
	aftLen := len(retailerStr)
	if preLen != aftLen {
		return false
	}
	retailer := strings.Split(retailerStr, " ")

	// itemInfo
	itemInfo := itemsPrep(receipt.Items)
	if overrideToTrue == true {
		return true
	}
	itemDesList := itemInfo[0]
	itemPriceList := itemInfo[1]


	/**
		behavior tendency %
	**/
	// test: retailer
	if isLLMtext(retailer) {
		behaviorsCount++
	}

	// test: itemDesList
	cnt := 0
	for _, description := range itemDesList {
		if isLLMtext(strings.Split(description, " ")) {
			cnt++
			behaviorsCount += 2
		}
	}
	totalTest += len(itemDesList)

	if float64(cnt) / float64(len(itemDesList)) > float64(percent) {
		return true
	}

	// test: itemPriceList
	if isLLMnumStr(itemPriceList) {
		behaviorsCount++
	}
	totalTest += len(itemPriceList)

	return float64(behaviorsCount) / float64(totalTest) > float64(percent)
}




