for i = 1, 10, 2
  do
    print(i)
  end
local obj = { }
obj[1] = 3.14
obj[2] = "nihaop"
-- print(obj[1])
-- print(obj["1"])
print(obj[2.0])
-- assert(obj[1] == 13.14)