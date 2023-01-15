from typing import List
from collections import defaultdict
from math import comb

class Solution:
    def numberOfGoodPaths(self, vals: List[int], edges: List[List[int]]) -> int:
        graph = {}

        def find(node):
            graph.setdefault(node, node)
            if node != graph[node]:
                graph[node] = find(graph[node])
            return graph[node]

        def union(node1 ,node2):
            graph[find(node1)] = find(node2)
        
        tree = defaultdict(list)
        val_node = defaultdict(set)
        for node1, node2 in edges:
            tree[node1].append(node2)
            tree[node2].append(node1)
            val_node[vals[node1]].add(node1)
            val_node[vals[node2]].add(node2)

        res = len(vals)

        for val in sorted(val_node.keys()):
            for node in val_node[val]:
                for neighbors in tree[node]:
                    if vals[neighbors] <= val:
                        union(node,neighbors)

            group_count = defaultdict(int)
            for node in val_node[val]:
                group_count[find(node)] += 1
                
            for root in group_count.keys():
                # Calculate the combination of nodes that can connect, equal to:
                # res += group_count[root] * (group_count[root]-1) // 2
                res += comb(group_count[root], 2)
        
        return res
