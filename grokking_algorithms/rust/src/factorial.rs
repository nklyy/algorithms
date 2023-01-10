#[allow(dead_code)]
pub fn factorial(x: usize) -> usize {
    if x == 1 {
        return 1;
    }

    return x * factorial(x - 1);
}

#[allow(dead_code)]
pub fn factorial_loop(x: usize) -> usize {
    if x == 0 || x == 1 {
        return 1;
    }

    let mut r = 1;
    for i in 1..(x + 1) {
        r = r * i;
    }

    r
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_factorial_1() {
        let x = 3;
        let result = factorial(x);

        assert_eq!(result, 6)
    }

    #[test]
    fn test_factorial_2() {
        let x = 5;
        let result = factorial(x);

        assert_eq!(result, 120)
    }

    #[test]
    fn test_factorial_loop_1() {
        let x = 3;
        let result = factorial_loop(x);

        assert_eq!(result, 6)
    }

    #[test]
    fn test_factorial_loop_2() {
        let x = 5;
        let result = factorial_loop(x);

        assert_eq!(result, 120)
    }

    #[test]
    fn test_factorial_loop_3() {
        let x = 10;
        let result = factorial_loop(x);

        assert_eq!(result, 3628800)
    }
}
