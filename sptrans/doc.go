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

       // Search by line id
       client.Stop.SearchByLine(1273)

       // Search by corridor id
       client.Stop.SearchByCorridor(8)

Corridors

       // All smart corridors
       client.Corridor.All()

Companies

       // All companies
       client.Company.All()

Vehicles Position

       // All Vehicles Position
       client.VehiclePosition.All()

       // Search by line id
       client.VehiclePosition.SearchByLine(1273)

       // Search by company id
       client.VehiclePosition.SearchByCompany(99)

Forecast

       // Search by stop id and line id
       client.Forecast.Search(2003, 2004)

       // Search by line id
       client.Forecast.SearchByLine(1273)

       // Search by stop id
       client.Forecast.SearchByStop(8)
*/
package sptrans
