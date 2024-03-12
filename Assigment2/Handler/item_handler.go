package handler

import (
	entity "Assigment2/Entity"
	"context"
	_ "context" // Impor "context" tidak digunakan dalam kode ini
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	_ "time" // Impor "time" tidak digunakan dalam kode ini

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // Driver PostgreSQL diimpor tetapi tidak digunakan dalam kode ini
)

// ItemHandlerInterface mendefinisikan antarmuka untuk handler item
type ItemHandlerInterface interface {
	ItemsHandler(w http.ResponseWriter, r *http.Request)
}

// ItemHandler adalah implementasi dari ItemHandlerInterface
type ItemHandler struct {
	db *sql.DB
}

// NewItemHandler membuat instans baru dari ItemHandlerInterface
func NewItemHandler(db *sql.DB) ItemHandlerInterface {
	return &ItemHandler{db: db}
}

// ItemsHandler menangani permintaan HTTP berdasarkan metode
func (h *ItemHandler) ItemsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	switch r.Method {
	case http.MethodGet:
		h.getItemsHandler(w, r)
	case http.MethodPost:
		h.createItemsHandler(w, r)
	case http.MethodPut:
		h.UpdateOrderById(w, r, id)
	case http.MethodDelete:
		h.DeleteOrderbyId(w, r, id)
	}
}

// getItemsHandler menangani permintaan GET untuk mendapatkan data pesanan
func (h *ItemHandler) getItemsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	queryString := `SELECT
		o.order_id as order_id
		,o.customer_name
		,o.ordered_at
		,json_agg(json_build_object(
			'item_id',i.item_id
			,'item_code',i.item_code
			,'description',i.description
			,'quantity',i.quantity
			,'order_id',i.order_id
		)) as items
	FROM orders o JOIN items i
	ON o.order_id = i.order_id
	GROUP BY o.order_id`
	rows, err := h.db.QueryContext(ctx, queryString)
	if err != nil {
		fmt.Println("query row error", err)
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var o entity.Order
		var itemsStr string
		if serr := rows.Scan(&o.Order_id, &o.Customer_Name, &o.Ordered_At, &itemsStr); serr != nil {
			fmt.Println("Scan error", serr)
		}
		var items []entity.Item
		if err := json.Unmarshal([]byte(itemsStr), &items); err != nil {
			fmt.Errorf("Error when parsing items")
		} else {
			o.Item = append(o.Item, items...)
		}
		orders = append(orders, &o)
	}
	jsonData, _ := json.Marshal(&orders)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

// createItemsHandler menangani permintaan POST untuk membuat pesanan baru
func (h *ItemHandler) createItemsHandler(w http.ResponseWriter, r *http.Request) {
	var newOrder entity.Order
	json.NewDecoder(r.Body).Decode(&newOrder)
	fmt.Println(newOrder)
	sqlStatment := `INSERT INTO orders
	(Customer_Name,Ordered_At)
	VALUES ($1 ,$2) RETURNING order_id ;`
	ctx := context.Background()
	var id int
	err := h.db.QueryRowContext(ctx, sqlStatment, newOrder.Customer_Name, time.Now()).Scan(&id)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(newOrder.Item); i++ {
		var items entity.Item
		items.ItemCode = newOrder.Item[i].ItemCode
		items.Description = newOrder.Item[i].Description
		items.Quantity = newOrder.Item[i].Quantity
		query := `INSERT INTO items 
		(item_code,description,quantity,order_id)
		VALUES ($1,$2,$3,$4) `

		_, err := h.db.Exec(query, items.ItemCode, items.Description, items.Quantity, id)
		if err != nil {
			panic(nil)
		}
	}

	w.Write([]byte(fmt.Sprint("Create user rows ")))
	return
}

// UpdateOrderById menangani permintaan PUT untuk memperbarui pesanan berdasarkan ID
func (h *ItemHandler) UpdateOrderById(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" { // get by id
		var newOrder entity.Order
		json.NewDecoder(r.Body).Decode(&newOrder)
		sqlstatment := `
		UPDATE orders SET customer_name = $1 , ordered_at = $2 WHERE order_id = $3;`

		res, err := h.db.Exec(sqlstatment,
			newOrder.Customer_Name,
			time.Now(),
			id,
		)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(newOrder.Item); i++ {
			var items entity.Item
			items.Item_Id = newOrder.Item[i].Item_Id
			items.ItemCode = newOrder.Item[i].ItemCode
			items.Description = newOrder.Item[i].Description
			items.Quantity = newOrder.Item[i].Quantity
			query := `UPDATE items SET item_code = $1, description = $2, quantity = $3 WHERE order_id = $4 AND item_id = $5`

			_, err := h.db.Exec(query, items.ItemCode, items.Description, items.Quantity, id, items.Item_Id)
			if err != nil {
				panic(nil)
			}
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}

		w.Write([]byte(fmt.Sprint("User  update ", count)))
		return
	}
}

// DeleteOrderbyId menangani permintaan DELETE untuk menghapus pesanan berdasarkan ID
func (h *ItemHandler) DeleteOrderbyId(w http.ResponseWriter, r *http.Request, id string) {
	sqlstament := `DELETE FROM orders WHERE Order_id = $1;`
	if idInt, err := strconv.Atoi(id); err == nil {
		sqlstament2 := `DELETE FROM items WHERE Order_id = $1;`
		_, err2 := h.db.Exec(sqlstament2, idInt)
		if err2 != nil {
			panic(err)
		}
		res, err := h.db.Exec(sqlstament, idInt)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}

		w.Write([]byte(fmt.Sprint("Delete user rows ", count)))
		return
	}
}
