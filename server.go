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

	return fmt.Sprintf("something: %s\n", out)
}

func web(w http.ResponseWriter, r *http.Request) {
	file := strings.Trim(r.RequestURI, "/")
	output := execCommand(file)

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
