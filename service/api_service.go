package service

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetListProduct(condition string,limit string) *sql.Rows{
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
							LEFT JOIN category ON category.category_id = product.category_id`+condition+` `+limit)
	if err != nil {
		panic(err)
	}
	return results
}

func GetTotalListProduct(condition string)  *sql.Rows{
	db, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
	}
	results_total, err := db.Query("SELECT COUNT(*) FROM product "+condition)
	if err != nil {
		log.Fatal(err)
	}
	return results_total
}

func InsertOrderBuy(address string,currentTime string)  (string){
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
	return strconv.Itoa(int(id))
}

func InsertProductOrderBuy(order_buy_id string,product_id string,product_amount string)  (bool){
	db, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO product_order_buy (order_buy_id, product_id, product_amount) VALUES ('"+order_buy_id+"', '"+product_id+"', '"+product_amount+"')")
	if err != nil {
		return false
	};
	return true
}
func GetListOrderBuy(condition string,limit string) *sql.Rows{
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
										LEFT JOIN product ON product.product_id = product_order_buy.product_id
										LEFT JOIN size ON size.size_id = product.size_id
										LEFT JOIN gender ON gender.gender_id = product.gender_id
										LEFT JOIN category ON category.category_id = product.category_id
										`+condition+` `+limit) 
	if err != nil {
		panic(err)
	}
	return results_order_buy
}

func GetTotalListOrderBuy(condition string)  *sql.Rows{
	db, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
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
									`+condition)
	if err != nil {
		panic(err)
	}
	return results_total
}