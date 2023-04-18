use std::{collections::HashMap, fmt::Debug, hash::Hash};

pub fn depth_first_search<T>(graph: &HashMap<T, Vec<T>>, v: &T, t: &T, visited: &mut Vec<T>) -> bool
where
    T: PartialOrd + Eq + Hash + Clone + Debug,
{
    if v == t {
        return true;
    }

    if visited.contains(&v) {
        return false;
    }

    visited.push(v.clone());

    if let Some(neighbors) = graph.get(v) {
        for neighbor in neighbors {
            if !visited.contains(neighbor) {
                if depth_first_search(graph, neighbor, t, visited) {
                    return true;
                }
            }
        }
    }

    return false;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_depth_first_search_success() {
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

        let mut visited: Vec<String> = Vec::new();
        let result = depth_first_search(&graph, &"A".to_string(), &"F".to_string(), &mut visited);

        assert_eq!(result, true)
    }

    #[test]
    fn test_depth_first_search_fail() {
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

        let mut visited: Vec<String> = Vec::new();
        let result = depth_first_search(&graph, &"B".to_string(), &"F".to_string(), &mut visited);

        assert_eq!(result, false)
    }
}
