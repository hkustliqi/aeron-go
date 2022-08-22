// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type StateTypeEnum int16
type StateTypeValues struct {
	ADMIN_USER             StateTypeEnum
	TRADING_USER           StateTypeEnum
	PORTFOLIO              StateTypeEnum
	PORTFOLIO_PERMISSION   StateTypeEnum
	ASSET                  StateTypeEnum
	INSTRUMENT             StateTypeEnum
	MARKET_DATA_USER       StateTypeEnum
	DROP_COPY_USER         StateTypeEnum
	DROP_COPY_SUBSCRIPTION StateTypeEnum
	MONITORED_USER         StateTypeEnum
	NullValue              StateTypeEnum
}

var StateType = StateTypeValues{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -32768}

func (s StateTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt16(_w, int16(s)); err != nil {
		return err
	}
	return nil
}

func (s *StateTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt16(_r, (*int16)(s)); err != nil {
		return err
	}
	return nil
}

func (s StateTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(StateType)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on StateType, unknown enumeration value %d", s)
}

func (*StateTypeEnum) EncodedLength() int64 {
	return 2
}

func (*StateTypeEnum) ADMIN_USERSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) ADMIN_USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ADMIN_USERSinceVersion()
}

func (*StateTypeEnum) ADMIN_USERDeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) TRADING_USERSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) TRADING_USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TRADING_USERSinceVersion()
}

func (*StateTypeEnum) TRADING_USERDeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) PORTFOLIOSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) PORTFOLIOInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PORTFOLIOSinceVersion()
}

func (*StateTypeEnum) PORTFOLIODeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) PORTFOLIO_PERMISSIONSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) PORTFOLIO_PERMISSIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PORTFOLIO_PERMISSIONSinceVersion()
}

func (*StateTypeEnum) PORTFOLIO_PERMISSIONDeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) ASSETSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) ASSETInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ASSETSinceVersion()
}

func (*StateTypeEnum) ASSETDeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) INSTRUMENTSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) INSTRUMENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.INSTRUMENTSinceVersion()
}

func (*StateTypeEnum) INSTRUMENTDeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) MARKET_DATA_USERSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) MARKET_DATA_USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MARKET_DATA_USERSinceVersion()
}

func (*StateTypeEnum) MARKET_DATA_USERDeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) DROP_COPY_USERSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) DROP_COPY_USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DROP_COPY_USERSinceVersion()
}

func (*StateTypeEnum) DROP_COPY_USERDeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) DROP_COPY_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) DROP_COPY_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DROP_COPY_SUBSCRIPTIONSinceVersion()
}

func (*StateTypeEnum) DROP_COPY_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}

func (*StateTypeEnum) MONITORED_USERSinceVersion() uint16 {
	return 0
}

func (s *StateTypeEnum) MONITORED_USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MONITORED_USERSinceVersion()
}

func (*StateTypeEnum) MONITORED_USERDeprecated() uint16 {
	return 0
}
