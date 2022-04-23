package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


type OrderReturn struct {
	Result			bool	`json:"results"`
}

type Order_buy struct {
	Order_buy_id	string	`json:"order_buy_id"`
	Address 		string	`json:"address"`
	Date 			string	`json:"date"`
	Status 			string	`json:"status"`
	Product_id 		string	`json:"product_id"`
	Product_name 	string	`json:"product_name"`
	Size 			string	`json:"size"`
	Gender 			string	`json:"gender"`
	Category 		string	`json:"category"`
	Product_amount 	string	`json:"product_amount"`
}

type Product struct {
	Product_id 		string	`json:"product_id"`
	Product_name 	string	`json:"product_name"`
	Size 			string	`json:"size"`
	Gender 			string	`json:"gender"`
	Category 		string	`json:"category"`
}

type DataProductPage struct {
	Page 			string	`json:"page"`
	Total_page 		string	`json:"total_page"`
	Product 		[]Product	`json:"products"`
}

type DataOderPage struct {
	Page 			string	`json:"page"`
	Total_page 		string	`json:"total_page"`
	Order_buy 		[]Order_buy	`json:"order_buys"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome homess!")
}

func productList(w http.ResponseWriter, r *http.Request) {
	gender := r.URL.Query()["gender"]
	category := r.URL.Query()["category"]
	size := r.URL.Query()["size"]
	page := r.URL.Query()["page"]
	per_page := r.URL.Query()["per_page"]
	condition := ""
	if(gender[0] != ""){
		if(condition != ""){
			condition += " AND product.gender_id = '"+gender[0]+"'"
		}else{
			condition += " WHERE product.gender_id = '"+gender[0]+"'"
		}
	}
	if(category[0] != ""){
		if(condition != ""){
			condition += " AND product.category_id = '"+category[0]+"'"
		}else{
			condition += " WHERE product.category_id = '"+category[0]+"'"
		}
	}
	if(size[0] != ""){
		if(condition != ""){
			condition += " AND product.size_id = '"+size[0]+"'"
		}else{
			condition += " WHERE product.size_id = '"+size[0]+"'"
		}
	}
	limit := ""
	data_per_page := 10
	if(page[0] != ""){
		if(per_page[0] !=""){
			page_index, _ := strconv.Atoi(page[0])
			index_per_page, _ := strconv.Atoi(per_page[0])
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page[0]
			data_per_page = index_per_page
		}else{
			page_index, _ := strconv.Atoi(page[0])
			index_per_page := 10
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page[0]
			data_per_page = index_per_page
		}
	}else{
		if(per_page[0]!=""){
			page_index := 1
			index_per_page, _ := strconv.Atoi(per_page[0])
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page[0]
			data_per_page = index_per_page
		}else{
			page_index := 1
			index_per_page := 10
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page[0]
			data_per_page = index_per_page
		}
	}
	
	db, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
	}
	results, err := db.Query(`SELECT
								product.product_id,
								product.product_name,
								size.size,
								gender.gender,
								category.category
							FROM
								product
							LEFT JOIN size ON size.size_id = product.size_id
							LEFT JOIN gender ON gender.gender_id = product.gender_id
							LEFT JOIN category ON category.category_id = product.category_id`+condition+" "+limit)
	if err != nil {
		panic(err)
	}
	defer results.Close()

	var data []Product
	for results.Next(){
		var dataPage Product
		err = results.Scan(&dataPage.Product_id,&dataPage.Product_name,&dataPage.Size,&dataPage.Gender,&dataPage.Category);
		if err != nil {
			panic(err)
		}
		data = append(data, dataPage)
	}
	results_total, err := db.Query("SELECT COUNT(*) FROM product "+condition)
	if err != nil {
		log.Fatal(err)
	}
	defer results_total.Close()
	var count int
	for results_total.Next() {   
		if err := results_total.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	var dataProductPage DataProductPage
	dataProductPage.Page = page[0]
	sum := math.Ceil(float64(count)/float64(data_per_page))
	dataProductPage.Total_page = strconv.FormatFloat(sum, 'f',0, 64)
	dataProductPage.Product = data

	jsonResp, err := json.Marshal(dataProductPage)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResp)
	
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	var product_id []string = r.PostForm["product_id"]
	var product_amount []string = r.PostForm["product_amount"]
	var address = r.FormValue("address")
	currentTime := time.Now().String()

	db, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
	}
	res, err := db.Exec("INSERT INTO order_buy (address, date, status) VALUES ('"+address+"', '"+currentTime+"', '0')")
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
    if err != nil {
        panic(err)
    }
	var order_buy_id string = strconv.Itoa(int(id))
	for i := 0; i < len(product_id) ; i++ {
		_, err = db.Exec("INSERT INTO product_order_buy (order_buy_id, product_id, product_amount) VALUES ('"+order_buy_id+"', '"+product_id[i]+"', '"+product_amount[i]+"')")
		if err != nil {
			panic(err)
		}
	}
	var orderReturn OrderReturn
	orderReturn.Result = true
	jsonResp, err := json.Marshal(orderReturn)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResp)
}

func orderList(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query()["start"]
	end := r.URL.Query()["end"]
	status := r.URL.Query()["status"]
	page := r.URL.Query()["page"]
	per_page := r.URL.Query()["per_page"]
	condition := ""
	if(start[0] != ""){
		if(condition != ""){
			if(end[0] != ""){
				condition += " AND order_buy.date >= '"+start[0]+"' AND order_buy.date <= '"+end[0]+"'"
			}else{
				condition += " AND order_buy.date >= '"+start[0]+"' AND order_buy.date <= '"+start[0]+"'"
			}
		}else{
			if(end[0] != ""){
				condition += " WHERE order_buy.date >= '"+start[0]+"' AND order_buy.date <= '"+end[0]+"'"
			}else{
				condition += " WHERE order_buy.date >= '"+start[0]+"' AND order_buy.date <= '"+start[0]+"'"
			}
		}
	}
	if(status[0] != ""){
		if(condition != ""){
			condition += " AND order_buy.status = '"+status[0]+"'"
		}else{
			condition += " WHERE order_buy.status = '"+status[0]+"'"
		}
	}
	limit := ""
	data_per_page := 10
	if(page[0] != ""){
		if(per_page[0] !=""){
			page_index, _ := strconv.Atoi(page[0])
			index_per_page, _ := strconv.Atoi(per_page[0])
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page[0]
			data_per_page = index_per_page
		}else{
			page_index, _ := strconv.Atoi(page[0])
			index_per_page := 10
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page[0]
			data_per_page = index_per_page
		}
	}else{
		if(per_page[0]!=""){
			page_index := 1
			index_per_page, _ := strconv.Atoi(per_page[0])
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page[0]
			data_per_page = index_per_page
		}else{
			page_index := 1
			index_per_page := 10
			page := strconv.Itoa((page_index-1)*index_per_page)
			limit += " LIMIT "+page+","+per_page[0]
			data_per_page = index_per_page
		}
	}
	
	db, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
	}
	results_order_buy, err := db.Query(`SELECT
											order_buy.order_buy_id,
											order_buy.address,
											order_buy.date,
											order_buy.status,
											product.product_id,
											product.product_name,
											size.size,
											gender.gender,
											category.category,
											product_order_buy.product_amount
										FROM
											order_buy
										LEFT JOIN product_order_buy ON product_order_buy.order_buy_id = order_buy.order_buy_id
										LEFT JOIN product ON product.product_id = product_order_buy.product_order_buy_id
										LEFT JOIN size ON size.size_id = product.size_id
										LEFT JOIN gender ON gender.gender_id = product.gender_id
										LEFT JOIN category ON category.category_id = product.category_id
										`+condition+` `+limit)
	if err != nil {
		panic(err)
	}
	defer results_order_buy.Close()
	var data_Order []Order_buy
	for results_order_buy.Next(){
		var data Order_buy
		err = results_order_buy.Scan(&data.Order_buy_id,&data.Address,&data.Date,&data.Status,&data.Product_id,&data.Product_name,&data.Size,&data.Gender,&data.Category,&data.Product_amount)
		if err != nil {
			panic(err)
		}
		data_Order = append(data_Order, data)
	}
	results_total, err := db.Query(`SELECT
										COUNT(*)
									FROM
										order_buy
									LEFT JOIN product_order_buy ON product_order_buy.order_buy_id = order_buy.order_buy_id
									LEFT JOIN product ON product.product_id = product_order_buy.product_order_buy_id
									LEFT JOIN size ON size.size_id = product.size_id
									LEFT JOIN gender ON gender.gender_id = product.gender_id
									LEFT JOIN category ON category.category_id = product.category_id
									`+condition+` `+limit)
	if err != nil {
		log.Fatal(err)
	}
	defer results_total.Close()
	var count int
	for results_total.Next() {   
		if err := results_total.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	

	var dataOderPage DataOderPage
	dataOderPage.Page = page[0]
	sum := math.Ceil(float64(count)/float64(data_per_page))
	dataOderPage.Total_page = strconv.FormatFloat(sum, 'f',0, 64)
	dataOderPage.Order_buy = data_Order

	jsonResp, err := json.Marshal(dataOderPage)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResp)
	
}

func main() {
	
	db, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("DB connect success ...");

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/productList/", productList)
	router.HandleFunc("/createOrder/", createOrder)
	router.HandleFunc("/orderList/", orderList)
	
	fmt.Println("server running on prot 8080 ...");
	http.ListenAndServe(":8080", router)
	
}