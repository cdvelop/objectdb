module github.com/cdvelop/objectdb

go 1.20

require (
	github.com/cdvelop/input v0.0.34
	github.com/cdvelop/model v0.0.73
)

require (
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/timetools v0.0.19 // indirect
)

require (
	github.com/cdvelop/dbtools v0.0.41
	github.com/cdvelop/maps v0.0.6
	github.com/cdvelop/timeserver v0.0.1
	github.com/cdvelop/unixid v0.0.2
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/maps => ../maps

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/timeserver => ../timeserver

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/input => ../input
