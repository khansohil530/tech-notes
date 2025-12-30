class Solution:
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        A, B = nums1, nums2
        if len(A) > len(B):
            A, B = B, A

        half = (len(A)+len(B))//2
        left, right = 0, len(A)-1
        while True:
            midA = left + (right-left)//2
            midB = half - midA - 2
            Aleft = A[midA] if midA >= 0 else float('-inf')
            Aright = A[midA+1] if midA+1 < len(A) else float('inf')
            Bleft = B[midB] if midB >= 0 else float('-inf')
            Bright = B[midB+1] if midB+1 < len(B) else float('inf')
            if Aleft <= Bright and Bleft <= Aright:
                if (len(A)+len(B))%2 == 0:
                    return (max(Aleft,Bleft) + min(Aright, Bright))/2
                else:
                    return min(Aright, Bright)
            elif Aleft > Bright:
                right = midA-1
            else:
                left = midA+1
        