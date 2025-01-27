class Solution(object):
    def checkIfPrerequisite(self, numCourses, prerequisites, queries):
        """
        :type numCourses: int
        :type prerequisites: List[List[int]]
        :type queries: List[List[int]]
        :rtype: List[bool]
        """
        reachable = [[False] * numCourses for _ in range(numCourses)]
        
        for u, v in prerequisites:
            reachable[u][v] = True

        for k in range(numCourses):
            for i in range(numCourses):
                for j in range(numCourses):
                    if reachable[i][k] and reachable[k][j]:
                        reachable[i][j] = True
                        
        result = []
        for u, v in queries:
            result.append(reachable[u][v])
        
        return result