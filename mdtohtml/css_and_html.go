package mdtohtml

var HtmlPrefix = `
  <html>
    <head>
	  <meta charset="utf-8">
	<style>
	
	
html,
body,
div,
span,
applet,
object,
iframe,
h1,
h2,
h3,
h4,
h5,
h6,
p,
blockquote,
pre,
a,
abbr,
acronym,
address,
big,
cite,
code,
del,
dfn,
em,
img,
ins,
kbd,
q,
s,
samp,
small,
strike,
strong,
sub,
sup,
tt,
var,
b,
u,
i,
center,
dl,
dt,
dd,
ol,
ul,
li,
fieldset,
form,
label,
legend,
table,
caption,
tbody,
tfoot,
thead,
tr,
th,
td,
article,
aside,
canvas,
details,
embed,
figure,
figcaption,
footer,
header,
hgroup,
menu,
nav,
output,
ruby,
section,
summary,
time,
mark,
audio,
video {
	margin: 0;
	padding: 0;
	border: 0;
	font: inherit;
	vertical-align: baseline
}

article,
aside,
details,
figcaption,
figure,
footer,
header,
hgroup,
menu,
nav,
section {
	display: block
}

ol,
ul {
	list-style: none
}

table {
	border-collapse: collapse;
	border-spacing: 0
}

body {
	box-sizing: border-box;
	color: #373737;
	background: #212121;
	font-size: 16px;
	font-family: 'Myriad Pro', Calibri, Helvetica, Arial, sans-serif;
	line-height: 1.5;
	-webkit-font-smoothing: antialiased
}

h1,
h2,
h3,
h4,
h5,
h6 {
	margin: 10px 0;
	font-weight: 700;
	color: #222222;
	font-family: 'Lucida Grande', 'Calibri', Helvetica, Arial, sans-serif;
	letter-spacing: -1px
}

h1 {
	font-size: 36px;
	font-weight: 700
}

h2 {
	padding-bottom: 10px;
	font-size: 32px;
	background: url("../images/bg_hr.png") repeat-x bottom
}

h3 {
	font-size: 24px
}

h4 {
	font-size: 21px
}

h5 {
	font-size: 18px
}

h6 {
	font-size: 16px
}

p {
	margin: 10px 0 15px 0
}

footer p {
	color: #f2f2f2
}

a {
	text-decoration: none;
	color: #0F79D0;
	text-shadow: none;
	transition: color 0.5s ease;
	transition: text-shadow 0.5s ease;
	-webkit-transition: color 0.5s ease;
	-webkit-transition: text-shadow 0.5s ease;
	-moz-transition: color 0.5s ease;
	-moz-transition: text-shadow 0.5s ease;
	-o-transition: color 0.5s ease;
	-o-transition: text-shadow 0.5s ease;
	-ms-transition: color 0.5s ease;
	-ms-transition: text-shadow 0.5s ease
}

a:hover,
a:focus {
	text-decoration: underline
}

footer a {
	color: #F2F2F2;
	text-decoration: underline
}

em,
cite {
	font-style: italic
}

strong {
	font-weight: bold
}

img {
	position: relative;
	margin: 0 auto;
	max-width: 739px;
	padding: 5px;
	margin: 10px 0 10px 0;
	border: 1px solid #ebebeb;
	box-shadow: 0 0 5px #ebebeb;
	-webkit-box-shadow: 0 0 5px #ebebeb;
	-moz-box-shadow: 0 0 5px #ebebeb;
	-o-box-shadow: 0 0 5px #ebebeb;
	-ms-box-shadow: 0 0 5px #ebebeb
}

p img {
	display: inline;
	margin: 0;
	padding: 0;
	vertical-align: middle;
	text-align: center;
	border: none
}

pre,
code {
	color: #222;
	background-color: #fff;
	font-family: Monaco, "Bitstream Vera Sans Mono", "Lucida Console", Terminal, monospace;
	font-size: 14px;
	border-radius: 2px;
	-moz-border-radius: 2px;
	-webkit-border-radius: 2px
}

pre {
	padding: 10px;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
	overflow: auto
}

code {
	padding: 3px;
	margin: 0 3px;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.1)
}

pre code {
	display: block;
	box-shadow: none
}

blockquote {
	color: #666;
	margin-bottom: 20px;
	padding: 0 0 0 20px;
	border-left: 3px solid #bbb
}

ul,
ol,
dl {
	margin-bottom: 15px
}

ul {
	list-style-position: inside;
	list-style: disc;
	padding-left: 20px
}

ol {
	list-style-position: inside;
	list-style: decimal;
	padding-left: 20px
}

dl dt {
	font-weight: bold
}

dl dd {
	padding-left: 20px;
	font-style: italic
}

dl p {
	padding-left: 20px;
	font-style: italic
}

hr {
	height: 1px;
	margin-bottom: 5px;
	border: none;
	background: url("../images/bg_hr.png") repeat-x center
}

table {
	border: 1px solid #373737;
	margin-bottom: 20px;
	text-align: left
}

th {
	font-family: 'Lucida Grande', 'Helvetica Neue', Helvetica, Arial, sans-serif;
	padding: 10px;
	background: #373737;
	color: #fff
}

td {
	padding: 10px;
	border: 1px solid #373737
}

form {
	background: #f2f2f2;
	padding: 20px
}

.outer {
	width: 100%
}

.inner {
	position: relative;
	max-width: 640px;
	padding: 20px 10px;
	margin: 0 auto
}

	
	
	
	
	
	
	
	
	
	
	
	
	


	</style>
    </head>
	<body>
	  <article class="markdown-body entry-content" style="padding: 30px;">
  `

