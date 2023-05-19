#[allow(dead_code)]
// fn roman_to_int(s: String) -> i32 {
//     let symbols: HashMap<String, i32> = [
//         ("I", 1),
//         ("IV", 4),
//         ("V", 5),
//         ("IX", 9),
//         ("X", 10),
//         ("XL", 40),
//         ("L", 50),
//         ("XC", 90),
//         ("C", 100),
//         ("CD", 400),
//         ("D", 500),
//         ("CM", 900),
//         ("M", 1000),
//     ]
//     .iter()
//     .map(|(k, v)| (String::from(*k), *v))
//     .collect();

//     let mut res = 0;
//     let mut l = 0;

//     while l < s.len() {
//         if l + 1 >= s.len() {
//             res += symbols.get(&s.chars().nth(l).unwrap().to_string()).unwrap();
//             return res;
//         }

//         let key = format!(
//             "{}{}",
//             s.chars().nth(l).unwrap().to_string(),
//             s.chars().nth(l + 1).unwrap().to_string()
//         );

//         if symbols.contains_key(&key) {
//             res += symbols.get(&key).unwrap();
//             l += 2
//         } else {
//             res += symbols.get(&s.chars().nth(l).unwrap().to_string()).unwrap();
//             l += 1;
//         }
//     }

//     res
// }

fn roman_to_int(s: String) -> i32 {
    s.chars().rfold(0, |acc, c| {
        acc + match c {
            'I' if acc >= 5 => -1,
            'I' => 1,
            'V' => 5,
            'X' if acc >= 50 => -10,
            'X' => 10,
            'L' => 50,
            'C' if acc >= 500 => -100,
            'C' => 100,
            'D' => 500,
            'M' => 1000,
            _ => 0,
        }
    })
}

// 1890
#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn roman_to_int_1() {
        let result = roman_to_int("MCMXCIV".to_string());

        assert_eq!(result, 1994)
    }

    #[test]
    fn roman_to_int_2() {
        let result = roman_to_int("III".to_string());

        assert_eq!(result, 3)
    }
}
