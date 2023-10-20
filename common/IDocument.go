package common

type DocumentHandle interface {
	GetFixedSizeBlob(FixedSizeBlobID int) []byte
	GetVariableSizeBlob(UvarBlobID int) []byte
	// AssertFact Sets or clears a fact associated with the document. The FactHandle
	// must have been previously registered in the IFactSet.
	AssertFact(FactHandle Fact, value bool)
}

type IDocument interface {
	// GetPostingCount number of postings created to determined which shard to use
	GetPostingCount() int
	// GetSourceByteSize Used to compute ingestion rate (bytes/second) statistic.
	GetSourceByteSize() int
	// Ingest Ingests the contents of this document into the index
	Ingest(handle DocumentHandle) error

	Contains(term string) bool
	// OpenStream Opens a named stream for term additions. Subsequent calls to
	// AddTerm() will add terms to this stream.
	OpenStream(streamID int)
	// AddTerm Adds a term to the currently opened stream.
	AddTerm(term string)
	CloseStream()
	CloseDocument()
}
