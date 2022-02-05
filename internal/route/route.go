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
	//mux.HandleFunc("/suppliers/**/products/", repProvider.IndividualProduct)

	http.HandleFunc("/login", handlerProvider.Login)
	http.Handle("/profile", m.RequireAuthentication(http.HandlerFunc(handlerProvider.Profile)))
	http.HandleFunc("/refresh", handlerProvider.Refresh)
	http.HandleFunc("/registration", handlerProvider.Registration)
	http.Handle("/logout", m.RequireAuthentication(http.HandlerFunc(handlerProvider.Logout)))

}
