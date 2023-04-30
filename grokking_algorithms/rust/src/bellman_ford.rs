#[allow(dead_code)]
struct BellmanFordEdge {
    src: i64,
    dest: i64,
    weigh: i64,
}

#[allow(dead_code)]
impl BellmanFordEdge {
    fn new(src: i64, dest: i64, weigh: i64) -> Self {
        Self {
            src: src,
            dest: dest,
            weigh: weigh,
        }
    }
}

#[allow(dead_code)]
struct BellmanFordGraph {
    graph: Vec<BellmanFordEdge>,
    v: i64,
}

#[allow(dead_code)]
impl BellmanFordGraph {
    pub fn new(v: i64) -> Self {
        Self {
            graph: Vec::new(),
            v: v,
        }
    }
}

#[allow(dead_code)]
impl BellmanFordGraph {
    fn add_edge(&mut self, s: i64, d: i64, w: i64) {
        self.graph.push(BellmanFordEdge::new(s, d, w));
    }
}

#[allow(dead_code)]
impl BellmanFordGraph {
    fn bellman_ford(&self, src: i64) -> Vec<i64> {
        let mut distances: Vec<i64> = vec![0; self.v as usize];

        for i in 0..self.v {
            distances[i as usize] = i64::MAX;
        }

        distances[src as usize] = 0;

        for _ in 0..self.v - 1 {
            for edge in &self.graph {
                if distances[edge.src as usize] != i64::MAX
                    && distances[edge.dest as usize] > distances[edge.src as usize] + edge.weigh
                {
                    distances[edge.dest as usize] = distances[edge.src as usize] + edge.weigh;
                }
            }
        }

        for edge in &self.graph {
            if distances[edge.src as usize] != i64::MAX
                && distances[edge.dest as usize] > distances[edge.src as usize] + edge.weigh
            {
                println!("Graph contains negative weight cycle");
                return vec![];
            }
        }

        distances
    }
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_bellman_ford_1() {
        let mut graph = BellmanFordGraph::new(5);
        graph.add_edge(0, 1, -1);
        graph.add_edge(0, 2, 4);
        graph.add_edge(1, 2, 3);
        graph.add_edge(1, 3, 2);
        graph.add_edge(1, 4, 2);
        graph.add_edge(3, 2, 5);
        graph.add_edge(3, 1, 1);
        graph.add_edge(4, 3, -3);

        let result = graph.bellman_ford(0);

        assert_eq!(result, vec![0, -1, 2, -2, 1])
    }

    #[test]
    fn test_bellman_ford_2() {
        let mut graph = BellmanFordGraph::new(4);
        graph.add_edge(0, 1, 1);
        graph.add_edge(1, 2, -1);
        graph.add_edge(2, 3, -1);
        graph.add_edge(3, 0, -1);

        let result = graph.bellman_ford(0);

        assert_eq!(result, vec![])
    }
}
