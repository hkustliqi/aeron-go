// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type PortfolioBalanceUpdateRejectReasonEnum int8
type PortfolioBalanceUpdateRejectReasonValues struct {
	INSUFFICIENT_FUNDS     PortfolioBalanceUpdateRejectReasonEnum
	UNKNOWN_PORTFOLIO      PortfolioBalanceUpdateRejectReasonEnum
	UNKNOWN_ASSET          PortfolioBalanceUpdateRejectReasonEnum
	UNKNOWN_REQUEST_ACTION PortfolioBalanceUpdateRejectReasonEnum
	NullValue              PortfolioBalanceUpdateRejectReasonEnum
}

var PortfolioBalanceUpdateRejectReason = PortfolioBalanceUpdateRejectReasonValues{0, 1, 2, 4, -128}

func (p PortfolioBalanceUpdateRejectReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(p)); err != nil {
		return err
	}
	return nil
}

func (p *PortfolioBalanceUpdateRejectReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(p)); err != nil {
		return err
	}
	return nil
}

func (p PortfolioBalanceUpdateRejectReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PortfolioBalanceUpdateRejectReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PortfolioBalanceUpdateRejectReason, unknown enumeration value %d", p)
}

func (*PortfolioBalanceUpdateRejectReasonEnum) EncodedLength() int64 {
	return 1
}

func (*PortfolioBalanceUpdateRejectReasonEnum) INSUFFICIENT_FUNDSSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRejectReasonEnum) INSUFFICIENT_FUNDSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.INSUFFICIENT_FUNDSSinceVersion()
}

func (*PortfolioBalanceUpdateRejectReasonEnum) INSUFFICIENT_FUNDSDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_PORTFOLIOSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_PORTFOLIOInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.UNKNOWN_PORTFOLIOSinceVersion()
}

func (*PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_PORTFOLIODeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_ASSETSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_ASSETInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.UNKNOWN_ASSETSinceVersion()
}

func (*PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_ASSETDeprecated() uint16 {
	return 0
}

func (*PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_REQUEST_ACTIONSinceVersion() uint16 {
	return 0
}

func (p *PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_REQUEST_ACTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.UNKNOWN_REQUEST_ACTIONSinceVersion()
}

func (*PortfolioBalanceUpdateRejectReasonEnum) UNKNOWN_REQUEST_ACTIONDeprecated() uint16 {
	return 0
}
