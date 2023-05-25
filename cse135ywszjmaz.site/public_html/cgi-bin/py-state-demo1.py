#!/usr/bin/env python3

from http import cookies
import os
import cgi

print("Content-type:text/html\r\n\r\n")
print("<html><head><title>State Demo</title>")
print("</head>")
print("<body>")
print("<h1 align=center>State Demo Page 1</h1>\
    <hr/>\n")

# Retrieve cookie value
form = cgi.FieldStorage()
name = form.getvalue('name')
major = form.getvalue('major')

c = cookies.SimpleCookie()
c['name'] = name
c['major'] = major

c['name']['expires'] = 60 * 60 # expires one hour
c['major']['expires'] = 60 * 60 # expires one hour

c['name']['path'] = '/'
c['major']['path'] = '/'

print(c.output())

if name != None:
    print("""
        <h2 align=center>Hello, <strong>{}</strong>!</h2>
        <h2 align=center>Have a Great Grade on {}</h2>
    """.format(name, major))

# elif os.getenv("HTTP_COOKIE") and "destroyed" in os.getenv("HTTP_COOKIE"):
#     print("""
#         <h2>Cookie: </h2>
#         <p>{}</p>
#     """.format(os.getenv("HTTP_COOKIE")))
else:
    print("""
        <h2 align=center>Name is required.</h2>
        <h2 align=center>Please back to last step.</h2>
        <form method="post" action="/cgi-bin/py-state-demo.py" align=center>
            <button type="botton">Back</button>
        </form>
    """)
    

print("""
    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/cgi-bin/py-state-demo2.py">Session Page 2</a>
        </div>
    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/cgi-bin/py-state-demo.py">Python CGI Form</a>
    </div>
    <form method="post" action="/cgi-bin/py-state-demo3.py" align=center style="display:block; margin-top:20px;">
        <button type="submit">Destroy Session</button>
    </form>
    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/">Home</a>
    </div>
    </body></html>
""")
