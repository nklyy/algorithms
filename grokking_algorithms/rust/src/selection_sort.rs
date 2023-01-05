#[allow(dead_code)]
pub fn selection_sort<T: PartialOrd + Copy>(mut arr: Vec<T>) -> Vec<T> {
    for i in 0..arr.len() {
        for j in i + 1..arr.len() {
            if arr[j] < arr[i] {
                arr.swap(i, j)
            }
        }
    }
    arr
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn selection_sort_uint() {
        let num_slice = vec![12, 2, 77, 33, 1, 0];
        let result = selection_sort(num_slice);

        assert_eq!(result, vec![0, 1, 2, 12, 33, 77])
    }

    #[test]
    fn selection_sort_int() {
        let num_slice = vec![-12, -2, -77, -33, -1];
        let result = selection_sort(num_slice);

        assert_eq!(result, vec![-77, -33, -12, -2, -1])
    }

    #[test]
    fn selection_sort_float() {
        let num_slice = vec![0.2, 0.1, 0.5];
        let result = selection_sort(num_slice);

        assert_eq!(result, vec![0.1, 0.2, 0.5])
    }

    #[test]
    fn selection_sort_string() {
        let num_slice = vec!["c", "b", "f", "a"];
        let result = selection_sort(num_slice);

        assert_eq!(result, vec!["a", "b", "c", "f"])
    }
}
