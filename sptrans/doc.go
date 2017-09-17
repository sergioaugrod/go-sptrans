/*
Package go-sptrans is a Go client library for the SPTrans Olho Vivo API.

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
*/
package sptrans
