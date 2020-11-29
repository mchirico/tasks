package handles

import (
	"fmt"
	"github.com/mchirico/tasks/pkg/utils"

	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_RootGET(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	h := HANDLE{}
	e := utils.NewUT()
	h.Process = e.Status

	h.BaseRoot(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if !strings.Contains(string(body), "tasks/status") {
		t.Fatalf("GET on root failed: %s\n",body)
	}

}

func Test_RootPUT(t *testing.T) {
	req := httptest.NewRequest("PUT", "/", nil)
	w := httptest.NewRecorder()

	h := HANDLE{}
	e := utils.NewUT()
	h.Process = e.Status

	h.BaseRoot(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		t.Fatalf("PUT is causing error")
	}

}

func Test_RootPOST(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	h := HANDLE{}
	e := utils.NewUT()
	h.Process = e.Status

	h.BaseRoot(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if string(body) != "post" {
		t.Log(string(body))
		t.Fatalf("post response not what expected")
	}

}
