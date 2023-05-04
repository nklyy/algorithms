#[allow(dead_code)]
fn search_insert(lst: Vec<i64>, target: i64) -> i64 {
    if lst.is_empty() {
        return 0;
    }

    if target < lst[0] {
        return 0;
    }

    let mut left = 0;
    let mut right = lst.len() - 1;

    while left <= right {
        let mid = (left + right) / 2;

        if lst[mid] == target {
            return mid as i64;
        }

        if lst[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    left as i64
}

// #[allow(dead_code)]
// fn search_insert(lst: Vec<i64>, target: i64) -> i64 {
//     lst.binary_search(&target).unwrap_or_else(|x| x) as i64
// }

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_search_insert_1() {
        let lst = vec![1, 3, 5, 6];
        let result = search_insert(lst, 5);

        assert_eq!(result, 2)
    }

    #[test]
    fn test_search_insert_2() {
        let lst = vec![1, 3, 5, 6];
        let result = search_insert(lst, 7);

        assert_eq!(result, 4)
    }
}
