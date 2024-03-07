import os
import discord
import requests
from dotenv import load_dotenv

load_dotenv()
token = os.getenv('PYTHON')

def get_quote():
    raw = requests.get("https://zenquotes.io/api/random")
    res = raw.json()[0]
    return f'{res["q"]}\n-{res["a"]}'


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
    
    if message.content.startswith("!hello"):
        await message.channel.send(f"Ahoy hoy, {message.author}")

    if message.content.startswith("!quote"):
        await message.channel.send(get_quote())

client.run(token)