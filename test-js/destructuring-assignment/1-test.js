let name1 = 'j';
let age = -1;

// error
// ({name, age}) = {
//     name: 'mk',
//     age: 100,
// };


({name1, age} = {
    name1: 'mk',
    age: 100,
});


console.log(name1);
console.log(age);