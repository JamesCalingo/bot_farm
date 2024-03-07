const express = require("express")
const app = express()

const port = process.env.PORT || 3000

app.all("/", (req, res) => {
  res.send("Bot online.")
})

function pokeBot() {
  app.listen(port, () => {
    console.log("Server listening on " + port)
  })
}

module.exports = pokeBot