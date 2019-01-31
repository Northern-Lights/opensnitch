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

	service := ui.Service{}.WithScheme("unix").WithPath("/tmp/osui.sock")
	server := grpc.NewServer()
	ui.RegisterUIServer(server, &service)

	lis, err := net.Listen("unix", "/tmp/osui.sock")
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ch
		gui.Quit()
		server.GracefulStop()
		lis.Close()
	}()

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
