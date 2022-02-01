package main

import (
	"fmt"
	"github.com/Reversaidx/system_monitoring/collector"
	"log"
	"net/http"
	"os"
	"runtime"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	myOS, myArch := runtime.GOOS, runtime.GOARCH
	inContainer := "inside"
	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		inContainer = "outside"
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Hello, %s!\n", r.UserAgent())
	_, _ = fmt.Fprintf(w, "I'm running on %s/%s.\n", myOS, myArch)
	_, _ = fmt.Fprintf(w, "I'm running %s of a container.\n", inContainer)
	_, _ = fmt.Fprintf(w, collector.collectCpuLinux())
}
func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("kurwa")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
