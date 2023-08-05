let p = new Promise((resovel, reject) => {
    return 1;
});


let p1 = p.then((x) => x);
let p2 = p.then(console.log);
console.log(p1 === p);
console.log(p2 === p);
console.log(p.finally(console.log)=== p);