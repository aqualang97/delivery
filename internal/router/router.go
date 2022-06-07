package router

import (
	"database/sql"
	"delivery/internal/auth/middlware"
	"delivery/internal/controllers"
	"github.com/aqualang97/logger/v4"
	"net/http"
)

//type MyRouter struct {
//	C   *controllers.Controllers
//	mux *http.ServeMux
//}

func Router(
	c *controllers.Controllers,
	mux *http.ServeMux,
	conn *sql.DB,
	TX *sql.Tx,
	myLogger *logger.Logger,
) {

	m := middlware.NewMiddleware(c)
	m.CORS(m.RequireAuthentication(mux))
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
	mux.Handle("/logout", m.CORS(http.HandlerFunc(c.Auth.Logout)))

	//User
	mux.Handle("/profile", m.CORS(m.RequireAuthentication(http.HandlerFunc(c.User.Profile))))
	mux.HandleFunc("/refresh", m.CORS(http.HandlerFunc(c.User.Refresh)))

	//Cart
	mux.HandleFunc("/products-in-cart", m.CORS(http.HandlerFunc(c.User.GetProductsInCart)))
	mux.HandleFunc("/card_pay", m.CORS(http.HandlerFunc(c.User.SimulationCardPay)))
	mux.HandleFunc("/confirm", m.CORS(http.HandlerFunc(c.User.AddProductsFromCart)))
	mux.HandleFunc("/old-orders/", m.CORS(http.HandlerFunc(c.User.GetOldOrders)))

	//var supp models.Supplier

}
func settingsHeader(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

}
