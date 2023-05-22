package objectdb

type dataModelDBTest struct {
	Object         string
	Data           map[string]string
	Result         bool
	IdRecovered    string
	SkipValidation bool
}
