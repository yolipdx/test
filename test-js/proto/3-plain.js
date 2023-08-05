let dict0 = {};

dict0['__proto__'] = 123;
console.log(dict0);





let dict = Object.create(null);

dict['__proto__'] = 123;
console.log(dict);
