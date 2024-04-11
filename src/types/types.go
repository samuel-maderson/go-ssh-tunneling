package types

type Remote struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Local struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Bastion struct {
	Host string `json:"host"`
}

type Config struct {
	User    string  `json:"user"`
	Key     string  `json:"key"`
	Remote  Remote  `json:"remote"`
	Local   Local   `json:"local"`
	Bastion Bastion `json:"bastion"`
}
