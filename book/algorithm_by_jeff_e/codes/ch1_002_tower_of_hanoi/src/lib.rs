pub fn hanoi(n: u8, src: u8, dst: u8, tmp: u8) {
    if n > 0 {
        hanoi(n - 1, src, tmp, dst);
        println!("move disk {} from {} to {}", n, src, dst);
        hanoi(n - 1, tmp, dst, src);
    }
}

#[cfg(test)]
mod test {
    use super::hanoi;

    #[test]
    fn hanoi_test() {
        println!("");
        hanoi(4, 1, 3, 2);
    }
}
