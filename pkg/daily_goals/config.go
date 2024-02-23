package daily_goals

type PixelaCredential struct {
	Token    string
	Username string
}

type Config struct {
	PixelaCredential PixelaCredential
}
