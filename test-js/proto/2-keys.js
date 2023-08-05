let animal = {
    eats: true
};

let rabbit = {
    jumps: true,
    __proto__: animal
};


console.log(Object.keys(rabbit));
for (let k in rabbit) {
    console.log("k: ", k);
}