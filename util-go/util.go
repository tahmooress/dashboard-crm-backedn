package util

//Product is model of product
type Product struct {
	ProductName  string `json:"productName"`
	ProductBrand string `json:"productBrand"`
	ProductID    int    `json:"productID"`
	Price        int    `json:"price"`
}

//Store is model of sotre
type Store struct {
	StoreID      int    `json:"storeID"`
	StoreName    string `json:"storeName"`
	StoreAddress string `json:"storeAddress"`
	Phone        string `json:"phone"`
	Proxy        string `json:"proxy"`
	Mobile       string `json:"mobile"`
}

//Factor is model of factor
type Factor struct {
	FactorID  int `json:"factorID"`
	TransID   int `json:"transID"`
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

//Transaction is a model of transactions in db
type Transaction struct {
	TransID int      `json:"transID"`
	StoreID int      `json:"storeID"`
	Date    string   `json:"date"`
	Factors []Factor `json:"factors"`
}
