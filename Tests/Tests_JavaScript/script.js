const FizzBuzz = (number) => {
if ((number % 3) == 0 && number % 5 == 0) {
 return "FizzBuzz"
} else if (number % 3 == 0) {
    return "Fizz"
} else if (number === 5) {
    return "Buzz"
} 
return number
}

console.log(FizzBuzz(15))

module.exports = { FizzBuzz } 

