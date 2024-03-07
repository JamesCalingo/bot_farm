const Discord = require("discord.js")
const poke = require("./server")

const client = new Discord.Client()

let gameRunning = false

client.on("ready", ()=> {
  console.log(`logged in as ${client.user.tag}`);
});

function closeGame () {
  gameRunning = false
}

client.on("message", msg => {
  if(msg.author.bot) return

  if(msg.content === "$creategame" && gameRunning === false){
    msg.channel.send("Starting new game. Please enter $join to join.")
    gameRunning = true
    return
  }

    if(msg.content === "$creategame" && gameRunning === true){
    msg.channel.send("Game has already been started.")
  }

  if(msg.content === "$join" && gameRunning === true){
    msg.reply("You're in!")
  }

   if(msg.content === "$join" && gameRunning === false){
    msg.reply("There is currently no game running right now. Use $creategame to create a new game.")
  }

  if(msg.content === "$start" && gameRunning === true){
    msg.channel.send("Starting game now. Roles have been DMed to you. Have fun!")
    closeGame()
  }

  if(msg.content === "$closegame"){
    closeGame()
    msg.channel.send("Closing game.")
  }

if(msg.content === "$check"){
  msg.channel.send(gameRunning)
}

});

poke()

client.login(process.env.TOKEN);