import os
import discord
from dotenv import load_dotenv

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
        await message.channel.send("Ahoy hoy")

client.run(token)