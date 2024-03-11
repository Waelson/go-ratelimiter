package config

// EnvLoader define a interface para carregar variáveis de ambiente.
type EnvLoader interface {
	Load() error
}
