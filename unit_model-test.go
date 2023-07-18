package objectdb

type dataModelDBTest struct {
	Object         string
	Data           map[string]string
	ExpectedError  string
	IdRecovered    string
	SkipValidation bool
}
