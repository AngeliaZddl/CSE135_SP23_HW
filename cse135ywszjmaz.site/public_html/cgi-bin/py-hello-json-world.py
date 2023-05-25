#!/usr/bin/env python3
import json
import datetime
import os

# Set the content type to JSON
print("Content-type: application/json\n")

# Define the response dictionary
content = {
    "message": "Hello World from Python!",
    "date": "Today's date is {}".format(datetime.date.today().strftime("%Y-%m-%d")),
    "ipAddress": os.environ.get('REMOTE_ADDR')
}

# Convert the dictionary to a JSON string
json_string = json.dumps(content)

# Print the JSON string
print(json_string)