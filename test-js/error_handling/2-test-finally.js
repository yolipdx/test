function add(a, b) {
    try {
        // dd
        return a + b;
    } catch (err) {
        console.error(err);
    } finally {
        console.log('clean up');
    }
}

console.log(add(1, 2));