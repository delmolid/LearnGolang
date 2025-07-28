import { expect, test, it } from 'vitest'
import { esTuMajeur, sayHello } from './Hello'

test('Cas de démarrage #1 - Dire bonjour avec un prénom', () => {
  expect(sayHello("Laïla")).toBe("Bonjour, Laïla !")
 
})

test('Cas de démarrage #1 - Dire bonjour avec un prénom', () => {
  expect(sayHello(42)).toBe("Bonjour, 42 !")
})

test('Cas de démarrage #1 - Dire bonjour avec un prénom', () => {
  expect(sayHello()).toBe("Bonjour, undefined !")
  
})



test('Cas de démarrage #1 - Dire bonjour avec un prénom', () => {
  expect(sayHello("saluttttttttttttttt")).toBe("Bonjour, saluttttttttttttttt !")
  
})

test('Cas de démarrage #1 - Dire bonjour avec un prénom', () => {
  expect(sayHello(null)).toBe("Bonjour, null !")
  
})

test('Cas ou la personne est mineur', () => {
    expect(esTuMajeur(20)).toBe(" Tu es majeur ")

})

it('Cas ou la personne est mineur', () => {
    expect(esTuMajeur(10)).toBe(" Tu es mineur ")

})