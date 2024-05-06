package bot

import (
	"fmt"
	"slices"
)

type Game struct {
	isRunning bool
	players   []string
	roles     []string
}

// type Player struct {
// 	id   string
// 	role string
// }

func (g *Game) toggleLobby() {
	g.isRunning = !g.isRunning
}

func (g *Game) joinGame(id string) {
	g.players = append(g.players, id)
	fmt.Println(g.players)
}

func (g *Game) start() {
	g.toggleLobby()
	g.players = slices.Delete(g.players, 0, len(g.players))
}
