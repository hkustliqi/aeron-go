// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"math"
)

type EngUserRequestHeader struct {
	GatewayLogPosition int64
	OmsLogPosition     int64
	OmsTimestamp       int64
	CorrelationId      int64
	UserId             int64
}

func (e *EngUserRequestHeader) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, e.GatewayLogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.OmsLogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.OmsTimestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.UserId); err != nil {
		return err
	}
	return nil
}

func (e *EngUserRequestHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !e.GatewayLogPositionInActingVersion(actingVersion) {
		e.GatewayLogPosition = e.GatewayLogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.GatewayLogPosition); err != nil {
			return err
		}
	}
	if !e.OmsLogPositionInActingVersion(actingVersion) {
		e.OmsLogPosition = e.OmsLogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OmsLogPosition); err != nil {
			return err
		}
	}
	if !e.OmsTimestampInActingVersion(actingVersion) {
		e.OmsTimestamp = e.OmsTimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OmsTimestamp); err != nil {
			return err
		}
	}
	if !e.CorrelationIdInActingVersion(actingVersion) {
		e.CorrelationId = e.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CorrelationId); err != nil {
			return err
		}
	}
	if !e.UserIdInActingVersion(actingVersion) {
		e.UserId = e.UserIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.UserId); err != nil {
			return err
		}
	}
	return nil
}

func (e *EngUserRequestHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.GatewayLogPositionInActingVersion(actingVersion) {
		if e.GatewayLogPosition < e.GatewayLogPositionMinValue() || e.GatewayLogPosition > e.GatewayLogPositionMaxValue() {
			return fmt.Errorf("Range check failed on e.GatewayLogPosition (%v < %v > %v)", e.GatewayLogPositionMinValue(), e.GatewayLogPosition, e.GatewayLogPositionMaxValue())
		}
	}
	if e.OmsLogPositionInActingVersion(actingVersion) {
		if e.OmsLogPosition < e.OmsLogPositionMinValue() || e.OmsLogPosition > e.OmsLogPositionMaxValue() {
			return fmt.Errorf("Range check failed on e.OmsLogPosition (%v < %v > %v)", e.OmsLogPositionMinValue(), e.OmsLogPosition, e.OmsLogPositionMaxValue())
		}
	}
	if e.OmsTimestampInActingVersion(actingVersion) {
		if e.OmsTimestamp < e.OmsTimestampMinValue() || e.OmsTimestamp > e.OmsTimestampMaxValue() {
			return fmt.Errorf("Range check failed on e.OmsTimestamp (%v < %v > %v)", e.OmsTimestampMinValue(), e.OmsTimestamp, e.OmsTimestampMaxValue())
		}
	}
	if e.CorrelationIdInActingVersion(actingVersion) {
		if e.CorrelationId < e.CorrelationIdMinValue() || e.CorrelationId > e.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on e.CorrelationId (%v < %v > %v)", e.CorrelationIdMinValue(), e.CorrelationId, e.CorrelationIdMaxValue())
		}
	}
	if e.UserIdInActingVersion(actingVersion) {
		if e.UserId < e.UserIdMinValue() || e.UserId > e.UserIdMaxValue() {
			return fmt.Errorf("Range check failed on e.UserId (%v < %v > %v)", e.UserIdMinValue(), e.UserId, e.UserIdMaxValue())
		}
	}
	return nil
}

func EngUserRequestHeaderInit(e *EngUserRequestHeader) {
	return
}

func (*EngUserRequestHeader) EncodedLength() int64 {
	return 40
}

func (*EngUserRequestHeader) GatewayLogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngUserRequestHeader) GatewayLogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*EngUserRequestHeader) GatewayLogPositionNullValue() int64 {
	return math.MinInt64
}

func (*EngUserRequestHeader) GatewayLogPositionSinceVersion() uint16 {
	return 0
}

func (e *EngUserRequestHeader) GatewayLogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.GatewayLogPositionSinceVersion()
}

func (*EngUserRequestHeader) GatewayLogPositionDeprecated() uint16 {
	return 0
}

func (*EngUserRequestHeader) OmsLogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngUserRequestHeader) OmsLogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*EngUserRequestHeader) OmsLogPositionNullValue() int64 {
	return math.MinInt64
}

func (*EngUserRequestHeader) OmsLogPositionSinceVersion() uint16 {
	return 0
}

func (e *EngUserRequestHeader) OmsLogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OmsLogPositionSinceVersion()
}

func (*EngUserRequestHeader) OmsLogPositionDeprecated() uint16 {
	return 0
}

func (*EngUserRequestHeader) OmsTimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngUserRequestHeader) OmsTimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*EngUserRequestHeader) OmsTimestampNullValue() int64 {
	return math.MinInt64
}

func (*EngUserRequestHeader) OmsTimestampSinceVersion() uint16 {
	return 0
}

func (e *EngUserRequestHeader) OmsTimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OmsTimestampSinceVersion()
}

func (*EngUserRequestHeader) OmsTimestampDeprecated() uint16 {
	return 0
}

func (*EngUserRequestHeader) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngUserRequestHeader) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngUserRequestHeader) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*EngUserRequestHeader) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (e *EngUserRequestHeader) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CorrelationIdSinceVersion()
}

func (*EngUserRequestHeader) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*EngUserRequestHeader) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngUserRequestHeader) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngUserRequestHeader) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*EngUserRequestHeader) UserIdSinceVersion() uint16 {
	return 0
}

func (e *EngUserRequestHeader) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UserIdSinceVersion()
}

func (*EngUserRequestHeader) UserIdDeprecated() uint16 {
	return 0
}
