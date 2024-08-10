package model

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names
import (
	"net"
	"time"

	"github.com/miekg/dns"
)

type ResponseType int

func (t ResponseType) ToExtendedErrorCode() uint16 {
	switch t {
	case ResponseTypeRESOLVED:
		return dns.ExtendedErrorCodeOther
	case ResponseTypeCACHED:
		return dns.ExtendedErrorCodeCachedError
	case ResponseTypeCONDITIONAL:
		return dns.ExtendedErrorCodeForgedAnswer
	case ResponseTypeCUSTOMDNS:
		return dns.ExtendedErrorCodeForgedAnswer
	case ResponseTypeHOSTSFILE:
		return dns.ExtendedErrorCodeForgedAnswer
	case ResponseTypeNOTFQDN:
		return dns.ExtendedErrorCodeBlocked
	case ResponseTypeBLOCKED:
		return dns.ExtendedErrorCodeBlocked
	case ResponseTypeFILTERED:
		return dns.ExtendedErrorCodeFiltered
	case ResponseTypeSPECIAL:
		return dns.ExtendedErrorCodeFiltered
	default:
		return dns.ExtendedErrorCodeOther
	}
}

// Response represents the response of a DNS query
type Response struct {
	Res    *dns.Msg
	Reason string
	RType  ResponseType
}

// RequestProtocol represents the server protocol ENUM(
// TCP // is the TCP protocol
// UDP // is the UDP protocol
// )
type RequestProtocol uint8

// Request represents client's DNS request
type Request struct {
	ClientIP        net.IP
	RequestClientID string
	Protocol        RequestProtocol
	ClientNames     []string
	Req             *dns.Msg
	RequestTS       time.Time
}
