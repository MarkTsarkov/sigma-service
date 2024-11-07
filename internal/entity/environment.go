package entity

type Environment struct {
	Port int    `envconfig:"PORT"`
	DB   string `envconfig:"DB"`
}