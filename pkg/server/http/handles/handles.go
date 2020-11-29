package handles

import (
	"bytes"
	"fmt"
	"github.com/mchirico/zcovid/pkg/echarts/gauge"
	"github.com/mchirico/zcovid/pkg/echarts/heatmap"
	"github.com/mchirico/zcovid/pkg/echarts/line"
	"log"

	"net/http"
)

var Count = 0
var CountStatus = 0

type HANDLE struct {
	Process func() string
	ProcessGmail func(email, value string,ttl int64) string
	Token string
}

func (h HANDLE) BaseRoot(w http.ResponseWriter, r *http.Request) {

	/*
		curl -H "Authorization: SomeToken" localhost:3000

	      tasks/status: 2020-11-29 19:50:01.9865971 +0000 UTC m=+6.274003001


	       SomeToken

	*/

	reqToken := r.Header.Get("Authorization")

	switch r.Method {
	case "GET":
		Count += 1

		msg := h.Process()
		msg += "\n\n" + reqToken
		w.Write([]byte(msg))
	case "POST":
		// msg := fmt.Sprintf("Hello world: POST: %v", r.FormValue("user"))
		w.Write([]byte("post"))
	default:
		w.Write([]byte(`"Sorry, only GET and POST methods are supported."`))
	}

}

func (h HANDLE) Gmail(w http.ResponseWriter, r *http.Request) {

	/*

	    curl -H "Authorization: SomeToken" -H "Email: bozo@s" -H "Value: 3" localhost:3000


	*/

	reqToken := r.Header.Get("Authorization")
    email := r.Header.Get("Email")
    value := r.Header.Get("Value")
	ipaddress := r.Header.Get("X-FORWARDED-FOR")

	switch r.Method {
	case "GET":
		Count += 1

		if reqToken != h.Token {
			msg := fmt.Sprintf("Bad token:%v\n" +
				"h.Token: %v\n" +
				"ipaddress: %v\n",reqToken,h.Token,ipaddress)
			w.Write([]byte(msg))
			return
		}

        log.Printf("h.ProcessGmail: %v\n",email)
		msg := h.ProcessGmail(email,ipaddress, 1200)
		msg += "data\n"
		msg += "email:" + email + "\n"
		msg += "value:" + value + "\n"
		msg += "ipaddress:->" + ipaddress + "<-\n"
		msg += "\n\n" + reqToken
		msg += fmt.Sprintf("\n%v\n",r.Header.Get("X-FORWARDED-FOR"))
		log.Printf("msg: %v\n",msg)
		w.Write([]byte(msg))
	case "POST":
		// msg := fmt.Sprintf("Hello world: POST: %v", r.FormValue("user"))
		w.Write([]byte("post"))
	default:
		w.Write([]byte(`"Sorry, only GET and POST methods are supported."`))
	}

}



func Gauge(w http.ResponseWriter, r *http.Request) {

	g := gauge.GaugeExamples{}
	buf := bytes.NewBufferString("")
	g.Examples(buf)
	w.Write(buf.Bytes())

}

func Line(w http.ResponseWriter, r *http.Request) {
	e := line.LineExamples{}
	buf := bytes.NewBufferString("")
	e.Examples(buf)
	w.Write(buf.Bytes())

}

func Heatmap(w http.ResponseWriter, r *http.Request) {

	h := heatmap.HeatmapExamples{}
	buf := bytes.NewBufferString("")
	h.Examples(buf)
	w.Write(buf.Bytes())

}
