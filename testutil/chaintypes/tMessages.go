package chaintypes

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Messages struct {
	x []_Bytes
}
type Messages = *_Messages
type _Messages__Maybe struct {
	m schema.Maybe
	v Messages
}
type MaybeMessages = *_Messages__Maybe

func (m MaybeMessages) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMessages) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMessages) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMessages) AsNode() ipld.Node {
	switch m.m {
	case schema.Maybe_Absent:
		return ipld.Absent
	case schema.Maybe_Null:
		return ipld.Null
	case schema.Maybe_Value:
		return m.v
	default:
		panic("unreachable")
	}
}
func (m MaybeMessages) Must() Messages {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var _ ipld.Node = (Messages)(&_Messages{})
var _ schema.TypedNode = (Messages)(&_Messages{})

func (Messages) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (Messages) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"chaintypes.Messages"}.LookupByString("")
}
func (n Messages) LookupByNode(k ipld.Node) (ipld.Node, error) {
	idx, err := k.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(idx)
}
func (n Messages) LookupByIndex(idx int) (ipld.Node, error) {
	if n.Length() <= idx {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfInt(idx)}
	}
	v := &n.x[idx]
	return v, nil
}
func (n Messages) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "chaintypes.Messages", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (Messages) MapIterator() ipld.MapIterator {
	return nil
}
func (n Messages) ListIterator() ipld.ListIterator {
	return &_Messages__ListItr{n, 0}
}

type _Messages__ListItr struct {
	n   Messages
	idx int
}

func (itr *_Messages__ListItr) Next() (idx int, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.x) {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	idx = itr.idx
	x := &itr.n.x[itr.idx]
	v = x
	itr.idx++
	return
}
func (itr *_Messages__ListItr) Done() bool {
	return itr.idx >= len(itr.n.x)
}

func (n Messages) Length() int {
	return len(n.x)
}
func (Messages) IsAbsent() bool {
	return false
}
func (Messages) IsNull() bool {
	return false
}
func (Messages) AsBool() (bool, error) {
	return mixins.List{"chaintypes.Messages"}.AsBool()
}
func (Messages) AsInt() (int, error) {
	return mixins.List{"chaintypes.Messages"}.AsInt()
}
func (Messages) AsFloat() (float64, error) {
	return mixins.List{"chaintypes.Messages"}.AsFloat()
}
func (Messages) AsString() (string, error) {
	return mixins.List{"chaintypes.Messages"}.AsString()
}
func (Messages) AsBytes() ([]byte, error) {
	return mixins.List{"chaintypes.Messages"}.AsBytes()
}
func (Messages) AsLink() (ipld.Link, error) {
	return mixins.List{"chaintypes.Messages"}.AsLink()
}
func (Messages) Prototype() ipld.NodePrototype {
	return _Messages__Prototype{}
}

type _Messages__Prototype struct{}

func (_Messages__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Messages__Builder
	nb.Reset()
	return &nb
}

type _Messages__Builder struct {
	_Messages__Assembler
}

func (nb *_Messages__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Messages__Builder) Reset() {
	var w _Messages
	var m schema.Maybe
	*nb = _Messages__Builder{_Messages__Assembler{w: &w, m: &m}}
}

type _Messages__Assembler struct {
	w     *_Messages
	m     *schema.Maybe
	state laState

	cm schema.Maybe
	va _Bytes__Assembler
}

func (na *_Messages__Assembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_Messages__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"chaintypes.Messages"}.BeginMap(0)
}
func (na *_Messages__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_Messages{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_Bytes, 0, sizeHint)
	}
	return na, nil
}
func (na *_Messages__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"chaintypes.Messages"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Messages__Assembler) AssignBool(bool) error {
	return mixins.ListAssembler{"chaintypes.Messages"}.AssignBool(false)
}
func (_Messages__Assembler) AssignInt(int) error {
	return mixins.ListAssembler{"chaintypes.Messages"}.AssignInt(0)
}
func (_Messages__Assembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"chaintypes.Messages"}.AssignFloat(0)
}
func (_Messages__Assembler) AssignString(string) error {
	return mixins.ListAssembler{"chaintypes.Messages"}.AssignString("")
}
func (_Messages__Assembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"chaintypes.Messages"}.AssignBytes(nil)
}
func (_Messages__Assembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"chaintypes.Messages"}.AssignLink(nil)
}
func (na *_Messages__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Messages); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_List {
		return ipld.ErrWrongKind{TypeName: "chaintypes.Messages", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
	}
	itr := v.ListIterator()
	for !itr.Done() {
		_, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Messages__Assembler) Prototype() ipld.NodePrototype {
	return _Messages__Prototype{}
}
func (la *_Messages__Assembler) valueFinishTidy() bool {
	switch la.cm {
	case schema.Maybe_Value:
		la.va.w = nil
		la.cm = schema.Maybe_Absent
		la.state = laState_initial
		la.va.reset()
		return true
	default:
		return false
	}
}
func (la *_Messages__Assembler) AssembleValue() ipld.NodeAssembler {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when still in the middle of assembling the previous value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	la.w.x = append(la.w.x, _Bytes{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.w = row
	la.va.m = &la.cm
	return &la.va
}
func (la *_Messages__Assembler) Finish() error {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	la.state = laState_finished
	*la.m = schema.Maybe_Value
	return nil
}
func (la *_Messages__Assembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _Bytes__Prototype{}
}
func (Messages) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Messages) Representation() ipld.Node {
	return (*_Messages__Repr)(n)
}

type _Messages__Repr _Messages

var _ ipld.Node = &_Messages__Repr{}

func (_Messages__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_Messages__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"chaintypes.Messages.Repr"}.LookupByString("")
}
func (nr *_Messages__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (Messages)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Bytes).Representation(), nil
}
func (nr *_Messages__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	v, err := (Messages)(nr).LookupByIndex(idx)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Bytes).Representation(), nil
}
func (n _Messages__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "chaintypes.Messages.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_Messages__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (nr *_Messages__Repr) ListIterator() ipld.ListIterator {
	return &_Messages__ReprListItr{(Messages)(nr), 0}
}

