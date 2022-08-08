package chapter3

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// To Do: Dibagian ini dicontohkan menggunakan collection data test untuk menguji fungsi.
func TestGetEnvWithCollection(t *testing.T) {
	tests := []struct {
		Key         string
		SetEnvValue string
		DefValue    any
		Expect      any
	}{
		{
			Key:         "EXAMPLE_KEY1",
			SetEnvValue: "5000",
			DefValue:    3000,
			Expect:      5000,
		},
		{
			Key:         "EXAMPLE_KEY2",
			SetEnvValue: "Hello World",
			DefValue:    "Empty String",
			Expect:      "Hello World",
		},
		{
			Key:         "EXAMPLE_KEY3",
			SetEnvValue: "True",
			DefValue:    false,
			Expect:      true,
		},
	}

	for _, test := range tests {
		if test.SetEnvValue != "" {
			os.Setenv(test.Key, test.SetEnvValue)
		}

		switch test.DefValue.(type) {
		case int:
			result := GetEnv(test.Key, test.DefValue.(int))
			assert.Equal(t, test.Expect.(int), result)

		case bool:
			result := GetEnv(test.Key, test.DefValue.(bool))
			assert.Equal(t, test.Expect.(bool), result)

		default:
			result := GetEnv(test.Key, test.DefValue.(string))
			assert.Equal(t, test.Expect.(string), result)

		}
	}
}
