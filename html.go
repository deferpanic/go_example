package main

var HtmlTemplate string = `<html>
<style>
@import url(https://fonts.googleapis.com/css?family=Roboto:700);
html,
body {
width: 100%;
height: 100%;
background-color: #171723;
overflow: hidden;
}

h1 {
position: absolute;
top: 40%;
left: 50%;
-webkit-transform: translate(-50%, -50%);
transform: translate(-50%, -50%);
color: rgba(255, 255, 255, 0.12);
font-size: 80px;
-webkit-animation: pulse-text 5s ease-in-out 5s;
animation: pulse-text 5s ease-in-out 5s;
font-family: 'Roboto', sans-serif;
}

.sound-pulses {
position: absolute;
width: 800px;
left: 0%;
margin-left: -400px;
}

.pulse {
width: 400px;
height: 400px;
background-image: url(data:image/gif;base64,{{.Image}});
border: 1px solid rgba(255, 255, 255, 0.12);
border-radius: 50%;
margin-left: 50%;
left: -200px;
position: absolute;
bottom: -160px;
-webkit-transform: scale(0.1, 0.1);
transform: scale(0.1, 0.1);
opacity: 0;
-webkit-animation: pulse 5000ms ease-out 10s;
animation: pulse 5000ms ease-out 10s;
}

.pulse:nth-of-type(2) {
-webkit-animation-delay: 600ms;
animation-delay: 600ms;
}

.pulse:nth-of-type(3) {
-webkit-animation-delay: 1200ms;
animation-delay: 1200ms;
}

@-webkit-keyframes pulse {
0% {
-webkit-transform: scale(0.1, 0.1);
transform: scale(0.1, 0.1);
opacity: 0;
}
50% {
opacity: 1;
}
100% {
-webkit-transform: scale(1.2, 1.2);
transform: scale(1.2, 1.2);
opacity: 0;
}
}

@keyframes pulse {
0% {
-webkit-transform: scale(0.1, 0.1);
transform: scale(0.1, 0.1);
opacity: 0;
}
50% {
opacity: 1;
}
100% {
-webkit-transform: scale(1.2, 1.2);
transform: scale(1.2, 1.2);
opacity: 0;
}
}
@-webkit-keyframes pulse-text {
0% {
opacity: 1;
}
50% {
opacity: 0;
}
100% {
opacity: 1;
}
}
@keyframes pulse-text {
0% {
opacity: 1;
}
50% {
opacity: 0;
}
100% {
opacity: 1;
}
}

</style>
<script>
window.onload = function(e) {
	for (i=0; i<10; i++) {
		var rtop = Math.floor(Math.random() * 600) + 100;
		var rleft = Math.floor(Math.random() * 900) + 100;
		var e = document.createElement('div');
		e.innerHTML = "<div class='sound-pulses' style='top: " 
		+ rtop + "px; left: " + rleft + "px'>" +
		"<div class='pulse'></div><div class='pulse'></div>" +
		"<div class='pulse'></div>";

		while(e.firstChild) {
			var el = document.getElementById('in');
			el.appendChild(e.firstChild);
		}
	}
};
</script>

<div id='in'>
</div>

<div class="sound-pulses">
<div class="pulse"></div>
<div class="pulse"></div>
<div class="pulse"></div>
</div>

<h1>GOLANG UNIKERNELS</h1>

</div>
</html>`
