const a = [
    1,
    [
        2,
        3,
        [
            4, 
            5, 
            [6, 7]
        ]
    ],
    8,
    [
        9,
        [10, 11],
        12,
    ]
];

console.log(a.flat(2));
console.log(a.flat(3));
console.log(a.flat(4));
console.log(a.flat(5));

const d = [
    1,  2, 3, 4,  5,
    6,  7, 8, 9, 10,
   11, 12
 ];

 console.log(d);