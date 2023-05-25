#!/usr/bin/env python3

# Body - HTML
print("Content-type:text/html\r\n\r\n")
print("<html><head><title>State Demo</title>")
print("</head>")
print("<body>")
print("<h1 align=center>State Demo</h1>\
    <hr/>\n")

print("""
    <h2 align=center>Hi, what's your name? :)</h2>
    <form method="get" action="/cgi-bin/py-state-demo1.py" align=center>
        <label for="name">Name:</label>
        <input type="text" id="name" name="name"><br><br>
    <h2 align=center>And your major? :)</h2>
        <label for="major">Major: </label>
        <select name = "major" style="font-size: x-large;">Major:
            <option value = "Maths" selected>Maths</option>
            <option value = "Physics">Physics</option>
            <option value = "CSE">CSE</option>
        </select>
        <p></p>
        <button type="submit" style="margin-top:20px; text-align:center;">Test Sessioning</button>
    </form>
    <div style="display:block; margin-top:20px; text-align:center;">
        <a href="/">Home</a>
    </div>
    </body></html>
""")
