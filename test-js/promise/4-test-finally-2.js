new Promise((resolve, reject) => {
    setTimeout(() => {
        resolve(100);
    }, 4000);
}).finally(() => {
    console.log("throw an error in finally");
    throw new Error("error in finally");
}).then(x => {
    console.log("then: ", x);
}).catch(e => {
    console.log("catch: ", e);
});