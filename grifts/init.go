package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/mclark4386/personal_site/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
