package config

import (
	"fmt"
	"strings"
)

const (
	// IPVersionDual is a IPVersion of type Dual.
	// IPv4 and IPv6
	IPVersionDual IPVersion = iota
	// IPVersionV4 is a IPVersion of type V4.
	// IPv4 only
	IPVersionV4
	// IPVersionV6 is a IPVersion of type V6.
	// IPv6 only
	IPVersionV6
)

var ErrInvalidIPVersion = fmt.Errorf("not a valid IPVersion, try [%s]", strings.Join(_IPVersionNames, ", "))

const _IPVersionName = "dualv4v6"

var _IPVersionNames = []string{
	_IPVersionName[0:4],
	_IPVersionName[4:6],
	_IPVersionName[6:8],
}

// IPVersionNames returns a list of possible string values of IPVersion.
func IPVersionNames() []string {
	tmp := make([]string, len(_IPVersionNames))
	copy(tmp, _IPVersionNames)
	return tmp
}

// IPVersionValues returns a list of the values for IPVersion
func IPVersionValues() []IPVersion {
	return []IPVersion{
		IPVersionDual,
		IPVersionV4,
		IPVersionV6,
	}
}

var _IPVersionMap = map[IPVersion]string{
	IPVersionDual: _IPVersionName[0:4],
	IPVersionV4:   _IPVersionName[4:6],
	IPVersionV6:   _IPVersionName[6:8],
}

// String implements the Stringer interface.
func (x IPVersion) String() string {
	if str, ok := _IPVersionMap[x]; ok {
		return str
	}
	return fmt.Sprintf("IPVersion(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x IPVersion) IsValid() bool {
	_, ok := _IPVersionMap[x]
	return ok
}

var _IPVersionValue = map[string]IPVersion{
	_IPVersionName[0:4]: IPVersionDual,
	_IPVersionName[4:6]: IPVersionV4,
	_IPVersionName[6:8]: IPVersionV6,
}

// ParseIPVersion attempts to convert a string to a IPVersion.
func ParseIPVersion(name string) (IPVersion, error) {
	if x, ok := _IPVersionValue[name]; ok {
		return x, nil
	}
	return IPVersion(0), fmt.Errorf("%s is %w", name, ErrInvalidIPVersion)
}

// MarshalText implements the text marshaller method.
func (x IPVersion) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *IPVersion) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseIPVersion(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// InitStrategyBlocking is a InitStrategy of type Blocking.
	// synchronously download blocking lists on startup
	InitStrategyBlocking InitStrategy = iota
	// InitStrategyFailOnError is a InitStrategy of type FailOnError.
	// synchronously download blocking lists on startup and shutdown on error
	InitStrategyFailOnError
	// InitStrategyFast is a InitStrategy of type Fast.
	// asyncronously download blocking lists on startup
	InitStrategyFast
)

var ErrInvalidInitStrategy = fmt.Errorf("not a valid InitStrategy, try [%s]", strings.Join(_InitStrategyNames, ", "))

const _InitStrategyName = "blockingfailOnErrorfast"

var _InitStrategyNames = []string{
	_InitStrategyName[0:8],
	_InitStrategyName[8:19],
	_InitStrategyName[19:23],
}

// InitStrategyNames returns a list of possible string values of InitStrategy.
func InitStrategyNames() []string {
	tmp := make([]string, len(_InitStrategyNames))
	copy(tmp, _InitStrategyNames)
	return tmp
}

// InitStrategyValues returns a list of the values for InitStrategy
func InitStrategyValues() []InitStrategy {
	return []InitStrategy{
		InitStrategyBlocking,
		InitStrategyFailOnError,
		InitStrategyFast,
	}
}

var _InitStrategyMap = map[InitStrategy]string{
	InitStrategyBlocking:    _InitStrategyName[0:8],
	InitStrategyFailOnError: _InitStrategyName[8:19],
	InitStrategyFast:        _InitStrategyName[19:23],
}

