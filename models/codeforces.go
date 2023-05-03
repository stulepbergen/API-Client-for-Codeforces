package models

type UserData struct {
	LastName      string `json:"lastName"`
	FirstName     string `json:"firstName"`
	Handle        string `json:"handle"`
	Rating        int    `json:"rating"`
	FriendOfCount int    `json:"friendOfCount"`
	Organization  string `json:"organization"`
	Rank          string `json:"rank"`
	MaxRating     int    `json:"maxRating"`
	MaxRank       string `json:"maxRank"`
}

type UsersInfo struct {
	Result []UserData `json:"result"`
}

type UserRatings struct {
	ContestId   int    `json:"contestId"`
	ContestName string `json:"contestName"`
	Handle      string `json:"handle"`
	Rank        int    `json:"rank"`
	OldRating   int    `json:"oldRating"`
	NewRating   int    `json:"newRating"`
}

type Handles struct {
	Handle string `json:"handle, omitempty"`
}

type Authors struct {
	ContestId int       `json:"contestId"`
	Members   []Handles `json:"members, omitempty"`
}

type ProblemInfo struct {
	ContestId int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Points    float64  `json:"points"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags, omitempty"`
}

type Problems struct {
	Problem ProblemInfo `json:"problem"`
	Author  Authors     `json:"author, omitempty"`
}

type ProblemSet struct {
	Result []Problems `json:"result"`
}

type UsersRating struct {
	Result []UserRatings `json:"result"`
}
