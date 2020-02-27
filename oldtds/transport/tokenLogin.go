package transport

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/SAP/go-ase/libase/libdsn"
)

const (
	TDS_MAXNAME   = 30
	TDS_NETBUF    = 4
	TDS_RPLEN     = 255
	TDS_VERSIZE   = 4
	TDS_PROGNLEN  = 10
	TDS_OLDSECURE = 2
	TDS_HA        = 6
	TDS_SECURE    = 2
	TDS_PKTLEN    = 6
)

type TokenLoginConfig struct {
	DSN       *libdsn.DsnInfo
	Hostname  string
	LHostProc string
	AppName   string
	ServName  string

	// TODO: Remove LibraryName/-Version in favour of a const?
	LibraryName    string
	LibraryVersion string

	Language string
	CharSet  string

	PacketSize uint16
}

func NewTokenLoginConfig(dsn *libdsn.DsnInfo) *TokenLoginConfig {
	conf := &TokenLoginConfig{}

	conf.DSN = dsn

	hostname, err := os.Hostname()
	if err == nil {
		conf.Hostname = hostname
	}

	conf.LHostProc = strconv.Itoa(os.Getpid())

	return conf
}

func (conf TokenLoginConfig) String() string {
	return fmt.Sprintf("%#v", conf)
}

var _ Message = (*TokenLogin)(nil)

type TokenLogin struct {
	Config *TokenLoginConfig

	buf        *bytes.Buffer
	tokenLogin *tokenLogin
}

func NewTokenLogin(config *TokenLoginConfig) (*TokenLogin, error) {
	token := &TokenLogin{
		Config: config,
	}

	return token, nil
}

func (token TokenLogin) String() string {
	return token.Config.String()
}

func (token *TokenLogin) Configure() error {
	token.buf = &bytes.Buffer{}

	// No error checking requires since bytes.Buffer.Write* methods
	// always return a nil error.

	// lhostname, lhostlen
	err := token.writeString(token.Config.Hostname, TDS_MAXNAME)
	if err != nil {
		return fmt.Errorf("error writing hostname: %v", err)
	}

	// lusername, lusernlen
	token.writeString(token.Config.DSN.Username, TDS_MAXNAME)

	// lpw, lpwnlen
	token.writeString(token.Config.DSN.Password, TDS_MAXNAME)

	// lhostproc, lhplen
	token.writeString(token.Config.LHostProc, TDS_MAXNAME)

	// lint2 -> little endian
	token.buf.WriteByte(0x2)
	// lint4 -> little endian
	token.buf.WriteByte(0x0)
	// lchar -> ASCII
	token.buf.WriteByte(0x6)
	// lflt -> little endian
	token.buf.WriteByte(0x4)
	// ldate -> little endian
	token.buf.WriteByte(0x8)

	// lusedb
	token.buf.WriteByte(0x0)
	// ldmpld
	token.buf.WriteByte(0x0)

	// only relevant for server-server comm
	// linterfacespare
	token.buf.WriteByte(0x0)
	// ltype
	token.buf.WriteByte(0x0)

	// deprecated
	// lbufsize
	token.writeString("", TDS_NETBUF)

	// lspare
	token.writeString("", 3)

	// lappname, lappnlen
	token.writeString(token.Config.AppName, TDS_MAXNAME)

	// lservname, lservnlen
	token.writeString(token.Config.ServName, TDS_MAXNAME)

	// only relevant for server-server comm
	// lrempw, lrempwlen
	token.writeString("", TDS_RPLEN)

	// ltds
	token.buf.Write([]byte{0x5, 0x0, 0x0, 0x0})

	// lprogname, lprognlen
	token.writeString(token.Config.LibraryName, TDS_PROGNLEN)

	// lprogvers
	// TODO TDS_VERSIZE unknown, guessing 4 based on docs
	token.buf.Write([]byte{0x0, 0x0, 0x0, 0x0})

	// lnoshort - do not convert short data types
	token.buf.WriteByte(0x0)

	// lflt4 little endian
	token.buf.WriteByte(0x12)
	// ldate4 little endian
	token.buf.WriteByte(0x16)

	// llanguage, llanglen
	token.writeString(token.Config.Language, TDS_MAXNAME)

	// lsetlang - notify of language changes
	token.buf.WriteByte(0x1)

	// loldsecure - deprecated
	token.buf.Write(make([]byte, TDS_OLDSECURE))
	// lseclogin - deprecated
	token.buf.WriteByte(0x0)
	// lsecbulk - deprecated
	token.buf.WriteByte(0x1)

	// lhalogin
	// TODO - values need to be determined by config to allow for
	// failover reconnects in clusters
	token.buf.WriteByte(0x1)
	// lhasessionid
	// TODO session id for HA failover, find out if this needs to be
	// user set or retrieved from the server
	token.buf.Write(make([]byte, TDS_HA))

	// lsecspare - unused
	// TODO TDS_SECURE unknown
	token.buf.Write(make([]byte, TDS_SECURE))

	// lcharset, lcharsetlen
	token.writeString(token.Config.CharSet, TDS_MAXNAME)

	// lsetcharset - notify of charset changes
	token.buf.WriteByte(0x1)

	// lpacketsize - 256 to 65535 bytes
	// TODO Choose default packet size
	if token.Config.PacketSize < 256 {
		return fmt.Errorf("packet size too low, must be at least 256 bytes")
	}
	token.writeString(strconv.Itoa(int(token.Config.PacketSize)), TDS_PKTLEN)

	// ldummy - apparently unused
	token.buf.Write(make([]byte, 4))

	return nil
}

