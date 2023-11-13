/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luyasr/hello-world/protocol"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "hello world API server",
	Long:  "hello world API server",
	Run: func(cmd *cobra.Command, args []string) {
		server := newServer()
		server.start()
	},
}

type server struct {
	http *protocol.HttpServer
	ch   chan os.Signal
}

func newServer() *server {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return &server{
		http: protocol.NewHttpServer(),
		ch:   sigs,
	}
}

func (s *server) start() {
	go func() {
		err := s.http.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	s.waitSigs(s.ch)
}

func (s *server) waitSigs(sigs chan os.Signal) {
	<-sigs
	log.Print("signal received, shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Print("server shutdown")
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
