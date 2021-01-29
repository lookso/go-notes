local tab = { 1, 2, 3 }

if next(tab) == nil then
    print("tab is nil")
else
    for k, v in ipairs(tab) do
        print("k-v", k, v)
    end
end
