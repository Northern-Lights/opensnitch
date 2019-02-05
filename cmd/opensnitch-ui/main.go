package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/evilsocket/opensnitch/ui"
	gui "github.com/evilsocket/opensnitch/ui/gotk3"
	"google.golang.org/grpc"
)

func sigHandler() <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	return ch
}

func cleanupOnStop(sig <-chan os.Signal, l net.Listener, s *grpc.Server, quitGUI func() error) {
	defer l.Close()
	defer s.GracefulStop()
	defer quitGUI()
	<-sig
}

func createServer() *grpc.Server {
	var service ui.Service
	server := grpc.NewServer()
	ui.RegisterUIServer(server, &service)
	return server
}

func main() {
	err := gui.Init(
		"net.evilsocket.opensnitch",
		os.ExpandEnv("${GOPATH}/src/github.com/evilsocket/opensnitch/ui/gotk3/ui.xml"),
	)
	if err != nil {
		panic(err)
	}

	err = gui.Run()
	if err != nil {
		panic(err)
	}

	server := createServer()

	lis, err := net.Listen("unix", "/tmp/osui.sock")
	if err != nil {
		panic(err)
	}

	sig := sigHandler()
	go cleanupOnStop(sig, lis, server, gui.Quit)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
