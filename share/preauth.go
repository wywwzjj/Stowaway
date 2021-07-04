package share

import (
	"errors"
	"io"
	"net"
	"time"

	"Stowaway/utils"
)

func ActivePreAuth(conn net.Conn, key string) error {
	var NOT_VALID = errors.New("Not valid secret,check the secret!")

	defer conn.SetReadDeadline(time.Time{})
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	secret := utils.GetStringMd5(key)
	conn.Write([]byte(secret[:16]))

	buffer := make([]byte, 16)
	count, err := io.ReadFull(conn, buffer)

	if timeoutErr, ok := err.(net.Error); ok && timeoutErr.Timeout() {
		conn.Close()
		return NOT_VALID
	}

	if err != nil {
		conn.Close()
		return NOT_VALID
	}

	if string(buffer[:count]) == secret[:16] {
		return nil
	}

	conn.Close()

	return NOT_VALID
}

func PassivePreAuth(conn net.Conn, key string) error {
	var NOT_VALID = errors.New("Not valid secret,check the secret!")

	defer conn.SetReadDeadline(time.Time{})
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	secret := utils.GetStringMd5(key)

	buffer := make([]byte, 16)
	count, err := io.ReadFull(conn, buffer)

	if timeoutErr, ok := err.(net.Error); ok && timeoutErr.Timeout() {
		conn.Close()
		return NOT_VALID
	}

	if err != nil {
		conn.Close()
		return NOT_VALID
	}

	if string(buffer[:count]) == secret[:16] {
		conn.Write([]byte(secret[:16]))
		return nil
	}

	conn.Close()

	return NOT_VALID
}
