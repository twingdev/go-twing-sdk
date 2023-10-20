package common

// IFactSet is an abstract base class or an interface for classes that
// maintain a collection of FactInfo objects that describe document facts.
// A document fact represents a predicate or assertion that is either true
// or false for a given document, e.g. document is in top 20% static rank,
// or document was indexed before 5pm.
//
// Document facts are typically implemented as private, rank 0 rows in the
// index. Document facts may be mutable or immutable. An immutable fact for
// a particular document is considered constant for the life of the
// document. A mutable fact may be changed at any time.
type IFactSet interface {
	DefineFact(friendlyName string, mutable bool)
	GetCount() int

	GetFactInfoByHandle(f FactHandle) *FactInfo
}

type IFact interface {
	GetFriendlyName(fact FactHandle) string
	GetHandle() FactHandle
	IsMutable(fact FactHandle) bool
}

type FactSet struct {
	IFactSet
}

type Fact struct {
	FactInfo
}

type FactHandle uint64

type FactInfo struct {
	FriendlyName string
	IsMutable    bool
	FactHandle   func()
}

func (f *FactSet) DefineFact(friendlyName string, mutable bool) FactHandle {
	return 0
}

func (f *Fact) GetFriendlyName() string {
	return f.FriendlyName
}

func (f *Fact) IsMutable() string {
	return f.IsMutable()
}
