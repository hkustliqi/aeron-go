// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type ManualSettlementTriggerResultEnum int8
type ManualSettlementTriggerResultValues struct {
	SUCCESS                     ManualSettlementTriggerResultEnum
	REJECT_UNKNOWN_INSTRUMENT   ManualSettlementTriggerResultEnum
	REJECT_SETTLEMENT_SCHEDULED ManualSettlementTriggerResultEnum
	REJECT_OTHER                ManualSettlementTriggerResultEnum
	NullValue                   ManualSettlementTriggerResultEnum
}

var ManualSettlementTriggerResult = ManualSettlementTriggerResultValues{0, 1, 2, 127, -128}

func (m ManualSettlementTriggerResultEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(m)); err != nil {
		return err
	}
	return nil
}

func (m *ManualSettlementTriggerResultEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(m)); err != nil {
		return err
	}
	return nil
}

func (m ManualSettlementTriggerResultEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ManualSettlementTriggerResult)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ManualSettlementTriggerResult, unknown enumeration value %d", m)
}

func (*ManualSettlementTriggerResultEnum) EncodedLength() int64 {
	return 1
}

func (*ManualSettlementTriggerResultEnum) SUCCESSSinceVersion() uint16 {
	return 0
}

func (m *ManualSettlementTriggerResultEnum) SUCCESSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SUCCESSSinceVersion()
}

func (*ManualSettlementTriggerResultEnum) SUCCESSDeprecated() uint16 {
	return 0
}

func (*ManualSettlementTriggerResultEnum) REJECT_UNKNOWN_INSTRUMENTSinceVersion() uint16 {
	return 0
}

func (m *ManualSettlementTriggerResultEnum) REJECT_UNKNOWN_INSTRUMENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.REJECT_UNKNOWN_INSTRUMENTSinceVersion()
}

func (*ManualSettlementTriggerResultEnum) REJECT_UNKNOWN_INSTRUMENTDeprecated() uint16 {
	return 0
}

func (*ManualSettlementTriggerResultEnum) REJECT_SETTLEMENT_SCHEDULEDSinceVersion() uint16 {
	return 0
}

func (m *ManualSettlementTriggerResultEnum) REJECT_SETTLEMENT_SCHEDULEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.REJECT_SETTLEMENT_SCHEDULEDSinceVersion()
}

func (*ManualSettlementTriggerResultEnum) REJECT_SETTLEMENT_SCHEDULEDDeprecated() uint16 {
	return 0
}

func (*ManualSettlementTriggerResultEnum) REJECT_OTHERSinceVersion() uint16 {
	return 0
}

func (m *ManualSettlementTriggerResultEnum) REJECT_OTHERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.REJECT_OTHERSinceVersion()
}

func (*ManualSettlementTriggerResultEnum) REJECT_OTHERDeprecated() uint16 {
	return 0
}
