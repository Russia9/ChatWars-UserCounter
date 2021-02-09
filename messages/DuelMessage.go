package messages

type DuelMessage struct {
	Winner    ArenaGamer `json:"winner"`
	Loser     ArenaGamer `json:"loser"`
	Challenge bool       `json:"isChallenge"`
	GuildDuel bool       `json:"isGuildDuel"`
}

type ArenaGamer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	Castle string `json:"castle"`
	Level  int    `json:"level"`
	HP     int    `json:"hp"`
}
