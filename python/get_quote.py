import requests

def get_quote():
    raw = requests.get("https://zenquotes.io/api/random")
    res = raw.json()[0]
    return f'\"{res["q"]}\"\n-{res["a"]}'