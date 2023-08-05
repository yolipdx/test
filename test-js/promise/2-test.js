let p0 = new Promise((resolve, reject) => {
    return 1;
}).then(() => {
    return 1;
});

let p1= new Promise((resolve, reject) => {
    return 1;
}).catch((err) => {
    return 1;
});

let p2 = new Promise((resolve, reject) => {
    return 1;
}).finally((err) => {
    return 1;
});

console.log(p0 instanceof Promise);
console.log(p1 instanceof Promise);
console.log(p2 instanceof Promise);