package receipts

import (
  "encoding/json"
)

// type Items struct {
//   ShortDescription  string `json:"shortDescription"`
//   Price             string `json:"price"`
// }

// type Receipt struct {
//   Retailer      string          `json:"retailer"`
// 	PurchaseDate  string          `json:"purchaseDate"`
//   PurchaseTime  string          `json:"purchaseTime"`
//   Items         []Items         `json:"items"`
//   Total         float64         `json:"total"`
//   ExtraFields   json.RawMessage `json:"extraFields"`
// }

type Items struct {
  ShortDescription  string `json:"shortDescription"`
  Price             string `json:"price"`
}

type Receipt struct {
  Retailer      string          `json:"retailer"`
	PurchaseDate  string          `json:"purchaseDate"`
  PurchaseTime  string          `json:"purchaseTime"`
  Items         []Items         `json:"items"`
  Total         string         `json:"total"`
  ExtraFields   json.RawMessage `json:"extraFields"`
}


