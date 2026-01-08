class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:
    DELIMITER = "$"

    def serialize(self, root):
        if not root: return ""
        res = []
        stk = [root]
        while stk:
            node = stk.pop()
            if node:
                val = node.val
                stk.append(node.right)
                stk.append(node.left)
            else:
                val = "N"

            res.append(str(val))

        return self.DELIMITER.join(res)

    def deserialize(self, data):
        if not data: return []
        res = data.split(self.DELIMITER)
        root = TreeNode(res[0])
        stk = [(root, False)] # node, left processed?
        idx = 1
        while stk and idx < len(res):
            node, processed = stk.pop()
            val = res[idx]
            if not processed: # process left first
                if val != "N":
                    node.left = TreeNode(int(val))
                    stk.append((node, True))
                    stk.append((node.left, False))
                else:
                    stk.append((node, True))
                idx+=1
            else: # process right else
                if val != "N":
                    node.right = TreeNode(int(val))
                    stk.append((node.right, False))
                idx+=1
        return root
