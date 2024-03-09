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

def add_challenge(url, list):
    if url is "":
        return "To add a challenge to my list, please type \"$add \" and then the url of the challenge you want to add"
    valid = re.search("^https://codingchallenges.fyi/challenges/", url)
    if not valid:
        return "This is not a valid challenge. Check your URL and try again."
    for challenge in list:
        if url == challenge["url"]:
            return "This challenge is already in the list.\n" + get_challenge_list(list)
    raw = requests.get(url)
    # There's a nasty edge case where we could potentially add a "junk challenge", so let's cover that
    if raw.status_code != 200:
        return "There was an error trying to add the challenge. Check the URL and try again."
    else:
        res = raw.text
        start = res.find("Build")
        end = res.find("|")
        name = res[start:end-1]
        return {'name': name, 'url': url}