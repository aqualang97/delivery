package router

import (
	"delivery/internal/auth/middlware"
	"delivery/internal/controllers"
	"net/http"
)

func Router(
	c *controllers.Controllers,
	mux *http.ServeMux,
) {
	m := middlware.NewMiddleware(c)
	//Menu
	mux.HandleFunc("/", c.Menu.Home)
	mux.HandleFunc("/suppliers", m.CORS(http.HandlerFunc(c.Menu.Suppliers)))
	mux.HandleFunc("/suppliers/", m.CORS(http.HandlerFunc(c.Menu.SupplierAndProdWithID)))
	mux.HandleFunc("/categories", m.CORS(http.HandlerFunc(c.Menu.Categories)))
	mux.HandleFunc("/categories/", m.CORS(http.HandlerFunc(c.Menu.ListOfProductsInSpecificCategory)))
	mux.HandleFunc("/all-products", m.CORS(http.HandlerFunc(c.Menu.ListOfAllProducts)))

	//Auth
	mux.HandleFunc("/login", m.CORS(http.HandlerFunc(c.Auth.Login)))
	mux.HandleFunc("/registration", m.CORS(http.HandlerFunc(c.Auth.Registration)))
	mux.Handle("/logout", m.RequireAuthentication(http.HandlerFunc(c.Auth.Logout)))

	//User
	mux.Handle("/profile", m.RequireAuthentication(http.HandlerFunc(c.User.Profile)))
	mux.HandleFunc("/refresh", m.CORS(http.HandlerFunc(c.User.Refresh)))

	//Cart
	mux.HandleFunc("/products-in-cart", m.CORS(http.HandlerFunc(c.User.GetProductsInCart)))
	mux.HandleFunc("/card_pay", m.CORS(http.HandlerFunc(c.User.SimulationCardPay)))
	mux.HandleFunc("/confirm", m.CORS(http.HandlerFunc(c.User.AddProductsFromCart)))
	mux.HandleFunc("/old-orders/", m.CORS(http.HandlerFunc(c.User.GetOldOrders)))
}
