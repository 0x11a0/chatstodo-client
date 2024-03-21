package backend

type User struct {
	Id int `json:"id"`
}

type Credential struct {
	Id               int    `json:"id"`
	CredentialId     string `json:"credentialID"`
	CredentialSecret string `json:"credentialSecret"`
}

type Platform struct {
	Id           int    `json:"id"`
	PlatformName string `json:"platformName"`
}

type Summary struct {
	Id    int      `json:"id"`
	Value string   `json:"value"`
	Tags  []string `json:"tags"`
}

type Task struct {
	Id       int      `json:"id"`
	Value    string   `json:"value"`
	Deadline string   `json:"deadline"`
	Tags     []string `json:"tags"`
}

type Event struct {
	Id        int      `json:"id"`
	Value     string   `json:"value"`
	Location  string   `json:"location"`
	DateStart string   `json:"dateStart"`
	DateEnd   string   `json:"dateEnd"`
	Tags      []string `json:"tags"`
}

type PlatformEntry struct {
	Platform    string              `json:"platform"`
	Credentials PlatformCredentials `json:"credentials"`
}

type PlatformCredentials struct {
	Token string `json:"token"`
}
