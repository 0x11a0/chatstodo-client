package backend

type User struct {
	Id int `json:"id"`
}

type Summary struct {
	Id    int      `json:"id"`
	Value string   `json:"value"`
	Tags  []string `json:"tags"`
}

type Task struct {
	Id              int    `json:"id"`
	Value           string `json:"value"`
	Deadline        string `json:"deadline"`
	LocalDeadline   string
	DisplayDeadline string
	Tags            []string `json:"tags"`
}

type Event struct {
	Id               int    `json:"id"`
	Value            string `json:"value"`
	Location         string `json:"location"`

	DateStart        string `json:"dateStart"`
	LocalDateStart   string
	DisplayDateStart string

	DateEnd          string `json:"dateEnd"`
	LocalDateEnd     string
	DisplayDateEnd   string
	Tags             []string `json:"tags"`
}

type PlatformEntry struct {
	Platform    string              `json:"platform"`
	Credentials PlatformCredentials `json:"credentials"`
}

type PlatformCredentials struct {
	Token string `json:"token"`
}
