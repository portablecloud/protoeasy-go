// Code generated by protoc-gen-gogo.
// source: pb/geo/geo.proto
// DO NOT EDIT!

package pbgeo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pb_money "go.pedge.io/pb/gogo/pb/money"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StreetDirection int32

const (
	StreetDirection_STREET_DIRECTION_NONE StreetDirection = 0
	StreetDirection_STREET_DIRECTION_N    StreetDirection = 1
	StreetDirection_STREET_DIRECTION_S    StreetDirection = 2
	StreetDirection_STREET_DIRECTION_E    StreetDirection = 3
	StreetDirection_STREET_DIRECTION_W    StreetDirection = 4
	StreetDirection_STREET_DIRECTION_NE   StreetDirection = 5
	StreetDirection_STREET_DIRECTION_SE   StreetDirection = 6
	StreetDirection_STREET_DIRECTION_NW   StreetDirection = 7
	StreetDirection_STREET_DIRECTION_SW   StreetDirection = 8
)

var StreetDirection_name = map[int32]string{
	0: "STREET_DIRECTION_NONE",
	1: "STREET_DIRECTION_N",
	2: "STREET_DIRECTION_S",
	3: "STREET_DIRECTION_E",
	4: "STREET_DIRECTION_W",
	5: "STREET_DIRECTION_NE",
	6: "STREET_DIRECTION_SE",
	7: "STREET_DIRECTION_NW",
	8: "STREET_DIRECTION_SW",
}
var StreetDirection_value = map[string]int32{
	"STREET_DIRECTION_NONE": 0,
	"STREET_DIRECTION_N":    1,
	"STREET_DIRECTION_S":    2,
	"STREET_DIRECTION_E":    3,
	"STREET_DIRECTION_W":    4,
	"STREET_DIRECTION_NE":   5,
	"STREET_DIRECTION_SE":   6,
	"STREET_DIRECTION_NW":   7,
	"STREET_DIRECTION_SW":   8,
}

func (x StreetDirection) String() string {
	return proto.EnumName(StreetDirection_name, int32(x))
}

type Country struct {
	Name         string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alpha_2Code  CountryAlpha2Code     `protobuf:"varint,2,opt,name=alpha_2_code,proto3,enum=pb.geo.CountryAlpha2Code" json:"alpha_2_code,omitempty"`
	Alpha_3Code  CountryAlpha3Code     `protobuf:"varint,3,opt,name=alpha_3_code,proto3,enum=pb.geo.CountryAlpha3Code" json:"alpha_3_code,omitempty"`
	NumericCode  uint32                `protobuf:"varint,4,opt,name=numeric_code,proto3" json:"numeric_code,omitempty"`
	CurrencyCode pb_money.CurrencyCode `protobuf:"varint,5,opt,name=currency_code,proto3,enum=pb.money.CurrencyCode" json:"currency_code,omitempty"`
}

func (m *Country) Reset()         { *m = Country{} }
func (m *Country) String() string { return proto.CompactTextString(m) }
func (*Country) ProtoMessage()    {}

type LatLng struct {
	LatPicos  int64 `protobuf:"varint,1,opt,name=lat_picos,proto3" json:"lat_picos,omitempty"`
	LongPicos int64 `protobuf:"varint,2,opt,name=long_picos,proto3" json:"long_picos,omitempty"`
}

func (m *LatLng) Reset()         { *m = LatLng{} }
func (m *LatLng) String() string { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()    {}

type PostalAddress struct {
	StreetNumber                     uint64            `protobuf:"varint,1,opt,name=street_number,proto3" json:"street_number,omitempty"`
	StreetNumberPostfix              string            `protobuf:"bytes,2,opt,name=street_number_postfix,proto3" json:"street_number_postfix,omitempty"`
	StreetName                       string            `protobuf:"bytes,3,opt,name=street_name,proto3" json:"street_name,omitempty"`
	PreStreetDirection               StreetDirection   `protobuf:"varint,4,opt,name=pre_street_direction,proto3,enum=pb.geo.StreetDirection" json:"pre_street_direction,omitempty"`
	PostStreetDirection              StreetDirection   `protobuf:"varint,5,opt,name=post_street_direction,proto3,enum=pb.geo.StreetDirection" json:"post_street_direction,omitempty"`
	StreetTypeAbbreviation           string            `protobuf:"bytes,6,opt,name=street_type_abbreviation,proto3" json:"street_type_abbreviation,omitempty"`
	SecondaryAddressTypeAbbreviation string            `protobuf:"bytes,7,opt,name=secondary_address_type_abbreviation,proto3" json:"secondary_address_type_abbreviation,omitempty"`
	SecondaryAddressValue            string            `protobuf:"bytes,8,opt,name=secondary_address_value,proto3" json:"secondary_address_value,omitempty"`
	LocalityName                     string            `protobuf:"bytes,9,opt,name=locality_name,proto3" json:"locality_name,omitempty"`
	RegionCode                       string            `protobuf:"bytes,10,opt,name=region_code,proto3" json:"region_code,omitempty"`
	PostalCode                       string            `protobuf:"bytes,11,opt,name=postal_code,proto3" json:"postal_code,omitempty"`
	CountryAlpha_2Code               CountryAlpha2Code `protobuf:"varint,12,opt,name=country_alpha_2_code,proto3,enum=pb.geo.CountryAlpha2Code" json:"country_alpha_2_code,omitempty"`
}

func (m *PostalAddress) Reset()         { *m = PostalAddress{} }
func (m *PostalAddress) String() string { return proto.CompactTextString(m) }
func (*PostalAddress) ProtoMessage()    {}

type SimplePostalAddress struct {
	StreetAddress      string            `protobuf:"bytes,1,opt,name=street_address,proto3" json:"street_address,omitempty"`
	StreetAddress_2    string            `protobuf:"bytes,2,opt,name=street_address_2,proto3" json:"street_address_2,omitempty"`
	LocalityName       string            `protobuf:"bytes,3,opt,name=locality_name,proto3" json:"locality_name,omitempty"`
	RegionCode         string            `protobuf:"bytes,4,opt,name=region_code,proto3" json:"region_code,omitempty"`
	PostalCode         string            `protobuf:"bytes,5,opt,name=postal_code,proto3" json:"postal_code,omitempty"`
	CountryAlpha_2Code CountryAlpha2Code `protobuf:"varint,6,opt,name=country_alpha_2_code,proto3,enum=pb.geo.CountryAlpha2Code" json:"country_alpha_2_code,omitempty"`
}

func (m *SimplePostalAddress) Reset()         { *m = SimplePostalAddress{} }
func (m *SimplePostalAddress) String() string { return proto.CompactTextString(m) }
func (*SimplePostalAddress) ProtoMessage()    {}

func init() {
	proto.RegisterType((*Country)(nil), "pb.geo.Country")
	proto.RegisterType((*LatLng)(nil), "pb.geo.LatLng")
	proto.RegisterType((*PostalAddress)(nil), "pb.geo.PostalAddress")
	proto.RegisterType((*SimplePostalAddress)(nil), "pb.geo.SimplePostalAddress")
	proto.RegisterEnum("pb.geo.StreetDirection", StreetDirection_name, StreetDirection_value)
}
