#!/usr/bin/env python3
import datetime
import os

print("Content-type: text/html\n")
print("<html>")
print("<head><title>Hello, HTML World!</title></head>")
print("<body>")
print("<h1 align=center>Hello, Python World!</h1>\
    <hr/>\n")
print("<h2>Hello World!</h2>")
print("<p>This page was generated with the <strong>Python</strong> programming language.</p>")
print("<p>This program was run at: {}</p>".format(datetime.date.today()))
print("<p>Your current IP address is: {}</p>".format(os.environ.get('REMOTE_ADDR')))
print("</body>")
print("</html>")