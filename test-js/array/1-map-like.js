let b = [1, 2, 3, 4, undefined, undefined, 7];

let c = b.map((x, i) => {
    console.log(`index: ${i} => ${x}`);
    return x * 2;
});
console.log("c size: ", c.length);

console.log("\n\n");
b = [1, 2, 3, 4];
b[6] = 7;

c = b.map((x, i) => {
    console.log(`index: ${i} => ${x}`);
    return x * 2;
});
console.log("c size: ", c.length);


console.log("\n\nfor each:");
b.forEach((x, i) => {
    console.log(`index: ${i} => ${x}`);
    return x * 2;
});


console.log("\n\nfilter:");
b.filter((x, i) => {
    console.log(`index: ${i} => ${x}`);
    return x >= 2;
});



console.log("\n\nreduce:");
console.log(b);
const sum1 = b.reduce((acc, cur, i) => {
    console.log(`on index: ${i} => ${cur}`);
    return acc + cur;
});
console.log("sum: ", sum1);


console.log("\n\nreduce2:");
console.log(b);
const sum2 = b.reduce((acc, cur, i) => {
    console.log(`on index: ${i} => ${cur}`);
    return acc + cur;
}, 0);
console.log("sum: ", sum2);


console.log([].reduce((acc, cur) => acc + cur));