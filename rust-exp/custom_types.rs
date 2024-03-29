// An attribute to hide warnings for unused code.
#![allow(dead_code)]

#[derive(Debug)]
struct Person {
    name: String,
    age: u8,
}

// A unit struct.
struct Unit;

// A tuple struct.
struct Pair(i32, f32);

// A struct with two fields.
#[derive(Debug)]
struct Point {
    x: f32,
    y: f32,
}

// Structs can be reused as fields of another struct.
// Need to declare lifetime duration for this struct -> because of the
// heap type (Point) references -> the Rust rutime compiler didn't knew
// exactly when the Point struct will be allocated in the heap memory
// in the future -> `<'a>` is born to tell to the compiler that this prop's
// type will exist homogeneously at the compile time.

// #[derive(Debug)]
// struct Rectangle<'a> {
//     // A rectangle can be specified by where the top left and bottom right
//     // corners are in space.
//     top_left: &'a Point,
//     bottom_right: &'a Point,
// }

#[derive(Debug)]
struct Rectangle {
    top_left: Point,
    bottom_right: Point,
}

// fn rect_area(rect: &Rectangle) -> f32 {
//     let len: f32 = (*rect).top_left.y - (*rect).bottom_right.y;
//     let width: f32 = (*rect).bottom_right.x - (*rect).top_left.x;
//     let area = (len * width).abs();
//     area
// }

fn rect_area(rect: Rectangle) -> f32 {
    let len: f32 = rect.top_left.y - rect.bottom_right.y;
    let width: f32 = rect.bottom_right.x - rect.top_left.x;
    let area = (len * width).abs();
    area
}

// fn square<'a>(p: &Point, edge: f32) -> Rectangle<'a> {
//     // error[E0716]: Temporary value dropped while borrowed.
//     let top_left: &'a Point = &Point { x: p.x, y: p.y }; // --> Freed after use!
//     let bottom_right: &'a Point = &Point { x: p.x + edge, y: p.y - edge }; // --> Freed after use!
//
//     // Need to implement `std::AsRef<Point>` fo using associated function `.as_ref()`.
//     // error[E0515]: Cannot return value referencing local variable
//     // -> Both { top_left, bottom_right } are reference variables with different lifetime!
//     Rectangle {
//         top_left: top_left.as_ref(),
//         bottom_right: bottom_right.as_ref()
//     }
// }

fn square(p: Point, edge: f32) -> Rectangle {
    let top_left = Point { x: p.x, y: p.y };
    let bottom_right = Point { x: p.x + edge, y: p.y - edge };

    Rectangle {
        top_left: top_left,
        bottom_right: bottom_right
    }
}

fn main() {
    // Create struct with field init shorthand.
    let name = String::from("Peter");
    let age = 27;
    let peter = Person { name, age };

    // Print debug struct.
    println!("{:?}", peter);

    // Instantiate a `Point`
    let point: Point = Point { x: 10.3, y: 0.4 };

    // Access the fields of the point.
    println!("point coordinates: ({}, {})", point.x, point.y);

    // Make a new point by using struct update syntax to use the fields of our
    // other one.
    let bottom_right = Point { x: 5.2, ..point };
    println!("Bottom right = {:#?}", bottom_right);

    // `bottom_right.y` will be the same as `point.y` because we used that field.
    // from `point`
    println!("second point: ({}, {})", bottom_right.x, bottom_right.y);

    // Destructure the point using a `let` binding.
    let Point { x: left_edge, y: top_edge } = point;
    let top_left = Point { x: left_edge, y: top_edge };
    // let top_left = point; // Still Ok!
    println!("Top left = {:#?}", top_left);

    let rectangle = Rectangle {
        // struct instantiation is an expression too.
        top_left: top_left,
        bottom_right: bottom_right,
    };

    println!("Rectangle Area: {}", rect_area(rectangle));

    let n_rect = square(Point { x: 10.0, y: 11.0 }, 5.0);
    println!("Square Area: {}", rect_area(n_rect));

    // Instantiate a unit struct.
    let _unit = Unit;

    // Instantiate a tuple struct.
    let pair = Pair(1, 0.1);

    // Access the fields of a tuple struct.
    println!("pair contains {:?} and {:?}", pair.0, pair.1);

    // Destructure a tuple struct.
    let Pair(integer, decimal) = pair;

    println!("pair contains {:?} and {:?}", integer, decimal);
}