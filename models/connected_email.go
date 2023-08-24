package models

type ConnectedEmail struct {
	ID           int
	Email        string
	Pop3pHost    string
	Pop3Port     int
	Pop3Username string
	Pop3Password string
}

func (c *ConnectedEmail) AuthPop3() error {
	return nil
}

func (c *ConnectedEmail) ReadEmailsPop3() error {
	return nil
}
