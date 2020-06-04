package api

type ChallengeRequest struct {
	NumberPositions int    `json:"numero_casas"`
	Token           string `json:"token"`
	Cipher          string `json:"cifrado"`
	Decipher		string `json:"decifrado"`
	SHA1			string `json:"resumo_criptografico"`
}