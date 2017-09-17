/*
Package sptrans is a Go client library for the SPTrans Olho Vivo API.

Usage:

        import "github.com/sergioaugrod/go-sptrans/sptrans"

Authentication:

	token = "123456"
	client = sptrans.NewClient(token)
        client.Authenticate()

Search Route by description or number:

        client.Route.Search("Lapa")
        client.Route.Search("8000")

Search Route by direction:

        client.Route.SearchByDirection("Lapa", 1)
        client.Route.SearchByDirection("Lapa", 2)

Search Stop by name or address:

        client.Stop.Search("Afonso")
        client.Stop.Search("Rua Baltharzar da Veiga")

Search Stop by route code:

        client.Stop.SearchByRoute(1273)

Search Stop by corridor code:

	client.Stop.SearchByCorridor(8)
*/
package sptrans
