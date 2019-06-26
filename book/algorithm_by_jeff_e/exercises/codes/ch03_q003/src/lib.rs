// dream dollars
pub fn fewest_bills(denominations: &[usize], target: usize) -> Option<Vec<usize>> {
    let deno_size = denominations.len();
    assert!(deno_size > 0);

    let mut solutions: Vec<Option<Vec<usize>>> = vec![None; target + 1];
    solutions[0] = Some(vec![]);
    // 从最小面值起, 才可能有解
    // 最小面值的解为 1
    let min_denomination = denominations[0];
    if target < min_denomination {
        return None;
    }

    // 恰好为面额的值, 最少张数都为 1
    for deno in denominations {
        if *deno > target {
            break;
        }

        solutions[*deno] = Some(vec![*deno]);
    }

    for k in min_denomination..target + 1 {
        let mut solution: Option<Vec<usize>> = None;
        for sub_solution in 0..k / 2 {
            let (s1, s2) = (&solutions[sub_solution], &solutions[k - sub_solution]);

            if s1.is_some() && s2.is_some() {
                let mut bills: Vec<usize> = s1.as_ref().cloned().unwrap();
                bills.extend_from_slice(s2.as_ref().unwrap());

                match solution {
                    Some(ref fewest) => {
                        if fewest.len() > bills.len() {
                            (&mut bills[..]).sort();
                            solution = Some(bills);
                        }
                    }

                    None => {
                        (&mut bills[..]).sort();
                        solution = Some(bills);
                    }
                }
            }
        }

        solutions[k] = solution;
    }

    solutions.pop().unwrap()
}

#[cfg(test)]
mod tests {
    #[test]
    fn fewest_bills_test() {
        let denominations = vec![4, 7, 13, 28, 52, 91, 365];
        let target = 1320;

        let res = super::fewest_bills(&denominations[..], target);
        println!("{:?}", res);
        if let Some(solution) = res {
            assert_eq!(solution.into_iter().sum::<usize>(), target);
        }
    }
}
