package geo

import "testing"

func TestCircleCompare(t *testing.T) {
	// Different
	compareTestcases{
		{
			G1: &Circle{Radius: 1, Center: Point{0, 0}},
			G2: &Circle{Radius: 1, Center: Point{0, 2}},
		},
		{
			G1: &Circle{Radius: 1, Center: Point{0, 0}},
			G2: &Circle{Radius: 3, Center: Point{0, 0}},
		},
		{
			G1: &Circle{Radius: 1, Center: Point{0, 0}},
			G2: &Point{1, 1},
		},
		{
			G1: &Circle{Radius: 1, Center: Point{0, 0}},
			G2: &Line{{0, 0}, {1, 1}},
		},
		{
			G1: &Circle{Radius: 1, Center: Point{0, 0}},
			G2: &Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 1}, {0, 0}}},
		},
	}.fail(t)
}

func TestCircleMarshalJSON(t *testing.T) {
	marshalTestcases{
		{
			Input:    &Circle{Radius: 1.23, Center: Point{0, 0}},
			Expected: `{"type":"Circle","radius":1.23,"coordinates":[0,0]}`,
		},
	}.pass(t)
}

func TestCircleScan(t *testing.T) {
	// Pass
	for _, testcase := range []struct {
		WKT      interface{}
		Expected *Circle
	}{
		{
			WKT:      "CIRCULARSTRING(1 0, 0 1, -1 0, 0 -1, 1 0)",
			Expected: &Circle{Radius: 1, Center: Point{0, 0}},
		},
	} {
		c := &Circle{}
		if err := c.Scan(testcase.WKT); err != nil {
			t.Fatal(err)
		}
		if !c.Compare(testcase.Expected) {
			t.Fatalf("expected %v, got %v", c, testcase.Expected)
		}
	}

	// Fail
	for _, testcase := range []interface{}{
		"CIRCULARSTRING(1, 1, 2, 3, 4, 5)",              // bad comma
		[]byte("CIRFCRESOT(1.4 2.3, 3.6 8.2, 4.6 0.2)"), // typo
		[]byte("CIRCULARSTRING(1 0, 0 1)"),              // < 3 points
		7, // bad type
	} {
		c := &Circle{}
		if err := c.Scan(testcase); err == nil {
			t.Fatal("expected error, got nil")
		}
	}
}

func TestCircleValue(t *testing.T) {
	var (
		c        = &Circle{Radius: 2, Center: Point{0, 4}}
		expected = `CIRCULARSTRING(2 4, 0 6, -2 4, 0 2, 2 4)`
	)
	value, err := c.Value()
	if err != nil {
		t.Fatal(err)
	}
	got, ok := value.(string)
	if !ok {
		t.Fatalf("expected string, got %T", value)
	}
	if expected != got {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}