var xxx = `
 body {
    background-color: #FFF;
    color: #172B4D;
    font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen,Ubuntu,"Fira Sans","Droid Sans","Helvetica Neue",sans-serif;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 1.42857142857143;
    -ms-overflow-style: -ms-autohiding-scrollbar;
    text-decoration-skip: ink;
}
	 p {
    display: block;
	magin-left: 1em;
    margin-block-start: 1em;
    margin-block-end: 1em;
    margin-inline-start: 0px;
    margin-inline-end: 0px;
}

h1 {
    font-size: 1.74285714em;
    font-style: inherit;
    font-weight: 600;
    line-height: 1.25;
    letter-spacing: -.006em;
    margin-top: 24px;
    color: #172B4D;
}

h2 {
    font-size: 1.52857143em;
    font-style: inherit;
    font-weight: 500;
    letter-spacing: -.008em;
    line-height: 1.2;
    margin-top: 28px;
    color: #172B4D;
}

 h3 {
    font-size: 1.34285714em;
    font-style: inherit;
    font-weight: 600;
    line-height: 1.25;
    letter-spacing: -.006em;
    margin-top: 24px;
    color: #172B4D;
}

h4 {
    font-size: 1.1;
    font-style: inherit;
    font-weight: 500;
    line-height: 1.25;
    letter-spacing: -.006em;
    margin-top: 24px;
    color: #172B4D;
}

h5 {
    font-size: 1.0;
    font-style: inherit;
    font-weight: 400;
    line-height: 1.25;
    letter-spacing: -.006em;
    margin-top: 24px;
    color: #172B4D;
}


pre {
    background: #F4F5F7;
    border: 1px solid #DFE1E6;
    border-radius: 3px;
    overflow-x: auto;
    padding: 5px 10px;
    word-wrap: normal;
	font-family: monospace; 
	font-size: 12px;
	font-weight: 300;
	
}


blockquote  {
    margin: 12px 12px 12px 12px;
	color: Green;
	font-weight: 700;
}

a:link {
  text-decoration: none;
  color: blue;
  font-weight: 500;
}

a:visited {
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
  font-weight: 550;
}

a:active {
  text-decoration: underline;
}`

var HtmlJavscriptReload = `
  <script>
  function doReload(){   
    location.reload(); 
  }

  setInterval(doReload, loopDelay);
  </script>
  `

var HtmlSuffix = `
    </article>
	</body>
	</html>
  `
