// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsStateEvent struct {
	ClusterSessionId   int64
	LastUpdatedTime    int64
	CorrelationId      int64
	StateId            int64
	StateType          StateTypeEnum
	StateAction        StateActionEnum
	Status             StateStatusEnum
	Name               [80]byte
	StateMessageHeader MessageHeader
}

func (o *OmsStateEvent) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.ClusterSessionId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LastUpdatedTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.StateId); err != nil {
		return err
	}
	if err := o.StateType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.StateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Status.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Name[:]); err != nil {
		return err
	}
	if err := o.StateMessageHeader.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OmsStateEvent) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.ClusterSessionIdInActingVersion(actingVersion) {
		o.ClusterSessionId = o.ClusterSessionIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.ClusterSessionId); err != nil {
			return err
		}
	}
	if !o.LastUpdatedTimeInActingVersion(actingVersion) {
		o.LastUpdatedTime = o.LastUpdatedTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LastUpdatedTime); err != nil {
			return err
		}
	}
	if !o.CorrelationIdInActingVersion(actingVersion) {
		o.CorrelationId = o.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.CorrelationId); err != nil {
			return err
		}
	}
	if !o.StateIdInActingVersion(actingVersion) {
		o.StateId = o.StateIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.StateId); err != nil {
			return err
		}
	}
	if o.StateTypeInActingVersion(actingVersion) {
		if err := o.StateType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.StateActionInActingVersion(actingVersion) {
		if err := o.StateAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.StatusInActingVersion(actingVersion) {
		if err := o.Status.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.NameInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.Name[idx] = o.NameNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Name[:]); err != nil {
			return err
		}
	}
	if o.StateMessageHeaderInActingVersion(actingVersion) {
		if err := o.StateMessageHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsStateEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.ClusterSessionIdInActingVersion(actingVersion) {
		if o.ClusterSessionId < o.ClusterSessionIdMinValue() || o.ClusterSessionId > o.ClusterSessionIdMaxValue() {
			return fmt.Errorf("Range check failed on o.ClusterSessionId (%v < %v > %v)", o.ClusterSessionIdMinValue(), o.ClusterSessionId, o.ClusterSessionIdMaxValue())
		}
	}
	if o.LastUpdatedTimeInActingVersion(actingVersion) {
		if o.LastUpdatedTime < o.LastUpdatedTimeMinValue() || o.LastUpdatedTime > o.LastUpdatedTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.LastUpdatedTime (%v < %v > %v)", o.LastUpdatedTimeMinValue(), o.LastUpdatedTime, o.LastUpdatedTimeMaxValue())
		}
	}
	if o.CorrelationIdInActingVersion(actingVersion) {
		if o.CorrelationId < o.CorrelationIdMinValue() || o.CorrelationId > o.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on o.CorrelationId (%v < %v > %v)", o.CorrelationIdMinValue(), o.CorrelationId, o.CorrelationIdMaxValue())
		}
	}
	if o.StateIdInActingVersion(actingVersion) {
		if o.StateId < o.StateIdMinValue() || o.StateId > o.StateIdMaxValue() {
			return fmt.Errorf("Range check failed on o.StateId (%v < %v > %v)", o.StateIdMinValue(), o.StateId, o.StateIdMaxValue())
		}
	}
	if err := o.StateType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.StateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.Status.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.NameInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.Name[idx] < o.NameMinValue() || o.Name[idx] > o.NameMaxValue() {
				return fmt.Errorf("Range check failed on o.Name[%d] (%v < %v > %v)", idx, o.NameMinValue(), o.Name[idx], o.NameMaxValue())
			}
		}
	}
	for idx, ch := range o.Name {
		if ch > 127 {
			return fmt.Errorf("o.Name[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func OmsStateEventInit(o *OmsStateEvent) {
	return
}

func (*OmsStateEvent) SbeBlockLength() (blockLength uint16) {
	return 124
}

func (*OmsStateEvent) SbeTemplateId() (templateId uint16) {
	return 1012
}

func (*OmsStateEvent) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsStateEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsStateEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsStateEvent) ClusterSessionIdId() uint16 {
	return 1
}

func (*OmsStateEvent) ClusterSessionIdSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) ClusterSessionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClusterSessionIdSinceVersion()
}

func (*OmsStateEvent) ClusterSessionIdDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) ClusterSessionIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OmsStateEvent) ClusterSessionIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateEvent) ClusterSessionIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateEvent) ClusterSessionIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateEvent) LastUpdatedTimeId() uint16 {
	return 2
}

func (*OmsStateEvent) LastUpdatedTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) LastUpdatedTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastUpdatedTimeSinceVersion()
}

func (*OmsStateEvent) LastUpdatedTimeDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) LastUpdatedTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OmsStateEvent) LastUpdatedTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateEvent) LastUpdatedTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateEvent) LastUpdatedTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateEvent) CorrelationIdId() uint16 {
	return 3
}

func (*OmsStateEvent) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsStateEvent) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) CorrelationIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OmsStateEvent) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateEvent) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateEvent) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateEvent) StateIdId() uint16 {
	return 4
}

func (*OmsStateEvent) StateIdSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) StateIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateIdSinceVersion()
}

func (*OmsStateEvent) StateIdDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) StateIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OmsStateEvent) StateIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateEvent) StateIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateEvent) StateIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateEvent) StateTypeId() uint16 {
	return 5
}

func (*OmsStateEvent) StateTypeSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) StateTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateTypeSinceVersion()
}

func (*OmsStateEvent) StateTypeDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) StateTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OmsStateEvent) StateActionId() uint16 {
	return 6
}

func (*OmsStateEvent) StateActionSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) StateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateActionSinceVersion()
}

func (*OmsStateEvent) StateActionDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) StateActionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OmsStateEvent) StatusId() uint16 {
	return 7
}

func (*OmsStateEvent) StatusSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) StatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StatusSinceVersion()
}

func (*OmsStateEvent) StatusDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) StatusMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OmsStateEvent) NameId() uint16 {
	return 8
}

func (*OmsStateEvent) NameSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) NameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NameSinceVersion()
}

func (*OmsStateEvent) NameDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) NameMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OmsStateEvent) NameMinValue() byte {
	return byte(32)
}

func (*OmsStateEvent) NameMaxValue() byte {
	return byte(126)
}

func (*OmsStateEvent) NameNullValue() byte {
	return 0
}

func (o *OmsStateEvent) NameCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsStateEvent) StateMessageHeaderId() uint16 {
	return 9
}

func (*OmsStateEvent) StateMessageHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsStateEvent) StateMessageHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateMessageHeaderSinceVersion()
}

func (*OmsStateEvent) StateMessageHeaderDeprecated() uint16 {
	return 0
}

func (*OmsStateEvent) StateMessageHeaderMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}
