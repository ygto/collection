package Collection

import "github.com/ygto/value-object"

type collection struct {
	keys []string
	data map[string]interface{}
}

func NewCollection() *collection {
	sm := collection{}
	sm.data = make(map[string]interface{})
	sm.keys = make([]string, 0, 0)
	return &sm
}

func (sm *collection) Set(key string, value interface{}) {
	sm.keys = append(sm.keys, key)
	sm.data[key] = value
}

func (sm *collection) Get(key string) *valobj.Valobj {

	return valobj.Val(sm.data[key])
}

func (sm *collection) Keys() []string {
	return sm.keys
}

func (sm *collection) All() []*valobj.Valobj {
	arr := make([]*valobj.Valobj, 0, 0)
	for _, v := range sm.Keys() {
		_ = v
		arr = append(arr, sm.Get(v))
	}
	return arr
}




