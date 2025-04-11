package receipts

import (
  "encoding/json"
  // "strings"
  // "fmt"
  // "strconv"
  // "math"
  // "time"
)

type Items struct {
  ShortDescription  string `json:"shortDescription"`
  Price             string `json:"price"`
}

type Receipt struct {
  Retailer      string          `json:"retailer"`
	PurchaseDate  string          `json:"purchaseDate"`
  PurchaseTime  string          `json:"purchaseTime"`
  Items         []Items         `json:"items"`
  Total         float64         `json:"total"`
  ExtraFields   json.RawMessage `json:"extraFields"`
}

// var TotalPrice float64 = 0.00



// func GetItemInfo(items []Items) [][]string {
// // func getItemInfo(items [][]string) [][]string {
// 	res := [][]string{{}}

// 	for _, item := range items {
// 		descrip := strings.TrimSpace(item.ShortDescription)
// 		price := item.Price

// 		res = append(res, []string {descrip, price})

// 		priceFloat, err := strconv.ParseFloat(price, 64)
// 		if err != nil {
// 			fmt.Println("Error parsing price:", err)
// 			continue
// 		}
// 		TotalPrice += priceFloat
// 	}

// 	return res
// }

// func GetTotalPrice(id string) float64 {
// 	receipt := getCachedReceipt(id)
// 	if receipt.Total == float64(0.00) {
// 		total := math.Round(TotalPrice * 100)/ 100
// 		receipt.Total = total
// 	}

// 	return receipt.Total
// }
// type Items struct {
//   ShortDescription  string
//   Price             string
// }

// type Receipt struct {
//   Retailer      string
// 	PurchaseDate  time.Time
//   PurchaseTime  time.Time
//   Items         []Items
//   Total         float64
//   ExtraFields   json.RawMessage
// }


// type Receipt struct {
//   Retailer      string          `json:"retailer"`
// 	PurchaseDate  string       `json:"purchaseDate"`
//   PurchaseTime  string       `json:"purchaseTime"`
//   Items         []Items         `json:"items"`
//   Total         float64         `json:"total"`
//   ExtraFields   json.RawMessage `json:"extraFields"`
// }

  // var rawReceipt Receipt
  // err := json.Unmarshal(jsonData, &rawReceipt)
  // if err != nil {
  //   fmt.Println("Error unmarshalling JSON:", err)
  // }

// func ProcessReceipt(receipt Receipt) Receipt {

//   if receipt.PurchaseDate.IsZero() {
//     return receipt
//   }

//   date, err := time.Parse("2006-01-02", receipt.PurchaseDate.Format("2006-01-02"))
//   if err != nil {
// 		fmt.Println("Error parsing date:", err)
// 	}
//   receipt.PurchaseDate = date

//   time, err := time.Parse("15:04", receipt.PurchaseTime.Format("15:04"))
//   if err != nil {
// 		fmt.Println("Error parsing time:", err)
// 	}
//   receipt.PurchaseTime = time

//   return receipt
// }

