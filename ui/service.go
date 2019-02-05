package ui

import (
	"context"
	"fmt"

	engine "github.com/Northern-Lights/os-rules-engine"
	"github.com/evilsocket/opensnitch/network"
	"github.com/evilsocket/opensnitch/rules"
	gui "github.com/evilsocket/opensnitch/ui/gotk3"
	"github.com/gosimple/slug"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service is the implementation of the UI service defined in the gRPC proto
type Service struct{}

// AskRule implements the UI service's RPC, prompting the user for an action to
// be taken on a connection
func (s *Service) AskRule(ctx context.Context, conn *network.Connection) (resp *rules.Rule, err error) {
	resp = makeDefaultRule(conn.DstIp, conn.DstPort)

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
func (s *Service) Ping(ctx context.Context, ping *PingRequest) (pong *PingReply, err error) {
	pong = &PingReply{
		Id: ping.Id,
	}

	return
}

// makeDefaultRule creates a rule using the configuration's setting to allow or
// deny the connection
func makeDefaultRule(ip string, port uint32) *rules.Rule {
	var (
		a    = Config.Action
		d    = Config.Duration
		eval = engine.And(
			engine.IPAddr(ip),
			engine.Port(port),
		)
		name = slug.Make(fmt.Sprintf("%s %s %d", a, ip, port)) // FIXME: need naming convention
	)

	resp := &rules.Rule{
		Name:      name,
		Action:    a,
		Duration:  d,
		Condition: eval.Serialize(),
	}

	return resp
}