func (token TokenLogin) writeString(s string, padTo int) error {
	token.buf.WriteString(s)
	token.buf.WriteString(string(make([]rune, padTo-len(s))))
	token.buf.WriteByte(byte(len(s)))

	return nil
}

func (token TokenLogin) Packets() chan Packet {
	ch := make(chan Packet, 10)

	header := MessageHeader{
		MsgType:  TDS_BUF_LOGIN,
		Status:   0,
		Length:   MsgLength,
		Channel:  0,
		PacketNr: 0,
		Window:   0,
	}

	go func() {
		for token.buf.Len() > 0 {

			// If the remaining buffer is smaller than the allowed
			// message length set length according to buffer length
			if token.buf.Len() < MsgBodyLength {
				header.Length = uint16(token.buf.Len() + MsgHeaderLength)
				// Last packet, send EOM flag
				header.Status = TDS_BUFSTAT_EOM
			}

			// // Increase PacketNr
			// header.PacketNr++

			p := Packet{
				Header: header,
				Data:   make([]byte, header.Length-MsgHeaderLength),
			}

			token.buf.Read(p.Data)

			ch <- p
		}

		close(ch)
	}()

	return ch
}

type tokenLogin struct {
	Lhostname       [TDS_MAXNAME]byte
	Lhostnlen       byte
	Lusername       [TDS_MAXNAME]byte
	Lusernlen       byte
	Lpw             [TDS_MAXNAME]byte
	Lpwnlen         byte
	Lhostproc       [TDS_MAXNAME]byte
	Lhplen          byte
	Lint2           byte
	Lint4           byte
	Lchar           byte
	Lflt            byte
	Ldate           byte
	Lusedb          byte
	Ldmpld          byte
	Linterfacespare byte
	Ltype           byte
	Lbufsize        [TDS_NETBUF]byte
	Spare           [3]byte
	Lappname        [TDS_MAXNAME]byte
	Lappnlen        byte
	Lservname       [TDS_MAXNAME]byte
	Lservnlen       byte
	Lrempw          [TDS_RPLEN]byte
	Lrempwlen       byte
	Ltds            [TDS_VERSIZE]byte
	Lprogname       [TDS_PROGNLEN]byte
	Lprognlen       byte
	Lprogvers       [TDS_VERSIZE]byte
	Lnoshort        byte
	Lflt4           byte
	Ldate4          byte
	Llanguage       [TDS_MAXNAME]byte
	Llanglen        byte
	Lsetlang        byte
	Loldsecure      [TDS_OLDSECURE]byte
	Lseclogin       byte
	Lsecbulk        byte
	Lhalogin        byte
	Lhasessionid    [TDS_HA]byte
	Lsecspare       [TDS_SECURE]byte
	Lcharset        [TDS_MAXNAME]byte
	Lcharsetlen     byte
	Lsetcharset     byte
	Lpacketsize     [TDS_PKTLEN]byte
	Lpacksetsizelen byte
	Ldummy          [4]byte
}
