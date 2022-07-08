package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}

}
func handle(conn net.Conn) {

	defer conn.Close()
	request(conn)

}
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}
func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	fmt.Println("*****************Method:", method)
	fmt.Println("****************URI:", uri)
	if method == "GET" && uri == "/" {
		indexRes(conn)

	}
	if method == "GET" && uri == "/news" {
		newsRes(conn)

	}
	if method == "GET" && uri == "/about" {
		aboutRes(conn)

	}
	if method == "GET" && uri == "/contact" {
		contactRes(conn)

	}

}
func indexRes(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<body>
	<h1>Home</h1>
	
	<ul>
	  <li><a href="/">Home</a></li>
	  <li><a href="/news">News</a></li>
	  <li><a href="/contact">Contact</a></li>
	  <li><a href="/about">About</a></li>
	</ul>
	

	
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type:text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintln(conn, body)

}
func newsRes(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<style>
	body {font-family: Arial, Helvetica, sans-serif;}
	
	form {
	  border: 3px solid #f1f1f1;
	  font-family: Arial;
	}
	
	.container {
	  padding: 20px;
	  background-color: #f1f1f1;
	}
	
	input[type=text], input[type=submit] {
	  width: 100%;
	  padding: 12px;
	  margin: 8px 0;
	  display: inline-block;
	  border: 1px solid #ccc;
	  box-sizing: border-box;
	}
	
	input[type=checkbox] {
	  margin-top: 16px;
	}
	
	input[type=submit] {
	  background-color: #04AA6D;
	  color: white;
	  border: none;
	}
	
	input[type=submit]:hover {
	  opacity: 0.8;
	}
	</style>
	<body>
	
	<h2>CSS Newsletter</h2>
	
	<form action="/about">
	  <div class="container">
		<h2>Subscribe to our Newsletter</h2>
		<p>Lorem ipsum text about why you should subscribe to our newsletter blabla. Lorem ipsum text about why you should subscribe to our newsletter blabla.</p>
	  </div>
	
	  <div class="container" style="background-color:white">
		<input type="text" placeholder="Name" name="name" required>
		<input type="text" placeholder="Email address" name="mail" required>
		<label>
		  <input type="checkbox" checked="checked" name="subscribe"> Daily Newsletter
		</label>
	  </div>
	
	  <div class="container">
		<input type="submit" value="Subscribe">
	  </div>
	</form>
	
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type:text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintln(conn, body)

}
func aboutRes(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
	body {
	  font-family: Arial, Helvetica, sans-serif;
	  margin: 0;
	}
	
	html {
	  box-sizing: border-box;
	}
	
	*, *:before, *:after {
	  box-sizing: inherit;
	}
	
	.column {
	  float: left;
	  width: 33.3%;
	  margin-bottom: 16px;
	  padding: 0 8px;
	}
	
	.card {
	  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
	  margin: 8px;
	}
	
	.about-section {
	  padding: 50px;
	  text-align: center;
	  background-color: #474e5d;
	  color: white;
	}
	
	.container {
	  padding: 0 16px;
	}
	
	.container::after, .row::after {
	  content: "";
	  clear: both;
	  display: table;
	}
	
	.title {
	  color: grey;
	}
	
	.button {
	  border: none;
	  outline: 0;
	  display: inline-block;
	  padding: 8px;
	  color: white;
	  background-color: #000;
	  text-align: center;
	  cursor: pointer;
	  width: 100%;
	}
	
	.button:hover {
	  background-color: #555;
	}
	
	@media screen and (max-width: 650px) {
	  .column {
		width: 100%;
		display: block;
	  }
	}
	</style>
	</head>
	<body>
	
	<div class="about-section">
	  <h1>About Us Page</h1>
	  <p>Some text about who we are and what we do.</p>
	  <p>Resize the browser window to see that this page is responsive by the way.</p>
	</div>
	
	<h2 style="text-align:center">Our Team</h2>
	<div class="row">
	  <div class="column">
		<div class="card">
		  <img src="/w3images/team1.jpg" alt="Jane" style="width:100%">
		  <div class="container">
			<h2>Jane Doe</h2>
			<p class="title">CEO & Founder</p>
			<p>Some text that describes me lorem ipsum ipsum lorem.</p>
			<p>jane@example.com</p>
			<p><button class="button">Contact</button></p>
		  </div>
		</div>
	  </div>
	
	  <div class="column">
		<div class="card">
		  <img src="/w3images/team2.jpg" alt="Mike" style="width:100%">
		  <div class="container">
			<h2>Mike Ross</h2>
			<p class="title">Art Director</p>
			<p>Some text that describes me lorem ipsum ipsum lorem.</p>
			<p>mike@example.com</p>
			<p><button class="button">Contact</button></p>
		  </div>
		</div>
	  </div>
	  
	  <div class="column">
		<div class="card">
		  <img src="/w3images/team3.jpg" alt="John" style="width:100%">
		  <div class="container">
			<h2>John Doe</h2>
			<p class="title">Designer</p>
			<p>Some text that describes me lorem ipsum ipsum lorem.</p>
			<p>john@example.com</p>
			<p><button class="button">Contact</button></p>
		  </div>
		</div>
	  </div>
	</div>
	
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type:text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintln(conn, body)
}
func contactRes(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
	body {
	  font-family: Arial, Helvetica, sans-serif;
	}
	
	* {
	  box-sizing: border-box;
	}
	
	/* Style inputs */
	input[type=text], select, textarea {
	  width: 100%;
	  padding: 12px;
	  border: 1px solid #ccc;
	  margin-top: 6px;
	  margin-bottom: 16px;
	  resize: vertical;
	}
	
	input[type=submit] {
	  background-color: #04AA6D;
	  color: white;
	  padding: 12px 20px;
	  border: none;
	  cursor: pointer;
	}
	
	input[type=submit]:hover {
	  background-color: #45a049;
	}
	
	/* Style the container/contact section */
	.container {
	  border-radius: 5px;
	  background-color: #f2f2f2;
	  padding: 10px;
	}
	
	/* Create two columns that float next to eachother */
	.column {
	  float: left;
	  width: 50%;
	  margin-top: 6px;
	  padding: 20px;
	}
	
	/* Clear floats after the columns */
	.row:after {
	  content: "";
	  display: table;
	  clear: both;
	}
	
	/* Responsive layout - when the screen is less than 600px wide, make the two columns stack on top of each other instead of next to each other */
	@media screen and (max-width: 600px) {
	  .column, input[type=submit] {
		width: 100%;
		margin-top: 0;
	  }
	}
	</style>
	</head>
	<body>
	
	<h2>Responsive Contact Section</h2>
	<p>Resize the browser window to see the effect.</p>
	
	<div class="container">
	  <div style="text-align:center">
		<h2>Contact Us</h2>
		<p>Swing by for a cup of coffee, or leave us a message:</p>
	  </div>
	  <div class="row">
		<div class="column">
		  <img src="/w3images/map.jpg" style="width:100%">
		</div>
		<div class="column">
		  <form action="/action_page.php">
			<label for="fname">First Name</label>
			<input type="text" id="fname" name="firstname" placeholder="Your name..">
			<label for="lname">Last Name</label>
			<input type="text" id="lname" name="lastname" placeholder="Your last name..">
			<label for="country">Country</label>
			<select id="country" name="country">
			  <option value="australia">Australia</option>
			  <option value="canada">Canada</option>
			  <option value="usa">USA</option>
			</select>
			<label for="subject">Subject</label>
			<textarea id="subject" name="subject" placeholder="Write something.." style="height:170px"></textarea>
			<input type="submit" value="Submit">
		  </form>
		</div>
	  </div>
	</div>
	
	</body>
	</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type:text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintln(conn, body)
}
