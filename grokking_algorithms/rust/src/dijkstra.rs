use std::collections::VecDeque;

#[allow(dead_code)]
struct QPair {
    node: i64,
    w: i64,
}

#[allow(dead_code)]
impl QPair {
    fn new(node: i64, w: i64) -> Self {
        Self { node: node, w: w }
    }
}
#[allow(dead_code)]
#[derive(Clone, Copy, Debug)]
struct DijkstraEdge {
    to: i64,
    w: i64,
}

#[allow(dead_code)]
impl DijkstraEdge {
    fn new(to: i64, w: i64) -> Self {
        Self { to: to, w: w }
    }
}

#[allow(dead_code)]
struct DijkstraGraph {
    graph: Vec<Vec<DijkstraEdge>>,
    v: i64,
}

#[allow(dead_code)]
impl DijkstraGraph {
    pub fn new(v: i64) -> Self {
        Self {
            graph: vec![vec![]; v as usize],
            v: v,
        }
    }
}

#[allow(dead_code)]
impl DijkstraGraph {
    fn add_edge(&mut self, f: i64, s: i64, w: i64) {
        self.graph[f as usize].push(DijkstraEdge::new(s, w));
        self.graph[s as usize].push(DijkstraEdge::new(f, w));
    }

    fn dijkstra(&self, src: i64) -> Vec<i64> {
        let mut distances: Vec<i64> = vec![0; self.v as usize];

        for i in 0..self.v {
            distances[i as usize] = i64::MAX;
        }

        let mut queue: VecDeque<QPair> = VecDeque::new();
        queue.push_back(QPair::new(src, 0));

        distances[src as usize] = 0;

        while queue.len() > 0 {
            let q_pair = queue.pop_front().unwrap();

            for edge in self.graph[q_pair.node as usize].iter() {
                if distances[edge.to as usize] > distances[q_pair.node as usize] + edge.w {
                    distances[edge.to as usize] = distances[q_pair.node as usize] + edge.w;
                    queue.push_back(QPair::new(edge.to, distances[edge.to as usize]));
                }
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
    fn test_dijkstra_1() {
        let mut graph = DijkstraGraph::new(4);
        graph.add_edge(0, 1, 5);
        graph.add_edge(1, 2, 6);
        graph.add_edge(2, 3, 2);
        graph.add_edge(0, 2, 15);

        graph.dijkstra(0);

        let result = graph.dijkstra(0);

        assert_eq!(result, vec![0, 5, 11, 13])
    }

    #[test]
    fn test_dijkstra_2() {
        let mut graph = DijkstraGraph::new(4);
        graph.add_edge(0, 1, 24);
        graph.add_edge(0, 3, 20);
        graph.add_edge(2, 0, 3);
        graph.add_edge(3, 2, 12);

        graph.dijkstra(0);

        let result = graph.dijkstra(0);

        assert_eq!(result, vec![0, 24, 3, 15])
    }

    #[test]
    fn test_dijkstra_3() {
        let mut graph = DijkstraGraph::new(9);
        graph.add_edge(0, 1, 4);
        graph.add_edge(0, 7, 8);
        graph.add_edge(1, 2, 8);
        graph.add_edge(1, 7, 11);
        graph.add_edge(2, 3, 7);
        graph.add_edge(2, 8, 2);
        graph.add_edge(2, 5, 4);
        graph.add_edge(3, 4, 9);
        graph.add_edge(3, 5, 14);
        graph.add_edge(4, 5, 10);
        graph.add_edge(5, 6, 2);
        graph.add_edge(6, 7, 1);
        graph.add_edge(6, 8, 6);
        graph.add_edge(7, 8, 7);

        graph.dijkstra(0);

        let result = graph.dijkstra(0);

        assert_eq!(result, vec![0, 4, 12, 19, 21, 11, 9, 8, 14])
    }
}
