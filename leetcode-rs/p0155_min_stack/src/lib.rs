struct MinStack {
    mins: Vec<i32>,
    nums: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {
    /** initialize your data structure here. */
    fn new() -> Self {
        MinStack {
            mins: vec![],
            nums: vec![],
        }
    }

    fn push(&mut self, x: i32) {
        let mut min = x;
        let size = self.mins.len();
        if size > 0 {
            min = min.min(self.mins[size - 1]);
        }

        self.mins.push(min);
        self.nums.push(x);
    }

    fn pop(&mut self) {
        self.nums.pop();
        self.mins.pop();
    }

    fn top(&self) -> i32 {
        let size = self.nums.len();
        assert!(size > 0);
        self.nums[size - 1]
    }

    fn get_min(&self) -> i32 {
        let size = self.mins.len();
        assert!(size > 0);
        self.mins[size - 1]
    }
}
