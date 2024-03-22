package backend

type User struct {
	Id int `json:"id"`
}

type Summary struct {
	Id          int      `json:"id"`
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
	Platform    string              `json:"platform"`
	Credentials PlatformCredentials `json:"credentials"`
}

type PlatformCredentials struct {
	Token string `json:"token"`
}

type PlatformGroups struct {
	Platform string  `json:"platform"`
	Groups   []Group `json:"groups"`
}

type Group struct {
	Id        int    `json:"_id"`
	UserId    int    `json:"user_id"`
	GroupId   int    `json:"group_id"`
	GroupName string `json:"group_name"`
	Platform  string `json:"platform"`
	// json iso string
	CreatedAt string `json:"created_at"`
}
