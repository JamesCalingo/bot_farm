import random
import re
import requests

def get_random_challenge(list):
    challenge = list[random.randint(0, len(list))]
    return f"{challenge['name']}: {challenge['url']}"

def get_challenge_list(list):
    output = ""
    for challenge in list:
        output += f"{challenge['name']}: {challenge['url']}\n"
    return output

def add_challenge(url):
    if url == "":
        return "To add a challenge to my list, please type \"$add \" and then the url of the challenge you want to add"
    valid = re.search("^https://codingchallenges.fyi/challenges/", url)
    if not valid:
        return "This is not a valid challenge. Check your URL and try again."
    raw = requests.get(url)
    res = raw.text
    start = res.find("Build")
    end = res.find("|")
    name = res[start:end-1]
    return {'name': name, 'url': url}