<h1 align="center">
<strong>Gocambridge</strong>
</h1>

<h3 align="center">
    A simple and nice Cambridge One solver CLI written in Go
</h3>

<p align="center">
    <img alt="Gopher" height="150" src="https://raw.githubusercontent.com/egonelbre/gophers/master/vector/science/soldering.svg">
</p>

## Requirements

- Go >= 1.21
- Cambridge One Cookie

## Installation

1. Clone the repository:

```
git clone https://github.com/frandier/animebible.git
```
```
cd cambridge-go
```

2. Install the dependencies:

```
go get
```
## Usage

To run the CLI:

```
go run main.go -c <cambridge_cookie> -u <cambridge_unit_url>
```

Help:

```
go run main.go -h
```

## Example

```
go run main.go -c s%3AdPnfqAw8N... -u https://www.cambridgeone.org/nlp/apigateway/your_org/class/7e1e853c../product/ic...
```
## Contributing

If you'd like to contribute to this project, please fork the repository and create a pull request. 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

This project was built using the following technologies:

- [Go](https://go.dev/)
- [Cambridge One](https://www.cambridgeone.org/)

## Contact

If you have any questions or comments about this project, please feel free to reach out to me at @frandier.