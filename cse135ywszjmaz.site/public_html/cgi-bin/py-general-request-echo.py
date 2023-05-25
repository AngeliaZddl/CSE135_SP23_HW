#!/usr/bin/env python3

import cgi
import os

# Get the request method and protocol
request_method = os.environ.get("REQUEST_METHOD")
protocol = os.environ.get("SERVER_PROTOCOL")

# Get the query string
query_string = os.environ.get("QUERY_STRING")

# Get the message body
form = cgi.FieldStorage()
message_body = form.getvalue('message') if form.getvalue('message') else ""

# Print the general request echo
print("Content-Type: text/html\n")
print("<html><head><title>General Request Echo</title>")
print("</head>")
print("<body>")
print("<h1 align=center>General Request Echo</h1>\
    <hr/>\n")

print("<p><strong>Request Method</strong>: {}</p>".format(request_method))
print("<p><strong>Protocol</strong>: {}</p>".format(protocol))
print("<p><strong>Query</strong>: {}</p>".format(query_string))
print("<p><strong>Message Body</strong>: {}</p>".format(message_body))

print("</body></html>")