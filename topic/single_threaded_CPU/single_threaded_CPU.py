from typing import List
import heapq


class Solution:
    def getOrder(self, tasks: List[List[int]]) -> List[int]:
        tasks = sorted([(t[0], t[1], i) for i, t in enumerate(tasks)])
        res = []
        available_task = []
        idx = 0
        time = tasks[0][0]

        while len(res) < len(tasks):
            while (idx < len(tasks)) and (tasks[idx][0] <= time):
                heapq.heappush(available_task, (tasks[idx][1], tasks[idx][2]))
                idx += 1
            if available_task:
                processing_time, original_idx = heapq.heappop(available_task)
                time += processing_time
                res.append(original_idx)
            elif idx < len(tasks):
                time = tasks[idx][0]

        return res
