package votes

import (
	"net/url"

	"../waves"
)

// CreateVote создание голосование
func CreateVote(userID string, description string) string {
	data := url.Values{
		"userID":      {userID},
		"description": {description},
	}
	url1 := "http://51.144.105.164:3001/createSeed"

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
	url1 := "http://51.144.105.164:3001/vote"

	body := waves.Post(url1, data)

	return body
}

// TotalVote Подсчитать голоса
func TotalVote(voteNum string, whatVote string) string {
	data := url.Values{
		"voteNum":  {voteNum},
		"whatVote": {whatVote},
	}
	url1 := "http://51.144.105.164:3001/totalVotes"

	body := waves.Post(url1, data)

	return body
}
