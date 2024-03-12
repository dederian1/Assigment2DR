package entity

import "time"

// Order adalah struktur data untuk mewakili pesanan
type Order struct {
	Order_id      int       `json:"orderid"`
	Customer_Name string    `json:"customerName"`
	Ordered_At    time.Time `json:"ordereAt"` // Perlu diperbaiki: seharusnya "orderedAt" agar sesuai dengan nama JSON
	Item          []Item    `json:"items"`
}

// Item adalah struktur data untuk mewakili item dalam pesanan
type Item struct {
	Item_Id     int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_Id    int    `json:"orderid"`
}
