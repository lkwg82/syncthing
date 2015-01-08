// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package protocol

import (
	"bytes"
	"io"

	"github.com/calmh/xdr"
)

/*

IndexMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Length of Folder                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                   Folder (variable length)                    \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Number of Files                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\               Zero or more FileInfo Structures                \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Flags                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Options                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Option Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct IndexMessage {
	string Folder<64>;
	FileInfo Files<>;
	unsigned int Flags;
	Option Options<64>;
}

*/

func (o IndexMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o IndexMessage) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o IndexMessage) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o IndexMessage) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o IndexMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.Folder); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Folder", l, 64)
	}
	xw.WriteString(o.Folder)
	xw.WriteUint32(uint32(len(o.Files)))
	for i := range o.Files {
		_, err := o.Files[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	xw.WriteUint32(o.Flags)
	if l := len(o.Options); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Options", l, 64)
	}
	xw.WriteUint32(uint32(len(o.Options)))
	for i := range o.Options {
		_, err := o.Options[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	return xw.Tot(), xw.Error()
}

func (o *IndexMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *IndexMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *IndexMessage) decodeXDR(xr *xdr.Reader) error {
	o.Folder = xr.ReadStringMax(64)
	_FilesSize := int(xr.ReadUint32())
	o.Files = make([]FileInfo, _FilesSize)
	for i := range o.Files {
		(&o.Files[i]).decodeXDR(xr)
	}
	o.Flags = xr.ReadUint32()
	_OptionsSize := int(xr.ReadUint32())
	if _OptionsSize > 64 {
		return xdr.ElementSizeExceeded("Options", _OptionsSize, 64)
	}
	o.Options = make([]Option, _OptionsSize)
	for i := range o.Options {
		(&o.Options[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

FileInfo Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Name                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Name (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Flags                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                      Modified (64 bits)                       +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                       Version (64 bits)                       +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                    Local Version (64 bits)                    +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Blocks                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\               Zero or more BlockInfo Structures               \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct FileInfo {
	string Name<8192>;
	unsigned int Flags;
	hyper Modified;
	unsigned hyper Version;
	unsigned hyper LocalVersion;
	BlockInfo Blocks<>;
}

*/

func (o FileInfo) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o FileInfo) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o FileInfo) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o FileInfo) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o FileInfo) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.Name); l > 8192 {
		return xw.Tot(), xdr.ElementSizeExceeded("Name", l, 8192)
	}
	xw.WriteString(o.Name)
	xw.WriteUint32(o.Flags)
	xw.WriteUint64(uint64(o.Modified))
	xw.WriteUint64(o.Version)
	xw.WriteUint64(o.LocalVersion)
	xw.WriteUint32(uint32(len(o.Blocks)))
	for i := range o.Blocks {
		_, err := o.Blocks[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	return xw.Tot(), xw.Error()
}

func (o *FileInfo) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *FileInfo) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *FileInfo) decodeXDR(xr *xdr.Reader) error {
	o.Name = xr.ReadStringMax(8192)
	o.Flags = xr.ReadUint32()
	o.Modified = int64(xr.ReadUint64())
	o.Version = xr.ReadUint64()
	o.LocalVersion = xr.ReadUint64()
	_BlocksSize := int(xr.ReadUint32())
	o.Blocks = make([]BlockInfo, _BlocksSize)
	for i := range o.Blocks {
		(&o.Blocks[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

FileInfoTruncated Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Name                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Name (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Flags                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                      Modified (64 bits)                       +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                       Version (64 bits)                       +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                    Local Version (64 bits)                    +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                          Num Blocks                           |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct FileInfoTruncated {
	string Name<8192>;
	unsigned int Flags;
	hyper Modified;
	unsigned hyper Version;
	unsigned hyper LocalVersion;
	unsigned int NumBlocks;
}

*/

func (o FileInfoTruncated) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o FileInfoTruncated) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o FileInfoTruncated) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o FileInfoTruncated) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o FileInfoTruncated) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.Name); l > 8192 {
		return xw.Tot(), xdr.ElementSizeExceeded("Name", l, 8192)
	}
	xw.WriteString(o.Name)
	xw.WriteUint32(o.Flags)
	xw.WriteUint64(uint64(o.Modified))
	xw.WriteUint64(o.Version)
	xw.WriteUint64(o.LocalVersion)
	xw.WriteUint32(o.NumBlocks)
	return xw.Tot(), xw.Error()
}

func (o *FileInfoTruncated) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *FileInfoTruncated) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *FileInfoTruncated) decodeXDR(xr *xdr.Reader) error {
	o.Name = xr.ReadStringMax(8192)
	o.Flags = xr.ReadUint32()
	o.Modified = int64(xr.ReadUint64())
	o.Version = xr.ReadUint64()
	o.LocalVersion = xr.ReadUint64()
	o.NumBlocks = xr.ReadUint32()
	return xr.Error()
}

