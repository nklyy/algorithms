#[allow(dead_code)]
pub fn recursive_sum(lst: Vec<i64>) -> i64 {
    if lst.len() == 1 {
        return lst[0];
    }

    return lst[0] + recursive_sum(lst[1..].to_vec());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_recursive_sum_1() {
        let lst = vec![2, 4, 6];
        let result = recursive_sum(lst);

        assert_eq!(result, 12)
    }

    #[test]
    fn test_recursive_sum_2() {
        let lst = vec![-2, 4, -6];
        let result = recursive_sum(lst);

        assert_eq!(result, -4)
    }
}
