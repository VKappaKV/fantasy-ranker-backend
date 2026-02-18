package riot

type Account struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type RiotAPIError struct {
	Status struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	} `json:"status"`
}

type RiotMatch struct {
	Metadata struct {
		MatchID string `json:"matchId"`
	} `json:"metadata"`

	Info struct {
		GameDuration int64 `json:"gameDuration"`
		QueueID      int   `json:"queueId"`
		Participants []struct {
			PUUID    string `json:"puuid"`
			Champion string `json:"championName"`
			Kills    int    `json:"kills"`
			Deaths   int    `json:"deaths"`
			Assists  int    `json:"assists"`
			Win      bool   `json:"win"`
		} `json:"participants"`
	} `json:"info"`
}