type Rectangle struct {
    Width, Height float64
}

// Area method for Rectangle type
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Perimeter method for Rectangle type
func (r Rectangle) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}
