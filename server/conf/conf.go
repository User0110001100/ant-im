package conf

import (
	"encoding/json"
	"os"
)

type Conf struct {
	Address   string
	Port      int
	DBAddress string
	DBPort    int
	DBURI     string
}

func (c *Conf) GetServerConf() (string, int) {
	return c.Address, c.Port
}

// Return address & port & uri
func (c *Conf) GetDBConf() (string, int, string) {
	return c.DBAddress, c.DBPort, c.DBURI
}

func MakeConf(filename string) *Conf {
	var c Conf

	info, err := os.Stat(filename)

	if err != nil {
		return nil
	}

	f, err := os.Open(filename)

	if err != nil {
		return nil
	}

	defer f.Close()

	data := make([]byte, info.Size())
	_, err = f.Read(data)

	if err != nil {
		return nil
	}

	json.Unmarshal(data, &c)

	if err != nil {
		return nil
	}

	return &c
}
