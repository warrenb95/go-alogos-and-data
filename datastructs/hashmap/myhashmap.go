package hashmap

import (
	"bytes"
	"container/list"
	"encoding/gob"
)

type myHashMap []*list.List

type myNode struct {
	key   string
	value string
}

func New() myHashMap {
	return make(myHashMap, 100)
}

func (m myHashMap) Set(key, value string) {
	intKey, err := createKey(key)
	if err != nil {
		return
	}

	myKey := intKey % int64(len(m))

	node := myNode{
		key:   key,
		value: value,
	}

	if m[myKey] == nil {
		m[myKey] = list.New()
		m[myKey].Front().Value = node
	} else {

		front := m[myKey].Front()
		set := false

		for i := 0; i < m[myKey].Len(); i++ {
			element := front.Value.(*myNode)
			if element.key == key {
				element.value = value
				set = true
				break
			}
			front = front.Next()
		}
		if !set {
			m[myKey].InsertAfter(node, m[myKey].Back())
		}
	}
}

func (m myHashMap) Get(key string) interface{} {
	intKey, err := createKey(key)
	if err != nil {
		return nil
	}
	myKey := intKey % int64(len(m))

	myList := m[myKey]

	if myList.Len() == 0 {
		return nil
	} else if myList.Len() == 1 {
		return myList.Front().Value
	}

	front := myList.Front()

	for i := 0; i < myList.Len(); i++ {
		node := front.Value.(myNode)
		if node.key == key {
			return node.value
		}
	}

	return nil
}

func createKey(key interface{}) (int64, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return 0, err
	}
	byteKey := buf.Bytes()

	intKey := int64(0)

	for _, v := range byteKey {
		intKey += int64(v)
	}

	return intKey, nil
}
