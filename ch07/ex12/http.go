package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var itemList = template.Must(template.New("itemlist").Parse(`
<html>
<body>
<h1>List</h1>
<table>
	<tr>
		<th>Item</th>
		<th>Price</th>
	</tr>
	{{ range $item, $price := . }}
	<tr>
		<td>{{ $item }}</td>
		<td>{{ $price }}</td>
	</tr>
	{{ end }}
</table>
</body>
</html>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	itemList.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parameter %q is required\n", "item")
		return
	}

	// すでにあったら400
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item %q already exists\n", item)
		return
	}

	// priceが不正だったら400
	price := req.URL.Query().Get("price")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parameter %q is required\n", "price")
		return
	}
	pricef, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price %q\n", price)
		return
	}

	// okだったら作成
	db[item] = dollars(pricef)
	fmt.Fprintf(w, "item %q created\n", item)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parameter %q is required\n", "item")
		return
	}

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	// priceが不正だったら400
	price := req.URL.Query().Get("price")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parameter %q is required\n", "price")
		return
	}
	pricef, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price %q\n", price)
		return
	}

	// okだったら更新
	db[item] = dollars(pricef)
	fmt.Fprintf(w, "item %q updated\n", item)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parameter %q is required\n", "item")
		return
	}

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	// okだったら削除
	delete(db, item)
	fmt.Fprintf(w, "item %q deleted\n", item)
}
