package ipisp

import (
	"bufio"
	"golang.org/x/xerrors"
	"net"
	"sync"
	"time"
)

// BulkClient may be used to lookup a large amount of IPs or ASNs in a quick burst.
// Calls to WriteIP and WriteASN do not return errors since they write the requests to a buffer.
// The buffer is dispatched on the next call to Read.
type BulkClient struct {
	mu sync.Mutex

	conn net.Conn
	w    *bufio.Writer
	sc   *bufio.Scanner
}

const (
	cymruNetcatAddress = "whois.cymru.com:43"
)

var bulkEOL = []byte("\r\n")

// DialBulkClient opens up a WHOIS connection to the service.
func DialBulkClient() (*BulkClient, error) {
	conn, err := net.DialTimeout("tcp", cymruNetcatAddress, Timeout)
	if err != nil {
		return nil, xerrors.Errorf("dial %s: %v", cymruNetcatAddress, err)
	}

	_ = conn.SetDeadline(time.Now().Add(time.Second * 10))

	bw := bufio.NewWriter(conn)
	bw.Write([]byte("begin"))
	bw.Write(bulkEOL)
	bw.Write([]byte("verbose"))
	bw.Write(bulkEOL)

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
		conn: conn,
		w:    bw,
		sc:   sc,
	}, nil
}

func (bc *BulkClient) WriteIP(ip string) {
	bc.w.WriteString(ip)
	bc.w.Write(bulkEOL)
}

func (bc *BulkClient) WriteASN(asn int) {
	bc.w.WriteString(ASN(asn).String())
	bc.w.Write(bulkEOL)
}

// Read returns a single response from the interface, whether or not the connection is still valid, and
// the error.
func (bc *BulkClient) Read() (*Response, bool, error) {
	bc.w.Flush()
}

// Close gracefully terminates the client.
func (bc *BulkClient) Close() error {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// These are courtesy messages to indicate that the client disconnected normally. Thus, their errors are not
	// important.

	bc.conn.SetWriteDeadline(time.Now().Add(time.Second))
	bc.conn.Write([]byte("end"))
	bc.conn.Write(bulkEOL)
	return bc.conn.Close()
}