type _Messages__ReprListItr _Messages__ListItr

func (itr *_Messages__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	idx, v, err = (*_Messages__ListItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return idx, v.(Bytes).Representation(), nil
}
func (itr *_Messages__ReprListItr) Done() bool {
	return (*_Messages__ListItr)(itr).Done()
}

func (rn *_Messages__Repr) Length() int {
	return len(rn.x)
}
func (_Messages__Repr) IsAbsent() bool {
	return false
}
func (_Messages__Repr) IsNull() bool {
	return false
}
func (_Messages__Repr) AsBool() (bool, error) {
	return mixins.List{"chaintypes.Messages.Repr"}.AsBool()
}
func (_Messages__Repr) AsInt() (int, error) {
	return mixins.List{"chaintypes.Messages.Repr"}.AsInt()
}
func (_Messages__Repr) AsFloat() (float64, error) {
	return mixins.List{"chaintypes.Messages.Repr"}.AsFloat()
}
func (_Messages__Repr) AsString() (string, error) {
	return mixins.List{"chaintypes.Messages.Repr"}.AsString()
}
func (_Messages__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"chaintypes.Messages.Repr"}.AsBytes()
}
func (_Messages__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"chaintypes.Messages.Repr"}.AsLink()
}
func (_Messages__Repr) Prototype() ipld.NodePrototype {
	return _Messages__ReprPrototype{}
}

type _Messages__ReprPrototype struct{}

func (_Messages__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _Messages__ReprBuilder
	nb.Reset()
	return &nb
}

type _Messages__ReprBuilder struct {
	_Messages__ReprAssembler
}

func (nb *_Messages__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Messages__ReprBuilder) Reset() {
	var w _Messages
	var m schema.Maybe
	*nb = _Messages__ReprBuilder{_Messages__ReprAssembler{w: &w, m: &m}}
}

type _Messages__ReprAssembler struct {
	w     *_Messages
	m     *schema.Maybe
	state laState

	cm schema.Maybe
	va _Bytes__ReprAssembler
}

func (na *_Messages__ReprAssembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_Messages__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"chaintypes.Messages.Repr"}.BeginMap(0)
}
func (na *_Messages__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_Messages{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_Bytes, 0, sizeHint)
	}
	return na, nil
}
func (na *_Messages__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"chaintypes.Messages.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Messages__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"chaintypes.Messages.Repr"}.AssignBool(false)
}
func (_Messages__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"chaintypes.Messages.Repr"}.AssignInt(0)
}
func (_Messages__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"chaintypes.Messages.Repr"}.AssignFloat(0)
}
func (_Messages__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"chaintypes.Messages.Repr"}.AssignString("")
}
func (_Messages__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"chaintypes.Messages.Repr"}.AssignBytes(nil)
}
func (_Messages__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"chaintypes.Messages.Repr"}.AssignLink(nil)
}
func (na *_Messages__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Messages); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.ReprKind() != ipld.ReprKind_List {
		return ipld.ErrWrongKind{TypeName: "chaintypes.Messages.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
	}
	itr := v.ListIterator()
	for !itr.Done() {
		_, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Messages__ReprAssembler) Prototype() ipld.NodePrototype {
	return _Messages__ReprPrototype{}
}
func (la *_Messages__ReprAssembler) valueFinishTidy() bool {
	switch la.cm {
	case schema.Maybe_Value:
		la.va.w = nil
		la.cm = schema.Maybe_Absent
		la.state = laState_initial
		la.va.reset()
		return true
	default:
		return false
	}
}
func (la *_Messages__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when still in the middle of assembling the previous value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	la.w.x = append(la.w.x, _Bytes{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.w = row
	la.va.m = &la.cm
	return &la.va
}
func (la *_Messages__ReprAssembler) Finish() error {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	la.state = laState_finished
	*la.m = schema.Maybe_Value
	return nil
}
func (la *_Messages__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _Bytes__ReprPrototype{}
}
