use std::collections::HashMap;

// O(n2) - nested loops
#[allow(dead_code)]
pub fn two_sum_v1(nums: Vec<i64>, target: i64) -> Vec<i64> {
    for i in 0..nums.len() {
        for j in (i + 1)..nums.len() {
            if nums[i] + nums[j] == target {
                return vec![i as i64, j as i64];
            }
        }
    }

    vec![]
}

// O(n) - Two-pass Hash Table
#[allow(dead_code)]
pub fn two_sum_v2(nums: Vec<i64>, target: i64) -> Vec<i64> {
    let mut hash_map: HashMap<i64, i64> = HashMap::new();

    for i in 0..nums.len() {
        hash_map.insert(nums[i], i as i64);
    }

    for i in 0..nums.len() {
        let complement = target - nums[i];

        if let Some(v) = hash_map.get(&complement) {
            if *v != i as i64 {
                return vec![i as i64, *v];
            }
        }
    }

    vec![]
}

// O(n) - One-pass Hash Table
#[allow(dead_code)]
pub fn two_sum_v3(nums: Vec<i64>, target: i64) -> Vec<i64> {
    let mut hash_map: HashMap<i64, i64> = HashMap::new();

    for i in 0..nums.len() {
        let complement = target - nums[i];

        if let Some(v) = hash_map.get(&complement) {
            return vec![*v, i as i64];
        }

        hash_map.insert(nums[i], i as i64);
    }

    vec![]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum_v1() {
        let nums = vec![2, 7, 11, 15];
        let result = two_sum_v1(nums, 9);

        assert_eq!(result, vec![0, 1])
    }

    #[test]
    fn test_two_sum_v2() {
        let nums = vec![3, 2, 4];
        let result = two_sum_v2(nums, 6);

        assert_eq!(result, vec![1, 2])
    }

    #[test]
    fn test_two_sum_v3() {
        let nums = vec![2, 11, 7, 15];
        let result = two_sum_v3(nums, 9);

        assert_eq!(result, vec![0, 2])
    }
}
