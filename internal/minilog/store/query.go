package store

import (
	"github.com/remicaumette/minilog/pkg/record"
	"strings"
)

func (s *Store) Query(query string) ([]*record.Record, error) {
	var found []*record.Record

	iter := s.db.NewIterator(nil, nil)
	for iter.Next() {
		curr := record.FromBinary(iter.Value())
		// levenshtein.DistanceDamerauLevenshtein(curr.Line, query) > len(curr.Line) - (len(query) / 2)
		if strings.Contains(curr.Line, query) {
			found = append(found, curr)
		}
	}
	return found, nil
}
