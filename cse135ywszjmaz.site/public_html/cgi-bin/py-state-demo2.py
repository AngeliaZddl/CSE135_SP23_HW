#!/usr/bin/env python3

import os
import http.cookies as Cookie
import cgi

print("Content-type:text/html\r\n\r\n")
print("<html><head><title>State Demo</title>")
print("</head>")
print("<body>")
print("<h1 align=center>State Demo Page 2</h1>\
    <hr/>\n")

cookie = os.getenv("HTTP_COOKIE")

if cookie and "destroyed" in cookie:
    print("""
        <h2>Cookie: </h2>
        <p>{}</p>
    """.format(cookie))
else:
    print("<tr><td>Cookie:</td><td>None</td></tr>")

print("""
    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/cgi-bin/py-state-demo1.py">Session Page 1</a>
    </div>
    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/cgi-bin/py-state-demo.py">Python CGI Form</a>
    </div>
    <form method="post" action="/cgi-bin/py-state-demo3.py" align=center style="display:block; margin-top:20px;">
        <button type="botton">Destroy Session</button>
    </form>
    <a style="display:block; margin-top:20px; text-align:center;" href="/">Home</a>
    </body></html>
""")