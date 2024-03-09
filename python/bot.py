import os
import discord
import json
from dotenv import load_dotenv
from get_quote import get_quote
from challenges import get_random_challenge, get_challenge_list, add_challenge

file = open("../challenges.json")
challenges = json.load(file)

load_dotenv()
token = os.getenv('PYTHON')

intents = discord.Intents.default()
intents.message_content = True

client = discord.Client(intents=intents)

@client.event
async def on_ready():
    print(f"logged in as {client.user}")

@client.event
async def on_message(message):
    if message.author == client.user:
        return
    
    if message.content.startswith("$hello"):
        await message.channel.send(f"Ahoy hoy, {message.author}")

    if message.content.startswith("$quote"):
        await message.channel.send(get_quote())

    if message.content.startswith("$challenge"):
        await message.channel.send("YOU WANT A CHALLENGE? How about this one:\n" + get_random_challenge(challenges))

    if message.content.startswith("$list"):
        await message.channel.send("The (bad word) List:\n" + get_challenge_list(challenges))

    if message.content.startswith("$add"):
        space = message.content.find(" ")
        url = message.content[space+1:]
        new_challenge = add_challenge(url, challenges)
        if type(new_challenge) is str:
            await message.channel.send(new_challenge)
        else:
            challenges.append(new_challenge)
            added = f"Added: {new_challenge['name']} {new_challenge['url']}"
            await message.channel.send(added)

client.run(token)