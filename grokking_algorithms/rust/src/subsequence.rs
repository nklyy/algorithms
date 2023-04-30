use std::cmp;

#[allow(dead_code)]
fn subsequence(a: String, b: String) -> Vec<Vec<i64>> {
    let mut cell = vec![vec![0; b.len()]; a.len()];

    for i in 0..a.len() {
        for j in 0..b.len() {
            if a.chars().nth(i) == b.chars().nth(j) {
                if i > 0 && j > 0 {
                    cell[i][j] = cell[i - 1][j - 1] + 1
                } else {
                    cell[i][j] = 1
                }
            } else {
                if i > 0 && j > 0 {
                    cell[i][j] = cmp::max(cell[i - 1][j], cell[i][j - 1])
                } else {
                    cell[i][j] = 1
                }
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
        let result = subsequence("fish".to_string(), "fosh".to_string());

        assert_eq!(
            result,
            vec![
                vec![1, 1, 1, 1],
                vec![1, 1, 1, 1],
                vec![1, 1, 2, 2],
                vec![1, 1, 2, 3]
            ]
        )
    }
}
