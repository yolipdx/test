let p = new Promise((resolve, reject) => {
    setTimeout(() => {
        resolve(`resolved at ${Date.now()}`);
    },  5000);
}).finally( ()=> {
    console.log('clean up 0');
    return 0;
}).then((v) => {
    console.log('then: ', v);
    return 1;
}).catch((err) => {
    console.log("error: ", err);
    return 2;
}).finally(() => {
    console.log('clean up 1');
    return 3;
}).then((r) => {
    console.log('r: ', r);
});
