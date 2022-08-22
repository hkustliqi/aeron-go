// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type PortfolioPermissionTypeEnum int8
type PortfolioPermissionTypeValues struct {
	READ      PortfolioPermissionTypeEnum
	TRADE     PortfolioPermissionTypeEnum
	NullValue PortfolioPermissionTypeEnum
}

var PortfolioPermissionType = PortfolioPermissionTypeValues{1, 2, -128}

func (p PortfolioPermissionTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(p)); err != nil {
		return err
	}
	return nil
}

func (p *PortfolioPermissionTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(p)); err != nil {
		return err
	}
	return nil
}

func (p PortfolioPermissionTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PortfolioPermissionType)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PortfolioPermissionType, unknown enumeration value %d", p)
}

func (*PortfolioPermissionTypeEnum) EncodedLength() int64 {
	return 1
}

func (*PortfolioPermissionTypeEnum) READSinceVersion() uint16 {
	return 0
}

func (p *PortfolioPermissionTypeEnum) READInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.READSinceVersion()
}

func (*PortfolioPermissionTypeEnum) READDeprecated() uint16 {
	return 0
}

func (*PortfolioPermissionTypeEnum) TRADESinceVersion() uint16 {
	return 0
}

func (p *PortfolioPermissionTypeEnum) TRADEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.TRADESinceVersion()
}

func (*PortfolioPermissionTypeEnum) TRADEDeprecated() uint16 {
	return 0
}
