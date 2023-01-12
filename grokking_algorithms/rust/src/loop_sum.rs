#[allow(dead_code)]
fn loop_sum(lst: Vec<i64>) -> i64 {
    let mut r = 0;

    for v in lst {
        r += v;
    }

    r
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_loop_sum_1() {
        let lst = vec![2, 4, 6];
        let result = loop_sum(lst);

        assert_eq!(result, 12)
    }

    #[test]
    fn test_loop_sum_2() {
        let lst = vec![-2, 4, -6];
        let result = loop_sum(lst);

        assert_eq!(result, -4)
    }
}
