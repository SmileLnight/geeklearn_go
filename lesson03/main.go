package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	eg, ctx := errgroup.WithContext(context.Background())
	c := make(chan os.Signal, 1)
	eg.Go(func() (err error) {
		for {
			select {
			case <-ctx.Done():
				err = startServer(eg, &ctx)
				return
			case <-c:
				err = stopServer(eg, &ctx)
				return
			}
		}
	})

	err := eg.Wait()
	if err != nil {
		log.Printf("stop info: %v\n", err)
	}
	log.Println("server stop!")
}

func startServer(g *errgroup.Group, c *context.Context) error {

		nsm := http.NewServeMux()
		nsm.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, map[string]string{
				"msg": "test ok!",
			})
		})
		s := &http.Server{
			Addr:    ":8080",
			Handler: nsm,
		}
		g.Go(stopServer(s, c))
		return s.ListenAndServe()
}

func stopServer(s *http.Server, c *context.Context) error {
		stop := make(chan os.Signal)
		signal.Notify(stop)
		<-stop
		ctx, cancel := context.WithTimeout(*c, 10*time.Second)
		defer cancel()
		return s.Shutdown(ctx)
}