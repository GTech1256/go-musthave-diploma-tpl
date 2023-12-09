package config

import (
	"errors"
	"flag"
	"os"
)

type Config struct {
	// Port - Флаг -a=<ЗНАЧЕНИЕ> отвечает за адрес эндпоинта HTTP-сервера (по умолчанию localhost:8080).
	Port *string

	DatabaseURI *string

	// AccrualSystemAddress - Флаг -k=<ЗНАЧЕНИЕ> При наличии ключа агент должен вычислять хеш и передавать в HTTP-заголовке запроса с именем HashSHA256.
	AccrualSystemAddress *string
}

const EmptyStringKey = ""

var ErrRequireVariable = errors.New("переменная окружения обязательна")

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	var (
		// Hack для тестирования
		command                                                 = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		port                                                    = command.String("a", ":8081", "address and port to run memory")
		portEnv, portEnvPresent                                 = os.LookupEnv("RUN_ADDRESS")
		databaseURI                                             = command.String("d", EmptyStringKey, "the path to database connection")
		databaseURIEnv, databaseURIEnvPresent                   = os.LookupEnv("DATABASE_URI")
		accrualSystemAddress                                    = command.String("r", EmptyStringKey, "адрес системы расчёта начислений")
		accrualSystemAddressEnv, accrualSystemAddressEnvPresent = os.LookupEnv("ACCRUAL_SYSTEM_ADDRESS")
	)

	c.Port = port
	if portEnvPresent {
		c.Port = &portEnv
	}
	if c.Port == nil {
		return ErrRequireVariable
	}

	c.DatabaseURI = databaseURI
	if databaseURIEnvPresent {
		c.DatabaseURI = &databaseURIEnv
	}
	if c.DatabaseURI == nil {
		return ErrRequireVariable
	}

	c.AccrualSystemAddress = accrualSystemAddress
	if accrualSystemAddressEnvPresent {
		c.AccrualSystemAddress = &accrualSystemAddressEnv
	}
	if c.AccrualSystemAddress == nil {
		return ErrRequireVariable
	}

	// Тесты запускают несколько раз метод Load.
	// А несколько раз flag.Parse() нельзя вызывать
	// Из-за этого хак с командными флагами
	command.Parse(os.Args[1:])

	return nil
}
