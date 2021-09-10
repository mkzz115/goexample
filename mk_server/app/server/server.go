package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"time"
)

type ZHttpServer struct {
	Addr string
	Timeout time.Duration
	router *httprouter.Router
}

func NewZHttpServer(addr string, timeout int) *ZHttpServer {
	return &ZHttpServer{
		Addr: addr,
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func (s *ZHttpServer) SetRouter(r *httprouter.Router)  {
	s.router = r
}

func (s *ZHttpServer) StartServer() error {
	serv := &http.Server{
		Addr: s.Addr,
		Handler: s.router,
		ReadHeaderTimeout: s.Timeout,
	}
	return serv.ListenAndServe()
}

// handle
func NewHttpRouter() *httprouter.Router {
	r := httprouter.New()
	r.HandlerFunc(http.MethodGet, "/hello", Hello)
	return r
}

func Hello(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	println(buf)
	res := struct {
		Value string `json:"value"`
	}{"hello"}
	resData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resData)
}