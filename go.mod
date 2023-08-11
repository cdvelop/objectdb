module github.com/cdvelop/objectdb

go 1.20

require (
	github.com/cdvelop/input v0.0.15
	github.com/cdvelop/model v0.0.36
)

require (
	github.com/cdvelop/dbtools v0.0.25
	github.com/cdvelop/gotools v0.0.26
	golang.org/x/text v0.11.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/input => ../input
