// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"math"
)

type OmsAdminRequestHeader struct {
	CorrelationId int64
	RequestTime   int64
	AdminId       int64
}

func (o *OmsAdminRequestHeader) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, o.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.RequestTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.AdminId); err != nil {
		return err
	}
	return nil
}

func (o *OmsAdminRequestHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !o.CorrelationIdInActingVersion(actingVersion) {
		o.CorrelationId = o.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.CorrelationId); err != nil {
			return err
		}
	}
	if !o.RequestTimeInActingVersion(actingVersion) {
		o.RequestTime = o.RequestTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.RequestTime); err != nil {
			return err
		}
	}
	if !o.AdminIdInActingVersion(actingVersion) {
		o.AdminId = o.AdminIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.AdminId); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsAdminRequestHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.CorrelationIdInActingVersion(actingVersion) {
		if o.CorrelationId < o.CorrelationIdMinValue() || o.CorrelationId > o.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on o.CorrelationId (%v < %v > %v)", o.CorrelationIdMinValue(), o.CorrelationId, o.CorrelationIdMaxValue())
		}
	}
	if o.RequestTimeInActingVersion(actingVersion) {
		if o.RequestTime < o.RequestTimeMinValue() || o.RequestTime > o.RequestTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.RequestTime (%v < %v > %v)", o.RequestTimeMinValue(), o.RequestTime, o.RequestTimeMaxValue())
		}
	}
	if o.AdminIdInActingVersion(actingVersion) {
		if o.AdminId < o.AdminIdMinValue() || o.AdminId > o.AdminIdMaxValue() {
			return fmt.Errorf("Range check failed on o.AdminId (%v < %v > %v)", o.AdminIdMinValue(), o.AdminId, o.AdminIdMaxValue())
		}
	}
	return nil
}

func OmsAdminRequestHeaderInit(o *OmsAdminRequestHeader) {
	return
}

func (*OmsAdminRequestHeader) EncodedLength() int64 {
	return 24
}

func (*OmsAdminRequestHeader) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsAdminRequestHeader) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsAdminRequestHeader) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsAdminRequestHeader) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestHeader) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsAdminRequestHeader) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestHeader) RequestTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsAdminRequestHeader) RequestTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsAdminRequestHeader) RequestTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsAdminRequestHeader) RequestTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestHeader) RequestTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestTimeSinceVersion()
}

func (*OmsAdminRequestHeader) RequestTimeDeprecated() uint16 {
	return 0
}

func (*OmsAdminRequestHeader) AdminIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsAdminRequestHeader) AdminIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsAdminRequestHeader) AdminIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsAdminRequestHeader) AdminIdSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminRequestHeader) AdminIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AdminIdSinceVersion()
}

func (*OmsAdminRequestHeader) AdminIdDeprecated() uint16 {
	return 0
}
