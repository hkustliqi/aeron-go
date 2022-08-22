// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"io"
	"io/ioutil"
)

type AdminUserState struct {
	Role AdminUserRoleEnum
}

func (a *AdminUserState) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := a.RangeCheck(a.SbeSchemaVersion(), a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := a.Role.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (a *AdminUserState) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if a.RoleInActingVersion(actingVersion) {
		if err := a.Role.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > a.SbeSchemaVersion() && blockLength > a.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-a.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := a.RangeCheck(actingVersion, a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (a *AdminUserState) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := a.Role.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func AdminUserStateInit(a *AdminUserState) {
	return
}

func (*AdminUserState) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (*AdminUserState) SbeTemplateId() (templateId uint16) {
	return 13
}

func (*AdminUserState) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*AdminUserState) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*AdminUserState) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*AdminUserState) RoleId() uint16 {
	return 1
}

func (*AdminUserState) RoleSinceVersion() uint16 {
	return 0
}

func (a *AdminUserState) RoleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.RoleSinceVersion()
}

func (*AdminUserState) RoleDeprecated() uint16 {
	return 0
}

func (*AdminUserState) RoleMetaAttribute(meta int) string {
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
