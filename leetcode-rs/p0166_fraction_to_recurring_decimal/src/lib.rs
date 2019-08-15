pub struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn fraction_to_decimal(numerator: i32, denominator: i32) -> String {
        if numerator == 0 {
            return "0".to_owned();
        }

        let mut n = (numerator as i64).abs();
        let d = (denominator as i64).abs();

        let mut bytes = Vec::new();
        if let Some(sign) = Solution::get_sign(numerator, denominator) {
            bytes.push(sign);
        }

        let int_part: Vec<u8> = (n / d).to_string().bytes().collect();
        bytes.extend(int_part);

        n = n % d;
        if n == 0 {
            return String::from_utf8(bytes).unwrap();
        }

        bytes.push(b'.');
        let mut decimal_parts: Vec<u8> = Vec::new();
        let mut repeats: HashMap<i64, usize> = HashMap::new();

        loop {
            if n == 0 {
                break;
            }

            if let Some(start_at) = repeats.get(&n).cloned() {
                decimal_parts.push(b'(');
                (&mut decimal_parts[start_at..].rotate_right(1));
                decimal_parts.push(b')');
                break;
            }

            repeats.insert(n, decimal_parts.len());
            n *= 10;
            if n < d {
                decimal_parts.push(b'0');
                continue;
            }

            let digit = b'0' + (n / d) as u8;
            decimal_parts.push(digit);

            n = n % d;
        }

        bytes.extend(decimal_parts);
        String::from_utf8(bytes).unwrap()
    }

    fn get_sign(n: i32, d: i32) -> Option<u8> {
        if n == 0 || (n < 0 && d < 0) || (n > 0 && d > 0) {
            return None;
        }

        Some(b'-')
    }
}
