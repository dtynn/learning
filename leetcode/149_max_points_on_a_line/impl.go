package impl

// Point data structure
type Point struct {
	X int
	Y int
}

func maxPoints(points []Point) int {
	if len(points) == 0 {
		return 0
	}

	maxCount := 1
	linePoints := map[line]map[Point]struct{}{}
	horizonCount := map[int]int{}
	verticalCount := map[int]int{}
	pointCount := map[Point]int{}

	plast := points[len(points)-1]
	pointCount[plast] = 1
	horizonCount[plast.Y] = 1
	verticalCount[plast.X] = 1

	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]

		pcnt := pointCount[p1] + 1
		if pcnt > maxCount {
			maxCount = pcnt
		}
		pointCount[p1] = pcnt

		hcnt := horizonCount[p1.Y] + 1
		if hcnt > maxCount {
			maxCount = hcnt
		}
		horizonCount[p1.Y] = hcnt

		vcnt := verticalCount[p1.X] + 1
		if vcnt > maxCount {
			maxCount = vcnt
		}
		verticalCount[p1.X] = vcnt

		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			if p1 == p2 {
				continue
			}

			l := calcLine(p1, p2)
			if _, ok := linePoints[l]; !ok {
				linePoints[l] = map[Point]struct{}{}
			}

			linePoints[l][p1] = struct{}{}
			linePoints[l][p2] = struct{}{}
		}
	}

	for _, m := range linePoints {
		count := 0
		for p := range m {
			count += pointCount[p]
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func calcLine(p1, p2 Point) line {
	b := fraction{
		numerator:  p2.Y - p1.Y,
		denomiator: p2.X - p1.X,
	}.nom()

	p := choose(p1, p2)

	a := b.mul(-p.X).add(p.Y).nom()

	return line{
		a: a,
		b: b,
	}
}

func choose(p1, p2 Point) Point {
	if p1.X != 0 || p1.Y != 0 {
		return p1
	}

	return p2
}

type line struct {
	a fraction
	b fraction
}

type fraction struct {
	numerator  int
	denomiator int
}

func (f fraction) add(num int) fraction {
	return fraction{
		numerator:  f.numerator + f.denomiator*num,
		denomiator: f.denomiator,
	}
}

func (f fraction) minus(num int) fraction {
	return f.add(-num)
}

func (f fraction) mul(num int) fraction {
	return fraction{
		numerator:  f.numerator * num,
		denomiator: f.denomiator,
	}
}

func (f fraction) div(num int) fraction {
	return fraction{
		numerator:  f.numerator,
		denomiator: f.denomiator * num,
	}
}

func (f fraction) turn() fraction {
	return fraction{
		numerator:  f.denomiator,
		denomiator: f.numerator,
	}
}

func (f fraction) nom() fraction {
	numerator, denomiator := f.numerator, f.denomiator

	div := min(abs(numerator), abs(denomiator))
	for div > 1 {
		if numerator%div == 0 && denomiator%div == 0 {
			numerator /= div
			denomiator /= div

			div = min(abs(numerator), abs(denomiator))
			continue
		}

		div--
	}

	if denomiator < 0 {
		numerator = -numerator
		denomiator = -denomiator
	}

	return fraction{
		numerator:  numerator,
		denomiator: denomiator,
	}
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}
