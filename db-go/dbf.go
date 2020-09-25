package db

import (
	"database/sql"
	"fmt"
)

//RunDB open a postgres database
func RunDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", SetDbInfo())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "900587101"
	dbname   = "dashboard"
)

//SetDbInfo sets the required information for connecting to database
func SetDbInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

//required statment for working with DataBase
const (
	AddNewStore                   = `INSERT INTO stores (store_name,store_phone, store_address, store_proxy,proxy_mobile) VALUES ($1, $2, $3, $4, $5)`
	AllStores                     = `SELECT * FROM stores`
	UpdateStore                   = `UPDATE stores SET store_name=$2, store_phone = $3, store_address=$4, store_proxy=$5, proxy_mobile=$6 WHERE store_id=$1`
	SingleStore                   = `SELECT * FROM stores WHERE store_id=$1`
	AddNewProduct                 = `INSERT INTO products (product_id, product_name, product_brand, product_price) VALUES ($1, $2, $3, $4)`
	AllProducts                   = `SELECT * FROM products`
	AddNewTransactions            = `INSERT INTO transactions (store_id) VALUES ($1) RETURNING trans_id`
	AddNewFactor                  = `INSERT INTO factors (trans_id, product_id, quantity) VALUES ($1, $2, $3)`
	SelectAllTransIDs             = `SELECT * from transactions WHERE store_id = $1`
	SelectAllFactorsOfSingleTrans = `SELECT * FROM factors WHERE trans_id=$1`
)

//Env is enviorment variable for connecting to db
type Env struct {
	DB *sql.DB
}
