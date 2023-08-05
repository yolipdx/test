function add(a, b, c, d, e, f) {
    return Array.from(arguments).reduce((acc, cur) => acc + cur, 0);
}

console.log(add(1, 2, 3));

let add1 = add.bind(null, 1);
console.log(add1(4));


let add2 = add.bind(null, 1, 3);
console.log(add2(4));


let d = -100;
let add_d = add.bind(null, d);
console.log(add_d(4));

d = 100;
console.log(add_d(4));