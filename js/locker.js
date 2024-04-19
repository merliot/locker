import { WebSocketController, ViewMode } from './common.js'

export function run(prefix, url, viewMode) {
	const skeleton = new Skeleton(prefix, url, viewMode)
}

class Skeleton extends WebSocketController {

	open() {
		super.open()
		if (this.state.DeployParams === "") {
			return
		}
		this.setupBtn()
		this.showLed()
	}

	handle(msg) {
		switch(msg.Path) {
		case "click":
			this.saveClick(msg)
			break
		}
	}

	setupBtn() {
		if (this.viewMode === ViewMode.ViewFull) {
			this.btn = document.getElementById("led-btn")
			this.btn.onclick = () => { this.click() }
		}
	}

	updateBtn() {
		if (this.viewMode === ViewMode.ViewFull) {
			this.btn.textContent = this.state.Led.State ? "Turn Off LED" : "Turn On LED"
		}
	}

	showGpio() {
		if (this.viewMode === ViewMode.ViewFull) {
			let gpio = document.getElementById("gpio")
			if (this.state.Led.Gpio !== "") {
				gpio.classList.remove("hidden")
				gpio.textContent = this.state.Led.Gpio
			}
		}
	}

	showImg() {
		let img = document.getElementById("led-img")
		let onoff = this.state.Led.State ? "on" : "off"
		img.src = "images/led-" + onoff + ".png"
	}

	showLed() {
		this.showImg()
		this.showGpio()
		this.updateBtn()
	}

	saveClick(msg) {
		var led = this.state.Led
		led.State = msg.State
		this.showLed()
	}

	click() {
		if (!this.state.Online) {
			return
		}
		var led = this.state.Led
		led.State = !led.State
		this.webSocket.send(JSON.stringify({Path: "click", State: led.State}))
		this.showLed()
	}
}
