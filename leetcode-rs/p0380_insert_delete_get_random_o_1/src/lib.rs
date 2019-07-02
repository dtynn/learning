struct RandomizedSet {
    start_at: std::time::Instant,
    nums: Vec<Option<i32>>,
    empty_idx: Vec<usize>,
    num_map: std::collections::HashMap<i32, usize>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl RandomizedSet {
    /** Initialize your data structure here. */
    fn new() -> Self {
        RandomizedSet {
            start_at: std::time::Instant::now(),
            nums: vec![],
            empty_idx: vec![],
            num_map: std::collections::HashMap::new(),
        }
    }

    /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
    fn insert(&mut self, val: i32) -> bool {
        if self.num_map.contains_key(&val) {
            return false;
        }

        match self.empty_idx.pop() {
            Some(idx) => {
                self.num_map.insert(val, idx);
                self.nums[idx] = Some(val);
            }

            None => {
                let idx = self.nums.len();
                self.num_map.insert(val, idx);
                self.nums.push(Some(val));
            }
        }

        true
    }

    /** Removes a value from the set. Returns true if the set contained the specified element. */
    fn remove(&mut self, val: i32) -> bool {
        if let Some(idx) = self.num_map.remove(&val) {
            self.nums[idx] = None;
            self.empty_idx.push(idx);
            return true;
        }

        false
    }

    /** Get a random element from the set. */
    fn get_random(&mut self) -> i32 {
        let size = self.nums.len();
        let start = (self.start_at.elapsed().subsec_nanos() as usize) % size;
        for i in 0..size {
            let idx = (start + i) % size;
            if let Some(val) = self.nums[idx] {
                return val;
            }
        }

        unreachable!();
    }
}
