func test(var *a) {
    *a = 4
}

var b = 2;
test(&b);

return b;
