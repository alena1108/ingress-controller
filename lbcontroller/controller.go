package lbcontroller

import (
	"fmt"
	"github.com/rancher/ingress-controller/lbconfig"
	"github.com/rancher/ingress-controller/lbprovider"
)

type LBController interface {
	GetName() string
	Run(lbProvider lbprovider.LBProvider)
	Stop() error
	GetLBConfigs() []*lbconfig.LoadBalancerConfig
	IsHealthy() bool
}

var (
	controllers map[string]LBController
)

func GetController(name string) LBController {
	if controller, ok := controllers[name]; ok {
		return controller
	}
	return controllers["kubernetes"]
}

func RegisterController(name string, controller LBController) error {
	if controllers == nil {
		controllers = make(map[string]LBController)
	}
	if _, exists := controllers[name]; exists {
		return fmt.Errorf("controller already registered")
	}
	controllers[name] = controller
	return nil
}