// String implements the Stringer interface.
func (x InitStrategy) String() string {
	if str, ok := _InitStrategyMap[x]; ok {
		return str
	}
	return fmt.Sprintf("InitStrategy(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x InitStrategy) IsValid() bool {
	_, ok := _InitStrategyMap[x]
	return ok
}

var _InitStrategyValue = map[string]InitStrategy{
	_InitStrategyName[0:8]:   InitStrategyBlocking,
	_InitStrategyName[8:19]:  InitStrategyFailOnError,
	_InitStrategyName[19:23]: InitStrategyFast,
}

// ParseInitStrategy attempts to convert a string to a InitStrategy.
func ParseInitStrategy(name string) (InitStrategy, error) {
	if x, ok := _InitStrategyValue[name]; ok {
		return x, nil
	}
	return InitStrategy(0), fmt.Errorf("%s is %w", name, ErrInvalidInitStrategy)
}

// MarshalText implements the text marshaller method.
func (x InitStrategy) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *InitStrategy) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseInitStrategy(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// NetProtocolTcpUdp is a NetProtocol of type Tcp+Udp.
	// TCP and UDP protocols
	NetProtocolTcpUdp NetProtocol = iota
	// NetProtocolTcpTls is a NetProtocol of type Tcp-Tls.
	// TCP-TLS protocol
	NetProtocolTcpTls
	// NetProtocolHttps is a NetProtocol of type Https.
	// HTTPS protocol
	NetProtocolHttps
)

var ErrInvalidNetProtocol = fmt.Errorf("not a valid NetProtocol, try [%s]", strings.Join(_NetProtocolNames, ", "))

const _NetProtocolName = "tcp+udptcp-tlshttps"

var _NetProtocolNames = []string{
	_NetProtocolName[0:7],
	_NetProtocolName[7:14],
	_NetProtocolName[14:19],
}

// NetProtocolNames returns a list of possible string values of NetProtocol.
func NetProtocolNames() []string {
	tmp := make([]string, len(_NetProtocolNames))
	copy(tmp, _NetProtocolNames)
	return tmp
}

// NetProtocolValues returns a list of the values for NetProtocol
func NetProtocolValues() []NetProtocol {
	return []NetProtocol{
		NetProtocolTcpUdp,
		NetProtocolTcpTls,
		NetProtocolHttps,
	}
}

var _NetProtocolMap = map[NetProtocol]string{
	NetProtocolTcpUdp: _NetProtocolName[0:7],
	NetProtocolTcpTls: _NetProtocolName[7:14],
	NetProtocolHttps:  _NetProtocolName[14:19],
}

// String implements the Stringer interface.
func (x NetProtocol) String() string {
	if str, ok := _NetProtocolMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NetProtocol(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x NetProtocol) IsValid() bool {
	_, ok := _NetProtocolMap[x]
	return ok
}

var _NetProtocolValue = map[string]NetProtocol{
	_NetProtocolName[0:7]:   NetProtocolTcpUdp,
	_NetProtocolName[7:14]:  NetProtocolTcpTls,
	_NetProtocolName[14:19]: NetProtocolHttps,
}

// ParseNetProtocol attempts to convert a string to a NetProtocol.
func ParseNetProtocol(name string) (NetProtocol, error) {
	if x, ok := _NetProtocolValue[name]; ok {
		return x, nil
	}
	return NetProtocol(0), fmt.Errorf("%s is %w", name, ErrInvalidNetProtocol)
}

// MarshalText implements the text marshaller method.
func (x NetProtocol) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *NetProtocol) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseNetProtocol(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// QueryLogFieldClientIP is a QueryLogField of type clientIP.
	QueryLogFieldClientIP QueryLogField = "clientIP"
	// QueryLogFieldClientName is a QueryLogField of type clientName.
	QueryLogFieldClientName QueryLogField = "clientName"
	// QueryLogFieldResponseReason is a QueryLogField of type responseReason.
	QueryLogFieldResponseReason QueryLogField = "responseReason"
	// QueryLogFieldResponseAnswer is a QueryLogField of type responseAnswer.
	QueryLogFieldResponseAnswer QueryLogField = "responseAnswer"
	// QueryLogFieldQuestion is a QueryLogField of type question.
	QueryLogFieldQuestion QueryLogField = "question"
	// QueryLogFieldDuration is a QueryLogField of type duration.
	QueryLogFieldDuration QueryLogField = "duration"
)

