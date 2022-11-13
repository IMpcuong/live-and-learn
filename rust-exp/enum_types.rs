enum VeryVerboseEnumOfThingsToDoWithNumbers {
    Add,
    Subtract,
    Absolute,
}

type Shorten = VeryVerboseEnumOfThingsToDoWithNumbers;

impl Shorten {
    fn run(&self, x: i32, y: i32) -> i32 {
        match self {
            Self::Add => x + y,
            Self::Subtract =>
                if x > y {
                    x - y
                } else {
                    y - x
                },
            Self::Absolute => (x - y).abs(),
        }
    }
}

enum Status {
    Rich,
    Poor,
}

enum Work {
    Civilian,
    Soldier,
}

// enum with implicit discriminator (starts at 0).
enum Number {
    Zero,
    One,
    Two,
}

// enum with explicit discriminator.
enum Color {
    Red = 0xff0000,
    Green = 0x00ff00,
    Blue = 0x0000ff,
}

fn main() {
    let sample_add: i32 = Shorten::run(&Shorten::Add, 1, 2);
    let sample_sub: i32 = Shorten::run(&Shorten::Subtract, 1, 2);
    let sample_abs: i32 = Shorten::run(&Shorten::Absolute, 10, 20);
    println!("{}", sample_add);
    println!("{}", sample_sub);
    println!("{}", sample_abs);

    // Explicitly `use` each name so they are available without
    // manual scoping.
    use crate::Status::{Poor, Rich};
    // Automatically `use` each name inside `Work`.
    use crate::Work::*;

    match status {
        // Note the lack of scoping because of the explicit `use` above.
        Rich => println!("The rich have lots of money!"),
        Poor => println!("The poor have no money..."),
    }

    match work {
        // Note again the lack of scoping.
        Civilian => println!("Civilians work!"),
        Soldier  => println!("Soldiers fight!"),
    }

    // `enums` can be cast as integers.
    println!("zero is {}", Number::Zero as i32); // zero is 0
    println!("one is {}", Number::One as i32); // one is 1

    println!("roses are #{:06x}", Color::Red as i32); // roses are #ff0000
    println!("violets are #{:06x}", Color::Blue as i32); // violets are #0000ff
}