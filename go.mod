module github.com/cdvelop/objectdb

go 1.20

require (
	github.com/cdvelop/input v0.0.60
	github.com/cdvelop/model v0.0.77
)

require (
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/timetools v0.0.26 // indirect
)

require (
	github.com/cdvelop/dbtools v0.0.67
	github.com/cdvelop/maps v0.0.7
	github.com/cdvelop/timeserver v0.0.25
	github.com/cdvelop/unixid v0.0.26
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/maps => ../maps

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/timeserver => ../timeserver

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/input => ../input
