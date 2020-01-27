package gotest

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type testcase struct {
	Filename   string
	XFail      bool
	GoldenFile string
}

/* https://github.com/jstemmer/go-junit-report/blob/master/go-junit-report_test.go */

var testset = []testcase{
	{
		Filename:   "trx-sample-1.trx",
		XFail:      false,
		GoldenFile: "trx-sample-1.json",
	},
}

var update = flag.Bool("update", false, "update .golden files")

func TestParser(t *testing.T) {
	for _, test := range testset {
		if test.XFail {
			t.Logf("XFail: %s", test.Filename)
			continue
		}
		t.Logf("Running: %s", test.Filename)

		file, err := os.Open("testdata/" + test.Filename)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("Match structs", test.Filename)

		report, err := NewParser(file)
		if err != nil {
			t.Fatalf("Error parsing %s: %s", test.Filename, err)
		}

		goldenPath := filepath.Join("testdata", test.GoldenFile)
		if *update {
			t.Log("Update", goldenPath)
			data, _ := json.MarshalIndent(report, "", " ")
			if err := ioutil.WriteFile(goldenPath, data, 0644); err != nil {
				t.Fatalf("Failed to update golden file: %s", err)
			}
		}

		rawJson, err := ioutil.ReadFile(goldenPath)
		if err != nil {
			t.Fatal(err.Error())
		}
		var goldenReport JUnitReport
		json.Unmarshal(rawJson, &goldenReport)
		reflect.DeepEqual(*report, goldenReport)
	}
}