import fetch from 'node-fetch'

// Afficher la météo en appelant deux APIs : une pour récupérer les coordonnées de la ville, l'autre pour récupérer la météo
export async function displayMeteo(city) {
  const responseCoordinates = await fetchCoordinates(city)
  const coordinates = responseCoordinates.results[0]
  const data = await fetchWeather(coordinates.latitude, coordinates.longitude)
  console.log(
    `Les températures du jour de ${city} sont ${data.hourly.temperature_2m}`
  )
}

export async function fetchCoordinates(city) {
  const response = await fetch(
    `https://geocoding-api.open-meteo.com/v1/search?name=${city}`
  )
  return await response.json()
}

export async function fetchWeather(lat, lng) {
  const response = await fetch(
    `https://api.open-meteo.com/v1/forecast?latitude=${lat}&longitude=${lng}&hourly=temperature_2m&format=json&timeformat=unixtime`
  )
  return await response.json()
}

