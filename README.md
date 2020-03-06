# benchmarks

this package is about benchmarking full-text indexation, go-template management to assess the best option for qorporess dependencies

## Requirements

### With docker (quicker)
* Docker
* Docker-compose 

### With local environment (more accurate)
* Go v1.13+
* Elastic v6.4.2
* Manticore v3.3.0
* Sphinx v3.0.3

## Install

```bash
mkdir $GOPATH/src/github.com/qorpress
cd $GOPATH/src/github.com/qorpress
git clone --depth=1 https://github.com/qorpress/benchmarks
cd benchmarks
```

## Full-Text search engines

We aim to benchmark the best search engines for doing faceted search and provide filtered results to qorpress front-end.

### Candidates
* Elastic v6.4.2
* Manticore v3.3.0
* Sphinx v3.0.3

#### Docker (recommended)
```bash
cd ./full-text
docker-compose up --build
```

#### Locally
```bash
cd ./full-text
go test -v
```

## GO-Template generators

The goal is to define the fastest and more features friendly template engine in golang.

### Candidates

#### full featured template engines
- [Ace](https://github.com/yosssi/ace)
- [Amber](https://github.com/eknkc/amber)
- [Go](https://golang.org/pkg/html/template)
- [Handlebars](https://github.com/aymerick/raymond)
- removed - [Kasia](https://github.com/ziutek/kasia.go)
- [Mustache](https://github.com/hoisie/mustache)
- [Pongo2](https://github.com/flosch/pongo2)
- [Soy](https://github.com/robfig/soy)
- [Jet](https://github.com/CloudyKit/jet)

#### precompilation to Go code

- [ego](https://github.com/benbjohnson/ego)
- removed - [egon](https://github.com/commondream/egon)
- [egonslinso](https://github.com/SlinSo/egon)
- [ftmpl](https://github.com/tkrajina/ftmpl)
- [Gorazor](https://github.com/sipin/gorazor)
- [Quicktemplate](https://github.com/valyala/quicktemplate)
- [Hero](https://github.com/shiyanhui/hero)
- [Jade](https://github.com/Joker/jade)
