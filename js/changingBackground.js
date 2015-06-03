var backgrounds = [
	'images/wallpapers/729965-1440x900-[DesktopNexus.com].jpg',
	'images/wallpapers/813495-1440x900-[DesktopNexus.com].jpg',
	'images/wallpapers/821425-1440x900-[DesktopNexus.com].jpg',
	'images/wallpapers/1036809-1440x900-[DesktopNexus.com].jpg',
	'images/wallpapers/1062969-1440x900-[DesktopNexus.com].jpg'
];


function changeBackground(delta) {
	'use strict';
	window.requestAnimationFrame(changeBackground);
	var node = document.querySelector('html'),
			dateObject = new Date(),
			now = dateObject.getSeconds(),
			backgroundIndex = Math.floor(now / (60 / backgrounds.length));

	node.style.backgroundImage = 'url(' + backgrounds[backgroundIndex] + ')';
}

changeBackground(0);

var images = new Array();
function preload() {
	for (i = 0; i < 20; i++) {
		images[i] = new Image();
		images[i].src = 'images/8ball/' + (i+1) + '.png';
	}
}
preload();