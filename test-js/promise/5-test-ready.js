console.log('\nin main 0');

let p = new Promise((resolve) => {
    resolve(5);
});

p.then(console.log);

console.log('\nin main 1');