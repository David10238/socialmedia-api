package message

import "github.com/David10238/socialmedia-api/db"

type messageRequest struct {
	TargetId            uint `json:"target_id"`
	LowestLoadedMessage uint `json:"lowest_loaded_message"`
	SearchFriends       bool `json:"search_friends"`
}

type messageResponse struct {
	Messages []db.Message `json:"messages"`
	LowestId uint         `json:"lowest_id"`
}

func newMessageResponse(messages []db.Message) messageResponse {
	lowest := uint(0)
	if len(messages) != 0 {
		lowest = messages[len(messages)-1].ID
	}

	return messageResponse{
		Messages: messages,
		LowestId: lowest,
	}
}
