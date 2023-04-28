use std::collections::HashMap;

#[allow(dead_code)]
struct TopologyGraph {
    graph: HashMap<i64, Vec<i64>>,
    v: i64,
}

#[allow(dead_code)]
impl TopologyGraph {
    pub fn new(v: i64) -> Self {
        Self {
            graph: HashMap::new(),
            v: v,
        }
    }
}

#[allow(dead_code)]
impl TopologyGraph {
    fn add_edge(&mut self, u: i64, v: i64) {
        match self.graph.get_mut(&u) {
            Some(n) => n.push(v.clone()),
            None => {
                self.graph.insert(u, vec![v]);
            }
        }
    }

    fn topology_sort_util(&self, v: i64, visited: &mut Vec<bool>, stack: &mut Vec<i64>) {
        visited.insert(v as usize, true);

        if let Some(neighbors) = self.graph.get(&v) {
            for neighbor in neighbors {
                if !visited.get(*neighbor as usize).unwrap() {
                    self.topology_sort_util(*neighbor, visited, stack)
                }
            }
        }

        stack.push(v)
    }

    fn topology_sort(&self) -> Vec<i64> {
        let mut visited = vec![false; self.v as usize];
        let mut stack: Vec<i64> = Vec::new();

        for i in 0..self.v {
            if !*visited.get(i as usize).unwrap() {
                self.topology_sort_util(i, &mut visited, &mut stack);
            }
        }

        stack
    }
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_topology_sort_1() {
        let mut graph = TopologyGraph::new(6);
        graph.add_edge(5, 2);
        graph.add_edge(5, 0);
        graph.add_edge(4, 0);
        graph.add_edge(4, 1);
        graph.add_edge(2, 3);
        graph.add_edge(3, 1);

        let result = graph.topology_sort();

        assert_eq!(result, vec![0, 1, 3, 2, 4, 5])
    }

    #[test]
    fn test_topology_sort_2() {
        let mut graph = TopologyGraph::new(4);
        graph.add_edge(0, 1);
        graph.add_edge(0, 2);
        graph.add_edge(1, 2);
        graph.add_edge(2, 0);
        graph.add_edge(2, 3);
        graph.add_edge(3, 3);

        let result = graph.topology_sort();

        assert_eq!(result, vec![3, 2, 1, 0])
    }
}
