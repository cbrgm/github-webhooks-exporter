package main

import (
	"context"
	"github.com/alecthomas/kong"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"runtime"
	"time"
)

const (
	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
)

var (
	// Version of github-webhooks-exporter.
	Version string
	// Revision or Commit this binary was built from.
	Revision string
	// GoVersion running this binary.
	GoVersion = runtime.Version()
	// StartTime has the time this was started.
	StartTime = time.Now()
)

var config struct {
	WebWebhookAddr      string `name:"http.webhook-addr" default:"0.0.0.0:9212" help:"The address the webhook receiver is running on"`
	WebExporterAddr     string `name:"http.exporter-addr" default:"0.0.0.0:9213" help:"The address the exporter is running on"`
	WebPath             string `name:"http.path" default:"/metrics" help:"The path metrics will be exposed at"`
	GithubWebhookSecret string `name:"github.webhook-secret" help:"The Github webhook secret"`
	LogJSON             bool   `name:"log.json" default:"false" help:"Tell the exporter to log json and not key value pairs"`
	LogLevel            string `name:"log.level" default:"info" enum:"error,warn,info,debug" help:"The log level to use for filtering logs"`
}

func main() {

	_ = kong.Parse(&config,
		kong.Name("github-webhooks-exporter"),
	)

	levelFilter := map[string]level.Option{
		levelError: level.AllowError(),
		levelWarn:  level.AllowWarn(),
		levelInfo:  level.AllowInfo(),
		levelDebug: level.AllowDebug(),
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	if config.LogJSON {
		logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	}

	logger = level.NewFilter(logger, levelFilter[config.LogLevel])
	logger = log.With(logger,
		"ts", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	registry := prometheus.NewRegistry()
	registry.MustRegister(collectors.NewGoCollector(), collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	var gr run.Group
	{
		r := mux.NewRouter()
		r.Use(HTTPMetrics(registry))

		// routes
		r.Handle("/hook", HandleFunc(metricsHandler(registry))).Methods("POST")

		s := http.Server{
			Addr:    config.WebWebhookAddr,
			Handler: r,
		}

		gr.Add(func() error {
			level.Info(logger).Log("msg", "webhook receiver started", "addr", config.WebWebhookAddr)
			return s.ListenAndServe()
		}, func(err error) {
			_ = s.Shutdown(context.TODO())
		})
	}
	{
		r := mux.NewRouter()

		r.Handle(config.WebPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`<html>
			<head><title>Github Webhooks Exporter</title></head>
			<body>
			<h1>Github Webhooks Exporter</h1>
			<p><a href="` + config.WebPath + `">see metrics</a></p>
			</body>
			</html>`))
		})

		s := http.Server{
			Addr:    config.WebExporterAddr,
			Handler: r,
		}

		gr.Add(func() error {
			level.Info(logger).Log("msg", "exporter started", "addr", config.WebExporterAddr)
			return s.ListenAndServe()
		}, func(err error) {
			_ = s.Shutdown(context.TODO())
		})

		if err := gr.Run(); err != nil {
			level.Error(logger).Log("msg", "http listenandserve error", "err", err)
			os.Exit(1)
		}
	}
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) (int, error)

func HandleFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode, err := h(w, r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
			return
		}
		if statusCode != http.StatusOK {
			w.WriteHeader(statusCode)
		}
	}
}

func HTTPMetrics(registry *prometheus.Registry) func(next http.Handler) http.Handler {
	duration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "latency of http requests",
	}, nil)

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "http requests total",
	}, []string{"code", "method"})

	registry.MustRegister(duration, counter)

	return func(next http.Handler) http.Handler {
		return promhttp.InstrumentHandlerDuration(duration, promhttp.InstrumentHandlerCounter(counter, next))
	}
}
