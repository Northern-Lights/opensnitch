package ui

import (
	"context"
	"fmt"
	"net/url"

	"github.com/evilsocket/opensnitch/daemon/rule"
	"github.com/gosimple/slug"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gui "github.com/evilsocket/opensnitch/ui/gotk3"
	"github.com/evilsocket/opensnitch/ui/protocol"
)

// Service is the implementation of the UI service defined in the gRPC proto
type Service struct {
	url url.URL
}

// WithScheme returns a Service with the new scheme
// (e.g. http, https, unix, etc.)
func (s Service) WithScheme(scheme string) Service {
	s.url.Scheme = scheme
	return s
}

// WithHost returns the Service with the new host. If you have a port to
// specify, you can do that here
func (s Service) WithHost(host string) Service {
	s.url.Host = host
	return s
}

// WithPath returns the Service with the new path. This is can be useful for
// on-filesystem sockets using the unix scheme, for example
func (s Service) WithPath(path string) Service {
	s.url.Path = path
	return s
}

// AskRule implements the UI service's RPC, prompting the user for an action to
// be taken on a connection
func (s *Service) AskRule(ctx context.Context, conn *protocol.Connection) (resp *protocol.Rule, err error) {
	resp = makeDefaultRule(conn.DstIp)

	pctx, cancel := context.WithTimeout(ctx, Config.Timeout)
	defer cancel()
	recv := gui.Prompt(pctx, conn)
	if recv == nil {
		err = status.Errorf(codes.Unavailable, "ui: couldn't get response")
		return
	}

	select {
	case r := <-recv:
		if r != nil {
			resp = r
		} else {
			err = status.Errorf(codes.Internal, "ui: error getting response")
		}

	case <-ctx.Done():
		err = status.Errorf(codes.Canceled, "ui: client disconnected")
	}

	return
}

// Ping implements the UI service's RPC, saving statistics that can later be
// displayed in the stats window
func (s *Service) Ping(ctx context.Context, ping *protocol.PingRequest) (pong *protocol.PingReply, err error) {
	pong = &protocol.PingReply{
		Id: ping.Id,
	}

	return
}

// makeDefaultRule creates a rule using the configuration's setting to allow or
// deny the connection. That action is applied to the destination domain for the
// duration "until quit"
func makeDefaultRule(ip string) *protocol.Rule {
	var (
		a  = Config.Action
		d  = rule.Restart
		op = &protocol.Operator{
			Type:    string(rule.Simple),
			Operand: string(rule.OpDstIP),
			Data:    ip,
		}
		name = slug.Make(fmt.Sprintf("%s %s %s", a, op.Type, op.Data))
	)

	resp := &protocol.Rule{
		Name:     name,
		Action:   string(a),
		Duration: string(d),
		Operator: op,
	}

	return resp
}
