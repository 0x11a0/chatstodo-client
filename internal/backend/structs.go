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
	LocalDeadline   string
	DisplayDeadline string
	Tags            []string `json:"tags"`
	DisplayTags     string
}

type Event struct {
	Id       int    `json:"id"`
	Value    string `json:"value"`
	Location string `json:"location"`

	// json
	DateStart        string `json:"dateStart"`
	// html
	HTMLDateStart   string
	// pretty
	DisplayDateStart string

	// json
	DateEnd        string `json:"dateEnd"`
	// html
	HTMLDateEnd   string
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