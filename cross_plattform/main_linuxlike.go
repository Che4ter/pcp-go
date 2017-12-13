package main


// +build linux darwin netbsd openbsd freebsd
func GetCommand() []string {
	return []string{"ifconfig", "-a"}
}