var ErrInvalidQueryLogField = fmt.Errorf("not a valid QueryLogField, try [%s]", strings.Join(_QueryLogFieldNames, ", "))

var _QueryLogFieldNames = []string{
	string(QueryLogFieldClientIP),
	string(QueryLogFieldClientName),
	string(QueryLogFieldResponseReason),
	string(QueryLogFieldResponseAnswer),
	string(QueryLogFieldQuestion),
	string(QueryLogFieldDuration),
}

// QueryLogFieldNames returns a list of possible string values of QueryLogField.
func QueryLogFieldNames() []string {
	tmp := make([]string, len(_QueryLogFieldNames))
	copy(tmp, _QueryLogFieldNames)
	return tmp
}

// QueryLogFieldValues returns a list of the values for QueryLogField
func QueryLogFieldValues() []QueryLogField {
	return []QueryLogField{
		QueryLogFieldClientIP,
		QueryLogFieldClientName,
		QueryLogFieldResponseReason,
		QueryLogFieldResponseAnswer,
		QueryLogFieldQuestion,
		QueryLogFieldDuration,
	}
}

// String implements the Stringer interface.
func (x QueryLogField) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x QueryLogField) IsValid() bool {
	_, err := ParseQueryLogField(string(x))
	return err == nil
}

var _QueryLogFieldValue = map[string]QueryLogField{
	"clientIP":       QueryLogFieldClientIP,
	"clientName":     QueryLogFieldClientName,
	"responseReason": QueryLogFieldResponseReason,
	"responseAnswer": QueryLogFieldResponseAnswer,
	"question":       QueryLogFieldQuestion,
	"duration":       QueryLogFieldDuration,
}

// ParseQueryLogField attempts to convert a string to a QueryLogField.
func ParseQueryLogField(name string) (QueryLogField, error) {
	if x, ok := _QueryLogFieldValue[name]; ok {
		return x, nil
	}
	return QueryLogField(""), fmt.Errorf("%s is %w", name, ErrInvalidQueryLogField)
}

