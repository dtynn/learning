fn merge(v: &mut [i32], split_at: usize) {
    let size = v.len();
    if size <= 1 {
        return;
    }

    let mut tmp = Vec::with_capacity(v.len());
    let mut idx_left = 0;
    let mut idx_right = split_at;

    for _ in 0..size {
        if idx_right >= size {
            tmp.push(v[idx_left]);
            idx_left += 1;
            continue;
        }

        if idx_left >= split_at {
            tmp.push(v[idx_right]);
            idx_right += 1;
            continue;
        }

        if v[idx_left] > v[idx_right] {
            tmp.push(v[idx_right]);
            idx_right += 1;
        } else {
            tmp.push(v[idx_left]);
            idx_left += 1;
        }
    }

    for i in 0..size {
        v[i] = tmp[i];
    }
}

pub fn mergesort(v: &mut [i32]) {
    let size = v.len();
    if size <= 1 {
        return;
    }

    let split_at = size / 2;
    mergesort(&mut v[..split_at]);
    mergesort(&mut v[split_at..]);
    merge(&mut v[..], split_at);
}

#[cfg(test)]
mod test {
    use super::mergesort;

    #[test]
    fn mergesort_test() {
        let mut v1 = vec![0, 9, 7, 6, 3, 5, 2, 1, 4, 8];
        mergesort(&mut v1);
        assert_eq!(v1, vec![0, 1, 2, 3, 4, 5, 6, 7, 8, 9]);

        let mut v2 = vec![0, 7, 6, 3, 5, 2, 1, 4, 8];
        mergesort(&mut v2);
        assert_eq!(v2, vec![0, 1, 2, 3, 4, 5, 6, 7, 8]);

        let mut v3 = vec![8];
        mergesort(&mut v3);
        assert_eq!(v3, vec![8]);

        let mut v4 = vec![];
        mergesort(&mut v4);
        assert_eq!(v4, vec![]);
    }
}
