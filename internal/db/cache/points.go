package cache

import (
  "fmt"
	// "encoding/json"
	"strconv"
	"math"
  // "os"
  "time"
	// "reflect"
	"strings"
	"project/internal/db/receipts"
)

// var points = 0

func roundToTwoDecimal(price float64) float64 {
	return math.Round(price*100) / 100
}

func retailerPoints(receipt receipts.Receipt) int {
	if receipt.Retailer != "" {
		vendorPlainStr := strings.TrimSpace(receipt.Retailer)
		// return len(vendorPlainStr)
		points := len(vendorPlainStr)
		fmt.Println("retailer: ", points)
		return points
	}

	return 0
}

func itemPoints(items [][]string) int {
	if len(items) == 0 {
		return 0
	}

	pts := 0
	fmt.Println("len", len(items))
	for _, item := range items {
		if len(item) < 2 {
			continue
		}
		fmt.Println("~~")
		// descrip := item.ShortDescription
		// price := item.Price
		descrip := item[0]
		price := item[1]
		if descrip != "" && len(descrip) % 3 == 0 {
			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("Error converting price:", err)
				continue
			}
			fmt.Println("p: ", p)
			pts += int(math.Ceil(p * 0.2))
		}
	}
	fmt.Println("itemDesPoints: ", pts)

	return pts
}



func purchaseDatePoints(receipt receipts.Receipt) int {
	dateStr := receipt.PurchaseDate
	if dateStr == "" {
		fmt.Println("empty date points: ", 0)
		return 0
	}

	layout := "2006-01-02"

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date: ", err)
		return 0
	}

	if date.Day() % 2 == 1 {
		fmt.Println("date points: ", 6)
		return 6
	} else {
		fmt.Println("date points: ", 0)
		return 0
	}
}


func purchaseTimePoints(receipt receipts.Receipt) int {
	timeStr := receipt.PurchaseTime
	if timeStr == "" {
		fmt.Println("empty time points: ", 0)
		return 0
	}

	layout := "15:04"

	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time: ", err)
		return 0
	}

	time1, _ := time.Parse("15:04", "14:00")
	time2, _ := time.Parse("15:04", "16:00")

	if t.After(time1) && t.Before(time2) {
		fmt.Println("time points: ", 10)
		return 10
	} else {
		fmt.Println("time points: ", 0)
		return 0
	}
}


func everyTwoItemPoints(items [][]string) int {
	points := int(math.Floor(float64(len(items)) / float64(2))) * 5
	fmt.Println("everyTwoItemPoints: ", points)
	return points
}

func receiptTotalPricePoints(receipt receipts.Receipt, id string) int {
	pts := 0
	if receipt.Total == float64(0) {
		receipt.Total = getTotalPrice(id)
	}

	roundedTotal := roundToTwoDecimal(receipt.Total)

	if fmt.Sprintf("%.2f", roundedTotal) == fmt.Sprintf("%.2f", receipt.Total) {
		pts += 50
	}

	if int(receipt.Total * 100) % 25 == 0 {
		pts += 25
	}
	fmt.Println("receiptTotalPricePoints: ", pts)
	return pts
}

func llmPoints(receipt receipts.Receipt, id string) int {
	if isLLMGenerated(receipt, id) == true && (receipt.Total > float64(10.00)) {
		fmt.Println("llmPoinits: ", 5)
		return 5
	}
	fmt.Println("llmPoinits: ", 0)
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
	// points += llmPoints(items[0], targetId)

	return points
}

