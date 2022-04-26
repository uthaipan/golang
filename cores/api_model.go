package cores

type OrderReturn struct {
	Result bool `json:"results"`
}

type Order_buy struct {
	Order_buy_id   string `json:"order_buy_id"`
	Address        string `json:"address"`
	Date           string `json:"date"`
	Status         string `json:"status"`
	Product_id     string `json:"product_id"`
	Product_name   string `json:"product_name"`
	Size           string `json:"size"`
	Gender         string `json:"gender"`
	Category       string `json:"category"`
	Product_amount string `json:"product_amount"`
}

type Product struct {
	Product_id   string `json:"product_id"`
	Product_name string `json:"product_name"`
	Size         string `json:"size"`
	Gender       string `json:"gender"`
	Category     string `json:"category"`
}

type DataProductPage struct {
	Page       string    `json:"page"`
	Total_page string    `json:"total_page"`
	Product    []Product `json:"products"`
}

type DataOderPage struct {
	Page       string      `json:"page"`
	Total_page string      `json:"total_page"`
	Order_buy  []Order_buy `json:"order_buys"`
}