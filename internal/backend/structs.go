package backend

type User struct {
	Id int `json:"id"`
}

type Summary struct {
	Id          string   `json:"id"`
	Value       string   `json:"value"`
	Tags        []string `json:"tags"`
	DisplayTags string
}

type Task struct {
	Id              int    `json:"id"`
	Value           string `json:"value"`
	Deadline        string `json:"deadline"`
	HTMLDeadline    string
	DisplayDeadline string
	Tags            []string `json:"tags"`
	DisplayTags     string
}

type Event struct {
	Id       int    `json:"id"`
	Value    string `json:"value"`
	Location string `json:"location"`

	// json
	DateStart string `json:"dateStart"`
	// html
	HTMLDateStart string
	// pretty
	DisplayDateStart string

	// json
	DateEnd string `json:"dateEnd"`
	// html
	HTMLDateEnd string
	// pretty
	DisplayDateEnd string

	Tags        []string `json:"tags"`
	DisplayTags string
}

type PlatformEntry struct {
	Id           int    `json:"id"`
	PlatformName string `json:"platformName"`
	CredentialId string `json:"credentialId"`

	LastProcessed string `json:"lastProcessed"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`

	UserId string `json:"UserId"`
}

type PlatformGroups struct {
	Platform string  `json:"platform"`
	Groups   []Group `json:"groups"`
}

type Group struct {
	Id        string `json:"_id"`
	UserId    string `json:"user_id"`
	GroupId   string `json:"group_id"`
	GroupName string `json:"group_name"`
	Platform  string `json:"platform"`
	// json iso string
	CreatedAt string `json:"created_at"`
}
