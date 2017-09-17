# go-sptrans [![Build Status](https://travis-ci.org/sergioaugrod/go-sptrans.svg?branch=master)](https://travis-ci.org/sergioaugrod/go-sptrans) [![GoDoc](https://godoc.org/github.com/sergioaugrod/go-sptrans?status.png)](https://godoc.org/github.com/sergioaugrod/go-sptrans) [![Go Report Card](https://goreportcard.com/badge/github.com/sergioaugrod/go-sptrans)](https://goreportcard.com/report/github.com/sergioaugrod/go-sptrans)


go-sptrans is a Go client library for the [SPTrans Olho Vivo API](http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo.aspx).

### Features

- [x] Routes 
- [ ] Stops
- [ ] Corridors 
- [ ] Companies
- [ ] Position of vehicles
- [ ] Arrival Forecast 
- [ ] Speed on lanes


## Install

go get github.com/sergioaugrod/go-sptrans/sptrans

## Usage

### Authentication

```go
token = "123456"
client = sptrans.NewClient(token)
client.Authenticate()
```

### Search Route by description or number

```go
client.Route.Search("Lapa")
client.Route.Search("8000")
```

### Search Route by direction

```go
client.Route.SearchByDirection("Lapa", 1)
client.Route.SearchByDirection("Lapa", 2)
```

## Contributing

1. Clone it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## License

[MIT License](LICENSE)
