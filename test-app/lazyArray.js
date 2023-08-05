// 第二轮：LazyArray 大致的题目如下
// LazyArray a;
// a.map(std::function<int>(int) func).indexOf(num)
// 要考虑到多个map的call，这种情况要把多个function chain起来

// 实现lazyArray 就是实现java里面的array.map(a->XXX).indexOf(); 要求map操作推迟到最后才做。所以叫lazy

class LazyArray extends Array {
    constructor() {

    }

    map(f) {

    }

    indexOf(num) {
        return -1;
    }
}

