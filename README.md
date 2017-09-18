# go-sptrans [![Build Status](https://travis-ci.org/sergioaugrod/go-sptrans.svg?branch=master)](https://travis-ci.org/sergioaugrod/go-sptrans) [![GoDoc](https://godoc.org/github.com/sergioaugrod/go-sptrans?status.png)](https://godoc.org/github.com/sergioaugrod/go-sptrans) [![Go Report Card](https://goreportcard.com/badge/github.com/sergioaugrod/go-sptrans)](https://goreportcard.com/report/github.com/sergioaugrod/go-sptrans)


go-sptrans is a Go client library for the [SPTrans Olho Vivo API](http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo.aspx).

### Features

- [x] Lines
- [x] Stops
- [x] Corridors
- [x] Companies
- [ ] Position of vehicles
- [ ] Arrival Forecast
- [ ] Speed on lanes

## Documentation

See the documentation at [godoc](https://godoc.org/github.com/sergioaugrod/go-sptrans).

## Install

go get github.com/sergioaugrod/go-sptrans/sptrans

## Usage

### Import

```go
import "github.com/sergioaugrod/go-sptrans/sptrans"
```

### Authentication

```go
token = "123456"
client = sptrans.NewClient(token)
client.Authenticate()
```

### Endpoints

#### Lines

```go
// Search by line description or number
client.Line.Search("Lapa")
client.Line.Search("8000")

// Search by line direction
client.Line.SearchByDirection("Lapa", 1)
```

#### Stops

```go
// Search by stop name or address
client.Stop.Search("Afonso")
client.Stop.Search("Rua Baltharzar da Veiga")

// Search by line code
client.Stop.SearchByLine(1273)

// Search by corridor code
client.Stop.SearchByCorridor(8)
```

#### Corridors

```go
// All smart corridors
client.Corridor.All()
```

#### Companies

```go
// All companies
client.Company.All()
```

## Contributing

1. Clone it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## License

[MIT License](LICENSE)
