package types

//go:generate runtemplate -o keyvalue_list_gen.go -tpl threadsafe/list.tpl Type=KeyValue  Stringer:false  MapTo:string
//go:generate gofmt -w keyvalue_list_gen.go
type KeyValue struct {
	Key   string
	Value string
}
