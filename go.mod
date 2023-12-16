module github.com/cdvelop/objectdb

go 1.20

require (
	github.com/cdvelop/dbtools v0.0.76
	github.com/cdvelop/input v0.0.73
	github.com/cdvelop/maps v0.0.8
	github.com/cdvelop/model v0.0.102
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/timeserver v0.0.31
	github.com/cdvelop/unixid v0.0.39
)

require github.com/cdvelop/timetools v0.0.32 // indirect

replace github.com/cdvelop/model => ../model
