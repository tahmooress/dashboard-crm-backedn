package main

import (
	dbf "dashboard/db-go"
	rh "dashboard/handlers-go"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	/// setup router
	r := mux.NewRouter()
	method := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "UPDATE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"})
	db, err := dbf.RunDB()
	if err != nil {
		panic(err)
	}
	env := &dbf.Env{DB: db}
	r.Handle("/addProduct", rh.AddProduct(env)).Methods("POST")
	r.Handle("/addStore", rh.AddStore(env)).Methods("POST")
	r.Handle("/products", rh.Products(env)).Methods("GET")
	r.Handle("/stores", rh.Stores(env)).Methods("GET")
	r.Handle("/storeEdit", rh.HandleStoreEdit(env)).Methods("POST")
	r.Handle("/singleStore", rh.HandleStore(env)).Methods("GET")
	r.Handle("/addTrans", rh.HandleAddTrans(env)).Methods("POST")
	r.Handle("/factors", rh.AllTransOfStore(env)).Methods("GET")
	//runing and listening to server on port 8000
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, method, origin)(r)))
	defer db.Close()
}
