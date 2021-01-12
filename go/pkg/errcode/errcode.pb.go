// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errcode.proto

package errcode

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	Undefined                                  ErrCode = 0
	TODO                                       ErrCode = 666
	ErrNotImplemented                          ErrCode = 777
	ErrInternal                                ErrCode = 888
	ErrInvalidInput                            ErrCode = 100
	ErrInvalidRange                            ErrCode = 101
	ErrMissingInput                            ErrCode = 102
	ErrSerialization                           ErrCode = 103
	ErrDeserialization                         ErrCode = 104
	ErrStreamRead                              ErrCode = 105
	ErrStreamWrite                             ErrCode = 106
	ErrStreamTransform                         ErrCode = 110
	ErrStreamSendAndClose                      ErrCode = 111
	ErrStreamHeaderWrite                       ErrCode = 112
	ErrStreamHeaderRead                        ErrCode = 115
	ErrStreamSink                              ErrCode = 113
	ErrStreamCloseAndRecv                      ErrCode = 114
	ErrMissingMapKey                           ErrCode = 107
	ErrDBWrite                                 ErrCode = 108
	ErrDBRead                                  ErrCode = 109
	ErrCryptoRandomGeneration                  ErrCode = 200
	ErrCryptoKeyGeneration                     ErrCode = 201
	ErrCryptoNonceGeneration                   ErrCode = 202
	ErrCryptoSignature                         ErrCode = 203
	ErrCryptoSignatureVerification             ErrCode = 204
	ErrCryptoDecrypt                           ErrCode = 205
	ErrCryptoDecryptPayload                    ErrCode = 206
	ErrCryptoEncrypt                           ErrCode = 207
	ErrCryptoKeyConversion                     ErrCode = 208
	ErrCryptoCipherInit                        ErrCode = 209
	ErrCryptoKeyDerivation                     ErrCode = 210
	ErrMap                                     ErrCode = 300
	ErrForEach                                 ErrCode = 301
	ErrKeystoreGet                             ErrCode = 400
	ErrKeystorePut                             ErrCode = 401
	ErrNotFound                                ErrCode = 404
	ErrOrbitDBInit                             ErrCode = 1000
	ErrOrbitDBOpen                             ErrCode = 1001
	ErrOrbitDBAppend                           ErrCode = 1002
	ErrOrbitDBDeserialization                  ErrCode = 1003
	ErrOrbitDBStoreCast                        ErrCode = 1004
	ErrIPFSAdd                                 ErrCode = 1050
	ErrIPFSGet                                 ErrCode = 1051
	ErrHandshakeOwnEphemeralKeyGenSend         ErrCode = 1100
	ErrHandshakePeerEphemeralKeyRecv           ErrCode = 1101
	ErrHandshakeRequesterAuthenticateBoxKeyGen ErrCode = 1102
	ErrHandshakeResponderAcceptBoxKeyGen       ErrCode = 1103
	ErrHandshakeRequesterHello                 ErrCode = 1104
	ErrHandshakeResponderHello                 ErrCode = 1105
	ErrHandshakeRequesterAuthenticate          ErrCode = 1106
	ErrHandshakeResponderAccept                ErrCode = 1107
	ErrHandshakeRequesterAcknowledge           ErrCode = 1108
	ErrContactRequestSameAccount               ErrCode = 1200
	ErrContactRequestContactAlreadyAdded       ErrCode = 1201
	ErrContactRequestContactBlocked            ErrCode = 1202
	ErrContactRequestContactUndefined          ErrCode = 1203
	ErrContactRequestIncomingAlreadyReceived   ErrCode = 1204
	ErrGroupMemberLogEventOpen                 ErrCode = 1300
	ErrGroupMemberLogEventSignature            ErrCode = 1301
	ErrGroupMemberUnknownGroupID               ErrCode = 1302
	ErrGroupSecretOtherDestMember              ErrCode = 1303
	ErrGroupSecretAlreadySentToMember          ErrCode = 1304
	ErrGroupInvalidType                        ErrCode = 1305
	ErrGroupMissing                            ErrCode = 1306
	ErrGroupActivate                           ErrCode = 1307
	ErrGroupDeactivate                         ErrCode = 1308
	ErrGroupInfo                               ErrCode = 1309
	// Event errors
	ErrEventListMetadata                   ErrCode = 1400
	ErrEventListMessage                    ErrCode = 1401
	ErrMessageKeyPersistencePut            ErrCode = 1500
	ErrMessageKeyPersistenceGet            ErrCode = 1501
	ErrBridgeInterrupted                   ErrCode = 1600
	ErrBridgeNotRunning                    ErrCode = 1601
	ErrMessengerInvalidDeepLink            ErrCode = 2000
	ErrMessengerDeepLinkRequiresPassphrase ErrCode = 2001
	ErrMessengerDeepLinkInvalidPassphrase  ErrCode = 2002
	ErrDBEntryAlreadyExists                ErrCode = 2100
	ErrDBAddConversation                   ErrCode = 2101
	ErrDBAddContactRequestOutgoingSent     ErrCode = 2102
	ErrDBAddContactRequestOutgoingEnqueud  ErrCode = 2103
	ErrDBAddContactRequestIncomingReceived ErrCode = 2104
	ErrDBAddContactRequestIncomingAccepted ErrCode = 2105
	ErrDBAddGroupMemberDeviceAdded         ErrCode = 2106
	ErrDBMultipleRecords                   ErrCode = 2107
	ErrReplayProcessGroupMetadata          ErrCode = 2200
	ErrReplayProcessGroupMessage           ErrCode = 2201
	ErrAttachmentPrepare                   ErrCode = 2300
	ErrAttachmentRetrieve                  ErrCode = 2301
	ErrProtocolSend                        ErrCode = 2302
	// Test Error
	ErrTestEcho                          ErrCode = 2401
	ErrTestEchoRecv                      ErrCode = 2402
	ErrTestEchoSend                      ErrCode = 2403
	ErrCLINoTermcaps                     ErrCode = 3001
	ErrServicesAuth                      ErrCode = 4000
	ErrServicesAuthNotInitialized        ErrCode = 4001
	ErrServicesAuthWrongState            ErrCode = 4002
	ErrServicesAuthInvalidResponse       ErrCode = 4003
	ErrServicesAuthServer                ErrCode = 4004
	ErrServicesAuthCodeChallenge         ErrCode = 4005
	ErrServicesAuthServiceInvalidToken   ErrCode = 4006
	ErrServicesAuthServiceNotSupported   ErrCode = 4007
	ErrServicesAuthUnknownToken          ErrCode = 4008
	ErrServicesAuthInvalidURL            ErrCode = 4009
	ErrServiceReplication                ErrCode = 4100
	ErrServiceReplicationServer          ErrCode = 4101
	ErrServiceReplicationMissingEndpoint ErrCode = 4102
	ErrBertyAccount                      ErrCode = 5000
	ErrBertyAccountNoIDSpecified         ErrCode = 5001
	ErrBertyAccountAlreadyOpened         ErrCode = 5002
	ErrBertyAccountInvalidIDFormat       ErrCode = 5003
	ErrBertyAccountLoggerDecorator       ErrCode = 5004
	ErrBertyAccountGRPCClient            ErrCode = 5005
	ErrBertyAccountOpenAccount           ErrCode = 5006
	ErrBertyAccountDataNotFound          ErrCode = 5007
	ErrBertyAccountMetadataUpdate        ErrCode = 5008
	ErrBertyAccountManagerOpen           ErrCode = 5009
	ErrBertyAccountManagerClose          ErrCode = 5010
	ErrBertyAccountInvalidCLIArgs        ErrCode = 5011
	ErrBertyAccountFSError               ErrCode = 5012
	ErrBertyAccountAlreadyExists         ErrCode = 5013
	ErrBertyAccountNoBackupSpecified     ErrCode = 5014
	ErrBertyAccountIDGenFailed           ErrCode = 5015
	ErrBertyAccountCreationFailed        ErrCode = 5016
)

