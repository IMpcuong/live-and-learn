use std::mem;

// Tuples can be used as function arguments and as return values.
fn reverse(pair: (i32, bool)) -> (bool, i32) {
    // `let` can be used to bind the members of a tuple to variables
    let (integer, boolean) = pair;

    (boolean, integer)
}

// The following struct is for the activity.
#[derive(Debug)]
struct Matrix(f32, f32, f32, f32);

impl std::fmt::Display for Matrix {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        let first_half: &str = &format!("( {} {} )", self.0, self.1);
        let second_half: &str = &format!("( {} {} )", self.2, self.3);
        write!(f, "{}\n{}", first_half, second_half)
    }
}

fn transpose(matrix: &mut Matrix) -> Matrix {
    let matrix: Matrix = Matrix(matrix.0, matrix.2, matrix.1, matrix.3);
    matrix
    // Implicit return:
    // Matrix(matrix.0, matrix.2, matrix.1, matrix.3)
}

// This function borrows a slice.
fn analyze_slice(slice: &[i32]) {
    println!("first element of the slice: {}", slice[0]);
    println!("the slice has {} elements", slice.len());
}

fn main() {
    /// Exp1:
    // A tuple with a bunch of different types.
    let long_tuple = (1u8, 2u16, 3u32, 4u64,
                      -1i8, -2i16, -3i32, -4i64,
                      0.1f32, 0.2f64,
                      'a', true);

    // Values can be extracted from the tuple using tuple indexing.
    println!("long tuple first value: {}", long_tuple.0);
    println!("long tuple second value: {}", long_tuple.1);

    let num = 2;
    let string: &str = "";
    let tuple_test: (i32, String) = (123, String::from("Something"));
    let (num, string) = &tuple_test;
    println!("{} / {}", num, string); // 123 / Something

    // Tuples can be tuple members.
    let tuple_of_tuples = ((1u8, 2u16, 2u32), (4u64, -1i8), -2i16);

    // Tuples are printable.
    println!("tuple of tuples: {:?}", tuple_of_tuples);

    // But long Tuples (more than 12 elements) cannot be printed:
    // let too_long_tuple = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13);
    // println!("too long tuple: {:?}", too_long_tuple);
    // TODO: Uncomment the above 2 lines to see the compiler error.

    let pair = (1, true);
    println!("pair is {:?}", pair);

    println!("the reversed pair is {:?}", reverse(pair));

    // To create one element tuples, the comma is required to tell them apart
    // from a literal surrounded by parentheses.
    println!("one element tuple: {:?}", (5u32,));
    println!("just an integer: {:?}", (5u32));

    // Tuples can be destructured to create bindings.
    let tuple = (1, "hello", 4.5, true);

    let (a, b, c, d) = tuple;
    println!("{:?}, {:?}, {:?}, {:?}", a, b, c, d);

    let matrix = Matrix(1.1, 1.2, 2.1, 2.2);
    println!("{:?}", matrix);

    let mut n_matrix = Matrix(1.1, 1.2, 2.1, 2.2);
    println!("Matrix:\n{}", n_matrix);
    println!("Transpose:\n{}", transpose(&mut n_matrix));

    /// Exp2:
    // Fixed-size array (type signature is superfluous).
    let xs: [i32; 5] = [1, 2, 3, 4, 5];

    // All elements can be initialized to the same value.
    let ys: [i32; 500] = [0; 500];

    // Indexing starts at 0.
    println!("first element of the array: {}", xs[0]);

    // `len` returns the count of elements in the array.
    println!("number of elements in array: {}", xs.len());

    // Arrays are stack allocated
    println!("array occupies {} bytes", mem::size_of_val(&xs));

    // Arrays can be automatically borrowed as slices.
    println!("borrow the whole array as a slice");
    analyze_slice(&xs);

    // Slices can point to a section of an array:
    // They are of the form [starting_index..ending_index]
    // starting_index is the first position in the slice
    // ending_index is one more than the last position in the slice.
    println!("borrow a section of the array as a slice");
    analyze_slice(&ys[1 .. 4]);

    // Example of empty slice `&[]`
    let empty_array: [u32; 0] = [];
    assert_eq!(&empty_array, &[]);
    assert_eq!(&empty_array, &[][..]); // same but more verbose

    // Arrays can be safely accessed using `.get`, which returns an
    // `Option`. This can be matched as shown below, or used with
    // `.expect()` if you would like the program to exit with a nice
    // message instead of happily continue.
    for i in 0..xs.len() + 1 { // OOPS! one element too far.
        // match xs.get(i) {
        //     Some(xval) => println!("{}: {}", i, xval),
        //     None => println!("Slow down! {} is too far!", i),
        // }
        let val = xs.get(i).expect("err");
        println!("{i}: {val}");
    }

    // Out of bound indexing causes compile error:
    // println!("{}", xs[5]);
}