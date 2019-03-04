fn partition(v: &mut [i32], pivot: usize) -> usize {
    let size = v.len();
    let mut r = 0usize;
    v.swap(0, pivot);

    let mut head = 1usize;
    let mut tail = size - 1;
    while head <= tail {
        if v[head] < v[0] {
            r = head;
            head += 1;
        } else {
            v.swap(head, tail);
            tail -= 1;
        }
    }

    v.swap(0, r);
    r
}

pub fn quicksort(v: &mut [i32]) {
    let size = v.len();
    if size <= 1 {
        return;
    }

    // choose a pivot
    let pivot = 0;

    let r = partition(v, pivot);
    quicksort(&mut v[..r]);
    quicksort(&mut v[r + 1..]);
}

#[cfg(test)]
mod test {
    use super::quicksort;

    #[test]
    fn quicksort_test() {
        let mut v1 = vec![0, 9, 7, 6, 3, 5, 2, 1, 4, 8];
        quicksort(&mut v1);
        assert_eq!(v1, vec![0, 1, 2, 3, 4, 5, 6, 7, 8, 9]);

        let mut v2 = vec![0, 7, 6, 3, 5, 2, 1, 4, 8];
        quicksort(&mut v2);
        assert_eq!(v2, vec![0, 1, 2, 3, 4, 5, 6, 7, 8]);

        let mut v3 = vec![8];
        quicksort(&mut v3);
        assert_eq!(v3, vec![8]);

        let mut v4 = vec![];
        quicksort(&mut v4);
        assert_eq!(v4, vec![]);

        let mut v5 = vec![7, 9];
        quicksort(&mut v5);
        assert_eq!(v5, vec![7, 9]);

        let mut v6 = vec![9, 7];
        quicksort(&mut v6);
        assert_eq!(v6, vec![7, 9]);
    }

}
