import { test, expect } from "vitest";
import { FizzBuzz } from "script.js"

test('Cas de démarrage #1', () => {
  expect(FizzBuzz(1)).toBe(1)
})

test('Cas de démarrage #2', () => {
  expect(FizzBuzz(2)).toBe(2)
})

test ('Test FizzBuzz : Vérification si le nombre est divisible par 3 ', () => {
    expect(FizzBuzz(9)).toBe("Fizz")
}) 

test ('Test FizzBuzz : Vérification si le nombre est égale a 5 ', () => {
    expect(FizzBuzz(5)).toBe("Buzz")
})

test ('Test FizzBuzz : Vérification si le nombre est divisible par 3 et le reste est 5', () => {
    expect(FizzBuzz(15)).toBe("FizzBuzz")
})

test('Cas de démarrage #6', () => {
  expect(FizzBuzz(6)).toBe("Fizz")
})

