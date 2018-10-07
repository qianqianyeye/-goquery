package types

type ParserResult struct {
	Item []interface{}
	Request []Request
}
type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}
