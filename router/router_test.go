package router

import (
	tb "github.com/tucnak/telebot"
	"testing"
)

func TestRouterAdd(t *testing.T) {
	a := New()
	myFunc := func(m *tb.Message) string { return "nice :)" }
	a.Add("/test", myFunc)

	if a.routes["/test"] == nil {
		t.Errorf("New command /test was not added to the routes correctly")
	}
}
