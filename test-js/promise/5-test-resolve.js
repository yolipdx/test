console.log("=> in main");
Promise.resolve(1).then(x => {
    console.log('process resolved promise with v: ', x);
});
console.log("=> out main");