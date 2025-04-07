package api

import (
	"awesomeProject/service/cart"
	"awesomeProject/service/order"
	"awesomeProject/service/product"
	"awesomeProject/service/user"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}
func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)
	cartHAndler := cart.NewHandler(orderStore, productStore, userStore)
	cartHAndler.RegisterRoutes(subrouter)

	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
