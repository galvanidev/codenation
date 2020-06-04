package api

import (
	"bytes"
	"cipher/util"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	BASE_URL = "https://api.codenation.dev/v1"
)

type api struct {
	token string
}

func NewApi(token string) *api {
	api := api{}
	api.token = token
	return &api
}

func (api *api) GetData() ChallengeResponse {

	req, err := http.NewRequest("GET", BASE_URL + "/challenge/dev-ps/generate-data?token=" + api.token, nil)
	util.Check(err)

	bodyBytes := doRequest(req)

	var res ChallengeResponse
	err = json.Unmarshal(bodyBytes, &res)
	util.Check(err)

	return res;
}

func (api *api) SendFile(fileName string) {

	file, err := os.Open(fileName)
	util.Check(err)

	fileContent, err := ioutil.ReadAll(file)
	util.Check(err)

	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	multipartWriter, err := writer.CreateFormFile("answer", fileName)
	util.Check(err)
	multipartWriter.Write(fileContent)
	io.Copy(multipartWriter, file)
	writer.Close()

	req, err := http.NewRequest("POST", BASE_URL+ "/challenge/dev-ps/submit-solution?token=" + api.token, &requestBody)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	util.Check(err)

	log.Println("Request:\n", &requestBody)

	doRequest(req)
}

func doRequest(request *http.Request) []byte {
	client := http.Client{}
	response, err := client.Do(request)
	util.Check(err)
	validation(response)

	bytes, err := ioutil.ReadAll(response.Body)
	util.Check(err)
	defer response.Body.Close()

	log.Println(string(bytes))

	return bytes
}

func validation(response* http.Response) {
	if response.StatusCode != 200 {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		util.Check(err)
		log.Fatal(response.Status, string(bodyBytes))
	}
}