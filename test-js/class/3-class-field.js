class Dog {
    type = "dog";

    constructor(name) {
        this.name = name;
    }
}

let d = new Dog("kok");
console.log(Object.keys(d));
console.log(Object.keys(d.__proto__));