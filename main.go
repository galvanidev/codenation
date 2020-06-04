package main

import (
	"cipher/api"
	"cipher/caesar"
	"cipher/util"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"os"
)

func main() {

	codenation := api.NewApi(os.Getenv("TOKEN_API"))
	util.SaveFile("answer.json", codenation.GetData())

	file := util.ReadFile("answer.json")

	var res api.ChallengeResponse
	err := json.Unmarshal(file, &res)
	util.Check(err)

	c := caesar.NewCaesar(res.NumberPositions)
	decipher := c.Decipher(res.Cipher)

	var req api.ChallengeRequest

	h := sha1.New()
	h.Write([]byte(decipher))

	req.NumberPositions = res.NumberPositions
	req.Token = res.Token
	req.Cipher = res.Cipher
	req.Decipher = decipher
	req.SHA1 = string(hex.EncodeToString(h.Sum(nil)))

	util.SaveFile("answer.json", req)
	codenation.SendFile("answer.json")

}