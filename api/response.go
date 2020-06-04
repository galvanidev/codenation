package api

type ChallengeResponse struct {
	NumberPositions int    `json:"numero_casas"`
	Token           string `json:"token"`
	Cipher          string `json:"cifrado"`
}