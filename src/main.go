package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/ukko/cowsay/src/page"

	log "github.com/Sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	http.HandleFunc("/", handleFortune)

	log.Info("Start listen :", os.Getenv("HTTP_PORT"))

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handleFortune(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	cmd := exec.Command("/usr/games/fortune")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	log.Debug("handleFortune, out: ", out)

	tmpl := template.Must(template.ParseFiles("tmpl/layout.html"))

	data := page.New()
	data.PageContent = out.String()
	data.PageGenerated = time.Now().Sub(t).String()

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}

	log.Debug("handleFortune: render page")
}
