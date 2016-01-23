package run

import (
	"github.com/halverneus/sample/router"
	"github.com/halverneus/sample/settings"
)

// Service running function.
func Service(ignored ...string) (err error) {
	return router.Serve(settings.Get.BindAddress)
}
