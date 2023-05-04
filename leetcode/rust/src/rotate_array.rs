#[allow(dead_code)]
fn rotate_array(lst: &mut Vec<i64>, k: i64) {
    let k = k as usize % lst.len();

    lst.reverse();
    lst[..k].reverse();
    lst[k..].reverse();
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_rotate_array_1() {
        let mut lst = vec![1, 2, 3, 4, 5, 6, 7];
        rotate_array(&mut lst, 3);

        assert_eq!(lst, vec![5, 6, 7, 1, 2, 3, 4])
    }

    #[test]
    fn test_rotate_array_2() {
        let mut lst = vec![-1, -100, 3, 99];
        rotate_array(&mut lst, 2);

        assert_eq!(lst, vec![3, 99, -1, -100])
    }
}
