import json
import sys
import urllib.request

url = "https://www.toptal.com/developers/gitignore/api/list?format=json"

req = urllib.request.Request(url)
req.add_header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:85.0) Gecko/20100101 Firefox/85.0")

data = {}

with urllib.request.urlopen(req) as resp:
    data = json.load(resp)

new_data = {}

for key in data:
    new_data[key] = data[key]["contents"]

with open(sys.argv[1], "w") as f:
    json.dump(new_data, f, indent=4)