var ErrCode_name = map[int32]string{
	0:    "Undefined",
	666:  "TODO",
	777:  "ErrNotImplemented",
	888:  "ErrInternal",
	100:  "ErrInvalidInput",
	101:  "ErrInvalidRange",
	102:  "ErrMissingInput",
	103:  "ErrSerialization",
	104:  "ErrDeserialization",
	105:  "ErrStreamRead",
	106:  "ErrStreamWrite",
	110:  "ErrStreamTransform",
	111:  "ErrStreamSendAndClose",
	112:  "ErrStreamHeaderWrite",
	115:  "ErrStreamHeaderRead",
	113:  "ErrStreamSink",
	114:  "ErrStreamCloseAndRecv",
	107:  "ErrMissingMapKey",
	108:  "ErrDBWrite",
	109:  "ErrDBRead",
	200:  "ErrCryptoRandomGeneration",
	201:  "ErrCryptoKeyGeneration",
	202:  "ErrCryptoNonceGeneration",
	203:  "ErrCryptoSignature",
	204:  "ErrCryptoSignatureVerification",
	205:  "ErrCryptoDecrypt",
	206:  "ErrCryptoDecryptPayload",
	207:  "ErrCryptoEncrypt",
	208:  "ErrCryptoKeyConversion",
	209:  "ErrCryptoCipherInit",
	210:  "ErrCryptoKeyDerivation",
	300:  "ErrMap",
	301:  "ErrForEach",
	400:  "ErrKeystoreGet",
	401:  "ErrKeystorePut",
	404:  "ErrNotFound",
	1000: "ErrOrbitDBInit",
	1001: "ErrOrbitDBOpen",
	1002: "ErrOrbitDBAppend",
	1003: "ErrOrbitDBDeserialization",
	1004: "ErrOrbitDBStoreCast",
	1050: "ErrIPFSAdd",
	1051: "ErrIPFSGet",
	1100: "ErrHandshakeOwnEphemeralKeyGenSend",
	1101: "ErrHandshakePeerEphemeralKeyRecv",
	1102: "ErrHandshakeRequesterAuthenticateBoxKeyGen",
	1103: "ErrHandshakeResponderAcceptBoxKeyGen",
	1104: "ErrHandshakeRequesterHello",
	1105: "ErrHandshakeResponderHello",
	1106: "ErrHandshakeRequesterAuthenticate",
	1107: "ErrHandshakeResponderAccept",
	1108: "ErrHandshakeRequesterAcknowledge",
	1200: "ErrContactRequestSameAccount",
	1201: "ErrContactRequestContactAlreadyAdded",
	1202: "ErrContactRequestContactBlocked",
	1203: "ErrContactRequestContactUndefined",
	1204: "ErrContactRequestIncomingAlreadyReceived",
	1300: "ErrGroupMemberLogEventOpen",
	1301: "ErrGroupMemberLogEventSignature",
	1302: "ErrGroupMemberUnknownGroupID",
	1303: "ErrGroupSecretOtherDestMember",
	1304: "ErrGroupSecretAlreadySentToMember",
	1305: "ErrGroupInvalidType",
	1306: "ErrGroupMissing",
	1307: "ErrGroupActivate",
	1308: "ErrGroupDeactivate",
	1309: "ErrGroupInfo",
	1400: "ErrEventListMetadata",
	1401: "ErrEventListMessage",
	1500: "ErrMessageKeyPersistencePut",
	1501: "ErrMessageKeyPersistenceGet",
	1600: "ErrBridgeInterrupted",
	1601: "ErrBridgeNotRunning",
	2000: "ErrMessengerInvalidDeepLink",
	2001: "ErrMessengerDeepLinkRequiresPassphrase",
	2002: "ErrMessengerDeepLinkInvalidPassphrase",
	2100: "ErrDBEntryAlreadyExists",
	2101: "ErrDBAddConversation",
	2102: "ErrDBAddContactRequestOutgoingSent",
	2103: "ErrDBAddContactRequestOutgoingEnqueud",
	2104: "ErrDBAddContactRequestIncomingReceived",
	2105: "ErrDBAddContactRequestIncomingAccepted",
	2106: "ErrDBAddGroupMemberDeviceAdded",
	2107: "ErrDBMultipleRecords",
	2200: "ErrReplayProcessGroupMetadata",
	2201: "ErrReplayProcessGroupMessage",
	2300: "ErrAttachmentPrepare",
	2301: "ErrAttachmentRetrieve",
	2302: "ErrProtocolSend",
	2401: "ErrTestEcho",
	2402: "ErrTestEchoRecv",
	2403: "ErrTestEchoSend",
	3001: "ErrCLINoTermcaps",
	4000: "ErrServicesAuth",
	4001: "ErrServicesAuthNotInitialized",
	4002: "ErrServicesAuthWrongState",
	4003: "ErrServicesAuthInvalidResponse",
	4004: "ErrServicesAuthServer",
	4005: "ErrServicesAuthCodeChallenge",
	4006: "ErrServicesAuthServiceInvalidToken",
	4007: "ErrServicesAuthServiceNotSupported",
	4008: "ErrServicesAuthUnknownToken",
	4009: "ErrServicesAuthInvalidURL",
	4100: "ErrServiceReplication",
	4101: "ErrServiceReplicationServer",
	4102: "ErrServiceReplicationMissingEndpoint",
	5000: "ErrBertyAccount",
	5001: "ErrBertyAccountNoIDSpecified",
	5002: "ErrBertyAccountAlreadyOpened",
	5003: "ErrBertyAccountInvalidIDFormat",
	5004: "ErrBertyAccountLoggerDecorator",
	5005: "ErrBertyAccountGRPCClient",
	5006: "ErrBertyAccountOpenAccount",
	5007: "ErrBertyAccountDataNotFound",
	5008: "ErrBertyAccountMetadataUpdate",
	5009: "ErrBertyAccountManagerOpen",
	5010: "ErrBertyAccountManagerClose",
	5011: "ErrBertyAccountInvalidCLIArgs",
	5012: "ErrBertyAccountFSError",
	5013: "ErrBertyAccountAlreadyExists",
	5014: "ErrBertyAccountNoBackupSpecified",
	5015: "ErrBertyAccountIDGenFailed",
	5016: "ErrBertyAccountCreationFailed",
}

