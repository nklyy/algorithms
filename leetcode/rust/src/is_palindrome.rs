#[allow(dead_code)]
fn is_palindrome(x: i32) -> bool {
    if x < 0 {
        return false;
    }

    let mut res = 0;
    let mut n = x.clone();

    while n != 0 {
        let remain = n % 10;
        res = (res * 10) + remain;

        n /= 10;
    }

    x == res

    // one line solution
    // x.to_string().chars().rev().eq(x.to_string().chars())
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_is_palindrome_1() {
        let result = is_palindrome(121);

        assert_eq!(result, true)
    }

    #[test]
    fn test_is_palindrome_2() {
        let result = is_palindrome(123);

        assert_eq!(result, false)
    }
}
