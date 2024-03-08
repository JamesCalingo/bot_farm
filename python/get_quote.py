import requests

def get_quote():
    raw = requests.get("https://zenquotes.io/api/random")
    # Our API returns a list of length 1, so we have to get the data from an array index
    res = raw.json()[0]
    return f'\"{res["q"]}\"\n-{res["a"]}'