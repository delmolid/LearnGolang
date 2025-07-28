// function Hello (prenom) {
//    return ("Bonjour, " + prenom)
// }

// console.log(Hello("Molid"))

function sayHello(name) {
return `Bonjour, ${name} !`;
}

const esTuMajeur = (age) => {
   if (age > 18) {
      return " Tu es majeur "
   } else {
      return " Tu es mineur "
   }
}

module.exports = { sayHello, esTuMajeur }
