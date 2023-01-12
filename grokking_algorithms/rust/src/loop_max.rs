#[allow(dead_code)]
fn loop_max(lst: Vec<i64>) -> i64 {
    let mut max = lst[0];

    for i in lst {
        if i > max {
            max = i
        }
    }

    max
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_loop_max_1() {
        let lst = vec![2, 4, 6];
        let result = loop_max(lst);

        assert_eq!(result, 6)
    }

    #[test]
    fn test_loop_max_2() {
        let lst = vec![-2, 4, -6];
        let result = loop_max(lst);

        assert_eq!(result, 4)
    }
}
