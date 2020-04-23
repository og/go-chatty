package cha_test

import (
	cha "github.com/og/go-chatty"
	gis "github.com/og/x/test"
	"testing"
)

type User3 struct {
	ID string
	Son User3Son
}
func (user *User3) Chatty () {
	user.ID = cha.UUID()
}
type User3Son struct {
	ID string
	ID2 string
}
func (son *User3Son) Chatty () {
	son.ID = cha.NameIncrID("user3")
	son.ID2 = cha.NameIncrID("user3")
}


func TestSafeMock3(t *testing.T) {
	user := User3{}
	cha.Mock(&user)
	is := gis.New(t)
	is.Eql(len(user.ID), 36)
	is.Eql(user.Son.ID, "1")
	is.Eql(user.Son.ID2, "2")
}
