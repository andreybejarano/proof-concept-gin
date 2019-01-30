# PROOF CONCEPT GOLANG WITH GIN-GONIC

> This proyect is a API-REST starter kit for GoLang with framework Gin-gonic

## Dependencies
- [GoLang v1.11.4 or later](https://golang.org/doc/) Development language
- [gin-gonic v1.3.0](https://github.com/gin-gonic/gin) API-REST framework
- [Realize v2.0.3](https://github.com/oxequa/realize) Develoment live reload for golang
- [Gorm v1.9.2](https://github.com/jinzhu/gorm) ORM Database

## Fisrt steps
> Install and config Golang [Golang Getting Started](https://golang.org/doc/install)
> Install and config Postgress

## Config workspace Golang
> Fist you need create golang workspace [Setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
> On `~/.bash_profile` or `~/bashrc` write in end file 
```shell
export GOPATH=~/gocode
export PATH="$PATH:$GOPATH/bin"
```

## Start
> Clone project on workspace Go lang
``` shell
$ git clone git@github.com:andreybejarano/proof-concept-gin.git
```
> Set config database on `.env` file

## Install
> Install dependencies
``` shell
$ go get
```

## Start development
``` shell
$ realize start
```

### Run migrations 
> For create migrations with schema on data base this tool create all tables by schema:
``` shell
$ sh migrations.sh
```

