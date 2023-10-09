module github.com/cdvelop/objectdb

go 1.20

require (
	github.com/cdvelop/input v0.0.34
	github.com/cdvelop/model v0.0.50
)

require (
	github.com/cdvelop/dbtools v0.0.41
	github.com/cdvelop/gotools v0.0.42
	github.com/cdvelop/timeserver v0.0.1
	github.com/cdvelop/unixid v0.0.2
	golang.org/x/text v0.13.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timeserver => ../timeserver

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/input => ../input
