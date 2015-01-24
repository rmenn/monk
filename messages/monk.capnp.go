package messages

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/glycerine/go-capnproto"
	"unsafe"
)

type Pupil C.Struct

func NewPupil(s *C.Segment) Pupil      { return Pupil(s.NewStruct(0, 2)) }
func NewRootPupil(s *C.Segment) Pupil  { return Pupil(s.NewRootStruct(0, 2)) }
func AutoNewPupil(s *C.Segment) Pupil  { return Pupil(s.NewStructAR(0, 2)) }
func ReadRootPupil(s *C.Segment) Pupil { return Pupil(s.Root(0).ToStruct()) }
func (s Pupil) Uuid() string           { return C.Struct(s).GetObject(0).ToText() }
func (s Pupil) SetUuid(v string)       { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Pupil) Url() string            { return C.Struct(s).GetObject(1).ToText() }
func (s Pupil) SetUrl(v string)        { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s Pupil) MarshalJSON() (bs []byte, err error) { return }

type Pupil_List C.PointerList

func NewPupilList(s *C.Segment, sz int) Pupil_List { return Pupil_List(s.NewCompositeList(0, 2, sz)) }
func (s Pupil_List) Len() int                      { return C.PointerList(s).Len() }
func (s Pupil_List) At(i int) Pupil                { return Pupil(C.PointerList(s).At(i).ToStruct()) }
func (s Pupil_List) ToArray() []Pupil              { return *(*[]Pupil)(unsafe.Pointer(C.PointerList(s).ToArray())) }
func (s Pupil_List) Set(i int, item Pupil)         { C.PointerList(s).Set(i, C.Object(item)) }

type Response C.Struct

func NewResponse(s *C.Segment) Response      { return Response(s.NewStruct(8, 0)) }
func NewRootResponse(s *C.Segment) Response  { return Response(s.NewRootStruct(8, 0)) }
func AutoNewResponse(s *C.Segment) Response  { return Response(s.NewStructAR(8, 0)) }
func ReadRootResponse(s *C.Segment) Response { return Response(s.Root(0).ToStruct()) }
func (s Response) Success() bool             { return C.Struct(s).Get1(0) }
func (s Response) SetSuccess(v bool)         { C.Struct(s).Set1(0, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s Response) MarshalJSON() (bs []byte, err error) { return }

type Response_List C.PointerList

func NewResponseList(s *C.Segment, sz int) Response_List {
	return Response_List(s.NewCompositeList(8, 0, sz))
}
func (s Response_List) Len() int          { return C.PointerList(s).Len() }
func (s Response_List) At(i int) Response { return Response(C.PointerList(s).At(i).ToStruct()) }
func (s Response_List) ToArray() []Response {
	return *(*[]Response)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s Response_List) Set(i int, item Response) { C.PointerList(s).Set(i, C.Object(item)) }
