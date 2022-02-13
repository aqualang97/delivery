package route

import (
	handProv "delivery/internal/auth/handle_provide"
	"delivery/internal/auth/middlware"
	"delivery/internal/repositories_provider"
	"net/http"
)

func Router(
	handlerProvider handProv.HandlerProvider,
	repProvider repositories_provider.RepositoriesProvider,
	mux *http.ServeMux,
) {
	m := middlware.NewMiddleware(&handlerProvider)
	mux.HandleFunc("/suppliers", repProvider.Suppliers)
	mux.HandleFunc("/suppliers/", repProvider.SupplierAndProdWithID)

	mux.HandleFunc("/categories", repProvider.Categories)
	mux.HandleFunc("/categories/", repProvider.ListOfProductsInSpecificCategory)
	mux.HandleFunc("/all-products", repProvider.ListOfAllProducts)
	mux.HandleFunc("/login", handlerProvider.Login)
	mux.Handle("/profile", m.RequireAuthentication(http.HandlerFunc(handlerProvider.Profile)))
	mux.HandleFunc("/refresh", handlerProvider.Refresh)
	mux.HandleFunc("/registration", handlerProvider.Registration)
	mux.Handle("/logout", m.RequireAuthentication(http.HandlerFunc(handlerProvider.Logout)))

}
