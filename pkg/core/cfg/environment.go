package cfg

const (
	KeyEnvironment Key = "ENVIRONMENT" // Ключ: среда выполнения приложения.
	KeyPrefix      Key = "PREFIX"      // Ключ: префикс конфигурации приложения. Определяет пространство имен конфига.

	EnvDev Environment = "dev" // Среда разработки.
)

// Environment тип среды выполнения приложения.
type Environment string

func (e Environment) String() string {
	return string(e)
}
