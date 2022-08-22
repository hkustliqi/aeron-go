// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsStateChangeReject struct {
	ClusterSessionId int64
	LastUpdatedTime  int64
	CorrelationId    int64
	StateId          int64
	StateType        StateTypeEnum
	StateAction      StateActionEnum
	Status           StateStatusEnum
	Name             [80]byte
	ErrorMessage     [80]byte
}

func (o *OmsStateChangeReject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteBytes(_w, o.ErrorMessage[:]); err != nil {
		return err
	}
	return nil
}

func (o *OmsStateChangeReject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !o.ErrorMessageInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.ErrorMessage[idx] = o.ErrorMessageNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.ErrorMessage[:]); err != nil {
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

func (o *OmsStateChangeReject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if o.ErrorMessageInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.ErrorMessage[idx] < o.ErrorMessageMinValue() || o.ErrorMessage[idx] > o.ErrorMessageMaxValue() {
				return fmt.Errorf("Range check failed on o.ErrorMessage[%d] (%v < %v > %v)", idx, o.ErrorMessageMinValue(), o.ErrorMessage[idx], o.ErrorMessageMaxValue())
			}
		}
	}
	for idx, ch := range o.ErrorMessage {
		if ch > 127 {
			return fmt.Errorf("o.ErrorMessage[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func OmsStateChangeRejectInit(o *OmsStateChangeReject) {
	return
}

func (*OmsStateChangeReject) SbeBlockLength() (blockLength uint16) {
	return 196
}

func (*OmsStateChangeReject) SbeTemplateId() (templateId uint16) {
	return 1011
}

func (*OmsStateChangeReject) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsStateChangeReject) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsStateChangeReject) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsStateChangeReject) ClusterSessionIdId() uint16 {
	return 1
}

func (*OmsStateChangeReject) ClusterSessionIdSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) ClusterSessionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClusterSessionIdSinceVersion()
}

func (*OmsStateChangeReject) ClusterSessionIdDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) ClusterSessionIdMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) ClusterSessionIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateChangeReject) ClusterSessionIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateChangeReject) ClusterSessionIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateChangeReject) LastUpdatedTimeId() uint16 {
	return 2
}

func (*OmsStateChangeReject) LastUpdatedTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) LastUpdatedTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastUpdatedTimeSinceVersion()
}

func (*OmsStateChangeReject) LastUpdatedTimeDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) LastUpdatedTimeMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) LastUpdatedTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateChangeReject) LastUpdatedTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateChangeReject) LastUpdatedTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateChangeReject) CorrelationIdId() uint16 {
	return 3
}

func (*OmsStateChangeReject) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsStateChangeReject) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) CorrelationIdMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateChangeReject) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateChangeReject) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateChangeReject) StateIdId() uint16 {
	return 4
}

func (*OmsStateChangeReject) StateIdSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) StateIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateIdSinceVersion()
}

func (*OmsStateChangeReject) StateIdDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) StateIdMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) StateIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateChangeReject) StateIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateChangeReject) StateIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateChangeReject) StateTypeId() uint16 {
	return 5
}

func (*OmsStateChangeReject) StateTypeSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) StateTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateTypeSinceVersion()
}

func (*OmsStateChangeReject) StateTypeDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) StateTypeMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) StateActionId() uint16 {
	return 6
}

func (*OmsStateChangeReject) StateActionSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) StateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateActionSinceVersion()
}

func (*OmsStateChangeReject) StateActionDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) StateActionMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) StatusId() uint16 {
	return 7
}

func (*OmsStateChangeReject) StatusSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) StatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StatusSinceVersion()
}

func (*OmsStateChangeReject) StatusDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) StatusMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) NameId() uint16 {
	return 8
}

func (*OmsStateChangeReject) NameSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) NameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NameSinceVersion()
}

func (*OmsStateChangeReject) NameDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) NameMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) NameMinValue() byte {
	return byte(32)
}

func (*OmsStateChangeReject) NameMaxValue() byte {
	return byte(126)
}

func (*OmsStateChangeReject) NameNullValue() byte {
	return 0
}

func (o *OmsStateChangeReject) NameCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsStateChangeReject) ErrorMessageId() uint16 {
	return 9
}

func (*OmsStateChangeReject) ErrorMessageSinceVersion() uint16 {
	return 0
}

func (o *OmsStateChangeReject) ErrorMessageInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ErrorMessageSinceVersion()
}

func (*OmsStateChangeReject) ErrorMessageDeprecated() uint16 {
	return 0
}

func (*OmsStateChangeReject) ErrorMessageMetaAttribute(meta int) string {
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

func (*OmsStateChangeReject) ErrorMessageMinValue() byte {
	return byte(32)
}

func (*OmsStateChangeReject) ErrorMessageMaxValue() byte {
	return byte(126)
}

func (*OmsStateChangeReject) ErrorMessageNullValue() byte {
	return 0
}

func (o *OmsStateChangeReject) ErrorMessageCharacterEncoding() string {
	return "US-ASCII"
}
