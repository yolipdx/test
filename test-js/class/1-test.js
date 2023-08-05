class Dog {
    age = 10;
    #speed = 20;

    constructor(name) {
        this.name = name;
    }


    bark() {
        console.log(`dog ${this.name}:@${this.age} with speed:${this.#speed} is barking`);
    }
    
}


const d = new Dog("kok");
d.bark();
console.log(d.age);
// console.log(d.#speed);