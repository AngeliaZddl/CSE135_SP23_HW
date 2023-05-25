#!/usr/bin/env python3

import cgi

# Set response headers
print("Content-type: text/html")
print()
#!/usr/bin/env python3

# Get the POST data
form = cgi.FieldStorage()
print("<html><head><title>Hello, HTML World!</title></head>")
print("<body><h1 align=center>POST Request Echo:</h1>\
        <hr/>\n")
# Check if data was submitted
if form:
    # Output the data
    print("<h1 align=center>Your Info:</h1>")
    # keys = list(form.keys())
    # for key in keys:
    #     print("<h2 align=center>{}: {}</h2>".format(key, form[key].value))
    print("<h2 align=center>Name: {}</h2>".format(form['Name'].value))
    print("<h2 align=center>Email: {}</h2>".format(form['Email'].value))
    print("<h2 align=center>Phone: {}</h2>".format(form['Phone'].value))

else:
    # Output a form for submitting data
    print("""
        <h1 align=center>Your Info:</h1>
        <form method="post" align=center>
            <label for="Name">Name:</label>
            <input type="text" id="Name" name="Name"><br>
            <label for="Email" style="margin-top:20px;">Email:</label>
            <input type="email" id="Email" name="Email" style="margin-top:20px;"><br>
            <label for="Phone" style="margin-top:20px;">Phone:</label>
            <input type="tel" id="Phone" name="Phone" style="margin-top:20px;"><br>
            <input type="submit" value="Submit" style="margin-top:20px;">
        </form>
    """)
    
print("</body></html>")