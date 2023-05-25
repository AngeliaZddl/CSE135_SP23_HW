#!/usr/bin/env python3

import os
import sys
import urllib.parse

# Get the query string from environment variables
query_string = os.environ.get("QUERY_STRING", "")

# Parse the query string
params = urllib.parse.parse_qsl(query_string)

# Print the response HTML
print("Content-type:text/html\r\n\r\n")
print("<html>")
print("<head>")
print("<title>GET Echo</title>")
print("</head>")
print("<body>")
print("<h1 align=center>GET Requst Echo</h1>\
    <hr/>\n")
print("")
print("<p>Query String:</p>")
print("<ul>")
for name, value in params:
    print(f"<li><b>{name}:</b> {value}</li>")
print("</ul>")
print("</body>")
print("</html>")