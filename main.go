package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

const (
	defaultPort   = "80"
	defaultGithub = "github.com"
)

type Config struct {
	Port         int    `default:80`
	GithubAPI    string `default:"github.com"`
	AccessToken  string `required:"true"`
	Organization string `required:"true"`
}

var logger = func(method, uri, name string, start time.Time) {
	log.Printf("method:%q uri:%q name:%q  time:%q", method, uri, name, time.Since(start))
}

func main() {

	var c Config
	err := envconfig.Process("gbot", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("starting github bot with ", Port, "...")
	router := httprouter.New()

	router.GET("/", Logging(Healthcheck, "healhcheck"))
	router.GET("/webhook", Logging(Webhook, "organization"))

	log.Fatal(http.ListenAndServe(Port, router))

}

func Logging(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		logger(r.Method, r.URL.Path, name, start)
	}
}
