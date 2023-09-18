package main

w = Wheel{Circle{Point{8,8},5},20}

w = Wheel{
    Circle: Circle{
        Point:  Point{X: 8, Y: 8},
        Radius: 5,
    },
    Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
}

fmt.Printf("%#v\n", w)

w.X = 42

fmt.Printf("%#v\n", w)
