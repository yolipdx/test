new Promise((resolve, reject) => {
    resolve(1);    
}).finally(() => {
    console.log("clean up: finally");
    return 'v in finally';
}).then(console.log);

let a = Promise.resolve(1);
let b = a.finally(( () => console.log("with finally")));