// go run ./cmd
// go run -tags prime ./cmd
// tinygo flash -target xxx ./cmd

package main

import (
	"github.com/merliot/dean"
	"github.com/merliot/device/runner"
	"github.com/merliot/locker"
)

var (
	id           = dean.GetEnv("ID", "locker01")
	name         = dean.GetEnv("NAME", "Locker")
	deployParams = dean.GetEnv("DEPLOY_PARAMS", "")
	wsScheme     = dean.GetEnv("WS_SCHEME", "ws://")
	port         = dean.GetEnv("PORT", "8000")
	portPrime    = dean.GetEnv("PORT_PRIME", "8001")
	user         = dean.GetEnv("USER", "")
	passwd       = dean.GetEnv("PASSWD", "")
	dialURLs     = dean.GetEnv("DIAL_URLS", "")
	ssids        = dean.GetEnv("WIFI_SSIDS", "")
	passphrases  = dean.GetEnv("WIFI_PASSPHRASES", "")
)

func main() {
	l := locker.New(id, "locker", name).(*locker.Locker)
	l.SetDeployParams(deployParams)
	l.SetWifiAuth(ssids, passphrases)
	l.SetDialURLs(dialURLs)
	l.SetWsScheme(wsScheme)
	runner.Run(l, port, portPrime, user, passwd, dialURLs, wsScheme)
}
