package gotk3

import (
	"context"
	"fmt"

	engine "github.com/Northern-Lights/os-rules-engine"
	"github.com/evilsocket/opensnitch/network"
	"github.com/evilsocket/opensnitch/rules"
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
	conn *network.Connection
	rule chan *rules.Rule
}{}

type promptCheckButton struct {
	id  string
	obj *gtk.CheckButton
}

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

// check buttons for target selection
var (
	cbProcess = promptCheckButton{id: "check_process"}
	cbPort    = promptCheckButton{id: "check_port"}
	cbIP      = promptCheckButton{id: "check_ip"}
	cbHost    = promptCheckButton{id: "check_host"}
	cbDomain  = promptCheckButton{id: "check_domain"}
)

// radio buttons for duration selection
var (
	rbQuit    = promptRadioButton{id: "radio_quit"}
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
func initPrompt() error {
	var err error
	promptBuilder, err = gtk.BuilderNewFromFile(promptUIXMLBuilder)
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

	var checkButtons = []*promptCheckButton{
		&cbProcess, &cbPort, &cbIP, &cbHost, &cbDomain,
	}
	for _, btn := range checkButtons {
		btn.obj = getCheckButton(promptBuilder, btn.id)
		if btn.obj == nil {
			return fmt.Errorf(`gotk3: couldn't get check button "%s"`, btn.id)
		}
	}

	var radioButtons = []*promptRadioButton{
		&rbQuit, &rbForever, &rbSession, &rbOnce,
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
func Prompt(ctx context.Context, conn *network.Connection) <-chan *rules.Rule {
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
	req.rule = make(chan *rules.Rule, 1)

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

	a := getAction(resp)
	d := getDuration()
	procCondition := getProcessConditionPart(d)
	cond := getCondition(procCondition)
	name := slug.Make(fmt.Sprintf("%s %s %v", a, d, cond)) // FIXME: need naming convention

	r := &rules.Rule{
		Name:      name,
		Action:    a,
		Duration:  d,
		Condition: cond,
	}

	req.rule <- r
}

// signal for window closed; that's an error, and caller should do default rule
func close() {
	connDeny()
}

// show displays the prompt to let the user allow or deny the connection
func show(conn *network.Connection) {
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

		cbPort.obj.SetLabel(fmt.Sprintf("Port %d", conn.DstPort))
		cbIP.obj.SetLabel(fmt.Sprintf("IP %s", conn.DstIp))
		cbHost.obj.SetLabel(fmt.Sprintf("Host %s", conn.DstHost))
		cbDomain.obj.SetLabel(fmt.Sprintf("Domain %s", conn.DstHost)) // TODO: *.x.com

		restoreButtonState()

		dialog.ShowAll()
	})
}

func restoreButtonState() {
	rbQuit.obj.SetActive(true)
	cbProcess.obj.SetActive(true)
	cbPort.obj.SetActive(true)
	cbIP.obj.SetActive(true)
	cbHost.obj.SetActive(false)
	cbDomain.obj.SetActive(false)
}

func getAction(resp gtk.ResponseType) rules.Action {
	switch resp {
	case gtk.RESPONSE_ACCEPT:
		return rules.Action_ALLOW
	case gtk.RESPONSE_REJECT:
		return rules.Action_DENY
	}

	panic(fmt.Errorf("Expected ACCEPT or REJECT; got %d", resp))
}

func getDuration() rules.Duration {
	switch {
	case rbOnce.obj.GetActive():
		return rules.Duration_ONCE
	case rbSession.obj.GetActive():
		return rules.Duration_FIREWALL_SESSION
	case rbQuit.obj.GetActive():
		return rules.Duration_PROCESS_SESSION
	case rbForever.obj.GetActive():
		return rules.Duration_ALWAYS
	}

	panic(fmt.Errorf("No duration selected"))
}

func getCondition(processCondition rules.EvaluatorSerializer) *rules.Expression {
	var eval rules.EvaluatorSerializer

	if cbProcess.obj.GetActive() {
		eval = processCondition
	}
	if cbPort.obj.GetActive() {
		port := engine.Port(req.conn.DstPort)
		if eval == nil {
			eval = port
		} else {
			eval = engine.And(eval, port)
		}
	}
	if cbIP.obj.GetActive() {
		ip := engine.IPAddr(req.conn.DstIp)
		if eval == nil {
			eval = ip
		} else {
			eval = engine.And(eval, ip)
		}
	}
	if cbHost.obj.GetActive() {
		host := engine.Host(req.conn.DstHost)
		if eval == nil {
			eval = host
		} else {
			eval = engine.And(eval, host)
		}
	}
	if cbDomain.obj.GetActive() {
		// TODO: differentiate w/ host
		domain := engine.Host(req.conn.DstHost)
		if eval == nil {
			eval = domain
		} else {
			eval = engine.And(eval, domain)
		}
	}

	// if there is still nothing, it's a "true"
	if eval == nil {
		eval = engine.Bool(true)
	}

	return eval.Serialize()
}

// duration tells us whether to use process path or PID
func getProcessConditionPart(d rules.Duration) (eval rules.EvaluatorSerializer) {
	if usePathInsteadOfPID(d) {
		eval = engine.ProcPath(req.conn.ProcessPath)
	} else {
		eval = engine.PID(req.conn.ProcessId)
	}
	return
}

func usePathInsteadOfPID(d rules.Duration) bool {
	return d == rules.Duration_ALWAYS || d == rules.Duration_FIREWALL_SESSION
}
