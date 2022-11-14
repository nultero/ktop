package state

import "testing"

func Test_getBounds(t *testing.T) {

	testQuadrants := []Quad{QTopLeft, QTopRight, QBtmLeft, QBtmRight}

	dummySizes := [][]int{
		{120, 120},
		{120, 119},
		{119, 120},
		{119, 119},
	}
	rangeOfVals := [][][4]int{
		{ //top left
			{0, 60, 0, 60},
			{0, 60, 0, 59},
			{0, 59, 0, 60},
			{0, 59, 0, 59},
		},
		{ //top right
			{61, 120, 0, 60},
			{61, 120, 0, 59},
			{60, 119, 0, 60},
			{60, 119, 0, 59},
		},
		{ //btm left
			{0, 60, 61, 120},
			{0, 60, 60, 119},
			{0, 59, 61, 120},
			{0, 59, 60, 119},
		},
		{ //btm right
			{61, 120, 61, 120},
			{61, 120, 60, 119},
			{60, 119, 61, 120},
			{60, 119, 60, 119},
		},
	}

	for nth, tq := range testQuadrants {
		expectedVals := rangeOfVals[nth]
		for idx, dsz := range dummySizes {
			bounds := tq.GetBounds(dsz[0], dsz[1])
			for i, targetVal := range expectedVals[idx] {
				if bounds[i] != targetVal {
					t.Errorf("wanted %v,  got %v for boundary %d in quadrant %d", targetVal, bounds[i], i, tq)
				}
			}
		}
	}
}
