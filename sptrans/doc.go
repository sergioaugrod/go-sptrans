/*
Package sptrans is a Go client library for the SPTrans Olho Vivo API.

Usage:

	import "github.com/sergioaugrod/go-sptrans/sptrans"

Authentication:

	token = "123456"
	client = sptrans.NewClient(token)
	client.Authenticate()

Endpoints:

Lines

       // Search by line description or number
       client.Line.Search("Lapa")
       client.Line.Search("8000")

       // Search by line direction
       client.Line.SearchByDirection("Lapa", 1)

Stops

       // Search by stop name or address
       client.Stop.Search("Afonso")
       client.Stop.Search("Rua Baltharzar da Veiga")

       // Search by line code
       client.Stop.SearchByLine(1273)

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
