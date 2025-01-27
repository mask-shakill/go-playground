class Solution:
    def findTheDifference(self, s, t):
        sum_t = sum(ord(c) for c in t)
        sum_s = sum(ord(c) for c in s)
        return chr(sum_t - sum_s)
if __name__ == "__main__":
    s1, t1 = "abcd", "abcde"
    s2, t2 = "", "y"
    
    solution = Solution()
    
    print(solution.findTheDifference(s1, t1))  
    print(solution.findTheDifference(s2, t2))  
