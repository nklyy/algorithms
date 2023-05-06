#[allow(dead_code)]
fn first_bad_version(n: i64) -> i64 {
    let mut left = 0;
    let mut right = n;

    while left < right {
        let mid = left + (right - left) / 2;

        if is_bad_version(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    left
}

fn is_bad_version(_v: i64) -> bool {
    true
}
