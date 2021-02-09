package messages

type AuctionItem struct {
	LotID string `json:"lotId"`
	ItemName string `json:"itemName"`
	SellerTag string `json:"sellerTag"`
	SellerName string `json:"sellerName"`
	Quality string `json:"quality"`
	SellerCastle string `json:"sellerCastle"`
	Condition string `json:"condition"`
	EndAt string `json:"endAt"`
	StartAt string `json:"startedAt"`
	BuyerCastle string `json:"buyerCastle"`
	Status string `json:"status"`
	FinishedAt string `json:"finishedAt"`
	BuyerTag string `json:"buyerTag"`
	BuyerName string `json:"buyerName"`
	Price int `json:"price"`
	Stats map[string]int
}