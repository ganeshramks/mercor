module main-lambda

go 1.24.2

require (
	common-libs v0.0.0-00010101000000-000000000000
	github.com/beego/beego/v2 v2.3.7
	github.com/go-sql-driver/mysql v1.9.2
	github.com/spf13/cast v1.7.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)

replace common-libs => ../common-libs