fn main() {
    // A mutable variable's value can be changed.
    let mut mutable = 12; // Mutable `i32`
    mutable = 21;

    // Error! The type of a variable can't be changed: `mutable = true;`
    mutable = true.to_string().len();
    println!("{mutable}");

    // Integer addition:
    println!("1 + 2 = {}", 1u32 + 2); // 1 + 2 = 3

    // Integer subtraction:
    println!("1 - 2 = {}", 1i32 - 2); // 1 - 2 = -1
    // TODO: Try changing `1i32` to `1u32` to see why the type is important

    // Short-circuiting boolean logic:
    println!("true AND false is {}", true && false); // true AND false is false
    println!("true  OR false is {}", true || false); // true OR false is true
    println!("NOT true is {}", !true); // NOT true is false

    // Bitwise operations:
    println!("0011 AND 0101 is {:04b}", 0b0011u32 & 0b0101); // 0011 AND 0101 is 0001
    println!("0011  OR 0101 is {:04b}", 0b0011u32 | 0b0101); // 0011  OR 0101 is 0111
    println!("0011 XOR 0101 is {:04b}", 0b0011u32 ^ 0b0101); // 0011 XOR 0101 is 0110
    println!("1 << 5 is {}", 1u32 << 5); // 1 << 5 is 32
    println!("0x80 >> 2 is 0x{:x}", 0x80u32 >> 2); // 0x80 >> 2 is 0x20

    let hex: u32 = 0x80 as u32;
    println!("{hex}"); // 128
    println!("0x80 >> 2 || 128/2^2 is {}", hex >> 2); // 0x80 >> 2 is 32

    // Use underscores to improve readability!
    println!("One million is written as {}", 1_000_000u32); // One million is written as 100000
}