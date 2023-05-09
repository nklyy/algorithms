#[allow(dead_code)]
fn binary_search(lst: Vec<i32>, target: i32) -> i32 {
    let mut left = 0;
    let mut right = lst.len() - 1;

    while left <= right {
        let mid = (left + right) / 2;

        if lst[mid] == target {
            return mid as i32;
        }

        if lst[mid] > target {
            right = mid - 1;
        } else {
            left = mid + 1;
        }
    }

    -1
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_binary_search_1() {
        let lst = vec![2, 4, 5];
        let result = binary_search(lst, 2);

        assert_eq!(result, 0)
    }

    #[test]
    fn test_binary_search_2() {
        let lst = vec![2, 4, 5];
        let result = binary_search(lst, 6);

        assert_eq!(result, -1)
    }
}
