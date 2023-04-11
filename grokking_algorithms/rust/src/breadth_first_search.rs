use std::{collections::HashMap, hash::Hash};

#[allow(dead_code)]
fn breadth_first_search<T>(graph: HashMap<T, Vec<T>>, v: T, t: T) -> bool
where
    T: PartialOrd + Eq + Hash + Clone,
{
    let mut queue: Vec<T> = graph.get(&v).unwrap().to_vec();
    let mut visited: Vec<T> = Vec::new();

    while queue.len() != 0 {
        let neighbor = queue.pop().unwrap();

        if !visited.contains(&neighbor) {
            if neighbor == t {
                return true;
            }

            queue.extend(graph.get(&neighbor).unwrap().to_vec());
            visited.push(neighbor);
        }
    }

    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_breadth_first_search_success() {
        let graph: HashMap<String, Vec<String>> = [
            ("A", vec!["E", "B", "C"]),
            ("E", vec!["I", "W"]),
            ("B", vec!["W", "D"]),
            ("C", vec!["U"]),
            ("I", vec!["Y", "U"]),
            ("W", vec![]),
            ("D", vec![]),
            ("Y", vec!["O"]),
            ("U", vec!["F"]),
            ("O", vec![]),
            ("F", vec![]),
        ]
        .iter()
        .map(|(k, v)| {
            (
                String::from(*k),
                v.iter().map(|s| String::from(*s)).collect(),
            )
        })
        .collect();

        let result = breadth_first_search(graph, "A".to_string(), "F".to_string());

        assert_eq!(result, true)
    }

    #[test]
    fn test_breadth_first_search_fail() {
        let graph: HashMap<String, Vec<String>> = [
            ("A", vec!["E", "B", "C"]),
            ("E", vec!["I", "W"]),
            ("B", vec!["W", "D"]),
            ("C", vec!["U"]),
            ("I", vec!["Y", "U"]),
            ("W", vec![]),
            ("D", vec![]),
            ("Y", vec!["O"]),
            ("U", vec!["F"]),
            ("O", vec![]),
            ("F", vec![]),
        ]
        .iter()
        .map(|(k, v)| {
            (
                String::from(*k),
                v.iter().map(|s| String::from(*s)).collect(),
            )
        })
        .collect();

        let result = breadth_first_search(graph, "B".to_string(), "F".to_string());

        assert_eq!(result, false)
    }
}
