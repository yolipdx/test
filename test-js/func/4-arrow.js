function defer(f, ms) {

    return function(...args) {
        setTimeout(() => {
            f.apply(this, args)
        }, ms);
    };
};

function sayHi(who) {
    console.log('Hello, ' + who);
}
  
let sayHiDeferred = defer(sayHi, 2000);
sayHiDeferred("John"); // Hello, John after 2 seconds