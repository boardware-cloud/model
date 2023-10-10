package argus

import (
	"testing"
)

func TestArgus(t *testing.T) {
	a := new(Argus)
	a.SetMonitor(&HttpMonitor{
		Type: "HTTP",
	})
	http, ok := a.Monitor().(*HttpMonitor)
	t.Log(http, ok)
	b := new(Argus)
	b.SetMonitor(&PingMonitor{
		Type: "PING",
	})
	ping, ok := b.Monitor().(*PingMonitor)
	t.Log(ping, ok)
}
