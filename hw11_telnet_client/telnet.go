package main

import (
	"errors"
	"io"
	"net"
	"time"
)

var ErrClientWithoutConnection = errors.New("connection of the client is empty")

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type SimpleTelnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func (c *SimpleTelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *SimpleTelnetClient) Close() error {
	return c.in.Close()
}

func (c *SimpleTelnetClient) Send() error {
	if c.conn == nil {
		return ErrClientWithoutConnection
	}
	if _, err := io.Copy(c.conn, c.in); err != nil {
		return err
	}
	return nil
}

func (c *SimpleTelnetClient) Receive() error {
	if c.conn == nil {
		return ErrClientWithoutConnection
	}
	if _, err := io.Copy(c.out, c.conn); err != nil {
		return err
	}
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &SimpleTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
