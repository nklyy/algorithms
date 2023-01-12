#[allow(dead_code)]
pub fn recursive_max(lst: Vec<i64>) -> i64 {
    if lst.len() == 2 {
        if lst[0] > lst[1] {
            return lst[0];
        }

        return lst[1];
    }

    let max = recursive_max(lst[1..].to_vec());
    if lst[0] > max {
        return lst[0];
    }

    max
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_recursive_max_1() {
        let lst = vec![2, 4, 6];
        let result = recursive_max(lst);

        assert_eq!(result, 6)
    }

    #[test]
    fn test_recursive_max_2() {
        let lst = vec![-2, 4, -6];
        let result = recursive_max(lst);

        assert_eq!(result, 4)
    }
}
