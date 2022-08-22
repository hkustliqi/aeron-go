// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type PortfolioBalanceUpdateActionEnum int8
type PortfolioBalanceUpdateActionValues struct {
	DEPOSIT   PortfolioBalanceUpdateActionEnum
	WITHDRAW  PortfolioBalanceUpdateActionEnum
	NullValue PortfolioBalanceUpdateActionEnum
}

var PortfolioBalanceUpdateAction = PortfolioBalanceUpdateActionValues{0, 1, -128}

func (p PortfolioBalanceUpdateActionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(p)); err != nil {
		return err
	}
	return nil
}

func (p *PortfolioBalanceUpdateActionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(p)); err != nil {
		return err
	}
	return nil
}

func (p PortfolioBalanceUpdateActionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PortfolioBalanceUpdateAction)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PortfolioBalanceUpdateAction, unknown enumeration value %d", p)
}

func (*PortfolioBalanceUpdateActionEnum) EncodedLength() int64 {
	return 1
}

func (*PortfolioBalanceUpdateActionEnum) DEPOSITSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateActionEnum) DEPOSITInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.DEPOSITSinceVersion()
}

func (*PortfolioBalanceUpdateActionEnum) DEPOSITDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateActionEnum) WITHDRAWSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateActionEnum) WITHDRAWInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.WITHDRAWSinceVersion()
}

func (*PortfolioBalanceUpdateActionEnum) WITHDRAWDeprecated() uint16 {
	return 0
}
