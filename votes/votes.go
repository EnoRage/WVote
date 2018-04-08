package votes

import (
	"net/url"

	"../waves"
)

// var serverIP string = "51.144.105.164"
var serverIP string = "localhost"

// CreateVote создание голосование
func CreateVote(userID string, description string, endTime string) string {
	data := url.Values{
		"userID":      {userID},
		"description": {description},
		"endTime":     {endTime},
	}
	url1 := "http://" + serverIP + ":3001/createVote"

	body := waves.Post(url1, data)

	return body
}

// SendDataTx создание голосование
func SendDataTx(userID string, encrSeed string, voteNum string, vote string) string {
	data := url.Values{
		"userID":   {userID},
		"encrSeed": {encrSeed},
		"voteNum":  {voteNum},
		"vote":     {vote},
	}
	url1 := "http://" + serverIP + ":3001/sendDataTx"

	body := waves.Post(url1, data)

	return body
}

// Vote проголосовать
func Vote(voteNum string, address string, vote string) string {
	data := url.Values{
		"voteNum": {voteNum},
		"address": {address},
		"vote":    {vote},
	}
	url1 := "http://" + serverIP + ":3001/vote"

	body := waves.Post(url1, data)

	return body
}

// TotalVote Подсчитать голоса
func TotalVote(voteNum string, whatVote string) string {
	data := url.Values{
		"voteNum":  {voteNum},
		"whatVote": {whatVote},
	}
	url1 := "http://" + serverIP + ":3001/totalVotes"

	body := waves.Post(url1, data)

	return body
}
