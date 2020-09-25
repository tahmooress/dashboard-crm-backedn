package handlers

import (
	dbf "dashboard/db-go"
	"dashboard/util-go"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//AddProduct is a handler to add insert new product to db
func AddProduct(env *dbf.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var product util.Product
		fmt.Println(product)
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &product)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v", product)
		_, err = env.DB.Exec(dbf.AddNewProduct, product.ProductID, product.ProductName, product.ProductBrand, product.Price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			fmt.Println(err)
		}
		// rows, err := result.RowsAffected()
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// }
		w.Write([]byte("successfuly added"))
	})
}

//Products is a handlers for products
func Products(env *dbf.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var product util.Product
		var ps []util.Product
		rows, err := env.DB.Query(dbf.AllProducts)
		defer rows.Close()
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		for rows.Next() {
			if err := rows.Scan(&product.ProductID, &product.ProductName, &product.ProductBrand, &product.Price); err != nil {
				w.Write([]byte(err.Error()))
			}
			fmt.Println(product)
			ps = append(ps, product)
		}
		json.NewEncoder(w).Encode(ps)
	})
}

//AddStore is ....
func AddStore(env *dbf.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var store util.Store
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &store)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		_, err = env.DB.Exec(dbf.AddNewStore, store.StoreName, store.Phone, store.StoreAddress, store.Proxy, store.Mobile)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write([]byte("successfully added"))
	})
}

//Stores is a ...
func Stores(env *dbf.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var store util.Store
		rows, err := env.DB.Query(dbf.AllStores)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		var st []util.Store
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&store.StoreID, &store.StoreName, &store.Phone, &store.StoreAddress, &store.Proxy, &store.Mobile)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			st = append(st, store)
		}
		json.NewEncoder(w).Encode(st)
	})
}

//HandleStoreEdit is a handler for ....
func HandleStoreEdit(env *dbf.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var store util.Store
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &store)
		fmt.Println(store)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		_, err = env.DB.Exec(dbf.UpdateStore, store.StoreID, store.StoreName, store.Phone, store.StoreAddress, store.Proxy, store.Mobile)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write([]byte("successfuly updated"))
	})
}

//HandleStore is a handler for ...
func HandleStore(env *dbf.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, ok := r.URL.Query()["id"]
		if !ok || len(id[0]) < 1 {
			w.Write([]byte("wrong id number"))
		}
		row, err := env.DB.Query(dbf.SingleStore, id[0])
		defer row.Close()
		var store util.Store
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		for row.Next() {
			if err := row.Scan(&store.StoreID, &store.StoreName, &store.Phone, &store.StoreAddress, &store.Proxy, &store.Mobile); err != nil {
				w.Write([]byte(err.Error()))
			}
		}
		json.NewEncoder(w).Encode(store)
	})
}

//HandleAddTrans is a handler for ...
func HandleAddTrans(env *dbf.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, ok := r.URL.Query()["id"]
		if !ok || len(id[0]) < 1 {
			w.Write([]byte("wrong id number"))
			return
		}
		i, err := strconv.Atoi(id[0])
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(i)
		var factor []util.Factor
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &factor)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		transID := 0
		err = env.DB.QueryRow(dbf.AddNewTransactions, i).Scan(&transID)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		for _, v := range factor {
			_, err := env.DB.Exec(dbf.AddNewFactor, transID, v.ProductID, v.Quantity)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}
		w.Write([]byte("factors added successfuly"))
	})
}

//AllTransOfStore is a handler for all the transactions of a paricular store
func AllTransOfStore(env *dbf.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, ok := r.URL.Query()["id"]
		if !ok || len(id[0]) < 1 {
			w.Write([]byte("wrong id number"))
			return
		}
		i, err := strconv.Atoi(id[0])
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		var trans []util.Transaction
		rows, err := env.DB.Query(dbf.SelectAllTransIDs, i)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		defer rows.Close()
		for rows.Next() {
			var temp util.Transaction
			err := rows.Scan(&temp.TransID, &temp.StoreID, &temp.Date)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			r, err := env.DB.Query(dbf.SelectAllFactorsOfSingleTrans, temp.TransID)
			defer r.Close()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			for r.Next() {
				var tf util.Factor
				err := r.Scan(&tf.FactorID, &tf.TransID, &tf.ProductID, &tf.Quantity)
				if err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				temp.Factors = append(temp.Factors, tf)
			}
			trans = append(trans, temp)
		}
		json.NewEncoder(w).Encode(trans)
	})
}
