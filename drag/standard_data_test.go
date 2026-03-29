package drag

import (
	"fmt"
	"strings"
	"testing"
)

func TestStandardTableReferencesCoverEverySupportedTable(t *testing.T) {
	tables := []Table{
		TableG1,
		TableG2,
		TableG5,
		TableG6,
		TableG7,
		TableG8,
		TableGS,
		TableRA4,
	}

	for _, table := range tables {
		t.Run(string(table), func(t *testing.T) {
			reference, ok := standardTableReferences[table]
			if !ok {
				t.Fatalf("standardTableReferences[%q] missing", table)
			}

			if reference.sourceURL == "" {
				t.Fatalf("sourceURL for %q is empty", table)
			}

			if len(reference.points) < 2 {
				t.Fatalf("len(points) for %q = %d, want at least 2", table, len(reference.points))
			}

			for i := 1; i < len(reference.points); i++ {
				if reference.points[i-1].Mach >= reference.points[i].Mach {
					t.Fatalf("points for %q are not strictly increasing at index %d: %v >= %v", table, i, reference.points[i-1].Mach, reference.points[i].Mach)
				}
			}
		})
	}
}

func TestStandardTableReferencesUseDistinctPointSets(t *testing.T) {
	fingerprints := make(map[string]Table)

	for table, reference := range standardTableReferences {
		fingerprint := fingerprintPoints(reference.points)
		if prior, exists := fingerprints[fingerprint]; exists {
			t.Fatalf("tables %q and %q share identical point sets", prior, table)
		}
		fingerprints[fingerprint] = table
	}
}

func fingerprintPoints(points []CurvePoint) string {
	var builder strings.Builder
	for _, point := range points {
		fmt.Fprintf(&builder, "%.6f:%.6f;", point.Mach, point.Drag)
	}
	return builder.String()
}
