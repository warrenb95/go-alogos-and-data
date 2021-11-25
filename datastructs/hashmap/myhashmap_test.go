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
		m    func() myHashMap
		args args
	}{
		{
			name: "should add a new value",
			m: func() myHashMap {
				return New()
			},
			args: args{
				key:   "test-key",
				value: "test-value",
			},
		},
		{
			name: "should overwrite an exsiting value",
			m: func() myHashMap {
				theMap := New()
				theMap.Set("test-key", "test-value")
				return theMap
			},
			args: args{
				key:   "test-key",
				value: "not what i set",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			theMap := tt.m()

			theMap.Set(tt.args.key, tt.args.value)
			value := theMap.Get(tt.args.key).(string)
			assert.Equal(t, tt.args.value, value)
		})
	}
}
