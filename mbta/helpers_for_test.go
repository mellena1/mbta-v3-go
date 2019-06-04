package mbta

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var testDataPath = "testdata"

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func httpPathToTestData(path string) string {
	str := strings.Trim(path, "/")
	str = strings.ReplaceAll(str, "/", "_")
	str = str + ".json"
	return filepath.Join(testDataPath, str)
}

func handlerForServer(t *testing.T, path string) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		exp, _ := url.Parse(path)
		act := req.URL
		equals(t, exp.String(), act.String())

		fpath := httpPathToTestData(req.URL.Path)
		resp, err := ioutil.ReadFile(fpath)
		ok(t, err)
		rw.Write(resp)
	})
}

func strPtr(s string) *string {
	return &s
}
