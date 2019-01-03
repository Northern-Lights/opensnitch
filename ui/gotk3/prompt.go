package gotk3

import (
	"context"
	"fmt"

	"github.com/evilsocket/opensnitch/daemon/rule"

	"github.com/evilsocket/opensnitch/ui/protocol"
	"github.com/gosimple/slug"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const idDialog = "dlg_prompt"

var (
	promptBuilder *gtk.Builder
	dialog        *gtk.Dialog
)

// state singleton
var req = struct {
	ctx  context.Context
	conn *protocol.Connection
	rule chan *protocol.Rule
}{}

type promptLabel struct {
	id  string
	obj *gtk.Label
}

type promptRadioButton struct {
	id  string
	obj *gtk.RadioButton
}

// labels in the app info section
var (
	labelAppName = promptLabel{id: "app_name"}
	labelAppPath = promptLabel{id: "app_path"}
	labelAppInfo = promptLabel{id: "app_info_story"}
	// labelAppIcon = promptLabel{id: "app_icon"} // FIXME: this is a gtk.IconView
)

// labels in the connection info section
var (
	labelSrcIP = promptLabel{id: "label_src_ip"}
	labelDstIP = promptLabel{id: "label_dst_ip"}
	labelUID   = promptLabel{id: "label_uid"}
	labelPID   = promptLabel{id: "label_pid"}
)

// radio buttons for target selection
var (
	rbProcess  = promptRadioButton{id: "radio_process"}
	rbPort     = promptRadioButton{id: "radio_port"}
	rbDomainIP = promptRadioButton{id: "radio_domain_ip"}
)

// radio buttons for duration selection
var (
	rbForever = promptRadioButton{id: "radio_forever"}
	rbSession = promptRadioButton{id: "radio_session"}
	rbOnce    = promptRadioButton{id: "radio_once"}
)

var signals = map[string]interface{}{
	"conn_allow":   connAllow,
	"conn_deny":    connDeny,
	"dlg_close":    close,
	"dlg_resp":     dlgResponse,
	"delete-event": func() bool { fmt.Println("delete-event: why are we doing this?"); return true },
}

var lock = make(chan struct{}, 1)

// initPrompt initializes the prompt using the specified UI builder file
func initPrompt(uiFilePath string) error {
	var err error
	promptBuilder, err = gtk.BuilderNewFromFile(uiFilePath)
	if err != nil {
		return err
	}

	dialog = getDialog(promptBuilder, idDialog)
	if dialog == nil {
		return fmt.Errorf(`gotk3: couldn't get dialog "%s"`, idDialog)
	}

	var labels = []*promptLabel{
		&labelAppName, &labelAppPath, &labelAppInfo, // &labelAppIcon, // FIXME:
		&labelSrcIP, &labelDstIP, &labelUID, &labelPID,
	}
	for _, label := range labels {
		label.obj = getLabel(promptBuilder, label.id)
		if label.obj == nil {
			return fmt.Errorf(`gotk3: couldn't get label "%s"`, label.id)
		}
	}

	var radioButtons = []*promptRadioButton{
		&rbProcess, &rbPort, &rbDomainIP,
		&rbForever, &rbSession, &rbOnce,
	}
	for _, btn := range radioButtons {
		btn.obj = getRadioButton(promptBuilder, btn.id)
		if btn.obj == nil {
			return fmt.Errorf(`gotk3: couldn't get radio button "%s"`, btn.id)
		}
	}

	promptBuilder.ConnectSignals(signals)

	// enable
	lock <- struct{}{}

	return nil
}

// Prompt shows the dialog to ask the user what to do with the connection
func Prompt(ctx context.Context, conn *protocol.Connection) <-chan *protocol.Rule {
	select {
	case <-lock:
		// acquired; will proceed
	case <-ctx.Done():
		// we are no longer needed
		return nil
	}

	// set the state
	req.ctx = ctx
	req.conn = conn
	req.rule = make(chan *protocol.Rule, 1)

	show(conn)

	return req.rule
}

func connAllow() {
	dialog.Response(gtk.RESPONSE_ACCEPT)
	dialog.Hide()
}

func connDeny() {
	dialog.Response(gtk.RESPONSE_REJECT)
	dialog.Hide()
}

// this is where the final rule is made and sent back over the channel
func dlgResponse(dlg *gtk.Dialog, resp gtk.ResponseType) {
	// dlg is closed if we are here; we may unlock
	lock <- struct{}{}

	var (
		a    = getAction(resp)
		d    = getDuration()
		o    = getOperator()
		name = slug.Make(fmt.Sprintf("%s %s %s", a, o.Type, o.Data))
	)

	r := &protocol.Rule{
		Name:     name,
		Action:   string(a),
		Duration: string(d),
		Operator: o,
	}

	req.rule <- r
}

// signal for window closed; that's an error, and caller should do default rule
func close() {
	connDeny()
}

// show displays the prompt to let the user allow or deny the connection
func show(conn *protocol.Connection) {
	// See: https://github.com/gotk3/gotk3-examples/blob/master/gtk-examples/goroutines/goroutines.go
	glib.IdleAdd(func() {
		labelAppName.obj.SetText(
			fmt.Sprintf("%s (%s)", conn.ProcessArgs[0], conn.Protocol))
		labelAppPath.obj.SetText(conn.ProcessPath)
		labelAppInfo.obj.SetText(
			fmt.Sprintf("%s wants to connect to %s on %s port %d",
				conn.ProcessPath, conn.DstHost, conn.Protocol, conn.DstPort))
		labelSrcIP.obj.SetText(
			fmt.Sprintf("%s:%d", conn.SrcIp, conn.SrcPort))
		labelDstIP.obj.SetText(
			fmt.Sprintf("%s (%s:%d)", conn.DstHost, conn.DstIp, conn.DstPort))
		labelUID.obj.SetText(
			fmt.Sprintf("%d", conn.UserId))
		labelPID.obj.SetText(
			fmt.Sprintf("%d", conn.ProcessId))
		dialog.ShowAll()
	})
}

func getAction(resp gtk.ResponseType) rule.Action {
	switch resp {
	case gtk.RESPONSE_ACCEPT:
		return rule.Allow
	case gtk.RESPONSE_REJECT:
		return rule.Deny
	}

	panic(fmt.Errorf("Expected ACCEPT or REJECT; got %d", resp))
}

func getDuration() rule.Duration {
	switch {
	case rbOnce.obj.GetActive():
		return rule.Once
	case rbSession.obj.GetActive():
		return rule.Restart
	case rbForever.obj.GetActive():
		return rule.Always
	}

	panic(fmt.Errorf("No duration selected"))
}

func getOperator() *protocol.Operator {
	switch {
	case rbProcess.obj.GetActive():
		return &protocol.Operator{
			Type:    string(rule.Simple),
			Operand: string(rule.OpProcessPath),
			Data:    req.conn.ProcessPath,
		}
	case rbPort.obj.GetActive():
		return &protocol.Operator{
			Type:    string(rule.Simple),
			Operand: string(rule.OpDstPort),
			Data:    fmt.Sprintf("%d", req.conn.DstPort),
		}
	case rbDomainIP.obj.GetActive():
		return &protocol.Operator{
			Type:    string(rule.Simple),
			Operand: string(rule.OpDstHost),
			Data:    req.conn.DstHost,
		}
		// TODO: differentiate between IP and domain
		// TODO: one for IP + port
	}

	panic(fmt.Errorf("No operator selected"))
}
