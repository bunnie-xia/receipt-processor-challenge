package cache

import (
  "fmt"
	"strconv"
	"math"
  "time"
	"strings"
	"project/internal/db/receipts"
	// "project/internal/db/llm"
)


func retailerPoints(receipt receipts.Receipt) int {
	if receipt.Retailer != "" {
		vendorPlainStr := strings.TrimSpace(receipt.Retailer)
		points := len(vendorPlainStr)

		return points
	}

	return 0
}


func itemPoints(items [][]string) int {
	if len(items) == 0 {
		return 0
	}

	pts := 0
	for _, item := range items {
		if len(item) < 2 {
			continue
		}

		descrip := item[0]
		price := item[1]
		if descrip != "" && len(descrip) % 3 == 0 {
			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				continue
			}
			pts += int(math.Ceil(p * 0.2))
		}
	}

	return pts
}


func purchaseDatePoints(receipt receipts.Receipt) int {
	dateStr := receipt.PurchaseDate
	if dateStr == "" {
		return 0
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0
	}

	if date.Day() % 2 == 1 {
		return 6
	} else {
		return 0
	}
}


func purchaseTimePoints(receipt receipts.Receipt) int {
	timeStr := receipt.PurchaseTime
	if timeStr == "" {
		return 0
	}

	layout := "15:04"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0
	}

	time1, _ := time.Parse("15:04", "14:00")
	time2, _ := time.Parse("15:04", "16:00")

	if t.After(time1) && t.Before(time2) {
		return 10
	} else {
		return 0
	}
}


func everyTwoItemPoints(items [][]string) int {
	points := int(math.Floor(float64(len(items)) / float64(2))) * 5
	return points
}


func receiptTotalPricePoints(receipt receipts.Receipt, id string) int {
	pts := 0
	// roundedTotal := 0.00
	// if receipt.Total == "" {
	// 	roundedTotal = roundedCalTotal(id)
	// }

	roundedTotalStr := roundedRcptTotal(id)

	if fmt.Sprintf("%.2f", roundedTotalStr) == fmt.Sprintf("%.2f", receipt.Total) {
		pts += 50
	}

	ttl, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
			fmt.Println("Error parsing total to float:", err)
			return pts
	}

	if int(ttl * 100) % 25 == 0 {
		pts += 25
	}

	return pts
}


func llmPoints(receipt receipts.Receipt, id string) int {
	ttl, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
			fmt.Println("Error parsing total to float:", err)
	}

	if isLLMGenerated(receipt, id) == true && (float64(ttl) > float64(10.00)) {
		return 5
	}

	return 0
}


func CalculatePoints(targetId string) int {
	receipt := getCachedReceipt(targetId)
	items := getItemInfo(receipt.Items)

	points := 0
	points += retailerPoints(receipt)
	points += itemPoints(items)
	points += purchaseDatePoints(receipt)
	points += purchaseTimePoints(receipt)
	points += everyTwoItemPoints(items)
	points += receiptTotalPricePoints(receipt, targetId)
	points += llmPoints(receipt, targetId)

	return points
}

