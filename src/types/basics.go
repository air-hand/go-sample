package types

//go:generate runtemplate -o keyvalue_list_gen.go -tpl threadsafe/list.tpl Type=KeyValue  Stringer:false  MapTo:string
type KeyValue struct {
	Key   string
	Value string
}
