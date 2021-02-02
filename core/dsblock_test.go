package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_DeserializeFromJsonToDsBlockT(t *testing.T) {
	dsJson,err := ioutil.ReadFile("dsblock.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	var dsBlockT DsBlockT
	err2 := json.Unmarshal(dsJson,&dsBlockT)
	if err2 != nil {
		t.Fatal(err2.Error())
	}

	fmt.Println(dsBlockT)
}