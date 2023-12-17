package helper

import "net"

func GetFreePort() uint {
	var l *net.TCPListener
	for {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return GetFreePort()
		} else {
			l, err = net.ListenTCP("tcp", addr)
			if err != nil {
				return GetFreePort()
			}
			defer l.Close()

			break
		}
	}

	return uint(l.Addr().(*net.TCPAddr).Port)
}
