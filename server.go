package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func execCommand(cmd string) string {
	cmdWithPath := fmt.Sprintf("./bin/%s", cmd)

	if cmd == "" {
		return "Eponine."
	}

	_, err := os.Stat(cmdWithPath)
	if err != nil {
		return "Ain't no command by that name here."
	}

	out, err := exec.Command(cmdWithPath).Output()

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func web(w http.ResponseWriter, r *http.Request) {
	file := strings.Trim(r.RequestURI, "/")
	output := execCommand(file)
	log.Println(r.RequestURI)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, output)
}

func images(w http.ResponseWriter, r *http.Request) {
	file := strings.TrimPrefix(r.RequestURI, "/images/")
	file = strings.Trim(file, "/")

	http.ServeFile(w, r, "data/images/"+file)
}

func main() {
	http.HandleFunc("/images/", images)
	http.HandleFunc("/", web)

	http.ListenAndServe(":24601", nil)
}
