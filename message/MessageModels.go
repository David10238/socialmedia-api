package message

type message struct {
	Id       int    `json:"id"`
	AuthorId int    `json:"author_id"`
	Message  string `json:"message"`
	IsPublic bool   `json:"is_public"`
}

type messageRequest struct {
	TargetId            int `json:"target_id"`
	LowestLoadedMessage int `json:"lowest_loaded_message"`
}

type messageResponse struct {
	Messages []message `json:"messages"`
	LowestId int       `json:"lowest_id"`
}

func newMessageResponse(messages []message) messageResponse {
	return messageResponse{
		Messages: messages,
		LowestId: messages[len(messages)-1].Id,
	}
}
