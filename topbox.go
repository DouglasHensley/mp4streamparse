package mp4streamparse

import (
	"fmt"
)

// TopBox mp4 root atom
type TopBox struct {
	Box
	ftyp *FtypBox
	styp *StypBox
	moov *MoovBox
	moof *MoofBox
	mdat *MdatBox
}

func (b TopBox) String() string {
	strMsg := "<< TopBox >>"
	strMsg = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", strMsg, b.ftyp, b.styp, b.moov, b.moof, b.mdat)
	return strMsg
}
