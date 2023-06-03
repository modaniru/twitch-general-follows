package twitch

type OauthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type ValidToken struct {
	ClientId  string   `json:"client_id"`
	Login     string   `json:"login"`
	Scopes    []string `json:"scopes"`
	UserId    string   `json:"user_id"`
	ExpiresIn int      `json:"expires_in"`
}

type UserInfo struct {
	Id              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
	CreatedAt       string `json:"created_at"`
}

type UserCollection struct {
	Data []UserInfo `json:"data"`
}

type FollowInfo struct {
	FromId     string `json:"from_id"`
	FromLogin  string `json:"from_login"`
	FromName   string `json:"from_name"`
	ToId       string `json:"to_id"`
	ToName     string `json:"to_name"`
	FollowedAt string `json:"followed_at"`
}

type Pagination struct {
	Cursor string `json:"cursor"`
}

type FollowsCollection struct {
	Total      int
	Data       []FollowInfo
	Pagination Pagination
}
