import random

def get_challenge(list):
    challenge = list[random.randint(0, len(list))]
    return f"{challenge['name']}: {challenge['url']}"

def get_challenge_list(list):
    output = ""
    for challenge in list:
        output += f"{challenge['name']}: {challenge['url']}\n"
    return output