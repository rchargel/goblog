package web

import (
	"encoding/csv"
	"io"
	"os"
)

type Route struct {
	method  string // GET,POST,etc
	path    string // relative to base
	handler string // the controller to call
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{make([]Route, 0, 10)}
}

func (r *Router) Init(ch chan bool) {
	file, err := os.Open("./conf/routes.csv")
	if err != nil {
		ch <- false
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == nil {
			if len(record) == 3 {
				route := Route{record[0], record[1], record[2]}
				r.routes = append(r.routes, route)
			} else {
				break
			}
		} else if err == io.EOF {
			break
		} else {
			ch <- false
			panic(err)
		}
	}
	ch <- true
}
