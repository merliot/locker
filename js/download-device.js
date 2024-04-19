const gpioDefaultCheckbox = document.getElementById('gpio-default');
const gpioSelector = document.querySelector('.gpio-selector');

gpioDefaultCheckbox.addEventListener('change', function() {
	gpioSelector.disabled = this.checked;
	if (this.checked) {
		gpioSelector.selectedIndex = 0;
	}
});

