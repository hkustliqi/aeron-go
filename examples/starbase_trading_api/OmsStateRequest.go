// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsStateRequest struct {
	RequestHeader       OmsAdminRequestHeader
	LatestLogPositionId int64
	StateId             int64
	StateType           StateTypeEnum
	StateAction         StateActionEnum
	Status              StateStatusEnum
	Name                [80]byte
	StateMessageHeader  MessageHeader
}

func (o *OmsStateRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.RequestHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LatestLogPositionId); err != nil {
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

func (o *OmsStateRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.RequestHeaderInActingVersion(actingVersion) {
		if err := o.RequestHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.LatestLogPositionIdInActingVersion(actingVersion) {
		o.LatestLogPositionId = o.LatestLogPositionIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LatestLogPositionId); err != nil {
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

func (o *OmsStateRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.LatestLogPositionIdInActingVersion(actingVersion) {
		if o.LatestLogPositionId < o.LatestLogPositionIdMinValue() || o.LatestLogPositionId > o.LatestLogPositionIdMaxValue() {
			return fmt.Errorf("Range check failed on o.LatestLogPositionId (%v < %v > %v)", o.LatestLogPositionIdMinValue(), o.LatestLogPositionId, o.LatestLogPositionIdMaxValue())
		}
	}
	if o.StateIdInActingVersion(actingVersion) {
		if o.StateId != o.StateIdNullValue() && (o.StateId < o.StateIdMinValue() || o.StateId > o.StateIdMaxValue()) {
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

func OmsStateRequestInit(o *OmsStateRequest) {
	o.StateId = math.MinInt64
	return
}

func (*OmsStateRequest) SbeBlockLength() (blockLength uint16) {
	return 132
}

func (*OmsStateRequest) SbeTemplateId() (templateId uint16) {
	return 1010
}

func (*OmsStateRequest) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsStateRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsStateRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsStateRequest) RequestHeaderId() uint16 {
	return 1
}

func (*OmsStateRequest) RequestHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsStateRequest) RequestHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestHeaderSinceVersion()
}

func (*OmsStateRequest) RequestHeaderDeprecated() uint16 {
	return 0
}

func (*OmsStateRequest) RequestHeaderMetaAttribute(meta int) string {
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

func (*OmsStateRequest) LatestLogPositionIdId() uint16 {
	return 2
}

func (*OmsStateRequest) LatestLogPositionIdSinceVersion() uint16 {
	return 0
}

func (o *OmsStateRequest) LatestLogPositionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LatestLogPositionIdSinceVersion()
}

func (*OmsStateRequest) LatestLogPositionIdDeprecated() uint16 {
	return 0
}

func (*OmsStateRequest) LatestLogPositionIdMetaAttribute(meta int) string {
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

func (*OmsStateRequest) LatestLogPositionIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateRequest) LatestLogPositionIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateRequest) LatestLogPositionIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateRequest) StateIdId() uint16 {
	return 3
}

func (*OmsStateRequest) StateIdSinceVersion() uint16 {
	return 0
}

func (o *OmsStateRequest) StateIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateIdSinceVersion()
}

func (*OmsStateRequest) StateIdDeprecated() uint16 {
	return 0
}

func (*OmsStateRequest) StateIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsStateRequest) StateIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsStateRequest) StateIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsStateRequest) StateIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsStateRequest) StateTypeId() uint16 {
	return 4
}

func (*OmsStateRequest) StateTypeSinceVersion() uint16 {
	return 0
}

func (o *OmsStateRequest) StateTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateTypeSinceVersion()
}

func (*OmsStateRequest) StateTypeDeprecated() uint16 {
	return 0
}

func (*OmsStateRequest) StateTypeMetaAttribute(meta int) string {
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

func (*OmsStateRequest) StateActionId() uint16 {
	return 5
}

func (*OmsStateRequest) StateActionSinceVersion() uint16 {
	return 0
}

func (o *OmsStateRequest) StateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateActionSinceVersion()
}

func (*OmsStateRequest) StateActionDeprecated() uint16 {
	return 0
}

func (*OmsStateRequest) StateActionMetaAttribute(meta int) string {
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

func (*OmsStateRequest) StatusId() uint16 {
	return 6
}

func (*OmsStateRequest) StatusSinceVersion() uint16 {
	return 0
}

func (o *OmsStateRequest) StatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StatusSinceVersion()
}

func (*OmsStateRequest) StatusDeprecated() uint16 {
	return 0
}

func (*OmsStateRequest) StatusMetaAttribute(meta int) string {
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

func (*OmsStateRequest) NameId() uint16 {
	return 7
}

func (*OmsStateRequest) NameSinceVersion() uint16 {
	return 0
}

func (o *OmsStateRequest) NameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NameSinceVersion()
}

func (*OmsStateRequest) NameDeprecated() uint16 {
	return 0
}

func (*OmsStateRequest) NameMetaAttribute(meta int) string {
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

func (*OmsStateRequest) NameMinValue() byte {
	return byte(32)
}

func (*OmsStateRequest) NameMaxValue() byte {
	return byte(126)
}

func (*OmsStateRequest) NameNullValue() byte {
	return 0
}

func (o *OmsStateRequest) NameCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsStateRequest) StateMessageHeaderId() uint16 {
	return 8
}

func (*OmsStateRequest) StateMessageHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsStateRequest) StateMessageHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StateMessageHeaderSinceVersion()
}

func (*OmsStateRequest) StateMessageHeaderDeprecated() uint16 {
	return 0
}

func (*OmsStateRequest) StateMessageHeaderMetaAttribute(meta int) string {
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
