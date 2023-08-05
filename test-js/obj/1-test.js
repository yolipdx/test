let animal = {
    type: 'animal',
};

let duck = {
    color: 'red',
    __proto__: animal,
};


for (let key in duck) {
    console.log(key, duck[key]);
    console.log(duck.hasOwnProperty(key));
}

console.log("object functions:");
console.log(Object.keys(duck));

let jj = new Map();
jj.set('name', 100);
console.log([...jj]);