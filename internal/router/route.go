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

	mux.HandleFunc("/suppliers", c.Menu.Suppliers)
	mux.HandleFunc("/suppliers/", c.Menu.SupplierAndProdWithID)
	mux.HandleFunc("/categories", c.Menu.Categories)
	mux.HandleFunc("/categories/", c.Menu.ListOfProductsInSpecificCategory)
	mux.HandleFunc("/all-products", c.Menu.ListOfAllProducts)

	//Auth
	mux.HandleFunc("/login", c.Auth.Login)
	mux.HandleFunc("/registration", c.Auth.Registration)
	mux.Handle("/logout", m.RequireAuthentication(http.HandlerFunc(c.Auth.Logout)))

	//User
	mux.Handle("/profile", m.RequireAuthentication(http.HandlerFunc(c.User.Profile)))
	mux.HandleFunc("/refresh", c.User.Refresh)

}
