package core

type DnsLookupError struct {
	IsDnsLookupError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewDnsLookupError(code string, msg string, ctx *Context) *DnsLookupError {
	return &DnsLookupError{
		IsDnsLookupError: true,
		Sdk:              "DnsLookup",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *DnsLookupError) Error() string {
	return e.Msg
}
