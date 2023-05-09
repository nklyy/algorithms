#[allow(dead_code)]
fn sorted_squares(lst: Vec<i32>) -> Vec<i32> {
    let mut left = 0;
    let mut right = lst.len() - 1;

    let mut res = vec![0; lst.len()];

    for i in (0..lst.len()).rev() {
        if i32::abs(lst[left]) > i32::abs(lst[right]) {
            res[i] = lst[left] * lst[left];

            left += 1;
        } else {
            res[i] = lst[right] * lst[right];

            right -= 1;
        }
    }

    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_sorted_squares_1() {
        let lst = vec![-4, -1, 0, 3, 10];
        let result = sorted_squares(lst);

        assert_eq!(result, vec![0, 1, 9, 16, 100])
    }
}
