package main

import (
	"encoding/json"
	"testing"
)

func TestSyntaxError(t *testing.T) {
	bd := []byte(data)
	var v struct {
		Timestamp    string `json:"timestamp"`
		ParentTotals struct {
			C int `json:"C"`
		} `json:"parent_totals"`
	}
	if err := json.Unmarshal(bd, &v); err != nil {
		//  "c": 52.57732",
		// "N":
		if err, ok := err.(*json.SyntaxError); ok {
			t.Log(string(bd[err.Offset-15 : err.Offset+15]))
			return
		}
		// 2009/11/10 23:00:00 invalid character '"' after object key:value pair
		t.Fatal(err)
	}
	t.Fail()
}

// "c": 52.57732", の部分に「"」が足りないためinvalidなJSON文字列。
const data = `
{
      "timestamp": "2019-02-24 07:07:03",
      "parent_totals": {
        "C": 0,
        "b": 0,
        "d": 0,
        "f": 9,
        "h": 102,
        "M": 0,
        "c": 52.57732",
        "N": 0,
        "p": 34,
        "m": 58,
        "diff": null,
        "s": 1,
        "n": 194
      }
}`
