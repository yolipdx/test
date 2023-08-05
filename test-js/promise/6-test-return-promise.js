let p = new Promise((resolve, reject) => {
    setTimeout(() => {
        console.log('resolve 10');
        resolve(10);
    }, 4000);
})
    .finally(() => {
        console.log("clean up 0");
        return new Promise((resolve) => {
            setTimeout(() => {
                console.log('timer in finally');
                resolve(2000);
            }, 5000)
        });
    })
    .then((x) => {
        console.log('x: ', x);
    });
;