package Collection_test

import (
	"fmt"
	"testing"

	"github.com/ygto/collection"
	"github.com/ygto/value-object"
)

type TestStruct struct {
	username string
	age      int
}

func TestStringSetGet(t *testing.T) {
	username := "John Doe"
	collection := Collection.NewCollection()
	collection.Set("username", username)
	if username != collection.Get("username").String() {
		t.Fail()
	}
}

func TestIntSetGet(t *testing.T) {
	age := 50
	collection := Collection.NewCollection()
	collection.Set("age", age)
	if age != collection.Get("age").Int() {
		t.Fail()
	}
}
func TestStructSetGet(t *testing.T) {
	username := "John Doe"
	age := 50
	user := TestStruct{username: username, age: age}
	collection := Collection.NewCollection()
	collection.Set("user", user)
	if user != collection.Get("user").Get().(TestStruct) {
		t.Fail()
	}
}

func TestOrder(t *testing.T) {
	size := 100
	collection := Collection.NewCollection()

	for i := 0; i < size; i++ {
		key := fmt.Sprintf("key_%d", i)
		collection.Set(key, i)
	}
	for i, v := range collection.All() {
		if i != v.Int() {
			t.Fail()
		}
	}
}

func TestFilter(t *testing.T) {
	size := 100
	limit := 50
	collection := Collection.NewCollection()

	for i := 0; i < size; i++ {
		key := fmt.Sprintf("key_%d", i)
		collection.Set(key, i)
	}
	filter := func(obj *valobj.Valobj) bool {
		if obj.Int() > limit {
			return true
		}

		return false
	}
	for _, v := range collection.Filter(filter) {
		if v.Int() <= limit {
			t.Fail()
		}
	}
}
