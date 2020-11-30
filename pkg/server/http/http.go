package http

import (
	"github.com/mchirico/tasks/pkg/server/http/handles"
	"github.com/mchirico/tasks/pkg/utils"
	"log"
	"net/http"
	"os"
	_ "time/tzdata"
)

func SetupHandles() {

	h := handles.HANDLE{}
	u := utils.NewUT()
	u2 := utils.NewUT()
	h.Process = u.Status
	h.ProcessGmail = u2.Email
	h.Token = os.Getenv("TASK_KEY")
	if h.Token == "" {
		h.Token = "aslskdjsaaa_NOALL"
	}

	http.HandleFunc("/", h.BaseRoot)
	http.HandleFunc("/gmail", h.Gmail)
	http.HandleFunc("/gauge", handles.Gauge)
	http.HandleFunc("/line", handles.Line)
	http.HandleFunc("/heatmap", handles.Heatmap)

}

func Server() {
	SetupHandles()
	log.Println("starting tasks server... :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
