package locker

import (
	"fmt"
	"net/http"

	"github.com/merliot/dean"
	"github.com/merliot/device"
	"github.com/merliot/device/led"
)

var targets = []string{"demo", "nano-rp2040", "wioterminal"}

type Locker struct {
	*device.Device
	Led led.Led
}

type MsgClick struct {
	dean.ThingMsg
	State bool
}

func New(id, model, name string) dean.Thinger {
	fmt.Println("NEW LOCKER\r")
	return &Locker{
		Device: device.New(id, model, name, fs, targets).(*device.Device),
	}
}

func (l *Locker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.API(w, r, l)
}

func (l *Locker) save(msg *dean.Msg) {
	msg.Unmarshal(l).Broadcast()
}

func (l *Locker) getState(msg *dean.Msg) {
	l.Path = "state"
	msg.Marshal(l).Reply()
}

func (l *Locker) click(msg *dean.Msg) {
	msg.Unmarshal(&l.Led)
	if l.IsMetal() {
		l.Led.Set(l.Led.State)
	}
	msg.Broadcast()
}

func (l *Locker) Subscribers() dean.Subscribers {
	return dean.Subscribers{
		"state":     l.save,
		"get/state": l.getState,
		"click":     l.click,
	}
}

func (l *Locker) parseParams() {
	l.Led.Gpio = l.ParamFirstValue("gpio")
}

func (l *Locker) configure() {
	l.Led.Configure()
}

func (l *Locker) Setup() {
	l.Device.Setup()
	l.parseParams()
	l.configure()
}

func (l *Locker) Run(i *dean.Injector) {
	select {}
}
