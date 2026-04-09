package drag

import "fmt"

const pir = 2.08551e-04

func (d Definition) RawDrag(mach float64) (float64, error) {
	switch d.kind {
	case KindStandardTable:
		return rawStandardTableDrag(d.table, mach)
	case KindCustomFunction:
		if d.rawFunction == nil {
			return 0, fmt.Errorf("custom drag function is nil")
		}
		return d.rawFunction(mach), nil
	default:
		return 0, fmt.Errorf("unsupported drag definition kind %d", d.kind)
	}
}

func (d Definition) ScaledDrag(mach float64) (float64, error) {
	raw, err := d.RawDrag(mach)
	if err != nil {
		return 0, err
	}

	return raw * pir, nil
}

func rawStandardTableDrag(table Table, mach float64) (float64, error) {
	reference, ok := standardTableReferences[table]
	if !ok {
		return 0, fmt.Errorf("standard drag data for %q not loaded", table)
	}

	for _, point := range reference.points {
		if point.Mach == mach {
			return point.Drag, nil
		}
	}

	return 0, fmt.Errorf("mach %.6f not found in standard drag table %q", mach, table)
}
