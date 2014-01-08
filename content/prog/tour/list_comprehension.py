squares = [x**2 for x in range(10)]
print squares

couples = [(x, y) for x in [1,2,3] for y in [3,1,4] if x != y]
print couples
