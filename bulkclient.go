package ipisp

import (
	"bufio"
	"context"
	"golang.org/x/xerrors"
	"net"
	"sync"
	"time"
)

type BulkClient struct {
	mu sync.Mutex

	c  net.Conn
	w  *bufio.Writer
	sc *bufio.Scanner
}

const (
	cymruNetcatAddress = "whois.cymru.com:43"
)

// DialBulkClient opens up a WHOIS connection to the service.
func DialBulkClient(ctx context.Context) (*BulkClient, error) {
	conn, err := net.DialTimeout("tcp", cymruNetcatAddress, Timeout)
	if err != nil {
		return nil, xerrors.Errorf("dial %s: %v", cymruNetcatAddress, err)
	}

	_ = conn.SetDeadline(time.Now().Add(time.Second * 10))

	bw := bufio.NewWriter(conn)
	bw.Write([]byte("begin"))
	bw.Write(ncEOL)
	bw.Write([]byte("verbose"))
	bw.Write(ncEOL)

	err = bw.Flush()
	if err != nil {
		return nil, xerrors.Errorf("write begin message: %v", err)
	}

	sc := bufio.NewScanner(conn)

	// Discard first hello line
	sc.Scan()
	if sc.Err() != nil {
		_ = conn.Close()
		return nil, xerrors.Errorf("discard first line: %v", sc.Err())
	}

	return &BulkClient{
		c:  conn,
		w:  bw,
		sc: sc,
	}, nil

}
