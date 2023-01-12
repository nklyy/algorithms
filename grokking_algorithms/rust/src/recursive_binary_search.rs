#[allow(dead_code)]
pub fn recursive_binary_search<T: PartialOrd>(
    lst: Vec<T>,
    item: T,
    low: usize,
    high: usize,
) -> Option<usize> {
    if low > high {
        return None;
    }

    let mid = (low + high) / 2;
    if lst[mid] == item {
        return Some(mid);
    }

    if lst[mid] < item {
        return recursive_binary_search(lst, item, mid + 1, high);
    }

    return recursive_binary_search(lst, item, low, mid - 1);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_recursive_binary_search_usize_found() {
        let lst = vec![
            1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
            25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
            47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
            69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
            91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
        ];
        let result = recursive_binary_search(lst.clone(), 30, 0, lst.len() - 1);

        assert_eq!(result, Some(29))
    }

    #[test]
    fn test_recursive_binary_search_usize_not_found() {
        let lst = vec![
            1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
            25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
            47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
            69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
            91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
        ];
        let result = recursive_binary_search(lst.clone(), 200, 0, lst.len() - 1);

        assert_eq!(result, None)
    }

    #[test]
    fn test_recursive_binary_search_string_found() {
        let lst = vec!["a", "b", "c"];
        let result = recursive_binary_search(lst.clone(), "c", 0, lst.len() - 1);

        assert_eq!(result, Some(2))
    }

    #[test]
    fn test_recursive_binary_search_string_not_found() {
        let lst = vec!["a", "b", "c"];
        let result = recursive_binary_search(lst.clone(), "d", 0, lst.len() - 1);

        assert_eq!(result, None)
    }
}
