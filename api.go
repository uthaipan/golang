package main

import (
	"encoding/json"
	"example/hello/controller"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request)  {
	test := controller.HomeLink()
	fmt.Fprintf(w, "%v",test)
}

func productList(w http.ResponseWriter, r *http.Request)  {
	gender := r.URL.Query()["gender"]
	category := r.URL.Query()["category"]
	size := r.URL.Query()["size"]
	page := r.URL.Query()["page"]
	per_page := r.URL.Query()["per_page"]

	jsonResp, err := json.Marshal(controller.ProductList(gender[0],category[0],size[0],page[0],per_page[0]))
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResp)
}

func createOrder(w http.ResponseWriter, r *http.Request)  {
	
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	var product_id []string = r.PostForm["product_id"]
	var product_amount []string = r.PostForm["product_amount"]
	var address = r.FormValue("address")
	currentTime := time.Now().String()

	jsonResp, err := json.Marshal(controller.CreateOrder(address,currentTime,product_id,product_amount))
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResp)
}

func orderList(w http.ResponseWriter, r *http.Request)  {
	start := r.URL.Query()["start"]
	end := r.URL.Query()["end"]
	status := r.URL.Query()["status"]
	page := r.URL.Query()["page"]
	per_page := r.URL.Query()["per_page"]

	jsonResp, err := json.Marshal(controller.OrderList(start[0],end[0],status[0],page[0],per_page[0]))
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResp)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/productList/", productList)
	router.HandleFunc("/createOrder/", createOrder)
	router.HandleFunc("/orderList/", orderList)
	
	fmt.Println("server running on prot 8080 ...");
	http.ListenAndServe(":8080", router)
	
}