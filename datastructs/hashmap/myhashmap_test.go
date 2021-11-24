package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myHashMap_Set(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		m    myHashMap
		args args
	}{
		{
			name: "should add a new value",
			m:    make(myHashMap, 100),
			args: args{
				key:   "test-key",
				value: "test-value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Set(tt.args.key, tt.args.value)
			value := tt.m.Get(tt.args.key)
			assert.Equal(t, tt.args.value, value)
		})
	}
}
