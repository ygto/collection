package Collection

import "github.com/ygto/value-object"

type collection struct {
	keys []string
	data map[string]interface{}
}

func NewCollection() *collection {
	col := collection{}
	col.data = make(map[string]interface{})
	col.keys = make([]string, 0, 0)
	return &col
}

func (col *collection) Has(key string) bool {
	_, has := col.data[key]
	return has
}

func (col *collection) Get(key string) *valobj.Valobj {

	return valobj.Val(col.data[key])
}

func (col *collection) Set(key string, value interface{}) {
	col.keys = append(col.keys, key)
	col.data[key] = value
}

func (col *collection) Keys() []string {
	return col.keys
}

func (col *collection) All() []*valobj.Valobj {
	return col.Filter(func(i *valobj.Valobj) bool {
		return true
	})
}

func (col *collection) Filter(f func(obj *valobj.Valobj) bool) []*valobj.Valobj {
	arr := make([]*valobj.Valobj, 0, 0)
	for _, v := range col.Keys() {
		obj := col.Get(v)
		if f(obj) {
			arr = append(arr, obj)
		}
	}
	return arr
}
