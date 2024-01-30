module github.com/cdvelop/objectdb

go 1.20

require (
	github.com/cdvelop/dbtools v0.0.83
	github.com/cdvelop/input v0.0.88
	github.com/cdvelop/maps v0.0.8
	github.com/cdvelop/model v0.0.123
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/timeserver v0.0.36
	github.com/cdvelop/unixid v0.0.53
)

require github.com/cdvelop/timetools v0.0.42 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/dbtools => ../dbtools
