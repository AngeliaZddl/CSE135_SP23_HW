#!/usr/bin/env python3

# Body - HTML
print("Content-type:text/html\r\n\r\n")
print("<html><head><title>State Demo</title>")

print("""
    <h1 align=center>Session Destroyed</h1>
    <hr/>
    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/cgi-bin/py-state-demo1.py">Session Page 1</a>
    </div>

    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/cgi-bin/py-state-demo2.py">Session Page 2</a>
    </div>

    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/cgi-bin/py-state-demo.py">Python CGI Form</a>
    </div>
    
    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/">Home</a>
    </div>
""")