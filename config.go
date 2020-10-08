package main

type config struct {
	GhostBaseURL string `env:"GHOST_BASE_URL,required"`
	GhostApiKey  string `env:"GHOST_API_KEY,required"`

	SendinblueApiKey string `env:"SENDINBLUE_API_KEY,required"`
	SendinblueListID int64  `env:"SENDINBLUE_LIST_ID,required"`
}