var ErrCode_value = map[string]int32{
	"Undefined":                                  0,
	"TODO":                                       666,
	"ErrNotImplemented":                          777,
	"ErrInternal":                                888,
	"ErrInvalidInput":                            100,
	"ErrInvalidRange":                            101,
	"ErrMissingInput":                            102,
	"ErrSerialization":                           103,
	"ErrDeserialization":                         104,
	"ErrStreamRead":                              105,
	"ErrStreamWrite":                             106,
	"ErrStreamTransform":                         110,
	"ErrStreamSendAndClose":                      111,
	"ErrStreamHeaderWrite":                       112,
	"ErrStreamHeaderRead":                        115,
	"ErrStreamSink":                              113,
	"ErrStreamCloseAndRecv":                      114,
	"ErrMissingMapKey":                           107,
	"ErrDBWrite":                                 108,
	"ErrDBRead":                                  109,
	"ErrCryptoRandomGeneration":                  200,
	"ErrCryptoKeyGeneration":                     201,
	"ErrCryptoNonceGeneration":                   202,
	"ErrCryptoSignature":                         203,
	"ErrCryptoSignatureVerification":             204,
	"ErrCryptoDecrypt":                           205,
	"ErrCryptoDecryptPayload":                    206,
	"ErrCryptoEncrypt":                           207,
	"ErrCryptoKeyConversion":                     208,
	"ErrCryptoCipherInit":                        209,
	"ErrCryptoKeyDerivation":                     210,
	"ErrMap":                                     300,
	"ErrForEach":                                 301,
	"ErrKeystoreGet":                             400,
	"ErrKeystorePut":                             401,
	"ErrNotFound":                                404,
	"ErrOrbitDBInit":                             1000,
	"ErrOrbitDBOpen":                             1001,
	"ErrOrbitDBAppend":                           1002,
	"ErrOrbitDBDeserialization":                  1003,
	"ErrOrbitDBStoreCast":                        1004,
	"ErrIPFSAdd":                                 1050,
	"ErrIPFSGet":                                 1051,
	"ErrHandshakeOwnEphemeralKeyGenSend":         1100,
	"ErrHandshakePeerEphemeralKeyRecv":           1101,
	"ErrHandshakeRequesterAuthenticateBoxKeyGen": 1102,
	"ErrHandshakeResponderAcceptBoxKeyGen":       1103,
	"ErrHandshakeRequesterHello":                 1104,
	"ErrHandshakeResponderHello":                 1105,
	"ErrHandshakeRequesterAuthenticate":          1106,
	"ErrHandshakeResponderAccept":                1107,
	"ErrHandshakeRequesterAcknowledge":           1108,
	"ErrContactRequestSameAccount":               1200,
	"ErrContactRequestContactAlreadyAdded":       1201,
	"ErrContactRequestContactBlocked":            1202,
	"ErrContactRequestContactUndefined":          1203,
	"ErrContactRequestIncomingAlreadyReceived":   1204,
	"ErrGroupMemberLogEventOpen":                 1300,
	"ErrGroupMemberLogEventSignature":            1301,
	"ErrGroupMemberUnknownGroupID":               1302,
	"ErrGroupSecretOtherDestMember":              1303,
	"ErrGroupSecretAlreadySentToMember":          1304,
	"ErrGroupInvalidType":                        1305,
	"ErrGroupMissing":                            1306,
	"ErrGroupActivate":                           1307,
	"ErrGroupDeactivate":                         1308,
	"ErrGroupInfo":                               1309,
	"ErrEventListMetadata":                       1400,
	"ErrEventListMessage":                        1401,
	"ErrMessageKeyPersistencePut":                1500,
	"ErrMessageKeyPersistenceGet":                1501,
	"ErrBridgeInterrupted":                       1600,
	"ErrBridgeNotRunning":                        1601,
	"ErrMessengerInvalidDeepLink":                2000,
	"ErrMessengerDeepLinkRequiresPassphrase":     2001,
	"ErrMessengerDeepLinkInvalidPassphrase":      2002,
	"ErrDBEntryAlreadyExists":                    2100,
	"ErrDBAddConversation":                       2101,
	"ErrDBAddContactRequestOutgoingSent":         2102,
	"ErrDBAddContactRequestOutgoingEnqueud":      2103,
	"ErrDBAddContactRequestIncomingReceived":     2104,
	"ErrDBAddContactRequestIncomingAccepted":     2105,
	"ErrDBAddGroupMemberDeviceAdded":             2106,
	"ErrDBMultipleRecords":                       2107,
	"ErrReplayProcessGroupMetadata":              2200,
	"ErrReplayProcessGroupMessage":               2201,
	"ErrAttachmentPrepare":                       2300,
	"ErrAttachmentRetrieve":                      2301,
	"ErrProtocolSend":                            2302,
	"ErrTestEcho":                                2401,
	"ErrTestEchoRecv":                            2402,
	"ErrTestEchoSend":                            2403,
	"ErrCLINoTermcaps":                           3001,
	"ErrServicesAuth":                            4000,
	"ErrServicesAuthNotInitialized":              4001,
	"ErrServicesAuthWrongState":                  4002,
	"ErrServicesAuthInvalidResponse":             4003,
	"ErrServicesAuthServer":                      4004,
	"ErrServicesAuthCodeChallenge":               4005,
	"ErrServicesAuthServiceInvalidToken":         4006,
	"ErrServicesAuthServiceNotSupported":         4007,
	"ErrServicesAuthUnknownToken":                4008,
	"ErrServicesAuthInvalidURL":                  4009,
	"ErrServiceReplication":                      4100,
	"ErrServiceReplicationServer":                4101,
	"ErrServiceReplicationMissingEndpoint":       4102,
	"ErrBertyAccount":                            5000,
	"ErrBertyAccountNoIDSpecified":               5001,
	"ErrBertyAccountAlreadyOpened":               5002,
	"ErrBertyAccountInvalidIDFormat":             5003,
	"ErrBertyAccountLoggerDecorator":             5004,
	"ErrBertyAccountGRPCClient":                  5005,
	"ErrBertyAccountOpenAccount":                 5006,
	"ErrBertyAccountDataNotFound":                5007,
	"ErrBertyAccountMetadataUpdate":              5008,
	"ErrBertyAccountManagerOpen":                 5009,
	"ErrBertyAccountManagerClose":                5010,
	"ErrBertyAccountInvalidCLIArgs":              5011,
	"ErrBertyAccountFSError":                     5012,
	"ErrBertyAccountAlreadyExists":               5013,
	"ErrBertyAccountNoBackupSpecified":           5014,
	"ErrBertyAccountIDGenFailed":                 5015,
	"ErrBertyAccountCreationFailed":              5016,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

type ErrDetails struct {
	Codes                []ErrCode `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=berty.errcode.ErrCode" json:"codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ErrDetails) Reset()         { *m = ErrDetails{} }
func (m *ErrDetails) String() string { return proto.CompactTextString(m) }
func (*ErrDetails) ProtoMessage()    {}
func (*ErrDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

func (m *ErrDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrDetails.Unmarshal(m, b)
}

func (m *ErrDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrDetails.Marshal(b, m, deterministic)
}

func (m *ErrDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrDetails.Merge(m, src)
}

func (m *ErrDetails) XXX_Size() int {
	return xxx_messageInfo_ErrDetails.Size(m)
}

func (m *ErrDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ErrDetails proto.InternalMessageInfo

func (m *ErrDetails) GetCodes() []ErrCode {
	if m != nil {
		return m.Codes
	}
	return nil
}

func init() {
	proto.RegisterEnum("berty.errcode.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*ErrDetails)(nil), "berty.errcode.ErrDetails")
}

func init() { proto.RegisterFile("errcode.proto", fileDescriptor_4240057316120df7) }

var fileDescriptor_4240057316120df7 = []byte{
	// 1731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0x59, 0x93, 0x5b, 0x47,
	0x15, 0x8e, 0x62, 0xe2, 0xd1, 0xb4, 0xb7, 0x33, 0x6d, 0xc7, 0x5b, 0x1c, 0x6b, 0x62, 0xe2, 0x28,
	0x18, 0xf0, 0x54, 0xc1, 0x1b, 0x6f, 0xda, 0x66, 0xac, 0xf2, 0x2c, 0x2a, 0x69, 0x4c, 0xaa, 0x78,
	0x6b, 0xdf, 0x3e, 0x73, 0xd5, 0xcc, 0x55, 0xf7, 0x75, 0xdf, 0xd6, 0x24, 0xe2, 0x19, 0x28, 0x76,
	0x12, 0x70, 0x12, 0xc7, 0x81, 0x2a, 0xf6, 0xa5, 0x0a, 0xaa, 0x58, 0xc2, 0x12, 0x78, 0x81, 0x37,
	0x96, 0x2c, 0xb6, 0xe1, 0x11, 0x1e, 0x42, 0x5e, 0xd8, 0x7e, 0x40, 0xa8, 0x02, 0x8a, 0xea, 0xbe,
	0x47, 0x33, 0x57, 0x33, 0x62, 0xe0, 0x49, 0x57, 0xe7, 0x7c, 0x7d, 0xb6, 0x3e, 0xfd, 0x9d, 0x6e,
	0x76, 0x04, 0xad, 0x8d, 0x8c, 0xc4, 0xcb, 0xa9, 0x35, 0xce, 0xf0, 0x23, 0xd7, 0xd1, 0xba, 0xd1,
	0x65, 0x12, 0x9e, 0x3d, 0x11, 0x9b, 0xd8, 0x04, 0xcd, 0x82, 0xff, 0xca, 0x41, 0x17, 0xde, 0xc7,
	0x58, 0xcb, 0xda, 0x26, 0x3a, 0xa1, 0x92, 0x8c, 0xbf, 0x8b, 0x3d, 0xe0, 0xb1, 0xd9, 0xe9, 0xd2,
	0xfc, 0x81, 0xc7, 0x8f, 0xbe, 0xe7, 0xe4, 0xe5, 0x09, 0x13, 0x97, 0x5b, 0xd6, 0x36, 0x8c, 0xc4,
	0x6e, 0x0e, 0xba, 0xf4, 0xe6, 0x39, 0x36, 0x43, 0x22, 0x7e, 0x84, 0xcd, 0x5e, 0xd3, 0x12, 0x37,
	0x94, 0x46, 0x09, 0xf7, 0xf1, 0x59, 0xf6, 0xb6, 0xf5, 0xb5, 0xe6, 0x1a, 0xdc, 0x7e, 0x80, 0x9f,
	0x64, 0x73, 0x2d, 0x6b, 0x57, 0x8d, 0x6b, 0x0f, 0xd2, 0x04, 0x07, 0xa8, 0x1d, 0x4a, 0xf8, 0xf8,
	0x41, 0x0e, 0xec, 0x50, 0xcb, 0xda, 0xb6, 0x76, 0x68, 0xb5, 0x48, 0xe0, 0xad, 0x83, 0xfc, 0x38,
	0x3b, 0x16, 0x24, 0x5b, 0x22, 0x51, 0xb2, 0xad, 0xd3, 0xa1, 0x03, 0x39, 0x29, 0xec, 0x0a, 0x1d,
	0x23, 0x20, 0x09, 0x57, 0x54, 0x96, 0x29, 0x1d, 0xe7, 0xc8, 0x0d, 0x7e, 0x82, 0x41, 0xcb, 0xda,
	0x1e, 0x5a, 0x25, 0x12, 0xf5, 0x21, 0xe1, 0x94, 0xd1, 0x10, 0xf3, 0x93, 0x8c, 0x87, 0x04, 0xb3,
	0x09, 0x79, 0x9f, 0xcf, 0xb1, 0x23, 0x1e, 0xed, 0x2c, 0x8a, 0x41, 0x17, 0x85, 0x04, 0xc5, 0x39,
	0x3b, 0xba, 0x2d, 0x7a, 0xc2, 0x2a, 0x87, 0xf0, 0x41, 0x5a, 0x9e, 0xcb, 0xd6, 0xad, 0xd0, 0xd9,
	0x86, 0xb1, 0x03, 0xd0, 0xfc, 0x0c, 0x7b, 0x70, 0x5b, 0xde, 0x43, 0x2d, 0x6b, 0x5a, 0x36, 0x12,
	0x93, 0x21, 0x18, 0x7e, 0x9a, 0x9d, 0xd8, 0x56, 0x5d, 0x41, 0x21, 0xd1, 0xe6, 0xc6, 0x52, 0x7e,
	0x8a, 0x1d, 0xdf, 0xa5, 0x09, 0x9e, 0xb3, 0x89, 0x60, 0x7a, 0x4a, 0x6f, 0xc2, 0x8d, 0x09, 0x07,
	0xc1, 0x72, 0x4d, 0xcb, 0x2e, 0x46, 0x5b, 0x60, 0x29, 0x51, 0xca, 0x7e, 0x45, 0xa4, 0x57, 0x71,
	0x04, 0x9b, 0xfc, 0x68, 0xbe, 0x93, 0xf5, 0xdc, 0x59, 0xe2, 0x77, 0x24, 0xfc, 0x0f, 0x2e, 0x06,
	0xfc, 0x3c, 0x3b, 0xe3, 0xf7, 0xca, 0x8e, 0x52, 0x67, 0xba, 0x42, 0x4b, 0x33, 0x58, 0x42, 0x8d,
	0x36, 0x2f, 0xc7, 0xaf, 0x4a, 0xfc, 0x21, 0x76, 0x72, 0x5b, 0x7f, 0x15, 0x47, 0x05, 0xe5, 0xaf,
	0x4b, 0xfc, 0x61, 0x76, 0x7a, 0x5b, 0xb9, 0x6a, 0x74, 0x84, 0x05, 0xf5, 0x6f, 0x4a, 0xfc, 0x54,
	0x28, 0x52, 0xae, 0xee, 0xa9, 0x58, 0x0b, 0x37, 0xb4, 0x08, 0xbf, 0x2d, 0xf1, 0xb7, 0xb3, 0xf3,
	0x7b, 0x15, 0xef, 0x47, 0xab, 0x36, 0x54, 0x94, 0xaf, 0x7e, 0xa5, 0xc4, 0x1f, 0x0c, 0xe9, 0xe4,
	0xa0, 0x26, 0x46, 0xfe, 0x17, 0x5e, 0x2d, 0xf1, 0x73, 0xec, 0xd4, 0x6e, 0x71, 0x47, 0x8c, 0x12,
	0x23, 0x24, 0xbc, 0x36, 0xb9, 0xa8, 0xa5, 0xf3, 0x45, 0xaf, 0xef, 0xc9, 0xa2, 0x61, 0xf4, 0x16,
	0xda, 0xcc, 0x3b, 0xba, 0x53, 0xe2, 0xa7, 0x43, 0xf9, 0x73, 0x65, 0x43, 0xa5, 0x7d, 0xb4, 0x6d,
	0xad, 0x1c, 0xdc, 0xdd, 0xb3, 0xac, 0x89, 0x56, 0x6d, 0xe5, 0xf1, 0xdd, 0x2b, 0xf1, 0x43, 0xec,
	0xa0, 0x2f, 0xb7, 0x48, 0xe1, 0x3b, 0xf7, 0xf3, 0x63, 0xa1, 0xca, 0x8b, 0xc6, 0xb6, 0x44, 0xd4,
	0x87, 0xef, 0xde, 0xcf, 0x8f, 0x87, 0xa6, 0xb9, 0x8a, 0xa3, 0xcc, 0x19, 0x8b, 0x4b, 0xe8, 0xe0,
	0xe9, 0x03, 0xbb, 0x84, 0x9d, 0xa1, 0x83, 0x67, 0x0e, 0x50, 0xc3, 0xaf, 0x1a, 0xb7, 0x68, 0x86,
	0x5a, 0xc2, 0xcd, 0x31, 0x6c, 0xcd, 0x5e, 0x57, 0xae, 0x59, 0x0f, 0xb1, 0xfc, 0x79, 0x66, 0x52,
	0xb8, 0x96, 0xa2, 0x86, 0xbf, 0xcc, 0x50, 0xba, 0x24, 0xac, 0xa5, 0x29, 0x6a, 0x09, 0x7f, 0x9d,
	0xa1, 0x4d, 0x25, 0xf1, 0xee, 0x1e, 0xff, 0xdb, 0x0c, 0x65, 0x4c, 0xfa, 0x9e, 0x8f, 0xa5, 0x21,
	0x32, 0x07, 0x7f, 0x9f, 0xa1, 0x3c, 0xda, 0x9d, 0xc5, 0x5e, 0x4d, 0x4a, 0xb8, 0x5d, 0x2e, 0x08,
	0x7c, 0x0e, 0x2f, 0x96, 0x79, 0x95, 0x5d, 0x68, 0x59, 0x7b, 0x45, 0x68, 0x99, 0xf5, 0xc5, 0x26,
	0xae, 0x3d, 0xa9, 0x5b, 0x69, 0x1f, 0x07, 0x68, 0x45, 0x92, 0xf7, 0x87, 0x6f, 0x7b, 0x78, 0xa5,
	0xcc, 0x2f, 0xb2, 0xf9, 0x22, 0xb0, 0x83, 0x68, 0x8b, 0xc8, 0xd0, 0xb4, 0xaf, 0x96, 0xf9, 0x02,
	0xbb, 0x54, 0x84, 0x75, 0xf1, 0xc6, 0x10, 0x33, 0x87, 0xb6, 0x36, 0x74, 0x7d, 0xd4, 0xce, 0x37,
	0x04, 0xd6, 0xcd, 0x53, 0xb9, 0x6d, 0x78, 0xad, 0xcc, 0xdf, 0xc1, 0x1e, 0x9d, 0x5c, 0x90, 0xa5,
	0x46, 0x4b, 0xb4, 0xb5, 0x28, 0xc2, 0xd4, 0xed, 0x40, 0x5f, 0x2f, 0xf3, 0x0a, 0x3b, 0x3b, 0xd5,
	0xf6, 0x15, 0x4c, 0x12, 0x03, 0x77, 0xa6, 0x00, 0xc8, 0x56, 0x0e, 0xb8, 0x5b, 0xe6, 0x8f, 0xb1,
	0x47, 0xfe, 0x67, 0x74, 0x70, 0xaf, 0xcc, 0xe7, 0xd9, 0x43, 0xfb, 0x04, 0x05, 0xbf, 0xdb, 0x53,
	0x8e, 0x1d, 0x4b, 0xd1, 0xa6, 0x36, 0x4f, 0x26, 0x28, 0x63, 0x84, 0xdf, 0x97, 0xf9, 0x23, 0xec,
	0x5c, 0xe0, 0x4e, 0xed, 0x44, 0xe4, 0x08, 0xd4, 0x13, 0x03, 0xac, 0x45, 0x91, 0x19, 0x6a, 0x07,
	0xdf, 0x9b, 0xa5, 0x02, 0x4c, 0x42, 0xe8, 0x5f, 0x2d, 0xb1, 0x28, 0xe4, 0xa8, 0x26, 0x25, 0x4a,
	0xf8, 0xfe, 0x2c, 0x7f, 0x94, 0x55, 0xfe, 0x1b, 0xb4, 0x9e, 0x98, 0x68, 0x13, 0x25, 0xfc, 0x60,
	0x96, 0x92, 0x9c, 0x8a, 0xda, 0x21, 0xef, 0x1f, 0xce, 0xf2, 0x77, 0xb3, 0xc7, 0xf7, 0xe0, 0xda,
	0x3a, 0x32, 0x03, 0xa5, 0x63, 0xf2, 0xdc, 0xc5, 0x08, 0xd5, 0x16, 0x4a, 0x78, 0x69, 0x96, 0x8a,
	0xbb, 0x64, 0xcd, 0x30, 0x5d, 0xc1, 0xc1, 0x75, 0xb4, 0xcb, 0x26, 0x6e, 0x6d, 0xa1, 0x76, 0xa1,
	0x7b, 0x6f, 0x32, 0x8a, 0x6e, 0x0a, 0x60, 0x87, 0x2c, 0x9e, 0x65, 0x54, 0x91, 0x02, 0xea, 0x9a,
	0xf6, 0x15, 0xd3, 0x41, 0xd2, 0x6e, 0xc2, 0x73, 0x8c, 0x5f, 0x60, 0x0f, 0x8f, 0x21, 0x3d, 0x8c,
	0x2c, 0xba, 0x35, 0xd7, 0x47, 0x4f, 0xee, 0x2e, 0x5f, 0x01, 0xcf, 0x33, 0x4a, 0xb2, 0x80, 0xa1,
	0x88, 0x7b, 0xa8, 0xdd, 0xba, 0x21, 0xdc, 0x2d, 0x46, 0x67, 0x23, 0x37, 0x9e, 0x4f, 0x97, 0xf5,
	0x51, 0x8a, 0xf0, 0x02, 0xe3, 0x27, 0xc2, 0x74, 0xc9, 0x03, 0xc9, 0x49, 0x16, 0x6e, 0x33, 0x3a,
	0x82, 0x41, 0x5a, 0x8b, 0x9c, 0xe7, 0x07, 0x84, 0x17, 0x19, 0x71, 0x5f, 0x10, 0x37, 0x51, 0x8c,
	0x15, 0x5f, 0x60, 0x7c, 0x8e, 0x1d, 0xde, 0xb1, 0xbf, 0x61, 0xe0, 0x8b, 0x8c, 0x9f, 0x09, 0x93,
	0x21, 0x64, 0xbe, 0xac, 0x7c, 0xcc, 0x4e, 0x48, 0xe1, 0x04, 0xbc, 0x35, 0x8e, 0xa6, 0xa0, 0xca,
	0x32, 0x11, 0x23, 0xfc, 0x83, 0x51, 0xc7, 0x91, 0xe0, 0x2a, 0x8e, 0x3a, 0x9e, 0xd0, 0x32, 0x87,
	0x3a, 0x0a, 0xc4, 0xf2, 0x87, 0x43, 0xfb, 0x21, 0xfc, 0x59, 0xfe, 0xe3, 0x21, 0x72, 0x5c, 0xb7,
	0x4a, 0xc6, 0x18, 0x26, 0xae, 0x1d, 0xa6, 0x7e, 0x0c, 0xff, 0xe2, 0x30, 0x39, 0xce, 0x55, 0xab,
	0xc6, 0x75, 0x87, 0x5a, 0xfb, 0x84, 0x7f, 0x79, 0xb8, 0x60, 0x16, 0x75, 0x8c, 0xe3, 0x11, 0xdc,
	0x44, 0x4c, 0x97, 0xfd, 0x88, 0xba, 0x73, 0x8c, 0xbf, 0x93, 0x3d, 0x56, 0x44, 0x8c, 0x55, 0xbe,
	0x63, 0x94, 0xc5, 0xac, 0x23, 0xb2, 0x2c, 0xed, 0x5b, 0x91, 0x21, 0xdc, 0x3d, 0xc6, 0x2f, 0xb1,
	0x8b, 0xd3, 0xc0, 0x64, 0xb6, 0x80, 0xbd, 0x77, 0x8c, 0xb8, 0xbf, 0x59, 0x6f, 0x69, 0x67, 0x47,
	0xb4, 0x7f, 0xad, 0xa7, 0x54, 0xe6, 0x32, 0x78, 0x09, 0x28, 0x9b, 0x66, 0xbd, 0x26, 0x25, 0x11,
	0x7c, 0x4e, 0x78, 0x3f, 0x02, 0x22, 0xad, 0xb1, 0xaa, 0xd0, 0xbe, 0x6b, 0x43, 0x17, 0x1b, 0xa5,
	0x63, 0xdf, 0x05, 0xf0, 0x63, 0xa0, 0x68, 0xf6, 0x01, 0xb6, 0xf4, 0x8d, 0x21, 0x0e, 0x25, 0xfc,
	0x04, 0x28, 0xcd, 0x29, 0xd8, 0xf1, 0x99, 0xd8, 0x3e, 0x0c, 0x3f, 0xfd, 0x3f, 0xc0, 0x39, 0x53,
	0xa0, 0x84, 0x97, 0x81, 0xe6, 0x63, 0x00, 0x17, 0xfa, 0xbe, 0x89, 0x5b, 0x2a, 0xc2, 0xfc, 0x6c,
	0xff, 0x6c, 0x27, 0xdd, 0x95, 0x61, 0xe2, 0x54, 0x9a, 0x60, 0x17, 0x23, 0x63, 0x65, 0x06, 0x3f,
	0x07, 0x3a, 0x0f, 0x5d, 0x4c, 0x13, 0x31, 0xea, 0x58, 0x13, 0x61, 0x96, 0x91, 0x1d, 0xea, 0xac,
	0x5b, 0x73, 0x74, 0xac, 0xa6, 0x61, 0xf2, 0x16, 0x7b, 0x61, 0x8e, 0x3c, 0xd4, 0x9c, 0x13, 0x51,
	0xdf, 0x5f, 0xd0, 0x3a, 0x16, 0x53, 0x61, 0x11, 0xfe, 0x39, 0xc7, 0xcf, 0x86, 0x6b, 0xc8, 0x8e,
	0xaa, 0x8b, 0xce, 0x2a, 0xdc, 0x42, 0xf8, 0xd7, 0x1c, 0x9d, 0x93, 0x8e, 0xbf, 0x47, 0x46, 0x26,
	0x09, 0xe3, 0xe0, 0xdf, 0x73, 0x34, 0xe6, 0xd6, 0x31, 0x73, 0xad, 0xa8, 0x6f, 0xe0, 0x0d, 0x4e,
	0xb8, 0xb1, 0x24, 0xcc, 0x83, 0x3f, 0xed, 0x96, 0x86, 0xd5, 0x6f, 0xf2, 0xf1, 0x5c, 0x5f, 0x6e,
	0xaf, 0x9a, 0x75, 0xb4, 0x83, 0x48, 0xa4, 0x19, 0xbc, 0x7c, 0x8a, 0xc0, 0x3d, 0xb4, 0xbe, 0x32,
	0x99, 0x27, 0x65, 0xf8, 0x52, 0x85, 0xd2, 0x2f, 0x4a, 0xfd, 0x35, 0x53, 0x2b, 0x17, 0xa6, 0x20,
	0x4a, 0xf8, 0x72, 0x85, 0x46, 0x64, 0x11, 0xf3, 0x84, 0x35, 0x3a, 0xee, 0x39, 0x7f, 0x4c, 0xbf,
	0x52, 0xa1, 0x2d, 0x28, 0xea, 0xc7, 0x77, 0xcd, 0x40, 0xed, 0x19, 0xc2, 0x57, 0x2b, 0x54, 0x85,
	0x22, 0xc8, 0x7f, 0xa3, 0x85, 0xaf, 0x55, 0xa8, 0xbe, 0x45, 0x9d, 0xbf, 0x10, 0x37, 0xfa, 0x22,
	0x49, 0x7c, 0xa3, 0xc3, 0xd7, 0x2b, 0xd4, 0x95, 0xbb, 0x97, 0xab, 0x08, 0xc7, 0xc4, 0x63, 0x36,
	0x51, 0xc3, 0x37, 0xf6, 0x01, 0xae, 0x1a, 0xd7, 0x1b, 0xa6, 0xa9, 0xb1, 0xbe, 0x71, 0xbe, 0x59,
	0xa1, 0xb3, 0x59, 0x04, 0x12, 0x59, 0xe6, 0xa6, 0xbe, 0x35, 0x2d, 0x6f, 0x72, 0x76, 0xad, 0xbb,
	0x0c, 0xdf, 0xde, 0x95, 0x92, 0xef, 0x8e, 0xf1, 0x8d, 0xec, 0xc3, 0xf3, 0x93, 0xd6, 0x0b, 0x3a,
	0x4a, 0xfa, 0x23, 0xf3, 0x34, 0x9a, 0xf6, 0x22, 0x88, 0x2f, 0x5b, 0x5a, 0xa6, 0x46, 0x69, 0x07,
	0x1f, 0x9d, 0xa7, 0xad, 0xab, 0xfb, 0x87, 0xc4, 0x78, 0xb6, 0x7d, 0xac, 0x4a, 0x55, 0x2b, 0x4a,
	0x57, 0x4d, 0xbb, 0xd9, 0x4b, 0x31, 0x52, 0x1b, 0xca, 0x3f, 0x10, 0xa6, 0x41, 0x88, 0x09, 0xfc,
	0x58, 0x41, 0x09, 0x9f, 0xa8, 0xd2, 0xe6, 0x15, 0x21, 0xe3, 0xd7, 0x43, 0x73, 0xd1, 0xd8, 0x81,
	0x70, 0xf0, 0xc9, 0x69, 0xa0, 0x65, 0x13, 0x07, 0x06, 0x8a, 0x8c, 0x15, 0xce, 0x58, 0xf8, 0x54,
	0x95, 0xca, 0x55, 0x04, 0x2d, 0x75, 0x3b, 0x8d, 0x46, 0xa2, 0x3c, 0x5f, 0x7c, 0xba, 0x4a, 0x33,
	0xae, 0xa8, 0xf7, 0x51, 0x8c, 0x13, 0xfa, 0x4c, 0x95, 0x6a, 0x56, 0x04, 0x34, 0x85, 0x13, 0xdb,
	0xb7, 0xbd, 0xcf, 0x56, 0xa9, 0x5b, 0x8b, 0x88, 0xf1, 0x31, 0xbd, 0x96, 0x4a, 0xdf, 0x8d, 0x4f,
	0x4f, 0x73, 0xb3, 0x22, 0xb4, 0x88, 0xd1, 0x86, 0x51, 0xfa, 0xcc, 0x34, 0x37, 0x04, 0xc8, 0x5f,
	0x1f, 0x9f, 0x9b, 0xe6, 0x86, 0x6a, 0xd2, 0x58, 0x6e, 0xd7, 0x6c, 0x9c, 0xc1, 0xe7, 0xab, 0x74,
	0xdf, 0x2d, 0x62, 0x16, 0x7b, 0x2d, 0x6b, 0x8d, 0x85, 0x9b, 0xfb, 0xd4, 0x9d, 0x18, 0xf8, 0xd9,
	0x2a, 0xdd, 0x71, 0x26, 0x77, 0xaf, 0x2e, 0xa2, 0xcd, 0x61, 0xba, 0xb3, 0x83, 0xcf, 0x4d, 0xcb,
	0xa6, 0xdd, 0x5c, 0x42, 0xbd, 0x28, 0x54, 0x82, 0x12, 0x9e, 0x9f, 0x16, 0x6b, 0xc3, 0x62, 0x68,
	0x24, 0xc2, 0xdc, 0xaa, 0xd6, 0x2f, 0xde, 0x79, 0xe3, 0xfc, 0x7d, 0x1f, 0xa8, 0xe4, 0x2f, 0x51,
	0x87, 0x51, 0x7f, 0x21, 0x7c, 0x2e, 0xc4, 0x66, 0x21, 0xdd, 0x8c, 0x17, 0xe8, 0x6d, 0x7a, 0xfd,
	0x60, 0x78, 0xcf, 0xbe, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0x6a, 0x50, 0xf3, 0x05,
	0x0f, 0x00, 0x00,
}
