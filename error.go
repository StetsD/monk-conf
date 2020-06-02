package config

type ErrorConfig string

func (ec ErrorConfig) Error() string {
	return string(ec)
}
