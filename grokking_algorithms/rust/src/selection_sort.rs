#[allow(dead_code)]
pub fn selection_sort<T: PartialOrd + Copy>(mut lst: Vec<T>) -> Vec<T> {
    for i in 0..lst.len() {
        for j in i + 1..lst.len() {
            if lst[j] < lst[i] {
                lst.swap(i, j)
            }
        }
    }
    lst
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_selection_sort_uint() {
        let lst = vec![12, 2, 77, 33, 1, 0];
        let result = selection_sort(lst);

        assert_eq!(result, vec![0, 1, 2, 12, 33, 77])
    }

    #[test]
    fn test_selection_sort_int() {
        let lst = vec![-12, -2, -77, -33, -1];
        let result = selection_sort(lst);

        assert_eq!(result, vec![-77, -33, -12, -2, -1])
    }

    #[test]
    fn test_selection_sort_float() {
        let lst = vec![0.2, 0.1, 0.5];
        let result = selection_sort(lst);

        assert_eq!(result, vec![0.1, 0.2, 0.5])
    }

    #[test]
    fn test_selection_sort_string() {
        let lst = vec!["c", "b", "f", "a"];
        let result = selection_sort(lst);

        assert_eq!(result, vec!["a", "b", "c", "f"])
    }
}
