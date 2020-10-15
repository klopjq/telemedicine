package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	logPath := "/tmp/request.log"
	httpPort := 5050
	setLogFile(logPath)

	http.HandleFunc("/", indexHandler)

	fmt.Printf("Listening on %v\n", httpPort)
	fmt.Printf("Logging to %v\n", logPath)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var code int
	var msg string
	switch r.Method {
	case "DELETE":
		code = http.StatusNoContent
	case "GET":
		code = http.StatusOK
		msg = fmt.Sprintf(`[{"subscriptionId":"%s"}]`, uuid.New())
	case "POST":
		code = http.StatusCreated
		msg = fmt.Sprintf(`{"subscriptionId":"%s"}`, uuid.New())
		if r.URL.Path == "/204/" {
			code = http.StatusNoContent
			msg = ""
		}
	default:
		code = http.StatusOK
		msg = fmt.Sprintf(`{"method":"%s"}`, r.Method)
	}

	w.WriteHeader(code)
	if len(msg) == 0 && code != http.StatusNoContent {
		msg = fmt.Sprintf(`{"code": %v, "message":"success"}`, code)
	}
	_, _ = w.Write([]byte(msg))
	log.Println(fmt.Sprintf("Dummy Response: %s", msg))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request []string
		// URL
		url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
		request = append(request, url)
		// Host
		request = append(request, fmt.Sprintf("Host: %v", r.Host))
		// RemoteAddr
		request = append(request, fmt.Sprintf("RemoteAddr: %v", r.RemoteAddr))
		// Headers
		for name, headers := range r.Header {
			for i := range headers {
				request = append(request, fmt.Sprintf("%v: %v", name, headers[i]))
			}
		}
		start := time.Now()
		handler.ServeHTTP(w, r)

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		request = append(request, string(data))
		request = append(request, fmt.Sprintf("RemoteAddr: %v", time.Since(start)))
		log.Printf("%s\n",
			strings.Join(request, "\n"),
		)
	})
}

func setLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}
		log.SetOutput(lf)
	}
}
