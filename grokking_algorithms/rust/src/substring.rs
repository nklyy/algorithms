#[allow(dead_code)]
fn substring(a: String, b: String) -> Vec<Vec<i64>> {
    let mut cell = vec![vec![0; b.len()]; a.len()];

    for i in 0..a.len() {
        for j in 0..b.len() {
            if a.chars().nth(i) == b.chars().nth(j) {
                if i > 0 && j > 0 {
                    cell[i][j] = cell[i - 1][j - 1] + 1
                } else {
                    cell[i][j] = 0
                }
            } else {
                cell[i][j] = 0
            }
        }
    }

    cell
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_substring_1() {
        let result = substring("fish".to_string(), "hish".to_string());

        assert_eq!(
            result,
            vec![
                vec![0, 0, 0, 0],
                vec![0, 1, 0, 0],
                vec![0, 0, 2, 0],
                vec![0, 0, 0, 3]
            ]
        )
    }
}
