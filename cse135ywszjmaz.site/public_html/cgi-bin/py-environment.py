#!/usr/bin/env python3
import os

print("Content-type: text/html\n\n")

print("<html><head><title>Environment Variables</title></head>")
print("<body><h1 align=center>Environment Variables</h1>\
    <hr/>\n")
print("<table>")

for key, value in os.environ.items():
    print("<tr><td>{0}</td><td>{1}</td></tr>".format(key, value))

print("</table></body></html>")