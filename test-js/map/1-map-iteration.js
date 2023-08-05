let m = new Map([
    ['name', 'jk'],
    ['age', 100],
]);

console.log(m);

for (let x of m) {
    console.log(x);
}

m.forEach((v, k) => {
    console.log('map item: ', k, v);
});
