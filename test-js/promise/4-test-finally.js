let p = new Promise((resolve, reject) => {
    setTimeout(() => {
        console.log('to resolve');
        resolve('done');
        // reject(new Error('err 0'));
    }, 4000);
});

p.finally(() => {
    console.log('clean up');
    // throw Error('err 1');
    return 'done in fanally';
})
.then(
    (x) => {
        console.log('then result: ', x);
    },
    (err) => {
        console.log('err: ', err);
    }
)
