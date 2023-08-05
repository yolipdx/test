async function run() {

    async function fetchData() {
        return [1, 2, 3];
    }

    let a = await fetchData();
    console.log('data: ', a);

}

run();