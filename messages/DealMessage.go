package messages

type DealMessage struct {
	SellerID     string `json:"sellerId"`
	SellerCastle string `json:"sellerCastle"`
	SellerName   string `json:"sellerName"`
	BuyerID      string `json:"buyerId"`
	BuyerCastle  string `json:"buyerCastle"`
	BuyerName    string `json:"buyerName"`
	Item         string `json:"item"`
	Quantity     int    `json:"qty"`
	Price        int    `json:"price"`
}
