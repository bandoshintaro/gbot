package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port         string `default:"8080"`
	GithubAPI    string
	AccessToken  string `required:"true"`
	Organization string `required:"true"`
}

var (
	config Config
	logger = func(method, uri, name string, start time.Time) {
		log.Printf("method:%q uri:%q name:%q  time:%q", method, uri, name, time.Since(start))
	}
)

func main() {

	err := envconfig.Process("gbot", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	port := ":" + config.Port

	log.Println("starting github bot with ", config.Port, "...")
	router := httprouter.New()

	router.GET("/", Logging(Healthcheck, "healhcheck"))
	router.POST("/webhook", Logging(Webhook, "organization"))

	log.Fatal(http.ListenAndServe(port, router))

}

func Logging(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		logger(r.Method, r.URL.Path, name, start)
	}
}
