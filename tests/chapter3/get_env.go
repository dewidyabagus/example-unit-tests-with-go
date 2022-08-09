package chapter3

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Typ interface {
	int | string | bool
}

// To Do: Used to retrieve value from environment variable, if not found (empty) will use the default value.
//        When error convert, return with default value.
func GetEnv[T Typ](key string, def T) T {
	var res interface{} = def

	valEnv := strings.TrimSpace(os.Getenv(key))
	if valEnv == "" {
		return res.(T)
	}

	switch reflect.TypeOf(res).Kind() {
	case reflect.Int:
		val, err := strconv.Atoi(valEnv)
		if err == nil {
			res = val
		}

	case reflect.Bool:
		val, err := strconv.ParseBool(strings.ToLower(valEnv))
		if err == nil {
			res = val
		}

	default:
		res = valEnv

	}

	return res.(T)
}
