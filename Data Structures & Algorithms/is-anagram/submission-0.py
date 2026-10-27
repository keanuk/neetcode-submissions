class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        sMap = {}
        for char in s:
            sMap[char] = sMap.get(char, 0) + 1
        for char in t:
            if char not in sMap:
                return False
            sMap[char] = sMap.get(char) - 1
        for v in sMap.values():
            if v != 0:
                return False
        return True        
