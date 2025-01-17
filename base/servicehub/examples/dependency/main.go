// Author: recallsong
// Email: songruiguo@qq.com

package main

import (
	"fmt"
	"os"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/base/servicehub/examples/dependency/dependency"
)

// define Represents the definition of provider and provides some information
type define struct{}

// Declare what services the provider provides
func (d *define) Services() []string { return []string{"hello"} }

// Declare which services the provider depends on
func (d *define) Dependencies() []string { return []string{"example-dependency"} }

// Describe information about this provider
func (d *define) Description() string { return "hello for example" }

// Return an instance representing the configuration
func (d *define) Config() interface{} { return &config{} }

// Return a provider creator
func (d *define) Creator() servicehub.Creator {
	return func() servicehub.Provider {
		return &provider{}
	}
}

type config struct {
	Name string `file:"name" default:"recallsong"`
}

type provider struct {
	Cfg *config
	Log logs.Logger
}

func (p *provider) Init(ctx servicehub.Context) error {
	dep := ctx.Service("example-dependency").(dependency.Interface)
	fmt.Println(dep.Hello(p.Cfg.Name))
	return nil
}

func init() {
	servicehub.RegisterProvider("hello-provider", &define{})
}

func main() {
	hub := servicehub.New()
	hub.Run("examples", "", os.Args...)
}