/*

BlockInfo Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Size                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Hash                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Hash (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct BlockInfo {
	unsigned int Size;
	opaque Hash<64>;
}

*/

func (o BlockInfo) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o BlockInfo) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o BlockInfo) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o BlockInfo) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o BlockInfo) encodeXDR(xw *xdr.Writer) (int, error) {
	xw.WriteUint32(o.Size)
	if l := len(o.Hash); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Hash", l, 64)
	}
	xw.WriteBytes(o.Hash)
	return xw.Tot(), xw.Error()
}

func (o *BlockInfo) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *BlockInfo) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *BlockInfo) decodeXDR(xr *xdr.Reader) error {
	o.Size = xr.ReadUint32()
	o.Hash = xr.ReadBytesMax(64)
	return xr.Error()
}

/*

RequestMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Length of Folder                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                   Folder (variable length)                    \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Name                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Name (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                       Offset (64 bits)                        +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Size                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Hash                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Hash (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Flags                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Options                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Option Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct RequestMessage {
	string Folder<64>;
	string Name<8192>;
	unsigned hyper Offset;
	unsigned int Size;
	opaque Hash<64>;
	unsigned int Flags;
	Option Options<64>;
}

*/

func (o RequestMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o RequestMessage) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o RequestMessage) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o RequestMessage) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o RequestMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.Folder); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Folder", l, 64)
	}
	xw.WriteString(o.Folder)
	if l := len(o.Name); l > 8192 {
		return xw.Tot(), xdr.ElementSizeExceeded("Name", l, 8192)
	}
	xw.WriteString(o.Name)
	xw.WriteUint64(o.Offset)
	xw.WriteUint32(o.Size)
	if l := len(o.Hash); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Hash", l, 64)
	}
	xw.WriteBytes(o.Hash)
	xw.WriteUint32(o.Flags)
	if l := len(o.Options); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Options", l, 64)
	}
	xw.WriteUint32(uint32(len(o.Options)))
	for i := range o.Options {
		_, err := o.Options[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	return xw.Tot(), xw.Error()
}

func (o *RequestMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *RequestMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *RequestMessage) decodeXDR(xr *xdr.Reader) error {
	o.Folder = xr.ReadStringMax(64)
	o.Name = xr.ReadStringMax(8192)
	o.Offset = xr.ReadUint64()
	o.Size = xr.ReadUint32()
	o.Hash = xr.ReadBytesMax(64)
	o.Flags = xr.ReadUint32()
	_OptionsSize := int(xr.ReadUint32())
	if _OptionsSize > 64 {
		return xdr.ElementSizeExceeded("Options", _OptionsSize, 64)
	}
	o.Options = make([]Option, _OptionsSize)
	for i := range o.Options {
		(&o.Options[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

ResponseMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Data                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Data (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Error                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct ResponseMessage {
	opaque Data<>;
	unsigned int Error;
}

*/

func (o ResponseMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o ResponseMessage) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o ResponseMessage) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o ResponseMessage) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o ResponseMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	xw.WriteBytes(o.Data)
	xw.WriteUint32(o.Error)
	return xw.Tot(), xw.Error()
}

func (o *ResponseMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *ResponseMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *ResponseMessage) decodeXDR(xr *xdr.Reader) error {
	o.Data = xr.ReadBytes()
	o.Error = xr.ReadUint32()
	return xr.Error()
}

/*

ClusterConfigMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                     Length of Client Name                     |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                 Client Name (variable length)                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                   Length of Client Version                    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\               Client Version (variable length)                \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Folders                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Folder Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Options                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Option Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct ClusterConfigMessage {
	string ClientName<64>;
	string ClientVersion<64>;
	Folder Folders<64>;
	Option Options<64>;
}

*/

func (o ClusterConfigMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o ClusterConfigMessage) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o ClusterConfigMessage) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o ClusterConfigMessage) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o ClusterConfigMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.ClientName); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("ClientName", l, 64)
	}
	xw.WriteString(o.ClientName)
	if l := len(o.ClientVersion); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("ClientVersion", l, 64)
	}
	xw.WriteString(o.ClientVersion)
	if l := len(o.Folders); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Folders", l, 64)
	}
	xw.WriteUint32(uint32(len(o.Folders)))
	for i := range o.Folders {
		_, err := o.Folders[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	if l := len(o.Options); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Options", l, 64)
	}
	xw.WriteUint32(uint32(len(o.Options)))
	for i := range o.Options {
		_, err := o.Options[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	return xw.Tot(), xw.Error()
}

func (o *ClusterConfigMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *ClusterConfigMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *ClusterConfigMessage) decodeXDR(xr *xdr.Reader) error {
	o.ClientName = xr.ReadStringMax(64)
	o.ClientVersion = xr.ReadStringMax(64)
	_FoldersSize := int(xr.ReadUint32())
	if _FoldersSize > 64 {
		return xdr.ElementSizeExceeded("Folders", _FoldersSize, 64)
	}
	o.Folders = make([]Folder, _FoldersSize)
	for i := range o.Folders {
		(&o.Folders[i]).decodeXDR(xr)
	}
	_OptionsSize := int(xr.ReadUint32())
	if _OptionsSize > 64 {
		return xdr.ElementSizeExceeded("Options", _OptionsSize, 64)
	}
	o.Options = make([]Option, _OptionsSize)
	for i := range o.Options {
		(&o.Options[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

Folder Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of ID                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     ID (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Devices                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Device Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Folder {
	string ID<64>;
	Device Devices<>;
}

*/

func (o Folder) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Folder) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Folder) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o Folder) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o Folder) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.ID); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("ID", l, 64)
	}
	xw.WriteString(o.ID)
	xw.WriteUint32(uint32(len(o.Devices)))
	for i := range o.Devices {
		_, err := o.Devices[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	return xw.Tot(), xw.Error()
}

func (o *Folder) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Folder) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Folder) decodeXDR(xr *xdr.Reader) error {
	o.ID = xr.ReadStringMax(64)
	_DevicesSize := int(xr.ReadUint32())
	o.Devices = make([]Device, _DevicesSize)
	for i := range o.Devices {
		(&o.Devices[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

Device Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of ID                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     ID (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Flags                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                  Max Local Version (64 bits)                  +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Device {
	opaque ID<32>;
	unsigned int Flags;
	unsigned hyper MaxLocalVersion;
}

*/

func (o Device) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Device) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Device) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o Device) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o Device) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.ID); l > 32 {
		return xw.Tot(), xdr.ElementSizeExceeded("ID", l, 32)
	}
	xw.WriteBytes(o.ID)
	xw.WriteUint32(o.Flags)
	xw.WriteUint64(o.MaxLocalVersion)
	return xw.Tot(), xw.Error()
}

