let user = {
    firstName: "John",
    sayHi() {
      console.log(`Hello, ${this.firstName}!`);
    }
  };
  
let s1 = user.sayHi.bind(user);
s1();

s1.apply({
    firstName: "kkkk",
});

user.sayHi.apply({
    firstName: "kkkk",
});