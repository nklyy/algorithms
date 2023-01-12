#[allow(dead_code)]
pub fn quick_sort<T: PartialOrd + Copy>(mut lst: Vec<T>) -> Vec<T> {
    if lst.len() < 2 {
        return lst;
    }

    if lst.len() == 2 {
        if lst[0] > lst[1] {
            lst.swap(1, 0);
            return lst;
        }

        return lst;
    }

    let pivot = lst[0];
    let mut less: Vec<T> = Vec::new();
    let mut greater: Vec<T> = Vec::new();

    for v in lst.iter().skip(1) {
        if *v < pivot {
            less.push(*v)
        } else {
            greater.push(*v)
        }
    }

    vec![quick_sort(less), vec![pivot], quick_sort(greater)].concat()
}

// pub fn quick_sort<T: PartialOrd + Copy>(lst: Vec<T>) -> Vec<T> {
//     if lst.len() < 2 {
//         return lst;
//     } else {
//         let pivot = lst[0];
//         let mut less: Vec<T> = Vec::new();
//         let mut greater: Vec<T> = Vec::new();

//         for v in lst.iter().skip(1) {
//             if *v < pivot {
//                 less.push(*v)
//             } else {
//                 greater.push(*v)
//             }
//         }

//         vec![quick_sort(less), vec![pivot], quick_sort(greater)].concat()
//     }
// }

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_quick_sort_uint() {
        let lst = vec![9, 03, 83, 9, 2, 0, 1, 65, 2, 822, 9, 11, 22, 3, 3, 3, 47];
        let result = quick_sort(lst);

        assert_eq!(
            result,
            vec![0, 1, 2, 2, 3, 3, 3, 3, 9, 9, 9, 11, 22, 47, 65, 83, 822]
        )
    }

    #[test]
    fn test_quick_sort_int() {
        let lst = vec![-12, -2, -77, -33, -1];
        let result = quick_sort(lst);

        assert_eq!(result, vec![-77, -33, -12, -2, -1])
    }

    #[test]
    fn test_quick_sort_float() {
        let lst = vec![0.2, 0.1, 0.5];
        let result = quick_sort(lst);

        assert_eq!(result, vec![0.1, 0.2, 0.5])
    }

    #[test]
    fn test_quick_sort_string() {
        let lst = vec!["c", "b", "f", "a"];
        let result = quick_sort(lst);

        assert_eq!(result, vec!["a", "b", "c", "f"])
    }
}
