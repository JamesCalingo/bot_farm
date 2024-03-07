const Discord = require("discord.js");
const fetch = require("node-fetch")

const client = new Discord.Client();

function getQuote() {
  return fetch("https://zenquotes.io/api/random")
  .then(res => {
    return res.json()
  })
  .then(data => {
    return data[0]["q"] + " -" + data[0]["a"]
  })
}

client.on("ready", ()=> {
  console.log(`logged in as ${client.user.tag}`);
});

client.on("message", msg => {
  if(msg.author.bot) return
  else if(msg.content.includes("qhote")){
    msg.channel.send("use $quote)")
  }
  else if(msg.content === "$quote"){
    getQuote()
    .then(quote => msg.channel.send(quote))
  }

  if(msg.content.startsWith("$creategame"))
});

client.login(process.env.TOKEN);
