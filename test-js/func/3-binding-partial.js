function add(a, b, c, d, e, f) {
    return Array.from(arguments).reduce((acc, cur) => acc + cur, 0);
}

function partial(fn) {

    const fnLen = fn.length;
    console.log('fnLen: ', fnLen);
    return function doWork(...args) {
        if (args.length >= fnLen) {
            return fn.apply(this, args.slice(0, fnLen - 1));
        }

        // // only partial args are supplied, a function shoud be created carrying on these args
        // return  (...furtherArgs) => {
        //     // return fn.call(this, ...args, ...furtherArgs);
        //     return doWork.call(this, ...args, ...furtherArgs);
        // };

        return doWork.bind(this, ...args);
    };

}

let p = partial(add);
// console.log(p(1, 2, 3, 4, 5, 6));
// console.log(p(1, 2, 3, 4, 5, 6, 7));
// console.log(p(1, 2, 3, 4, 5, 6, 7, 8));
console.log(p(1)(2)(3)(4)(5)(6));

// console.log(p(1, 2, 3, 4));
console.log(p(1, 2, 3, 4)(5, 6));
console.log(p(1, 2, 3)(4, 5, 6));
console.log(p(1, 2, 3)(4, 5)(6));