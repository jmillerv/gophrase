package handlers

import (
	"github.com/jmillerv/gophrase/config"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

const (
	baseHTML              = "assets/html/templates/base.layout.html"
	footerHTML            = "assets/html/templates/footer.partial.html"
	homeHTML              = "assets/html/templates/home.page.html"
	indexUrl              = "/"
	generatePassphraseUrl = "/generate"
	port                  = ":8080"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != indexUrl {
		http.NotFound(w, r)
		return
	}
	files := []string{
		homeHTML,
		baseHTML,
		footerHTML,
	}
	ts, err := template.ParseFS(config.Assets, files...)
	if err != nil {
		log.WithError(err).Error("unable to parse templates")
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func RunServer() {
	http.HandleFunc(indexUrl, home)
	http.HandleFunc(generatePassphraseUrl, NewPassphrase)
	log.WithField("port", port).Info("Listening on")
	log.Fatal(http.ListenAndServe(port, nil))
}