func (o *Device) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Device) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Device) decodeXDR(xr *xdr.Reader) error {
	o.ID = xr.ReadBytesMax(32)
	o.Flags = xr.ReadUint32()
	o.MaxLocalVersion = xr.ReadUint64()
	return xr.Error()
}

/*

Option Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of Key                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     Key (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Value                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Value (variable length)                    \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Option {
	string Key<64>;
	string Value<1024>;
}

*/

func (o Option) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Option) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Option) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o Option) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o Option) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.Key); l > 64 {
		return xw.Tot(), xdr.ElementSizeExceeded("Key", l, 64)
	}
	xw.WriteString(o.Key)
	if l := len(o.Value); l > 1024 {
		return xw.Tot(), xdr.ElementSizeExceeded("Value", l, 1024)
	}
	xw.WriteString(o.Value)
	return xw.Tot(), xw.Error()
}

func (o *Option) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Option) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Option) decodeXDR(xr *xdr.Reader) error {
	o.Key = xr.ReadStringMax(64)
	o.Value = xr.ReadStringMax(1024)
	return xr.Error()
}

/*

CloseMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Length of Reason                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                   Reason (variable length)                    \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Code                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct CloseMessage {
	string Reason<1024>;
	unsigned int Code;
}

*/

func (o CloseMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o CloseMessage) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o CloseMessage) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o CloseMessage) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o CloseMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	if l := len(o.Reason); l > 1024 {
		return xw.Tot(), xdr.ElementSizeExceeded("Reason", l, 1024)
	}
	xw.WriteString(o.Reason)
	xw.WriteUint32(o.Code)
	return xw.Tot(), xw.Error()
}

func (o *CloseMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *CloseMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *CloseMessage) decodeXDR(xr *xdr.Reader) error {
	o.Reason = xr.ReadStringMax(1024)
	o.Code = xr.ReadUint32()
	return xr.Error()
}

/*

EmptyMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct EmptyMessage {
}

*/

func (o EmptyMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o EmptyMessage) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o EmptyMessage) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o EmptyMessage) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.encodeXDR(xw)
	return []byte(aw), err
}

func (o EmptyMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	return xw.Tot(), xw.Error()
}

func (o *EmptyMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *EmptyMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *EmptyMessage) decodeXDR(xr *xdr.Reader) error {
	return xr.Error()
}
