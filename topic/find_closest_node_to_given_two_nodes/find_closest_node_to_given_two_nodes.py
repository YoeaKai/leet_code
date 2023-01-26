from typing import List


class Solution:
    def closestMeetingNode(self, edges: List[int], node1: int, node2: int) -> int:
        res = -1

        def dfs(node, distance):
            if distance[node] == -1:
                distance[node] = 0
            neighbor = edges[node]
            if neighbor != -1 and distance[neighbor] == -1:
                distance[neighbor] = distance[node] + 1
                dfs(neighbor, distance)

        distance1 = [-1] * len(edges)
        distance2 = [-1] * len(edges)
        dfs(node1, distance1)
        dfs(node2, distance2)

        min_distance = float("inf")
        for currNode in range(len(edges)):
            if distance1[currNode] != -1 and distance2[currNode] != -1 and min_distance > max(distance1[currNode], distance2[currNode]):
                min_distance = max(distance1[currNode], distance2[currNode])
                res = currNode

        return res
