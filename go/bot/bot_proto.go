package bot

// var challenges = Open("../challenges.json")

// switch {
// case message.Content == "&hello":
// 	greeting := "Hello " + message.Author.Username
// 	discord.ChannelMessageSend(message.ChannelID, greeting)
// case message.Content == "&quote":
// 	discord.ChannelMessageSend(message.ChannelID, GetQuote())
// case strings.Contains(message.Content, "&challenge"):
// 	discord.ChannelMessageSend(message.ChannelID, GetRandomChallenge(challenges))
// case strings.Contains(message.Content, "&list"):
// 	discord.ChannelMessageSend(message.ChannelID, GetChallengeList(challenges))
// case strings.HasPrefix(message.Content, "&add"):
// 	_, url, _ := strings.Cut(message.Content, " ")
// 	newChallenge, err := AddChallenge(url)
// 	if err != nil {
// 		discord.ChannelMessageSend(message.ChannelID, "This is not a valid URL.")
// 		break
// 	}
// 	added := "added" + " " + newChallenge.URL
// 	challenges = append(challenges, newChallenge)
// 	discord.ChannelMessageSend(message.ChannelID, added)
// }
