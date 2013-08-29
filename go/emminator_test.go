package testemminator

import (
	"emminator"
	"testing"
)

var (
	msglist = make([]string, 0)
)

func TestAll(t *testing.T) {
	m := emminator.NewEmitter()
	m.On("try", func() { msglist = append(msglist, "try1") })
	m.Once("try", func() { msglist = append(msglist, "try2") })
	m.Ready(func() { msglist = append(msglist, "ready1") })
	m.Emit("try")
	if len(msglist) < 1 || msglist[0] != "try1" {
		t.Error("on is not working")
	}
	if len(msglist[1]) < 2 || msglist[1] != "try2" {
		t.Error("once is not working")
	}
	m.Emit("try")
	if len(msglist) < 3 || msglist[2] != "try1" {
		t.Error("on is not working indefinitely")
	}
	if len(msglist) > 3 && msglist[3] == "try2" {
		t.Error("once is not working once")
	}
	if len(msglist) > 3 && msglist[3] == "ready1" {
		t.Error("ready works before readyState true")
	}
	m.Emit("ready")
	if len(msglist) < 4 || msglist[3] != "ready1" {
		t.Error("ready not triggered")
	}
	m.Ready(func() { msglist = append(msglist, "ready2") })
	if len(msglist) < 5 || msglist[4] != "ready2" {
		t.Error("ready not triggered immediately after readyState true")
	}
}
