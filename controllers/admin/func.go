package admin

import (
	"encoding/json"
	"os"
)

type jsonContext struct {
	data map[string]interface{}
}

func Json(res bool) *jsonContext {
	c := new(jsonContext)
	c.data = make(map[string]interface{})
	c.data["res"] = res
	return c
}

func (jc *jsonContext) Set(key string, v interface{}) *jsonContext {
	jc.data[key] = v
	return jc
}

func (jc *jsonContext) End() {
	bytes, err := json.MarshalIndent(jc, "", "    ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(bytes)
}
