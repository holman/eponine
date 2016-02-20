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

func main() {
	http.HandleFunc("/", web)
	http.ListenAndServe(":24601", nil)
}