// MarshalText implements the text marshaller method.
func (x QueryLogField) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *QueryLogField) UnmarshalText(text []byte) error {
	tmp, err := ParseQueryLogField(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// QueryLogTypeConsole is a QueryLogType of type Console.
	// use logger as fallback
	QueryLogTypeConsole QueryLogType = iota
	// QueryLogTypeNone is a QueryLogType of type None.
	// no logging
	QueryLogTypeNone
	// QueryLogTypeMysql is a QueryLogType of type Mysql.
	// MySQL or MariaDB database
	QueryLogTypeMysql
	// QueryLogTypePostgresql is a QueryLogType of type Postgresql.
	// PostgreSQL database
	QueryLogTypePostgresql
	// QueryLogTypeCsv is a QueryLogType of type Csv.
	// CSV file per day
	QueryLogTypeCsv
	// QueryLogTypeCsvClient is a QueryLogType of type Csv-Client.
	// CSV file per day and client
	QueryLogTypeCsvClient
	// QueryLogTypeTimescale is a QueryLogType of type Timescale.
	// Timescale database
	QueryLogTypeTimescale
)

var ErrInvalidQueryLogType = fmt.Errorf("not a valid QueryLogType, try [%s]", strings.Join(_QueryLogTypeNames, ", "))

const _QueryLogTypeName = "consolenonemysqlpostgresqlcsvcsv-clienttimescale"

var _QueryLogTypeNames = []string{
	_QueryLogTypeName[0:7],
	_QueryLogTypeName[7:11],
	_QueryLogTypeName[11:16],
	_QueryLogTypeName[16:26],
	_QueryLogTypeName[26:29],
	_QueryLogTypeName[29:39],
	_QueryLogTypeName[39:48],
}

// QueryLogTypeNames returns a list of possible string values of QueryLogType.
func QueryLogTypeNames() []string {
	tmp := make([]string, len(_QueryLogTypeNames))
	copy(tmp, _QueryLogTypeNames)
	return tmp
}

// QueryLogTypeValues returns a list of the values for QueryLogType
func QueryLogTypeValues() []QueryLogType {
	return []QueryLogType{
		QueryLogTypeConsole,
		QueryLogTypeNone,
		QueryLogTypeMysql,
		QueryLogTypePostgresql,
		QueryLogTypeCsv,
		QueryLogTypeCsvClient,
		QueryLogTypeTimescale,
	}
}

var _QueryLogTypeMap = map[QueryLogType]string{
	QueryLogTypeConsole:    _QueryLogTypeName[0:7],
	QueryLogTypeNone:       _QueryLogTypeName[7:11],
	QueryLogTypeMysql:      _QueryLogTypeName[11:16],
	QueryLogTypePostgresql: _QueryLogTypeName[16:26],
	QueryLogTypeCsv:        _QueryLogTypeName[26:29],
	QueryLogTypeCsvClient:  _QueryLogTypeName[29:39],
	QueryLogTypeTimescale:  _QueryLogTypeName[39:48],
}

// String implements the Stringer interface.
func (x QueryLogType) String() string {
	if str, ok := _QueryLogTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("QueryLogType(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x QueryLogType) IsValid() bool {
	_, ok := _QueryLogTypeMap[x]
	return ok
}

var _QueryLogTypeValue = map[string]QueryLogType{
	_QueryLogTypeName[0:7]:   QueryLogTypeConsole,
	_QueryLogTypeName[7:11]:  QueryLogTypeNone,
	_QueryLogTypeName[11:16]: QueryLogTypeMysql,
	_QueryLogTypeName[16:26]: QueryLogTypePostgresql,
	_QueryLogTypeName[26:29]: QueryLogTypeCsv,
	_QueryLogTypeName[29:39]: QueryLogTypeCsvClient,
	_QueryLogTypeName[39:48]: QueryLogTypeTimescale,
}

// ParseQueryLogType attempts to convert a string to a QueryLogType.
func ParseQueryLogType(name string) (QueryLogType, error) {
	if x, ok := _QueryLogTypeValue[name]; ok {
		return x, nil
	}
	return QueryLogType(0), fmt.Errorf("%s is %w", name, ErrInvalidQueryLogType)
}

// MarshalText implements the text marshaller method.
func (x QueryLogType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *QueryLogType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseQueryLogType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// TLSVersion10 is a TLSVersion of type 1.0.
	TLSVersion10 TLSVersion = iota + 769
	// TLSVersion11 is a TLSVersion of type 1.1.
	TLSVersion11
	// TLSVersion12 is a TLSVersion of type 1.2.
	TLSVersion12
	// TLSVersion13 is a TLSVersion of type 1.3.
	TLSVersion13
)

var ErrInvalidTLSVersion = fmt.Errorf("not a valid TLSVersion, try [%s]", strings.Join(_TLSVersionNames, ", "))

const _TLSVersionName = "1.01.11.21.3"

var _TLSVersionNames = []string{
	_TLSVersionName[0:3],
	_TLSVersionName[3:6],
	_TLSVersionName[6:9],
	_TLSVersionName[9:12],
}

// TLSVersionNames returns a list of possible string values of TLSVersion.
func TLSVersionNames() []string {
	tmp := make([]string, len(_TLSVersionNames))
	copy(tmp, _TLSVersionNames)
	return tmp
}

// TLSVersionValues returns a list of the values for TLSVersion
func TLSVersionValues() []TLSVersion {
	return []TLSVersion{
		TLSVersion10,
		TLSVersion11,
		TLSVersion12,
		TLSVersion13,
	}
}

var _TLSVersionMap = map[TLSVersion]string{
	TLSVersion10: _TLSVersionName[0:3],
	TLSVersion11: _TLSVersionName[3:6],
	TLSVersion12: _TLSVersionName[6:9],
	TLSVersion13: _TLSVersionName[9:12],
}

// String implements the Stringer interface.
func (x TLSVersion) String() string {
	if str, ok := _TLSVersionMap[x]; ok {
		return str
	}
	return fmt.Sprintf("TLSVersion(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x TLSVersion) IsValid() bool {
	_, ok := _TLSVersionMap[x]
	return ok
}

var _TLSVersionValue = map[string]TLSVersion{
	_TLSVersionName[0:3]:  TLSVersion10,
	_TLSVersionName[3:6]:  TLSVersion11,
	_TLSVersionName[6:9]:  TLSVersion12,
	_TLSVersionName[9:12]: TLSVersion13,
}

// ParseTLSVersion attempts to convert a string to a TLSVersion.
func ParseTLSVersion(name string) (TLSVersion, error) {
	if x, ok := _TLSVersionValue[name]; ok {
		return x, nil
	}
	return TLSVersion(0), fmt.Errorf("%s is %w", name, ErrInvalidTLSVersion)
}

// MarshalText implements the text marshaller method.
func (x TLSVersion) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *TLSVersion) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseTLSVersion(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// UpstreamStrategyParallelBest is a UpstreamStrategy of type Parallel_best.
	UpstreamStrategyParallelBest UpstreamStrategy = iota
	// UpstreamStrategyStrict is a UpstreamStrategy of type Strict.
	UpstreamStrategyStrict
	// UpstreamStrategyRandom is a UpstreamStrategy of type Random.
	UpstreamStrategyRandom
)

var ErrInvalidUpstreamStrategy = fmt.Errorf("not a valid UpstreamStrategy, try [%s]", strings.Join(_UpstreamStrategyNames, ", "))

const _UpstreamStrategyName = "parallel_beststrictrandom"

var _UpstreamStrategyNames = []string{
	_UpstreamStrategyName[0:13],
	_UpstreamStrategyName[13:19],
	_UpstreamStrategyName[19:25],
}

// UpstreamStrategyNames returns a list of possible string values of UpstreamStrategy.
func UpstreamStrategyNames() []string {
	tmp := make([]string, len(_UpstreamStrategyNames))
	copy(tmp, _UpstreamStrategyNames)
	return tmp
}

// UpstreamStrategyValues returns a list of the values for UpstreamStrategy
func UpstreamStrategyValues() []UpstreamStrategy {
	return []UpstreamStrategy{
		UpstreamStrategyParallelBest,
		UpstreamStrategyStrict,
		UpstreamStrategyRandom,
	}
}

var _UpstreamStrategyMap = map[UpstreamStrategy]string{
	UpstreamStrategyParallelBest: _UpstreamStrategyName[0:13],
	UpstreamStrategyStrict:       _UpstreamStrategyName[13:19],
	UpstreamStrategyRandom:       _UpstreamStrategyName[19:25],
}

// String implements the Stringer interface.
func (x UpstreamStrategy) String() string {
	if str, ok := _UpstreamStrategyMap[x]; ok {
		return str
	}
	return fmt.Sprintf("UpstreamStrategy(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x UpstreamStrategy) IsValid() bool {
	_, ok := _UpstreamStrategyMap[x]
	return ok
}

var _UpstreamStrategyValue = map[string]UpstreamStrategy{
	_UpstreamStrategyName[0:13]:  UpstreamStrategyParallelBest,
	_UpstreamStrategyName[13:19]: UpstreamStrategyStrict,
	_UpstreamStrategyName[19:25]: UpstreamStrategyRandom,
}

// ParseUpstreamStrategy attempts to convert a string to a UpstreamStrategy.
func ParseUpstreamStrategy(name string) (UpstreamStrategy, error) {
	if x, ok := _UpstreamStrategyValue[name]; ok {
		return x, nil
	}
	return UpstreamStrategy(0), fmt.Errorf("%s is %w", name, ErrInvalidUpstreamStrategy)
}

// MarshalText implements the text marshaller method.
func (x UpstreamStrategy) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *UpstreamStrategy) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseUpstreamStrategy(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
