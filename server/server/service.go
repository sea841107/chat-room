package server

import (
	"chatserver/user"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Service = &service{
	id:   "ServerService",
	port: "8000",
}

var routes []Route

type service struct {
	id   string
	port string
}

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

func init() {
	register("POST", "/chat/data", httpPost, nil)
	register("GET", "/chat/data/{value}", httpGet, nil)
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}

func (s *service) Start() {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}

	s.serve(r)
}

func (s *service) serve(router *mux.Router) {
	if err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), router); err != nil {
		panic(err)
	}
}

func httpPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Read Body Error", err)
		httpResponse(w, http.StatusBadRequest, nil)
		return
	}

	response := make(chan []byte, 1)

	go user.Service.ProcessMsg(msg, response)

	handleResponse(w, response)
}

func httpGet(w http.ResponseWriter, r *http.Request) {
	var msg []byte

	vars := mux.Vars(r)
	if v, ok := vars["value"]; ok {
		msg = []byte(v)
	} else {
		log.Println("Empty Value")
		httpResponse(w, http.StatusBadRequest, nil)
		return
	}

	response := make(chan []byte, 1)

	go user.Service.ProcessMsg(msg, response)

	handleResponse(w, response)
}

func handleResponse(w http.ResponseWriter, response chan []byte) {
	select {
	case data, ok := <-response:
		if !ok {
			httpResponse(w, http.StatusRequestTimeout, nil)
			return
		}

		if data != nil {
			httpResponse(w, http.StatusOK, data)
		} else {
			httpResponse(w, http.StatusBadRequest, nil)
		}

	case <-time.After(1000 * time.Second):
		httpResponse(w, http.StatusRequestTimeout, nil)
		close(response)
	}
}

func httpResponse(w http.ResponseWriter, header int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Accept, Origin, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	w.WriteHeader(header)
	w.Write(response)
}
