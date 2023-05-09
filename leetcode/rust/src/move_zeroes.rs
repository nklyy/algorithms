#[allow(dead_code)]
fn move_zeroes(lst: &mut Vec<i32>) {
    // let mut zeroes: Vec<i32> = Vec::new();

    // lst.retain(|x| {
    //     if *x != 0 {
    //         true
    //     } else {
    //         zeroes.push(0);
    //         false
    //     }
    // });

    // lst.append(&mut zeroes);

    let mut i = 0; // pos of the last non-zero
    let mut j = 0; // pos of the first element to check.

    while j < lst.len() {
        if lst[j] == 0 {
            j += 1;
        } else {
            lst.swap(i as usize, j);
            i += 1;
            j += 1;
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_move_zeroes_1() {
        let mut lst = vec![0, 1, 0, 3, 12];
        move_zeroes(&mut lst);

        assert_eq!(lst, vec![1, 3, 12, 0, 0])
    }
}
