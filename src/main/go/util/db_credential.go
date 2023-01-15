package util

type DbCredential struct {
	HostName  string
	UserName  string
	Password  string
	DefaultDB any
}

func CreateDbCredential(hostName, userName, password, defaultDB string) *DbCredential {
	return &DbCredential{
		hostName,
		userName,
		password,
		defaultDB}
}
