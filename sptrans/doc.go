/*
Package sptrans is a Go client library for the SPTrans Olho Vivo API.

Usage:

	import "github.com/sergioaugrod/go-sptrans/sptrans"

Authentication:

	token = "123456"
	client = sptrans.NewClient(token)
	client.Authenticate()

Endpoints:

Routes

	// Search by route description or number
	client.Route.Search("Lapa")
	client.Route.Search("8000")

       // Search by route direction
       client.Route.SearchByDirection("Lapa", 1)

Stops

       // Search by stop name or address
       client.Stop.Search("Afonso")
       client.Stop.Search("Rua Baltharzar da Veiga")

       // Search by route code
       client.Stop.SearchByRoute(1273)

       // Search by corridor code
       client.Stop.SearchByCorridor(8)

Corridors

       // All smart corridors
       client.Corridor.All()

Companies

       // All companies
       client.Company.All()
*/
package sptrans
