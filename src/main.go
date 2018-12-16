package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	log "github.com/Sirupsen/logrus"
)

func init()  {
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

type Page struct {
	PageTitle string
	PageAuthor string
	PageDescription string
	PageContent string
}

func handleFortune(w http.ResponseWriter, r *http.Request)  {
	cmd := exec.Command("/usr/games/fortune")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	
	log.Debug("handleFortune, out: ", out)

	tmpl := template.Must(template.ParseFiles("tmpl/layout.html"))

	data := Page{
		PageTitle: "Fortune",
		PageAuthor: "Max Kamashev",
		PageDescription: "Cow say fortunes",
		PageContent: out.String(),
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}

	log.Debug("handleFortune: render page")
}
