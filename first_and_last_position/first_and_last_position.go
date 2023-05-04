/*
input: array of integer nums sorted and target
result: [start, end] indexes of target

binary search function returns index of found value, or -1 if not found

if binary search on input returns -1 (M)
  return [-1, -1]

binary search [M+1 until last idx] (L)
if it returns -1
	end = M
else
  end = L

binary search [first idx until M-1] (R)
if it returns -1
  start = M
else
  start = R

return [start, end]
*/
