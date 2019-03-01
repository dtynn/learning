pub fn multiply(x: u32, y: u32) -> u32 {
    if x == 0 {
        return x;
    }

    let mut prod = multiply(x / 2, y + y);
    if x % 2 != 0 {
        prod += y;
    }

    prod
}
