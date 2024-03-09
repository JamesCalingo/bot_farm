# Bot farm

Where I build a bunch of Discord bots for testing a few things.

Based on John Crickett's [Build Your Own Discord Bot Challenge](https://codingchallenges.fyi/challenges/challenge-discord/)

## Meet the bots

There are three bots in this repository. Currently, I have them set so that you can't add them to your server unfortunately.

- Monty, a bot written in Python (and named after Monty Python just like the language)
- Gofer, a bot written in Go
- A super old JavaScript bot I worked on in 2020. It's in need of some updating.

## Commands

Both the Python and Go bots have the same set of basic commands:

- `hello` - the bot will reply with a short message which includes your username
- `quote` - the bot will reply with a quote at random from `https://zenquotes.io/api/random`
- `challenge` - the bot will reply with one of the [Coding Challenges](https://codingchallenges.fyi), selected at random
- `list` - the bot will reply with a (partial) list of the aforementioned Coding Challenges.
- `add` - if a link to a [Coding Challenge](https://codingchallenges.fyi) is included with this command, the bot will add it to its list of Coding Challenges. Otherwise, it will reply that it can't do anything.
  (NOTE: This is currently only active with the Python bot)

## The Process

I'll use a more "general" approach to describing how I did this.

The first thing we need: connection to Discord. This not only means setting up the bot on Discord itself, but connecting it to our code. For the Python bot, I used [discord.py](https://discordpy.readthedocs.io/en/stable/), and for the Go bot, [this DiscordGo package](https://github.com/bwmarrin/discordgo). For setting up the actual bots on Discord itself, the Python version has a 

Most languages have 

Second, we need to get quotes. The API I mentioned in the list 

Notably, it returns an array of length 1 with the data in it, so kinda have to "hack" our way into that 

Next, the stuff involving challenges. John Crickett provided a list of challenges from a while ago, and while it is simple and helpful, I noticed

He started the titles off with "Write", but his website uses the word "Build"! 


## The future

Currently, these bots have "on" and "off" buttons, so they're not always "on". I'd like to look into running them "full time", but alas, a lot of the free options that used to exist no longer do.

Also, I have some ideas for other features for these bots, as well as some thoughts for Slack versions.
