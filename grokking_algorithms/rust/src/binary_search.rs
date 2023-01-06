#[allow(dead_code)]
pub fn binary_search<T: PartialOrd>(lst: Vec<T>, item: T) -> Option<usize> {
    let mut low = 0;
    let mut high = lst.len() - 1;

    while low <= high {
        let mid = (low + high) / 2;

        if lst[mid] == item {
            return Some(mid);
        }

        // if taken value less than input we make shift to the right site (low = mid + 1)
        if lst[mid] < item {
            low = mid + 1;
        }

        // if taken value greater than input we make shift to the left site (high = mid - 1)
        if lst[mid] > item {
            high = mid - 1;
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_binary_search_usize_found() {
        let num_slice = vec![
            1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
            25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
            47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
            69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
            91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
        ];
        let result = binary_search(num_slice, 30);

        assert_eq!(result, Some(29))
    }

    #[test]
    fn test_binary_search_usize_not_found() {
        let num_slice = vec![
            1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
            25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
            47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
            69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
            91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
        ];
        let result = binary_search(num_slice, 200);

        assert_eq!(result, None)
    }

    #[test]
    fn test_binary_search_string_found() {
        let num_slice = vec!["a", "b", "c"];
        let result = binary_search(num_slice, "c");

        assert_eq!(result, Some(2))
    }

    #[test]
    fn test_binary_search_string_not_found() {
        let num_slice = vec!["a", "b", "c"];
        let result = binary_search(num_slice, "d");

        assert_eq!(result, None)
    }
}
