from typing import List


class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        unlock = [False] * len(rooms)
        unlock[0] = True
        keys = rooms[0]

        while keys:
            key = keys.pop()
            if unlock[key] == False:
                unlock[key] = True
                keys.extend(rooms[key])

        return False not in unlock
