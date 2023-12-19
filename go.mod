module github.com/cdvelop/objectdb

go 1.20

require (
	github.com/cdvelop/dbtools v0.0.80
	github.com/cdvelop/input v0.0.77
	github.com/cdvelop/maps v0.0.8
	github.com/cdvelop/model v0.0.106
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/timeserver v0.0.31
	github.com/cdvelop/unixid v0.0.46
)

require github.com/cdvelop/timetools v0.0.32 // indirect

replace github.com/cdvelop/model => ../model
