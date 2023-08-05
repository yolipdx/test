a = 100;

class Dog {
    constructor(name) {
        this.name = name;
        b = 100;
        // ReferenceError: b is not defined
        console.log("b in constructor: ", b);
    }
}

console.log("global a: ", globalThis.a);
console.log("global b: ", globalThis.b);

let d = new Dog("kok");