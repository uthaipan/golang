package controller

import (
	"log"
	"math"
	"strconv"

	"example/hello/cores"
	"example/hello/service"

	_ "github.com/go-sql-driver/mysql"
)
var test string
func HomeLink() string{
	test = "Welcome homess!"
	return test
}

func ProductList(gender string,category string,size,page string,per_page string) cores.DataProductPage {
	
	condition := ""
	if(gender != ""){
		if(condition != ""){
			condition += " AND product.gender_id = '"+gender+"'"
		}else{
			condition += " WHERE product.gender_id = '"+gender+"'"
		}
	}
	if(category != ""){
		if(condition != ""){
			condition += " AND product.category_id = '"+category+"'"
		}else{
			condition += " WHERE product.category_id = '"+category+"'"
		}
	}
	if(size != ""){
		if(condition != ""){
			condition += " AND product.size_id = '"+size+"'"
		}else{
			condition += " WHERE product.size_id = '"+size+"'"
		}
	}
	limit := ""
	data_per_page := 10
	if(page != ""){
		if(per_page !=""){
			page_index, _ := strconv.Atoi(page)
			index_per_page, _ := strconv.Atoi(per_page)
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page
			data_per_page = index_per_page
		}else{
			page_index, _ := strconv.Atoi(page)
			index_per_page := 10
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page
			data_per_page = index_per_page
		}
	}else{
		if(per_page!=""){
			page_index := 1
			index_per_page, _ := strconv.Atoi(per_page)
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page
			data_per_page = index_per_page
		}else{
			page_index := 1
			index_per_page := 10
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page
			data_per_page = index_per_page
		}
	}
	results := service.GetListProduct(condition,limit)
	var data []cores.Product
	for results.Next(){
		var dataPage cores.Product
		results.Scan(&dataPage.Product_id,&dataPage.Product_name,&dataPage.Size,&dataPage.Gender,&dataPage.Category);
		data = append(data, dataPage)
	}
	results_total := service.GetTotalListProduct(condition)
	var count int
	for results_total.Next() {   
		if err := results_total.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	var dataProductPage cores.DataProductPage
	dataProductPage.Page = page
	sum := math.Ceil(float64(count)/float64(data_per_page))
	dataProductPage.Total_page = strconv.FormatFloat(sum, 'f',0, 64)
	dataProductPage.Product = data
	return dataProductPage
}

func CreateOrder(address string,currentTime string,product_id []string,product_amount []string) cores.OrderReturn {
	order_buy_id := service.InsertOrderBuy(address,currentTime)
	var orderReturn cores.OrderReturn
	for i := 0; i < len(product_id) ; i++ {
		orderReturn.Result = service.InsertProductOrderBuy(order_buy_id,product_id[i],product_amount[i])
	}
	return orderReturn
}

func OrderList(start string,end string,status string,page string,per_page string) cores.DataOderPage{
	condition := ""
	if(start != ""){
		if(condition != ""){
			if(end != ""){
				condition += " AND order_buy.date >= '"+start+"' AND order_buy.date <= '"+end+"'"
			}else{
				condition += " AND order_buy.date >= '"+start+"' AND order_buy.date <= '"+start+"'"
			}
		}else{
			if(end != ""){
				condition += " WHERE order_buy.date >= '"+start+"' AND order_buy.date <= '"+end+"'"
			}else{
				condition += " WHERE order_buy.date >= '"+start+"' AND order_buy.date <= '"+start+"'"
			}
		}
	}
	if(status != ""){
		if(condition != ""){
			condition += " AND order_buy.status = '"+status+"'"
		}else{
			condition += " WHERE order_buy.status = '"+status+"'"
		}
	}
	limit := ""
	data_per_page := 10
	if(page != ""){
		if(per_page !=""){
			page_index, _ := strconv.Atoi(page)
			index_per_page, _ := strconv.Atoi(per_page)
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page
			data_per_page = index_per_page
		}else{
			page_index, _ := strconv.Atoi(page)
			index_per_page := 10
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page
			data_per_page = index_per_page
		}
	}else{
		if(per_page !=""){
			page_index := 1
			index_per_page, _ := strconv.Atoi(per_page)
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page
			data_per_page = index_per_page
		}else{
			page_index := 1
			index_per_page := 10
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page
			data_per_page = index_per_page
		}
	}
	
	results_order_buy := service.GetListOrderBuy(condition,limit)
	var data_Order []cores.Order_buy
	for results_order_buy.Next(){
		var data cores.Order_buy
		results_order_buy.Scan(&data.Order_buy_id,&data.Address,&data.Date,&data.Status,&data.Product_id,&data.Product_name,&data.Size,&data.Gender,&data.Category,&data.Product_amount)
		data_Order = append(data_Order, data)
	}
	results_total := service.GetTotalListOrderBuy(condition)
	var count int
	for results_total.Next() {   
		if err := results_total.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	var dataOderPage cores.DataOderPage
	dataOderPage.Page = page
	sum := math.Ceil(float64(count)/float64(data_per_page))
	dataOderPage.Total_page = strconv.FormatFloat(sum, 'f',0, 64)
	dataOderPage.Order_buy = data_Order

	return dataOderPage
}
