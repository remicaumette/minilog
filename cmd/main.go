package main

import (
	"context"
	"github.com/remicaumette/minilog/internal/minilog"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := minilog.New()

	go func() {
		if err := server.Start(); err != nil {
			logrus.WithError(err).Fatal("failed to start server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	logrus.Info("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("could not gracefully shutdown")
	}

		//file, err := os.Open("apache_logs")
		//if err != nil {
		//	logrus.WithError(err).Fatal("failed to open file")
		//}
		//sc := bufio.NewScanner(file)
		//dateRegex := regexp.MustCompile(`\[([^]]+)\]\s`)
		//dateLayout := "02/Jan/2006:15:04:05 -0700"
		//for sc.Scan() {
		//	line := sc.Text()
		//	rawDate := dateRegex.FindStringSubmatch(line)[1]
		//	t, err := time.Parse(dateLayout, rawDate)
		//	if err != nil {
		//		logrus.WithError(err).Fatal("failed to parse date")
		//	}
		//	if err := storeLine(db, t, line); err != nil {
		//		logrus.WithError(err).Fatal("failed to store line")
		//	}
		//}
}

